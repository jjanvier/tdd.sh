package main

type NewHandlerI interface {
	HandleNew(message string) (ExecutionResult, error)
}

type NewHandler struct {
	executor               CommandExecutorI
	commandFactory         CommandFactory
	executionResultFactory ExecutionResultFactory
}

func (handler NewHandler) HandleNew(message string) (ExecutionResult, error) {
	cmd := handler.commandFactory.CreateGitCommitEmpty(message)
	err := handler.executor.Execute(cmd)

	return handler.executionResultFactory.CreateExecutionResultSuccess([]Command{cmd}), err
}
