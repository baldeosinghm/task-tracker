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

	// userInput[1:] is a slice, so you'll have to parse the slice for the update command

	var task models.Task
	taskFile := filemanager.New("counter_state", "tasks.json")

	switch command {
	case "add":
		task.Add(taskDetails)
	case "list":
		task.ListAll(taskFile)
		// case "update":
		// 	task.Update(taskDetails)
	}
}
