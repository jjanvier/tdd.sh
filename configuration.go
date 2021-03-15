package main

import (
	"io/ioutil"
	"log"

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
