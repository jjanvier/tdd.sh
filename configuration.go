package main

import (
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

func (conf Configuration) GetCommand(alias string) Command {
	cmd := conf.Aliases[alias].Command
	args := strings.Fields(cmd)

	return Command{args[0], args[1:]}
}
