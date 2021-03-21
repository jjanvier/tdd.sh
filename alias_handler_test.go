package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

type successCommandExecutorMock struct {
	mock.Mock
}

func (executor successCommandExecutorMock) ExecuteWithOutput(cmd Command) error {
	return nil
}

func (executor successCommandExecutorMock) Execute(cmd Command) error {
	return nil
}

type errorCommandExecutorMock struct {
	mock.Mock
}

func (executor errorCommandExecutorMock) ExecuteWithOutput(cmd Command) error {
	return errors.New("an error occurred during the execution of the command")
}

func (executor errorCommandExecutorMock) Execute(cmd Command) error {
	return errors.New("an error occurred during the execution of the command")
}

func TestHandleAliasCommandWhenTestsPass(t *testing.T) {
	cmd := Command{"go", []string{"test", "-v"}}
	executor := new(successCommandExecutorMock)

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}}
	result := handler.HandleTestCommand(cmd)

	assert.Equal(t, "go test -v && git add . && git commit --reuse-message=HEAD", result.Command)
	assert.Equal(t, 0, result.ExitCode)
}

func TestHandleAliasCommandWhenTestsDoNotPass(t *testing.T) {
	cmd := Command{"go", []string{"test", "-v"}}
	executor := new(errorCommandExecutorMock)

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}}
	result := handler.HandleTestCommand(cmd)

	assert.Equal(t, "go test -v && git add . && git commit --reuse-message=HEAD", result.Command)
	assert.Equal(t, 1, result.ExitCode)
}

func TestCreateExecutionResultGreen(t *testing.T) {
	factory := ExecutionResultFactory{}
	result := factory.CreateExecutionResultGreen([]Command{
		{"toto", []string{"titi", "--tata"}},
		{"foo", []string{"bar", "baz"}},
	})

	assert.Equal(t, "toto titi --tata && foo bar baz", result.Command)
	assert.Equal(t, 0, result.ExitCode)
}

func TestCreateExecutionResultRed(t *testing.T) {
	factory := ExecutionResultFactory{}
	result := factory.CreateExecutionResultRed([]Command{
		{"toto", []string{"titi", "--tata"}},
		{"foo", []string{"bar", "baz"}},
	})

	assert.Equal(t, "toto titi --tata && foo bar baz", result.Command)
	assert.Equal(t, 1, result.ExitCode)
}
