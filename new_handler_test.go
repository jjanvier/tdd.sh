package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandleNew(t *testing.T) {
	executor := new(successCommandExecutorMock)
	executionResultFactory := ExecutionResultFactory{}
	commandFactory := CommandFactory{}
	newHandler := NewHandler{executor, commandFactory, executionResultFactory}

	result, _ := newHandler.HandleNew("here is my commit message")

	assert.Equal(t, "git commit --allow-empty -m here is my commit message", result.Command)
	assert.Equal(t, true, result.IsSuccess)
}
