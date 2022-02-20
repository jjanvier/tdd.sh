package execution

import (
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Name      string
	Arguments []string
}

type CommandFactory struct{}

func (factory CommandFactory) CreateGitAdd(pathSpec string) Command {
	if pathSpec == "." {
		return Command{"git", []string{"add", "."}}
	}

	// quite tricky here...
	// if we want something like "git add -- *.php doc/*", we need "*.php" and "doc/*"
	// to be treated as 2 different arguments of the git command
	// otherwise, we'll get something like "pathspec did not match any files"
	pathSpecs := strings.Fields(pathSpec)

	return Command{"git", append([]string{"add", "--"}, pathSpecs...)}
}

func (factory CommandFactory) CreateGitCommit() Command {
	return Command{"git", []string{"commit", "--reuse-message=HEAD"}}
}

func (factory CommandFactory) CreateGitCommitEmpty(message string) Command {
	return Command{"git", []string{"commit", "--allow-empty", "-m", message}}
}

func (factory CommandFactory) CreateGitCommitAmend() Command {
	return Command{"git", []string{"commit", "--amend", "--no-edit"}}
}

func (factory CommandFactory) CreateNotify(delay int, message string) Command {
	return Command{Name: os.Args[0], Arguments: []string{"notify", strconv.Itoa(delay), message}}
}

func (cmd Command) String() string {
	return strings.Join(append([]string{cmd.Name}, cmd.Arguments...), " ")
}
