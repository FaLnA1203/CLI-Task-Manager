package task

import (
	"fmt"
	"testing"
)

func DeleteTask(t *[]Task, id int) {
	if id < 1 || id > len(*t) {
		fmt.Println("Задача с таким ID не найдена")
		return
	}
	*t = append((*t)[:id-1], (*t)[id:]...)
	for i, _ := range *t {
		(*t)[i].Id = i + 1
	}
}

func TestComplete(t *testing.T) {
	currentTask := Task{
		Id:        1,
		TaskName:  "Изучить Go",
		Completed: false,
	}

	currentTask.Complete()

	if !currentTask.Completed {
		t.Error("после Complete задача должна быть завершена")
	}
}

func TestDelete(t *testing.T) {
	currentTasks1 := []Task{
		{Id: 1, TaskName: "Learn Go", Completed: false},
		{Id: 2, TaskName: "Learn Python", Completed: true},
		{Id: 3, TaskName: "Learn Java", Completed: true},
	}
	DeleteTask(&currentTasks1, 2)

	if len(currentTasks1) != 2 {
		t.Error("количество задач не уменьшилось")
	}
	for i, currentTask := range currentTasks1 {
		if currentTask.TaskName == "Learn Python" {
			t.Error("нужная задача не удалилась")
		}
		if currentTask.Id != i+1 {
			t.Error("ID задач не пересчитались")
		}
	}
	currentTasks2 := []Task{
		{Id: 1, TaskName: "Learn Go", Completed: false},
	}
	DeleteTask(&currentTasks2, 1)
	if len(currentTasks2) != 0 {
		t.Error("количество задач не уменьшилось")
	}
	DeleteTask(&currentTasks2, 2)
	t.Log("Программа продолжает работать")
	
}
