package handler

import (
	"errors"
	"fmt"
	"github.com/jjanvier/tdd/execution"
	"github.com/txgruppi/parseargs-go"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// The configuration file should be like:
const defaultConfigurationFile = `aliases:
    ut: # I use "ut" for Unit Tests. Personally, I define a "ut" alias for all my projects
        command: echo "change me"
        git:
            amend: true # commits will be amended when tests are green
        timer: 60 # you'll receive a small notification if your steps are still red after 60 seconds
    another-alias:
        command: echo "change me too"
        # if no "git" key is configured, commits won't be amended: the previous message will be reused
        # if no "timer" key is defined, no notification will pop
`

type AliasNotFoundError struct {
	Alias string
}

func (e *AliasNotFoundError) Error() string {
	return fmt.Sprintf("the alias \"%s\" does not exist", e.Alias)
}

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
	configFileExists := ConfigurationFileExists(path)
	if !configFileExists {
		return Configuration{}, errors.New("no configuration file found")
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Configuration{}, errors.New("impossible to read the configuration file")
	}

	conf := Configuration{}

	err2 := yaml.Unmarshal(content, &conf)
	if err2 != nil {
		return Configuration{}, errors.New("invalid configuration, please fix your configuration file")
	}

	if len(conf.Aliases) == 0 {
		return Configuration{}, errors.New("no alias found, please fix your configuration file")
	}

	return conf, nil
}

func (conf Configuration) getCommand(alias string) (execution.Command, error) {
	if _, ok := conf.Aliases[alias]; !ok {
		return execution.Command{}, &AliasNotFoundError{Alias: alias}
	}

	cmd := conf.Aliases[alias].Command
	args, err := parseargs.Parse(cmd)
	if err != nil {
		return execution.Command{}, fmt.Errorf("unable to parse the command of the alias \"%s\", please fix your configuration file", alias)
	}

	return execution.Command{Name: args[0], Arguments: args[1:]}, nil
}

func (conf Configuration) shouldAmendCommits(alias string) (bool, error) {
	if _, ok := conf.Aliases[alias]; !ok {
		return false, &AliasNotFoundError{Alias: alias}
	}

	return conf.Aliases[alias].Git.Amend, nil
}

func (conf Configuration) getTimer(alias string) (int, error) {
	if _, ok := conf.Aliases[alias]; !ok {
		return 0, &AliasNotFoundError{Alias: alias}
	}

	return conf.Aliases[alias].Timer, nil
}

func ConfigurationFileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
