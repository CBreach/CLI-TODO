package tasks

import "time"

type Task struct {
	ID      int
	Name    string
	Status  bool
	dueDate string
	created time.Time
}

func newTask(id int, name string, status bool, dueDate string) *Task {
	return &Task{
		ID:      id,
		Name:    name,
		Status:  status,
		dueDate: dueDate, //may need to add some logic in case due date is not probided... unsure on whether to do it here or elsewhere
		created: time.Now(),
	}
}
