package models

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/baldeosinghm/task-tracker.git/filemanager"
)

type Task struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Create a task
func (task Task) Add(details []string, taskFile filemanager.FileManager) error {
	id, err := taskFile.GenerateID()

	if err != nil {
		return err
	}

	task.ID = id
	task.Description = strings.Join(details, " ")
	task.Status = "todo" // statuses: todo, in-progress, done
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	task.save(taskFile) // Save task in JSON file

	fmt.Println("Task successfully created!")
	return nil
}

// List all tasks
func (task Task) ListAll(fm filemanager.FileManager) error {
	// Return if error file doesn't exist
	file, err := os.ReadFile(fm.OutputFilePath)

	if err != nil {
		return err
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)

	if err != nil {
		return err
	}

	prettifiedFile, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		return err
	}

	fmt.Println(string(prettifiedFile))
	return nil
}

// Save the task's fields to JSON file
func (task Task) save(fm filemanager.FileManager) error {
	if !fm.FileExists() {
		// Create a new slice of type "Task" and initializes it with a single element, Task
		fm.CreateFile([]Task{task})
		return nil
	}
	fm.UpdateFile(task)
	return nil
}

// Delete task

// Mark task as in progress

// Mark task as done

// List all tasks that are done

// List all tasks not done

// List all tasks in progress
