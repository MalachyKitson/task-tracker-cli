package internal

import (
	"encoding/json"
	"errors"
	"os"
)

const fileName = "config/tasks.json"

func LoadTasks() ([]Task, error) {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return []Task{}, nil
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}
