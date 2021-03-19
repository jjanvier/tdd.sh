package main

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

type Alias struct {
	Command string
	Timer   int
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
