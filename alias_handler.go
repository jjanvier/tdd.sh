package main

import (
	"github.com/manifoldco/promptui"
	"io"
	"strings"
)

const notificationMessage = "The time is up and the tests are still red! Maybe you should reset your changes and take a smaller step?"

type ExecutionResult struct {
	IsSuccess bool
	Command   string
}

type ExecutionResultFactory struct{}

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
	HandleTodo(message string) (ExecutionResult, error)
	HandleDo(stdin io.ReadCloser) (ExecutionResult, error)
	HandleDone() (ExecutionResult, error)
}

type AliasHandler struct {
	executor               CommandExecutorI
	commandFactory         CommandFactory
	executionResultFactory ExecutionResultFactory
	notificationsCenter    NotificationCenterI
	todo                   TodoList
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

	handler.notificationsCenter.Reset(alias)

	if handler.executor.ExecuteWithOutput(testCmd) != nil {
		timer, _ := conf.GetTimer(alias)
		handler.notificationsCenter.NotifyWithDelay(alias, timer, notificationMessage)
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

func (handler AliasHandler) HandleTodo(message string) (ExecutionResult, error) {
	err := handler.todo.Add(message)
	if err != nil {
		return handler.executionResultFactory.CreateExecutionResultFailure([]Command{}), err
	}

	return handler.executionResultFactory.CreateExecutionResultSuccess([]Command{}), nil
}

func (handler AliasHandler) HandleDo(stdin io.ReadCloser) (ExecutionResult, error) {
	todoList, err := handler.todo.GetItems()
	if err != nil {
		return handler.executionResultFactory.CreateExecutionResultFailure([]Command{}), err
	}

	prompt := promptui.Select{
		Label: "Here is your todo list, which task do you want to tackle?",
		Items: todoList,
		Stdin: stdin,
	}

	_, selected, err := prompt.Run()

	if selected == "" {
		return handler.executionResultFactory.CreateExecutionResultFailure([]Command{}), err
	}

	if err != nil {
		return handler.executionResultFactory.CreateExecutionResultFailure([]Command{}), err
	}

	return handler.HandleNew(selected)
}

func (handler AliasHandler) HandleDone() (ExecutionResult, error) {
	err := handler.todo.Clear()
	if err != nil {
		return handler.executionResultFactory.CreateExecutionResultFailure([]Command{}), err
	}

	return handler.executionResultFactory.CreateExecutionResultSuccess([]Command{}), nil
}
