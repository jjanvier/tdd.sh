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

type notificationsCenterMock struct {
	mock.Mock
}

func (center *notificationsCenterMock) Notify(message string) {
	center.Called()
}

func (center *notificationsCenterMock) Alert(message string) {
	center.Called()
}

func TestHandleAliasCommandWhenTestsPass(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"go test -v", 120, Git{false}}
	conf.Aliases = aliases

	executor := new(successCommandExecutorMock)

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}, NotificationsCenter{}}
	result, _ := handler.HandleTestCommand(conf, "foo")

	assert.Equal(t, "go test -v && git add . && git commit --reuse-message=HEAD", result.Command)
	assert.Equal(t, true, result.IsSuccess)
}

func TestHandleAliasCommandWhenTestsPassAndCommitsAreAmended(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"go test -v", 120, Git{true}}
	conf.Aliases = aliases

	executor := new(successCommandExecutorMock)

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}, NotificationsCenter{}}
	result, _ := handler.HandleTestCommand(conf, "foo")

	assert.Equal(t, "go test -v && git add . && git commit --amend --no-edit", result.Command)
	assert.Equal(t, true, result.IsSuccess)
}

func TestHandleAliasCommandWhenTestsDoNotPass(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"go test -v", 120, Git{false}}
	conf.Aliases = aliases

	executor := new(errorCommandExecutorMock)
	center := new(notificationsCenterMock)
	center.On("Alert").Once()

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}, center}
	result, _ := handler.HandleTestCommand(conf, "foo")

	assert.Equal(t, "go test -v && git add . && git commit --reuse-message=HEAD", result.Command)
	assert.Equal(t, false, result.IsSuccess)
	center.AssertExpectations(t)
}

func TestHandleNew(t *testing.T) {
	executor := new(successCommandExecutorMock)

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}, NotificationsCenter{}}
	result, _ := handler.HandleNew("here is my commit message")

	assert.Equal(t, "git commit --allow-empty -m here is my commit message", result.Command)
	assert.Equal(t, true, result.IsSuccess)
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
