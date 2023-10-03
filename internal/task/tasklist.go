package task

import (
	"encoding/json"
	"fmt"
	"os"
)

type TaskList struct {
	Tasks []*Task `json:"tasks"`
}

func NewTaskList() *TaskList {
	return &TaskList{
		Tasks: []*Task{},
	}
}

func (tl *TaskList) AddTask(description string) {
	id := len(tl.Tasks) + 1
	newTask := NewTask(id, description, false)
	tl.Tasks = append(tl.Tasks, newTask)
}

func (tl *TaskList) RemoveTask(id int) {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			return
		}
	}
}

func (tl *TaskList) GetTaskByID(id int) *Task {
	for _, task := range tl.Tasks {
		if task.ID == id {
			return task
		}
	}
	return nil
}

func (tl *TaskList) UpdateTaskDescription(id int, description string) {
	task := tl.GetTaskByID(id)
	if task != nil {
		task.Title = description
	}
}

func (tl *TaskList) CompleteTask(id int) {
	task := tl.GetTaskByID(id)
	if task != nil {
		task.Completed = true
	}
}

func (tl *TaskList) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (tl *TaskList) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, tl)
}

func (tl *TaskList) String() string {
	var result string
	for _, task := range tl.Tasks {
		result += fmt.Sprintf("%s\n", task.String())
	}
	return result
}

func (tl *TaskList) Length() int {
	return len(tl.Tasks)
}

func (tl *TaskList) GetTaskByIndex(index int) *Task {
	if index < 0 || index >= len(tl.Tasks) {
		return nil
	}
	return tl.Tasks[index]
}
