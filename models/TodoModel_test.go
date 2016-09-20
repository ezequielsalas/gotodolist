package models

import (
	"fmt"
)

/**
func TestAddTaskToList(t *testing.T) {

}
**/

func ExampleAddTaskToList() {

	task1 := Task{"My first task", "My simple description"}

	list := TodoList{}
	list.Name = "My list"

	list.AddTask(task1)

	fmt.Println(list)
	// Output:
	// {My list [{My first task My simple description}]}
}
