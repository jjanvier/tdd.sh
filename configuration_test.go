package main

import (
	"io/ioutil"
	"log"
	"os"
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

	confFile := createTmpFile(confContent)
	// TODO: use https://golang.org/pkg/testing/#B.Cleanup instead?
	defer removeTmpFile(confFile)

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

	expected := Command{"command1", []string{"arg1", "arg2", "--opt1"}}
	actual, _ := conf.GetCommand("foo")

	assert.Equal(t, expected, actual)
}

func TestGetCommandAliasNotFound(t *testing.T) {
	conf := Configuration{}
	_, actualError := conf.GetCommand("foo")

	assert.Error(t, actualError)
}

func TestShouldAmendCommits(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"command1 arg1 arg2 --opt1", 120, Git{false}}
	aliases["bar"] = Alias{"command2", 60, Git{true}}
	conf.Aliases = aliases

	notAmended, _ := conf.ShouldAmendCommits("foo")
	assert.False(t, notAmended)

	amended, _ := conf.ShouldAmendCommits("bar")
	assert.True(t, amended)
}

func TestShouldAmendCommitsAliasNotFound(t *testing.T) {
	conf := Configuration{}

	_, actualError := conf.ShouldAmendCommits("foo")
	assert.Error(t, actualError)
}

func TestGetTimer(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"command1 arg1 arg2 --opt1", 120, Git{false}}
	conf.Aliases = aliases

	actualTimer, _ := conf.GetTimer("foo")
	assert.Equal(t, 120, actualTimer)
}


func TestGetTimerAliasNotFound(t *testing.T) {
	conf := Configuration{}

	_, actualError := conf.GetTimer("foo")
	assert.Error(t, actualError)
}

func createTmpFile(content string) *os.File {
	data := []byte(content)
	tmpfile, err := ioutil.TempFile("/tmp", "tdd.sh-")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write(data); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	return tmpfile
}

func removeTmpFile(tmpfile *os.File) {
	os.Remove(tmpfile.Name())
}
