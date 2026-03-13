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
	case "list":
		task.ListAll(taskFile)
	}
}
