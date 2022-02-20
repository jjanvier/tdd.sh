package handler

import (
	"github.com/jjanvier/tdd/execution"
	"github.com/jjanvier/tdd/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfiguration(t *testing.T) {
	confContent := `aliases:
  foo:
    command: command1
    timer: 120
  bar:
    command: echo "a more complex command"
    timer: 500
    git:
      amend: true
      add: "*.php doc/*"
  baz:
    command: simplestalias
`

	confFile := helper.CreateTmpFile(confContent)
	// TODO: use https://golang.org/pkg/testing/#B.Cleanup instead?
	defer helper.RemoveTmpFile(confFile)

	actual, _ := LoadConfiguration(confFile.Name())

	expectedAliases := make(map[string]Alias)
	expectedAliases["foo"] = Alias{"command1", 120, Git{false, "."}}
	expectedAliases["bar"] = Alias{"echo \"a more complex command\"", 500, Git{true, "*.php doc/*"}}
	expectedAliases["baz"] = Alias{"simplestalias", 0, Git{false, "."}}
	expected := Configuration{}
	expected.Aliases = expectedAliases

	assert.Equal(t, expected, actual)
}

func TestLoadConfigurationErrorRoot(t *testing.T) {
	confContent := `aliasesWrong:
  foo:
    command: command1
    timer: 120
`

	confFile := helper.CreateTmpFile(confContent)
	defer helper.RemoveTmpFile(confFile)

	_, err := LoadConfiguration(confFile.Name())

	assert.Error(t, err)
}

func TestLoadConfigurationErrorNoCommand(t *testing.T) {
	confContent := `aliasesWrong:
  foo:
    timer: 120
`

	confFile := helper.CreateTmpFile(confContent)
	defer helper.RemoveTmpFile(confFile)

	_, err := LoadConfiguration(confFile.Name())

	assert.Error(t, err)
}

func TestLoadConfigurationFileNotExists(t *testing.T) {
	_, err := LoadConfiguration("/a/file/that/does/not/exist")

	assert.Error(t, err)
	assert.Equal(t, "no configuration file found", err.Error())
}

func TestGetCommand(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"command1 arg1 arg2 --opt1", 120, Git{false, "."}}
	conf.Aliases = aliases

	expected := execution.Command{Name: "command1", Arguments: []string{"arg1", "arg2", "--opt1"}}
	actual, _ := conf.getAliasCommand("foo")

	assert.Equal(t, expected, actual)
}

func TestGetCommandWithStringArgs(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"command1 \"double quotes arg\" 'simple quote arg' third-arg --option -o", 120, Git{false, "."}}
	conf.Aliases = aliases

	expected := execution.Command{Name: "command1", Arguments: []string{"double quotes arg", "simple quote arg", "third-arg", "--option", "-o"}}
	actual, _ := conf.getAliasCommand("foo")

	assert.Equal(t, expected, actual)
}

func TestGetCommandAliasNotFound(t *testing.T) {
	conf := Configuration{}
	_, actualError := conf.getAliasCommand("foo")

	assert.Error(t, actualError)
}

func TestShouldAmendCommits(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"command1 arg1 arg2 --opt1", 120, Git{false, "."}}
	aliases["bar"] = Alias{"command2", 60, Git{true, "."}}
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
	aliases["foo"] = Alias{"command1 arg1 arg2 --opt1", 120, Git{false, "."}}
	conf.Aliases = aliases

	actualTimer, _ := conf.getTimer("foo")
	assert.Equal(t, 120, actualTimer)
}

func TestGetTimerAliasNotFound(t *testing.T) {
	conf := Configuration{}

	_, actualError := conf.getTimer("foo")
	assert.Error(t, actualError)
}

func TestFileExists(t *testing.T) {
	confFile := helper.CreateTmpFile("")
	defer helper.RemoveTmpFile(confFile)

	assert.True(t, ConfigurationFileExists(confFile.Name()))
	assert.False(t, ConfigurationFileExists("/this/one/does/not/exist"))
}

func TestGetGitAddCommand(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"command1", 120, Git{false, "*.php doc/*"}}
	conf.Aliases = aliases

	commandFactory := execution.CommandFactory{}
	expected := commandFactory.CreateGitAdd("*.php doc/*")
	actual, _ := conf.getGitAddCommand("foo")

	assert.Equal(t, expected, actual)
}

func TestGetGitAddCommandAliasNotFound(t *testing.T) {
	conf := Configuration{}
	_, actualError := conf.getGitAddCommand("foo")

	assert.Error(t, actualError)
}
