package main

import (
	"github.com/jjanvier/tdd/execution"
	"github.com/jjanvier/tdd/handler"
	"io"
	"testing"

	"github.com/stretchr/testify/mock"
)

type aliasHandlerMock struct {
	mock.Mock
}
type newHandlerMock struct {
	mock.Mock
}
type todoHandlerMock struct {
	mock.Mock
}

func (m *aliasHandlerMock) HandleAlias(conf handler.Configuration, alias string) (execution.ExecutionResult, error) {
	m.Called()

	return execution.ExecutionResult{}, nil
}

func (m *newHandlerMock) HandleNew(message string) (execution.ExecutionResult, error) {
	m.Called()

	return execution.ExecutionResult{}, nil
}

func (m *todoHandlerMock) HandleTodo(message string) error {
	m.Called()

	return nil
}

func (m *todoHandlerMock) HandleDo(stdin io.ReadCloser) (execution.ExecutionResult, error) {
	m.Called()

	return execution.ExecutionResult{}, nil
}

func (m *todoHandlerMock) HandleDone() error {
	m.Called()

	return nil
}

func TestTddItHandlesTestCommand(t *testing.T) {
	conf := handler.Configuration{}
	aliases := make(map[string]handler.Alias)
	aliases["foo"] = handler.Alias{Command: "command1 arg1 arg2 --opt1", Timer: 120, Git: handler.Git{false}}
	conf.Aliases = aliases

	aliasHandler := new(aliasHandlerMock)
	aliasHandler.On("HandleAlias").Once()

	Tdd("foo", conf, aliasHandler, new(newHandlerMock), new(todoHandlerMock))
}

func TestTddItHandlesNewCommand(t *testing.T) {
	conf := handler.Configuration{}

	newHandler := new(newHandlerMock)
	newHandler.On("HandleNew").Once()

	Tdd("new", conf, new(aliasHandlerMock), newHandler, new(todoHandlerMock))
}

func TestTddItHandlesTodoCommand(t *testing.T) {
	conf := handler.Configuration{}

	handler := new(todoHandlerMock)
	handler.On("HandleTodo").Once()

	Tdd("todo", conf, new(aliasHandlerMock), new(newHandlerMock), handler)
}
