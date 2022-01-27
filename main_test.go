package main

import (
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

func (m *aliasHandlerMock) HandleAlias(conf Configuration, alias string) (ExecutionResult, error) {
	m.Called()

	return ExecutionResult{}, nil
}

func (m *newHandlerMock) HandleNew(message string) (ExecutionResult, error) {
	m.Called()

	return ExecutionResult{}, nil
}

func (m *todoHandlerMock) HandleTodo(message string) error {
	m.Called()

	return nil
}

func (m *todoHandlerMock) HandleDo(stdin io.ReadCloser) (ExecutionResult, error) {
	m.Called()

	return ExecutionResult{}, nil
}

func (m *todoHandlerMock) HandleDone() error {
	m.Called()

	return nil
}

func TestTddItHandlesTestCommand(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"command1 arg1 arg2 --opt1", 120, Git{false}}
	conf.Aliases = aliases

	aliasHandler := new(aliasHandlerMock)
	aliasHandler.On("HandleAlias").Once()

	Tdd("foo", conf, aliasHandler, new(newHandlerMock), new(todoHandlerMock))
}

func TestTddItHandlesNewCommand(t *testing.T) {
	conf := Configuration{}

	newHandler := new(newHandlerMock)
	newHandler.On("HandleNew").Once()

	Tdd("new", conf, new(aliasHandlerMock), newHandler, new(todoHandlerMock))
}

func TestTddItHandlesTodoCommand(t *testing.T) {
	conf := Configuration{}

	handler := new(todoHandlerMock)
	handler.On("HandleTodo").Once()

	Tdd("todo", conf, new(aliasHandlerMock), new(newHandlerMock), handler)
}
