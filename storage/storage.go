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
	var tasks []task.Task
	if os.IsNotExist(err){return tasks, nil}
	if err != nil {
		return nil, err
	}


	err = json.Unmarshal(data, &tasks)
	return tasks, err
}
