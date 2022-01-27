package execution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: not output in the test results would be better
func TestExecuteWithOutput(t *testing.T) {
	cmd := Command{"echo", []string{"foo"}}
	executor := CommandExecutor{}
	assert.NoError(t, executor.ExecuteWithOutput(cmd))
}

// TODO: not output in the test results would be better
func TestExecuteWithOutputOnError(t *testing.T) {
	cmd := Command{"ls", []string{"/this/does/not/exist"}}
	executor := CommandExecutor{}
	assert.Error(t, executor.ExecuteWithOutput(cmd))
}

// TODO: not output in the test results would be better
func TestExecute(t *testing.T) {
	cmd := Command{"echo", []string{"foo"}}
	executor := CommandExecutor{}
	assert.NoError(t, executor.Execute(cmd))
}

// TODO: not output in the test results would be better
func TestExecuteOnError(t *testing.T) {
	cmd := Command{"ls", []string{"/this/does/not/exist"}}
	executor := CommandExecutor{}
	assert.Error(t, executor.Execute(cmd))
}

func TestExecuteBackground(t *testing.T) {
	cmd := Command{"echo", []string{"foo"}}
	executor := CommandExecutor{}
	_, err := executor.ExecuteBackground(cmd)
	assert.NoError(t, err)
}

func TestExecuteBackgroundOnError(t *testing.T) {
	cmd := Command{"/this/command/does/not/exist", []string{}}
	executor := CommandExecutor{}
	pid, err := executor.ExecuteBackground(cmd)
	assert.Equal(t, -1, pid)
	assert.Error(t, err)
}
