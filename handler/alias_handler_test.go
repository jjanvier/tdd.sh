package handler

import (
	"github.com/jjanvier/tdd/execution"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

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
	center.On("Reset")

	handler := _createSuccessAliasHandler(center)
	result, _ := handler.HandleAlias(conf, "foo")

	assert.Equal(t, "go test -v && git add . && git commit --reuse-message=HEAD", result.Command)
	assert.Equal(t, true, result.IsSuccess)

	center.AssertNumberOfCalls(t, "Reset", 1)
}

func TestHandleAliasCommandWhenTestsPassAndCommitsAreAmended(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"go test -v", 120, Git{true}}
	conf.Aliases = aliases

	center := new(notificationsCenterMock)
	center.On("Reset")

	handler := _createSuccessAliasHandler(center)
	result, _ := handler.HandleAlias(conf, "foo")

	assert.Equal(t, "go test -v && git add . && git commit --amend --no-edit", result.Command)
	assert.Equal(t, true, result.IsSuccess)
	center.AssertNumberOfCalls(t, "Reset", 1)
}

func TestHandleAliasCommandWhenTestsDoNotPass(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"go test -v", 120, Git{false}}
	conf.Aliases = aliases

	center := new(notificationsCenterMock)
	center.On("Reset")
	center.On("NotifyWithDelay")

	handler := _createErrorAliasHandler(center)
	result, err := handler.HandleAlias(conf, "foo")

	assert.Equal(t, "go test -v", result.Command)
	assert.Equal(t, false, result.IsSuccess)
	assert.Nil(t, err)
	center.AssertNumberOfCalls(t, "Reset", 1)
	center.AssertNumberOfCalls(t, "NotifyWithDelay", 1)
}

func TestHandleAliasCommandWhenCommandDoNotExist(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"doesnotexit", 120, Git{false}}
	conf.Aliases = aliases

	center := new(notificationsCenterMock)
	center.On("Reset")

	handler := _createUnknownCommandErrorAliasHandler(center)
	result, err := handler.HandleAlias(conf, "foo")

	assert.Equal(t, "doesnotexit", result.Command)
	assert.Equal(t, false, result.IsSuccess)
	assert.IsType(t, &execution.UnknownCommandError{}, err)
	center.AssertNumberOfCalls(t, "Reset", 1)
	center.AssertNumberOfCalls(t, "NotifyWithDelay", 0)
}

func _createSuccessAliasHandler(center *notificationsCenterMock) AliasHandler {
	executor := new(execution.SuccessCommandExecutorMock)
	commandFactory := execution.CommandFactory{}
	executionResultFactory := execution.ExecutionResultFactory{}

	return AliasHandler{executor, commandFactory, executionResultFactory, center}
}

func _createErrorAliasHandler(center *notificationsCenterMock) AliasHandler {
	executor := new(execution.ErrorCommandExecutorMock)
	commandFactory := execution.CommandFactory{}
	executionResultFactory := execution.ExecutionResultFactory{}

	return AliasHandler{executor, commandFactory, executionResultFactory, center}
}

func _createUnknownCommandErrorAliasHandler(center *notificationsCenterMock) AliasHandler {
	executor := new(execution.UnknownCommandExecutorMock)
	commandFactory := execution.CommandFactory{}
	executionResultFactory := execution.ExecutionResultFactory{}

	return AliasHandler{executor, commandFactory, executionResultFactory, center}
}
