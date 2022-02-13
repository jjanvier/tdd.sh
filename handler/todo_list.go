package handler

import (
	"io/ioutil"
	"os"
	"strings"
)

type TodoList struct {
	Path string
}

func (list TodoList) Add(todos []string) error {
	todoFile, err := os.OpenFile(list.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer todoFile.Close()
	if err != nil {
		return err
	}

	for _, todo := range todos {
		_, err2 := todoFile.WriteString(todo + "\n")
		if err2 != nil {
			return err2
		}
	}

	return nil
}

func (list TodoList) GetItems() ([]string, error) {
	todoContent, err := ioutil.ReadFile(list.Path)

	if err != nil {
		return []string{}, err
	}

	items := strings.Split(string(todoContent), "\n")

	return removeEmptyStringsFromSlice(items), nil
}

func (list TodoList) Clear() error {
	todoFile, err := os.OpenFile(list.Path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	defer todoFile.Close()

	return err
}

func (list TodoList) Remove(index int) error {
	items, err := list.GetItems()
	if err != nil {
		return err
	}

	newItems := removeItemFromSlice(items, index)

	err = list.Clear()
	if err != nil {
		return err
	}

	err = list.Add(newItems)
	if err != nil {
		return err
	}

	return err
}

func removeItemFromSlice(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func removeEmptyStringsFromSlice(slice []string) []string {
	var result []string
	for _, str := range slice {
		if strings.TrimSpace(str) != "" {
			result = append(result, str)
		}
	}

	return result
}
