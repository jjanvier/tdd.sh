package main

import (
	"github.com/manifoldco/promptui"
	"io"
	"io/ioutil"
	"os"
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
	HandleTodo(message string, todoFile string) (ExecutionResult, error)
	HandleDo(todoFile string, stdin io.ReadCloser) (ExecutionResult, error)
	HandleDone(todoFile string) (ExecutionResult, error)
}

type AliasHandler struct {
	executor               CommandExecutorI
	commandFactory         CommandFactory
	executionResultFactory ExecutionResultFactory
	notificationsCenter    NotificationCenterI
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

func (handler AliasHandler) HandleTodo(message string, todoFilePath string) (ExecutionResult, error) {
	fakeTodoCommand := Command{message, []string{}}
	fakeTodoCommands := []Command{fakeTodoCommand}

	todoFile, err := os.OpenFile(todoFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer todoFile.Close()
	if err != nil {
		return handler.executionResultFactory.CreateExecutionResultFailure(fakeTodoCommands), err
	}

	_, err2 := todoFile.WriteString(message + "\n")
	if err2 != nil {
		return handler.executionResultFactory.CreateExecutionResultFailure(fakeTodoCommands), err2
	}

	return handler.executionResultFactory.CreateExecutionResultSuccess(fakeTodoCommands), nil
}

func (handler AliasHandler) HandleDo(todoFilePath string, stdin io.ReadCloser) (ExecutionResult, error) {
	todoContent, _ := ioutil.ReadFile(todoFilePath)
	todoList := strings.Split(string(todoContent), "\n")

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

func (handler AliasHandler) HandleDone(todoFilePath string) (ExecutionResult, error) {
	todoFile, err := os.OpenFile(todoFilePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	defer todoFile.Close()

	if err != nil {
		return handler.executionResultFactory.CreateExecutionResultFailure([]Command{}), err
	}

	return handler.executionResultFactory.CreateExecutionResultSuccess([]Command{}), nil
}
