package model

import "path"

type Tasks struct {
	Tasks map[string]*Task
}

var allTasks = &Tasks{Tasks: map[string]*Task{}}

func RegisterTask(task *Task) {
	allTasks.Tasks[path.Base(task.TargetFile)] = task
}

func GetAllTasks() *Tasks {
	return allTasks
}