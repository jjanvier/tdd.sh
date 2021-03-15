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
`

	confFile := createTmpFile(confContent)
	defer removeTmpFile(confFile)

	actual := Load(confFile.Name())

	expectedAliases := make(map[string]Alias)
	expectedAliases["foo"] = Alias{"command1", 120}
	expectedAliases["bar"] = Alias{"command2", 500}
	expected := Configuration{}
	expected.Aliases = expectedAliases

	assert.Equal(t, expected, actual)
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
