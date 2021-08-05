package main

import "fmt"

var index int
var tasks map[int]Task = map[int]Task{}

type Task struct {
	Title string
	Done  bool
}

func main() {
	New("task1")
	New("task2")
	New("task3")

	for k, v := range tasks {
		fmt.Println(k, v)
	}
}

func List() map[int]Task {
	return tasks
}

func New(task string) {
	defer func() {
		index++
	}()

	tasks[index] = Task{
		Title: task,
		Done:  false,
	}
}
