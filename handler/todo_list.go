package handler

import (
	"io/ioutil"
	"os"
	"strings"
)

type TodoList struct {
	Path string
}

func (list TodoList) Add(todo string) error {
	todoFile, err := os.OpenFile(list.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer todoFile.Close()
	if err != nil {
		return err
	}

	_, err2 := todoFile.WriteString(todo + "\n")
	if err2 != nil {
		return err2
	}

	return nil
}

func (list TodoList) GetItems() ([]string, error) {
	todoContent, err := ioutil.ReadFile(list.Path)

	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(todoContent), "\n"), nil
}

func (list TodoList) Clear() error {
	todoFile, err := os.OpenFile(list.Path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	defer todoFile.Close()

	return err
}
