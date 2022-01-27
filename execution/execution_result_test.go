package execution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateExecutionResultSuccess(t *testing.T) {
	factory := ExecutionResultFactory{}
	result := factory.Success([]Command{
		{"toto", []string{"titi", "--tata"}},
		{"foo", []string{"bar", "baz"}},
	})

	assert.Equal(t, "toto titi --tata && foo bar baz", result.Command)
	assert.Equal(t, true, result.IsSuccess)
}

func TestCreateExecutionResultFailure(t *testing.T) {
	factory := ExecutionResultFactory{}
	result := factory.Failure([]Command{
		{"toto", []string{"titi", "--tata"}},
		{"foo", []string{"bar", "baz"}},
	})

	assert.Equal(t, "toto titi --tata && foo bar baz", result.Command)
	assert.Equal(t, false, result.IsSuccess)
}
