package execution

import (
	"strings"
)

type ExecutionResult struct {
	IsSuccess bool
	Command   string
}

type ExecutionResultFactory struct{}

func (factory ExecutionResultFactory) Success(cmds []Command) ExecutionResult {
	return ExecutionResult{true, factory.joinCommands(cmds)}
}

func (factory ExecutionResultFactory) Failure(cmds []Command) ExecutionResult {
	return ExecutionResult{false, factory.joinCommands(cmds)}
}

func (factory ExecutionResultFactory) joinCommands(cmds []Command) string {
	var cmdsString []string
	for _, cmd := range cmds {
		cmdsString = append(cmdsString, cmd.String())
	}

	return strings.Join(cmdsString, " && ")
}
