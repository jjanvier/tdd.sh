package main

import (
	"strings"
)

type ExecutionResult struct {
	IsSuccess bool
	Command   string
}

type ExecutionResultFactory struct {}

func (factory ExecutionResultFactory) CreateExecutionResultSuccess(cmds []Command) ExecutionResult {
	return ExecutionResult{true, factory.joinCommands(cmds)}
}

func (factory ExecutionResultFactory) CreateExecutionResultFailure(cmds []Command) ExecutionResult {
	return ExecutionResult{false, factory.joinCommands(cmds)}
}

func (factory ExecutionResultFactory) joinCommands(cmds []Command) string {
	var cmdsString []string
	for _, cmd := range cmds {
		cmdsString = append(cmdsString, cmd.String())
	}

	return strings.Join(cmdsString, " && ")
}

type AliasHandlerI interface {
	HandleTestCommand(conf Configuration, alias string) (ExecutionResult, error)
	HandleNew(message string) (ExecutionResult, error)
}

type AliasHandler struct {
	executor CommandExecutorI
	commandFactory CommandFactory
	executionResultFactory ExecutionResultFactory
}

func (handler AliasHandler) HandleTestCommand(conf Configuration, alias string) (ExecutionResult, error) {
	testCmd, err := conf.GetCommand(alias)
	if err != nil {
		return ExecutionResult{}, err
	}

	gitAddCmd := handler.commandFactory.CreateGitAdd()
	gitCommitCmd := handler.commandFactory.CreateGitCommit()
	if amend, _ := conf.ShouldAmendCommits(alias); amend {
		gitCommitCmd = handler.commandFactory.CreateGitCommitAmend()
	}

	if handler.executor.ExecuteWithOutput(testCmd) != nil {
		return handler.executionResultFactory.CreateExecutionResultFailure([]Command{testCmd, gitAddCmd, gitCommitCmd}), nil
	}

	if err := handler.executor.ExecuteWithOutput(gitAddCmd); err != nil {
		return handler.executionResultFactory.CreateExecutionResultFailure([]Command{testCmd, gitAddCmd}), err
	}

	if err := handler.executor.ExecuteWithOutput(gitCommitCmd); err != nil {
		return handler.executionResultFactory.CreateExecutionResultFailure([]Command{testCmd, gitAddCmd, gitCommitCmd}), err
	}

	return handler.executionResultFactory.CreateExecutionResultSuccess([]Command{testCmd, gitAddCmd, gitCommitCmd}), nil
}

func (handler AliasHandler) HandleNew(message string) (ExecutionResult, error) {
	cmd := handler.commandFactory.CreateGitCommitEmpty(message)
	err := handler.executor.Execute(cmd)

	return handler.executionResultFactory.CreateExecutionResultSuccess([]Command{cmd}), err
}
