package models

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	filemanager "github.com/baldeosinghm/task-tracker.git/filemanager"
)

type Task struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Create a task
func (task Task) Add(details []string) error {
	task.Description = strings.Join(details, " ")

	// Store task in JSON file
	var taskFile filemanager.FileManager
	taskFile.InputFilePath, taskFile.OutputFilePath = "counter_state.txt", "tasks.json"

	id, err := taskFile.GenerateID()
	task.Status = "todo" // statuses: todo, in-progress, done

	if err != nil {
		return err
	}

	task.ID = id
	taskFile.SaveTask(task)

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

	var taskStruct Task
	err = json.Unmarshal(file, &taskStruct)

	if err != nil {
		return err
	}

	prettifiedFile, err := json.MarshalIndent(taskStruct, "", " ")

	if err != nil {
		return err
	}

	fmt.Println(string(prettifiedFile))
	return nil
}

// // Update task (change details)
// func (task Task) Update() error {
// 	// Read the file

// 	// Parse IDs for requested task's id and description

// 	// If after parsing there's no match, return error: unable to find id

// 	// Return description matched to id

// }

// Delete task

// Mark task as in progress

// Mark task as done

// List all tasks that are done

// List all tasks not done

// List all tasks in progress
