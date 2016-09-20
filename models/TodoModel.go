package models

/*
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
*/
type Task struct {
	Title       string
	Description string
}

type TodoList struct {
	Name string
	Task []Task
}

func (list *TodoList) AddTask(task Task) []Task {
	list.Task = append(list.Task, task)
	return list.Task
}

// func (list *TodoList) SaveTask(task Task) []Task {
// 	list.Task = append(list.Task, task)
// 	return list.Task
// }
