package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type aliasHandlerMock struct {
	mock.Mock
}

func (m *aliasHandlerMock) HandleTestCommand(testCmd Command) ExecutionResult {
	m.Called()

	return ExecutionResult{}
}

func (m *aliasHandlerMock) HandleNew(message string) ExecutionResult {
	m.Called()

	return ExecutionResult{}
}

func TestTddItHandlesTestCommand(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"command1 arg1 arg2 --opt1", 120}
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

func TestHello(t *testing.T) {
	assert.Equal(t, "Hello foo", Hello("foo"))
}
