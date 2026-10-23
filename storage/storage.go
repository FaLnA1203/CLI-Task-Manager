package storage

import (
	"TaskManager/task"
	"encoding/json"
	"os"
)

func Save(filename string, t []task.Task) error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func Load(filename string) ([]task.Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tasks []task.Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}
