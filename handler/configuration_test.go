package handler

import (
	"github.com/jjanvier/tdd/execution"
	"github.com/jjanvier/tdd/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	confContent := `aliases:
  foo:
    command: command1
    timer: 120
  bar:
    command: command2
    timer: 500
    git:
      amend: true
`

	confFile := helper.CreateTmpFile(confContent)
	// TODO: use https://golang.org/pkg/testing/#B.Cleanup instead?
	defer helper.RemoveTmpFile(confFile)

	actual := Load(confFile.Name())

	expectedAliases := make(map[string]Alias)
	expectedAliases["foo"] = Alias{"command1", 120, Git{false}}
	expectedAliases["bar"] = Alias{"command2", 500, Git{true}}
	expected := Configuration{}
	expected.Aliases = expectedAliases

	assert.Equal(t, expected, actual)
}

func TestGetCommand(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"command1 arg1 arg2 --opt1", 120, Git{false}}
	conf.Aliases = aliases

	expected := execution.Command{Name: "command1", Arguments: []string{"arg1", "arg2", "--opt1"}}
	actual, _ := conf.getCommand("foo")

	assert.Equal(t, expected, actual)
}

func TestGetCommandAliasNotFound(t *testing.T) {
	conf := Configuration{}
	_, actualError := conf.getCommand("foo")

	assert.Error(t, actualError)
}

func TestShouldAmendCommits(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"command1 arg1 arg2 --opt1", 120, Git{false}}
	aliases["bar"] = Alias{"command2", 60, Git{true}}
	conf.Aliases = aliases

	notAmended, _ := conf.shouldAmendCommits("foo")
	assert.False(t, notAmended)

	amended, _ := conf.shouldAmendCommits("bar")
	assert.True(t, amended)
}

func TestShouldAmendCommitsAliasNotFound(t *testing.T) {
	conf := Configuration{}

	_, actualError := conf.shouldAmendCommits("foo")
	assert.Error(t, actualError)
}

func TestGetTimer(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"command1 arg1 arg2 --opt1", 120, Git{false}}
	conf.Aliases = aliases

	actualTimer, _ := conf.getTimer("foo")
	assert.Equal(t, 120, actualTimer)
}

func TestGetTimerAliasNotFound(t *testing.T) {
	conf := Configuration{}

	_, actualError := conf.getTimer("foo")
	assert.Error(t, actualError)
}
