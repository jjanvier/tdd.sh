package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

const commandPid = 100

type successCommandExecutorMock struct {
	mock.Mock
}

func (executor successCommandExecutorMock) ExecuteWithOutput(cmd Command) error {
	return nil
}

func (executor successCommandExecutorMock) Execute(cmd Command) error {
	return nil
}

func (executor successCommandExecutorMock) ExecuteBackground(cmd Command) (int, error) {
	return commandPid, nil
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

func (executor errorCommandExecutorMock) ExecuteBackground(cmd Command) (int, error) {
	return 0, errors.New("an error occurred during the execution of the command")
}

type notificationsCenterMock struct {
	mock.Mock
}

func (center *notificationsCenterMock) NotifyWithDelay(alias string, delay int, message string) {
	center.Called()
}

func (center *notificationsCenterMock) Reset(alias string) {
	center.Called()
}

func TestHandleAliasCommandWhenTestsPass(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"go test -v", 120, Git{false}}
	conf.Aliases = aliases

	center := new(notificationsCenterMock)
	center.On("Reset").Once()

	handler := _createSuccessAliasHandler(center)
	result, _ := handler.HandleAlias(conf, "foo")

	assert.Equal(t, "go test -v && git add . && git commit --reuse-message=HEAD", result.Command)
	assert.Equal(t, true, result.IsSuccess)
	center.AssertExpectations(t)
}

func TestHandleAliasCommandWhenTestsPassAndCommitsAreAmended(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"go test -v", 120, Git{true}}
	conf.Aliases = aliases

	center := new(notificationsCenterMock)
	center.On("Reset").Once()

	handler := _createSuccessAliasHandler(center)
	result, _ := handler.HandleAlias(conf, "foo")

	assert.Equal(t, "go test -v && git add . && git commit --amend --no-edit", result.Command)
	assert.Equal(t, true, result.IsSuccess)
	center.AssertExpectations(t)
}

func TestHandleAliasCommandWhenTestsDoNotPass(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"go test -v", 120, Git{false}}
	conf.Aliases = aliases

	center := new(notificationsCenterMock)
	center.On("Reset").Once()
	center.On("NotifyWithDelay").Once()

	handler := _createErrorAliasHandler(center)
	result, _ := handler.HandleAlias(conf, "foo")

	assert.Equal(t, "go test -v && git add . && git commit --reuse-message=HEAD", result.Command)
	assert.Equal(t, false, result.IsSuccess)
	center.AssertExpectations(t)
}

func TestCreateExecutionResultSuccess(t *testing.T) {
	factory := ExecutionResultFactory{}
	result := factory.CreateExecutionResultSuccess([]Command{
		{"toto", []string{"titi", "--tata"}},
		{"foo", []string{"bar", "baz"}},
	})

	assert.Equal(t, "toto titi --tata && foo bar baz", result.Command)
	assert.Equal(t, true, result.IsSuccess)
}

func TestCreateExecutionResultFailure(t *testing.T) {
	factory := ExecutionResultFactory{}
	result := factory.CreateExecutionResultFailure([]Command{
		{"toto", []string{"titi", "--tata"}},
		{"foo", []string{"bar", "baz"}},
	})

	assert.Equal(t, "toto titi --tata && foo bar baz", result.Command)
	assert.Equal(t, false, result.IsSuccess)
}

func _createSuccessAliasHandler(center *notificationsCenterMock) AliasHandler {
	executor := new(successCommandExecutorMock)
	commandFactory := CommandFactory{}
	executionResultFactory := ExecutionResultFactory{}

	return AliasHandler{executor, commandFactory, executionResultFactory, center}
}

func _createErrorAliasHandler(center *notificationsCenterMock) AliasHandler {
	executor := new(errorCommandExecutorMock)
	commandFactory := CommandFactory{}
	executionResultFactory := ExecutionResultFactory{}

	return AliasHandler{executor, commandFactory, executionResultFactory, center}
}
