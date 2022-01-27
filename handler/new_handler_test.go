package handler

import (
	"github.com/jjanvier/tdd/execution"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandleNew(t *testing.T) {
	executor := new(execution.SuccessCommandExecutorMock)
	executionResultFactory := execution.ExecutionResultFactory{}
	commandFactory := execution.CommandFactory{}
	newHandler := NewHandler{executor, commandFactory, executionResultFactory}

	result, _ := newHandler.HandleNew("here is my commit message")

	assert.Equal(t, "git commit --allow-empty -m here is my commit message", result.Command)
	assert.Equal(t, true, result.IsSuccess)
}
