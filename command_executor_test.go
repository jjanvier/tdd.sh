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

func TestCreateGitAddCommand(t *testing.T) {
	factory := CommandFactory{}
	cmd := factory.CreateGitAdd()
	assert.Equal(t, "git add .", cmd.String())
}

func TestCreateGitCommit(t *testing.T) {
	factory := CommandFactory{}
	cmd := factory.CreateGitCommit()
	assert.Equal(t, "git commit --reuse-message=HEAD", cmd.String())
}

func TestCreateGitCommitEmpty(t *testing.T) {
	factory := CommandFactory{}
	cmd := factory.CreateGitCommitEmpty("my beautiful commit message")
	assert.Equal(t, "git commit --allow-empty -m my beautiful commit message", cmd.String())
}

func TestCreateGitCommitAmend(t *testing.T) {
	factory := CommandFactory{}
	cmd := factory.CreateGitCommitAmend()
	assert.Equal(t, "git commit --amend --no-edit", cmd.String())
}
