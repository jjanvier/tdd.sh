package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type aliasHandlerMock struct {
	mock.Mock
}

func (m *aliasHandlerMock) HandleTestCommand(conf Configuration, alias string) (ExecutionResult, error) {
	m.Called()

	return ExecutionResult{}, nil
}

func (m *aliasHandlerMock) HandleNew(message string) (ExecutionResult, error) {
	m.Called()

	return ExecutionResult{}, nil
}

func (m *aliasHandlerMock) HandleTodo(message string) (ExecutionResult, error) {
	m.Called()

	return ExecutionResult{}, nil
}

func TestTddItHandlesTestCommand(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"command1 arg1 arg2 --opt1", 120, Git{false}}
	conf.Aliases = aliases

	handler := new(aliasHandlerMock)
	handler.On("HandleTestCommand").Once()

	Tdd("foo", conf, handler)
}

func TestTddItHandlesNewCommand(t *testing.T) {
	conf := Configuration{}

	handler := new(aliasHandlerMock)
	handler.On("HandleNew").Once()

	Tdd("new", conf, handler)
}

func TestTddItHandlesTodoCommand(t *testing.T) {
	conf := Configuration{}

	handler := new(aliasHandlerMock)
	handler.On("HandleTodo").Once()

	Tdd("todo", conf, handler)
}
