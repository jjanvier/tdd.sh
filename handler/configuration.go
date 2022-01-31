package handler

import (
	"errors"
	"github.com/jjanvier/tdd/execution"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// The configuration file should be like:
const defaultConfigurationFile = `aliases:
    ut: # I use "ut" for Unit Tests. Personally, I define a "ut" alias for all my projects
        command: echo changeme
        git:
            amend: true # commits will be amended when tests are green
        timer: 60 # you'll receive a small notification if your steps are still red after 60 seconds
    another-alias:
        command: echo changemetoo
        # if no "git" key is configured, commits won't be amended: the previous message will be reused
        # if no "timer" key is defined, no notification will pop
`

type Git struct {
	Amend bool
}

type Alias struct {
	Command string
	Timer   int
	Git     Git
}

type Configuration struct {
	Aliases map[string]Alias
}

func LoadConfiguration(path string) (Configuration, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Configuration{}, err
	}

	conf := Configuration{}

	err2 := yaml.Unmarshal(content, &conf)
	if err2 != nil {
		return Configuration{}, err2
	}

	if len(conf.Aliases) == 0 {
		return Configuration{}, errors.New("No alias in the configuration file")
	}

	return conf, nil
}

func (conf Configuration) getCommand(alias string) (execution.Command, error) {
	if _, ok := conf.Aliases[alias]; !ok {
		return execution.Command{}, errors.New("The alias '" + alias + "' does not exist.")
	}

	cmd := conf.Aliases[alias].Command
	args := strings.Fields(cmd)

	return execution.Command{Name: args[0], Arguments: args[1:]}, nil
}

func (conf Configuration) shouldAmendCommits(alias string) (bool, error) {
	if _, ok := conf.Aliases[alias]; !ok {
		return false, errors.New("The alias '" + alias + "' does not exist.")
	}

	return conf.Aliases[alias].Git.Amend, nil
}

func (conf Configuration) getTimer(alias string) (int, error) {
	if _, ok := conf.Aliases[alias]; !ok {
		return 0, errors.New("The alias '" + alias + "' does not exist.")
	}

	return conf.Aliases[alias].Timer, nil
}

func ConfigurationFileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
