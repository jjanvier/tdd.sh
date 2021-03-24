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
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"go test -v", 120, Git{false}}
	conf.Aliases = aliases

	executor := new(successCommandExecutorMock)

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}}
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

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}}
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

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}}
	result, _ := handler.HandleTestCommand(conf, "foo")

	assert.Equal(t, "go test -v && git add . && git commit --reuse-message=HEAD", result.Command)
	assert.Equal(t, false, result.IsSuccess)
}

func TestHandleNew(t *testing.T) {
	executor := new(successCommandExecutorMock)

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}}
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
