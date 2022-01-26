package main

import (
	"io/ioutil"
	"os"
	"strings"
)

type TodoList struct {
	path string
}

func (list TodoList) Add(todo string) error {
	todoFile, err := os.OpenFile(list.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
	todoContent, err := ioutil.ReadFile(list.path)

	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(todoContent), "\n"), nil
}

func (list TodoList) Clear() error {
	todoFile, err := os.OpenFile(list.path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	defer todoFile.Close()

	return err
}
