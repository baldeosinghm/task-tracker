package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
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
	err = task.save(taskFile) // Save task in JSON file
	if err != nil {
		return err
	}

	fmt.Println("Task successfully created!")
	return nil
}

// Update a task's description
func Update(command []string, taskFile filemanager.FileManager) error {
	id, err := strconv.ParseFloat(command[0], 64)
	if err != nil {
		return errors.New("Failed to convert string ID to int.")
	}
	description := strings.Join(command[1:], " ")
	tasks, err := taskFile.Parser()
	if err != nil {
		return err
	}

	// Look through json file for task with matching ID
	for _, task := range tasks {
		if task["id"] == id {
			task["description"] = description // Update description of task
		}
	}

	// Update json file with updated task
	err = taskFile.WriteResult(tasks)
	if err != nil {
		return err
	}
	return nil
}

// Delete a task
func Delete(taskID []string, taskFile filemanager.FileManager) error {
	taskToDeleteID, err := strconv.ParseFloat(taskID[0], 64)
	if err != nil {
		return errors.New("Failed to convert string ID to float64.")
	}

	tasks, err := taskFile.Parser()

	// Parse json data
	var fileteredTasks []map[string]any
	for _, task := range tasks {
		if task["id"] != taskToDeleteID {
			fileteredTasks = append(fileteredTasks, task)
		}
	}

	// Marshall data to be stored in json file
	convertedData, err := json.MarshalIndent(fileteredTasks, "", " ")
	if err != nil {
		return errors.New("Failed to marshal JSON data.")
	}

	// Update file holding json data
	err = os.WriteFile(taskFile.OutputFilePath, convertedData, 0644)
	if err != nil {
		return errors.New("Failed to write to file.")
	}

	fmt.Printf("Task %v deleted!", taskID)

	return nil
}

// Mark task as in progress or done
func (task Task) UpdateStatus(status string, taskDetails []string, taskFile filemanager.FileManager) error {
	taskID, err := strconv.ParseFloat(taskDetails[0], 64)
	tasks, err := taskFile.Parser()
	if err != nil {
		return err
	}
	for _, task := range tasks {
		if task["id"] == taskID {
			switch status {
			case "mark-in-progress":
				task["status"] = "in-progress"
			case "mark-done":
				task["status"] = "done"
			}
		}
	}
	err = taskFile.WriteResult(tasks)
	if err != nil {
		return err
	}
	fmt.Println("Task status updated!")
	return nil
}

// List all tasks, those that are done, or are in progress
func List(taskDetails []string, taskFile filemanager.FileManager) error {
	tasks, err := taskFile.Parser()
	if err != nil {
		return err
	}
	prettifiedFile, err := json.MarshalIndent(tasks, "", " ")

	command := strings.Join(taskDetails, "")
	switch command {
	case "":
		if err != nil {
			return err
		}
		fmt.Println(string(prettifiedFile))
	case "done":
		for _, task := range tasks {
			if task["status"] == "done" {
				prettifiedFile, err := json.MarshalIndent(task, "", " ")
				if err != nil {
					return err
				} else {
					fmt.Println(string(prettifiedFile))
				}
			}
		}
	case "todo":
		for _, task := range tasks {
			if task["status"] == "todo" {
				prettifiedFile, err := json.MarshalIndent(task, "", " ")
				if err != nil {
					return err
				} else {
					fmt.Println(string(prettifiedFile))
				}
			}
		}
	case "in-progress":
		for _, task := range tasks {
			if task["status"] == "in-progress" {
				prettifiedFile, err := json.MarshalIndent(task, "", " ")
				if err != nil {
					return err
				} else {
					fmt.Println(string(prettifiedFile))
				}
			}
		}
	}
	return nil
}

// Save the task's fields to JSON file
func (task Task) save(fm filemanager.FileManager) error {
	if !fm.FileExists() {
		// Create a slice of type "Task" and initialize it with a single element, Task
		err := fm.CreateFile([]Task{task})
		if err != nil {
			return err
		}
	} else {
		err := fm.UpdateFile(task)
		if err != nil {
			return err
		}
	}
	return nil
}
