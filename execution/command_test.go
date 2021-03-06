package execution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	cmd := Command{"ls", []string{"-al", "-h"}}
	assert.Equal(t, "ls -al -h", cmd.String())
}

func TestCreateGitAdd(t *testing.T) {
	factory := CommandFactory{}
	cmd := factory.CreateGitAdd("*.php doc/*")
	assert.Equal(t, "git add -- *.php doc/*", cmd.String())
}

func TestCreateGitAddSimple(t *testing.T) {
	factory := CommandFactory{}
	cmd := factory.CreateGitAdd(".")
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

func TestCreateNotifyCommand(t *testing.T) {
	factory := CommandFactory{}
	cmd := factory.CreateNotify(25, "my message")
	assert.Contains(t, cmd.String(), " notify 25 my message")
}
