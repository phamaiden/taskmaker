package services

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const (
	TASK_STATUS_TODO        = "todo"
	TASK_STATUS_IN_PROGRESS = "in-progress"
	TASK_STATUS_DONE        = "done"
)

func NewTask(id int, desc string) *Task {
	return &Task{
		ID:          id,
		Description: desc,
		Status:      TASK_STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ListTasks(filter string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	switch filter {
	case "all":
		for _, task := range tasks {
			fmt.Printf("Task %v: %v (%v)\n", task.ID, task.Description, task.Status)
		}
	case "todo":
		for _, task := range tasks {
			if task.Status == TASK_STATUS_TODO {
				fmt.Printf("Task %v: %v\n", task.ID, task.Description)
			}
		}
	case "in-progress":
		for _, task := range tasks {
			if task.Status == TASK_STATUS_IN_PROGRESS {
				fmt.Printf("Task %v: %v\n", task.ID, task.Description)
			}
		}
	case "done":
		for _, task := range tasks {
			if task.Status == TASK_STATUS_DONE {
				fmt.Printf("Task %v: %v\n", task.ID, task.Description)
			}
		}
	}

	return nil
}

func AddTask(desc string) (int, error) {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return 0, err
	}

	var taskId int
	if len(tasks) > 0 {
		taskId = tasks[len(tasks)-1].ID + 1
	} else {
		taskId = 1
	}

	task := NewTask(taskId, desc)
	tasks = append(tasks, *task)

	return taskId, WriteTasksToFile(tasks)
}

func UpdateTask(taskId int, desc string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var updatedTasks []Task
	var foundTask bool = false
	for _, task := range tasks {
		if task.ID == taskId {
			foundTask = true
			task.Description = desc
			task.UpdatedAt = time.Now()
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !foundTask {
		return fmt.Errorf("Task ID: %v doesn't exist", taskId)
	}

	return WriteTasksToFile(updatedTasks)
}

func UpdateTaskStatus(taskId int, status string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var updatedTasks []Task
	var foundTask bool = false
	for _, task := range tasks {
		if task.ID == taskId {
			foundTask = true
			switch status {
			case TASK_STATUS_IN_PROGRESS:
				task.Status = TASK_STATUS_IN_PROGRESS
			case TASK_STATUS_DONE:
				task.Status = TASK_STATUS_DONE
			}
			task.UpdatedAt = time.Now()
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !foundTask {
		return fmt.Errorf("Task ID: %v not found", taskId)
	}

	return WriteTasksToFile(updatedTasks)
}

func DeleteTask(taskId int) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var foundTask bool = false
	for i, task := range tasks {
		if task.ID == taskId {
			foundTask = true
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}

	if !foundTask {
		return fmt.Errorf("Task ID: %v doesn't exist", taskId)
	}

	return WriteTasksToFile(tasks)
}
