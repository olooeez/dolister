package task

import "fmt"

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func NewTask(id int, title string, completed bool) *Task {
	return &Task{
		ID:        id,
		Title:     title,
		Completed: completed,
	}
}

func (t *Task) String() string {
	status := " "
	if t.Completed {
		status = "x"
	}

	return fmt.Sprintf("[%s] %s", status, t.Title)
}
