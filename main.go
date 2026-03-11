package main

import (
	"fmt"

	filemanager "github.com/baldeosinghm/task-tracker.git/fileManager"
	"github.com/baldeosinghm/task-tracker.git/models"
)

// Create main function
func main() {
	fmt.Println("Select an option: ")
	options := []string{
		"add",
		"update",
		"delete",
		"mark-in-progress",
		"mark-done",
		"list",
		"list done",
		"list in-progres"}
	for _, value := range options {
		fmt.Println(value)
	}

	var choice string
	fmt.Print("Your choice: ")
	fmt.Scan(&choice) // We can get the stored value w/ the ampersand operator

	var task models.Task
	fm := filemanager.New("counter_state", "tasks.json")

	switch choice {
	case "add":
		task.Add()
	case "list":
		task.DisplayAll(fm)
	}
}
