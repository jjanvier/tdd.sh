package handler

import (
	"github.com/jjanvier/tdd/helper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestGetItems(t *testing.T) {
	todoFile := helper.CreateTmpFile(`I should do that
also this should be done
really important to do that`)
	defer helper.RemoveTmpFile(todoFile)

	todoList := TodoList{Path: todoFile.Name()}

	actual, _ := todoList.GetItems()

	assert.Contains(t, actual, "I should do that")
	assert.Contains(t, actual, "also this should be done")
	assert.Contains(t, actual, "really important to do that")
	assert.Equal(t, 3, len(actual))
}

func TestGetItemsWithEmptyLines(t *testing.T) {
	todoFile := helper.CreateTmpFile(`I should do that
   
really important to do that
	

`)
	defer helper.RemoveTmpFile(todoFile)

	todoList := TodoList{Path: todoFile.Name()}

	actual, _ := todoList.GetItems()

	assert.Contains(t, actual, "I should do that")
	assert.Contains(t, actual, "really important to do that")
	assert.Equal(t, 2, len(actual))
}

func TestRemoveItem(t *testing.T) {
	todoFile := helper.CreateTmpFile(`I should do that
also this should be done
really important to do that`)
	defer helper.RemoveTmpFile(todoFile)

	todoList := TodoList{Path: todoFile.Name()}

	res := todoList.Remove(1)

	actualTodo, _ := ioutil.ReadFile(todoFile.Name())
	expectedTodo := `I should do that
really important to do that
`
	assert.Equal(t, expectedTodo, string(actualTodo))
	assert.NoError(t, res)
}
