package services

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadTasksFromFile() ([]Task, error) {
	filepath := "tasks/tasks.json"
	_, err := os.Stat(filepath)
	// create new file if it doesn't exist and return empty Task list
	if os.IsNotExist(err) {
		fmt.Println("File doesn't exist, creating file")
		err := os.WriteFile(filepath, []byte("[]"), 0644)

		if err != nil {
			fmt.Println("Error creating/writing to file:", err)
			return nil, err
		}

		return []Task{}, nil
	}

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	defer file.Close()

	tasks := []Task{}
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return nil, err
	}

	return tasks, nil
}

func WriteTasksToFile(tasks []Task) error {
	filepath := "tasks/tasks.json"
	file, err := os.Create(filepath) // if file exists, truncate/clear file and write new list of tasks
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding file:", err)
		return err
	}

	return nil
}
