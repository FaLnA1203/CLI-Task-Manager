package task

import "fmt"

type Task struct {
	Id        int
	taskName  string
	completed bool
}

func (t *Task) CreateTask(_id int, name string) {
	t.Id = _id
	t.taskName = name
	t.completed = false
}

func (t *Task) Complete() {
	t.completed = true
}

func (t *Task) PrintTask() {
	fmt.Print(t.Id, ". ", t.taskName)
	switch t.completed {
	case true:
		fmt.Println(" ✓")
	case false:
		fmt.Println(" ✖")
	}
}
