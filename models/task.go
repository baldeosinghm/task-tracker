package models

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"

	filemanager "github.com/baldeosinghm/task-tracker.git/fileManager"
)

type Task struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Create a task
func (task Task) Add() error {
	fmt.Print("add ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return err
	}

	task.Description = scanner.Text()

	// Store task in JSON file
	taskFile := filemanager.New("counter_state", "tasks.json")

	id, err := taskFile.UpdateCount()
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
func (task Task) DisplayAll(fm filemanager.FileManager) error {
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
