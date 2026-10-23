package task

import "fmt"

type Task struct {
	Id        int
	TaskName  string
	Completed bool
}

func (t *Task) CreateTask(_id int, name string) {
	t.Id = _id
	t.TaskName = name
	t.Completed = false
}

func (t *Task) Complete() {
	t.Completed = true
}

func (t *Task) PrintTask() {
	fmt.Print(t.Id, ". ", t.TaskName)
	switch t.Completed {
	case true:
		fmt.Println(" ✓")
	case false:
		fmt.Println(" ✖")
	}
}
