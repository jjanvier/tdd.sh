package main

import "strings"

type ExecutionResult struct {
	ExitCode int
	Command  string
}

type ExecutionResultFactory struct {}

func (factory ExecutionResultFactory) CreateExecutionResultGreen(cmds []Command) ExecutionResult {
	return ExecutionResult{0, factory.joinCommands(cmds)}
}

func (factory ExecutionResultFactory) CreateExecutionResultRed(cmds []Command) ExecutionResult {
	return ExecutionResult{1, factory.joinCommands(cmds)}
}

func (factory ExecutionResultFactory) joinCommands(cmds []Command) string {
	var cmdsString []string
	for _, cmd := range cmds {
		cmdsString = append(cmdsString, cmd.String())
	}

	return strings.Join(cmdsString, " && ")
}

type AliasHandlerI interface {
	HandleTestCommand(testCmd Command) ExecutionResult
	HandleNew(message string) ExecutionResult
}

type AliasHandler struct {
	executor CommandExecutorI
	commandFactory CommandFactory
	executionResultFactory ExecutionResultFactory
}

func (handler AliasHandler) HandleTestCommand(testCmd Command) ExecutionResult {
	gitAddCmd := handler.commandFactory.CreateGitAdd()
	gitCommitCmd := handler.commandFactory.CreateGitCommit()

	if handler.executor.ExecuteWithOutput(testCmd) != nil {
		return handler.executionResultFactory.CreateExecutionResultRed([]Command{testCmd, gitAddCmd, gitCommitCmd})
	}

	if handler.executor.ExecuteWithOutput(gitAddCmd) != nil || handler.executor.ExecuteWithOutput(gitCommitCmd) != nil {
		return handler.executionResultFactory.CreateExecutionResultRed([]Command{testCmd, gitAddCmd, gitCommitCmd})
	}

	return handler.executionResultFactory.CreateExecutionResultGreen([]Command{testCmd, gitAddCmd, gitCommitCmd})
}

func (handler AliasHandler) HandleNew(message string) ExecutionResult {
	cmd := handler.commandFactory.CreateGitCommitEmpty(message)
	handler.executor.Execute(cmd)

	return handler.executionResultFactory.CreateExecutionResultGreen([]Command{cmd})
}
