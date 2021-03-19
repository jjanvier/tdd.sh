package main

import (
	"log"
	"os"
)

const configurationFile = ".tdd.yml"

func main() {
	alias := os.Args[1]
	conf := Load(configurationFile)
	cmd, err := conf.GetCommand(alias)
	if err != nil {
		log.Fatal(err)
	}
	cmd.ExecuteWithOutput()
}

func Hello(name string) string {
	return "Hello " + name
}
