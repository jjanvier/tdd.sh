package main

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

// The configuration file should be like:
//
// aliases:
//   ut:
//     command: go test -v
//     timer: 120
//     git:
//     amend: true

type Git struct {
	Amend bool
}

type Alias struct {
	Command string
	Timer   int
	Git Git
}

type Configuration struct {
	Aliases map[string]Alias
}

func Load(path string) Configuration {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	conf := Configuration{}
	yaml.Unmarshal(content, &conf)

	return conf
}

func (conf Configuration) GetCommand(alias string) (Command, error) {
	if _, ok := conf.Aliases[alias]; !ok {
		return Command{}, errors.New("The alias '" + alias + "' does not exist.")
	}

	cmd := conf.Aliases[alias].Command
	args := strings.Fields(cmd)

	return Command{args[0], args[1:]}, nil
}

func (conf Configuration) ShouldAmendCommits(alias string) (bool, error) {
	if _, ok := conf.Aliases[alias]; !ok {
		return false, errors.New("The alias '" + alias + "' does not exist.")
	}

	return conf.Aliases[alias].Git.Amend, nil
}

func (conf Configuration) GetTimer(alias string) (int, error) {
	if _, ok := conf.Aliases[alias]; !ok {
		return 0, errors.New("The alias '" + alias + "' does not exist.")
	}

	return conf.Aliases[alias].Timer, nil
}
