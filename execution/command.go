package execution

import "strings"

type Command struct {
	Name      string
	Arguments []string
}

type CommandFactory struct{}

func (factory CommandFactory) CreateGitAdd() Command {
	return Command{"git", []string{"add", "."}}
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

func (cmd Command) String() string {
	return strings.Join(append([]string{cmd.Name}, cmd.Arguments...), " ")
}
