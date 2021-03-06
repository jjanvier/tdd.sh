package handler

import (
	"errors"
	"github.com/jjanvier/tdd/execution"
	"github.com/jjanvier/tdd/notification"
)

const notificationMessage = "The time is up and the tests are still red! Maybe you should reset your changes and take a smaller step?"

type AliasHandlerI interface {
	HandleAlias(conf Configuration, alias string) (execution.ExecutionResult, error)
}

type AliasHandler struct {
	Executor               execution.CommandExecutorI
	CommandFactory         execution.CommandFactory
	ExecutionResultFactory execution.ExecutionResultFactory
	NotificationsCenter    notification.NotificationCenterI
}

func (handler AliasHandler) HandleAlias(conf Configuration, alias string) (execution.ExecutionResult, error) {
	testCmd, err := conf.getAliasCommand(alias)
	if err != nil {
		return execution.ExecutionResult{}, err
	}

	gitAddCmd, err := conf.getGitAddCommand(alias)
	if err != nil {
		return execution.ExecutionResult{}, err
	}

	gitCommitCmd := handler.CommandFactory.CreateGitCommit()
	if amend, _ := conf.shouldAmendCommits(alias); amend {
		gitCommitCmd = handler.CommandFactory.CreateGitCommitAmend()
	}

	handler.NotificationsCenter.Reset(alias)

	if err := handler.Executor.ExecuteWithLiveOutput(testCmd); err != nil {
		timer, _ := conf.getTimer(alias)
		var unknownCommandError *execution.UnknownCommandError
		if errors.As(err, &unknownCommandError) {
			return handler.ExecutionResultFactory.Failure([]execution.Command{testCmd}), err
		}

		handler.NotificationsCenter.NotifyWithDelay(alias, timer, notificationMessage)
		return handler.ExecutionResultFactory.Failure([]execution.Command{testCmd}), nil
	}

	if err := handler.Executor.ExecuteWithLiveOutput(gitAddCmd); err != nil {
		return handler.ExecutionResultFactory.Failure([]execution.Command{gitAddCmd}), err
	}

	if err := handler.Executor.ExecuteWithLiveOutput(gitCommitCmd); err != nil {
		return handler.ExecutionResultFactory.Failure([]execution.Command{gitCommitCmd}), err
	}

	return handler.ExecutionResultFactory.Success([]execution.Command{testCmd, gitAddCmd, gitCommitCmd}), nil
}
