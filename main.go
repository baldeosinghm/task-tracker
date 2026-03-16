package main

import (
	filemanager "github.com/baldeosinghm/task-tracker.git/filemanager"
	"github.com/baldeosinghm/task-tracker.git/iomanager"
	"github.com/baldeosinghm/task-tracker.git/models"
)

// Create main function
func main() {
	iomanager.WelcomeMsg()
	userInput, _ := iomanager.InputParser()
	command, taskDetails := userInput[0], userInput[1:]

	var task models.Task
	taskFile := filemanager.New("counter_state.txt", "tasks.json")

	switch command {
	case "add":
		task.Add(taskDetails, taskFile)
	case "update":
		models.Update(taskDetails, taskFile)
	case "delete":
		models.Delete(taskDetails, taskFile)
	case "mark-in-progress", "mark-done":
		task.UpdateStatus(command, taskDetails, taskFile)
	case "list":
		models.List(taskDetails, taskFile)
	}
}
