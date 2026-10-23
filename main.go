package main

import (
	"TaskManager/storage"
	"TaskManager/task"
	"bufio"
	"fmt"
	"os"
)

const dataFile = "tasks.json"

func DeleteTask(t *[]task.Task, id int) {
	*t = append((*t)[:id-1], (*t)[id:]...)
	for i, _ := range *t {
		(*t)[i].Id -= 1
	}
}

func main() {
	t, _ := storage.Load(dataFile)
	var opt int
	var id int
	var task1 task.Task
	scanner := bufio.NewScanner(os.Stdin)
FOR:
	for {
		storage.Save(dataFile, t)
		fmt.Print("\033[2J")
		fmt.Println("Task Manager")
		fmt.Println("")
		fmt.Println("1. Add task\n2. List tasks\n3. Complete task\n4. Delete task\n5.Exit")
		fmt.Println("")
		fmt.Print("Enter number: ")
		fmt.Scan(&opt)
		switch opt {
		case 1:
			fmt.Print("\033[2J")
			fmt.Print("Task title: ")
			scanner.Scan()
			name := scanner.Text()
			task1.CreateTask(len(t)+1, name)
			t = append(t, task1)
		case 2:
			fmt.Print("\033[2J")
			for _, task := range t {
				task.PrintTask()
			}
			fmt.Println("Нажмите Enter, чтобы продолжить...")
			fmt.Scanln()
			fmt.Println("Клавиша нажата!")
		case 3:
			fmt.Print("\033[2J")
			fmt.Print("Task id: ")
			fmt.Scan(&id)
			t[id-1].Complete()
		case 4:
			fmt.Print("\033[2J")
			fmt.Print("Task id: ")
			fmt.Scan(&id)
			DeleteTask(&t, id)
		case 5:
			break FOR
		default:
			fmt.Println("There is no such command.")
		}
	}
}
