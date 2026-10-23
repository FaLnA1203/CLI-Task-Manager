package storage

import (
	"TaskManager/task"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	currentTasks, err := Load("rand")
	if err != nil {
		t.Fatal(err)
	}
	if len(currentTasks) != 0 {
		t.Error("expected 0 tasks, got ", len(currentTasks))
	}
}

func TestSave(t *testing.T) {
	currentTasks1 := []task.Task{
		{Id: 1, TaskName: "Learn Go", Completed: false},
		{Id: 2, TaskName: "Learn Python", Completed: true},
		{Id: 3, TaskName: "Learn Java", Completed: true},
	}
	Save("testing", currentTasks1)
	var currentTasks2 []task.Task
	currentTasks2, err := Load("testing")
	if err != nil {
		t.Fatal(err)
	}
	for i, _ := range currentTasks1 {
		if currentTasks2[i].TaskName != currentTasks1[i].TaskName {
			t.Error("expected ", currentTasks1[i].TaskName, "got ", currentTasks2[i].TaskName)
		}
	}
	os.Remove("testing.json")
}
