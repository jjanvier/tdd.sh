package handler

import "github.com/jjanvier/tdd/execution"

type NewHandlerI interface {
	HandleNew(message string) (execution.ExecutionResult, error)
}

type NewHandler struct {
	Executor               execution.CommandExecutorI
	CommandFactory         execution.CommandFactory
	ExecutionResultFactory execution.ExecutionResultFactory
}

func (handler NewHandler) HandleNew(message string) (execution.ExecutionResult, error) {
	cmd := handler.CommandFactory.CreateGitCommitEmpty(message)
	err := handler.Executor.Execute(cmd)

	return handler.ExecutionResultFactory.Success([]execution.Command{cmd}), err
}
