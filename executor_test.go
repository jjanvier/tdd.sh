package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	cmd := Command{"ls", []string{"-al", "-h"}}
	assert.Equal(t, "ls -al -h", cmd.String())
}

// TODO: not output in the test results would be better
func TestExecuteWithOutput(t *testing.T) {
	cmd := Command{"echo", []string{"foo"}}
	assert.NoError(t, cmd.ExecuteWithOutput())
}

// TODO: not output in the test results would be better
func TestExecuteWithOutputOnError(t *testing.T) {
	cmd := Command{"ls", []string{"/this/does/not/exist"}}
	assert.Error(t, cmd.ExecuteWithOutput())
}
