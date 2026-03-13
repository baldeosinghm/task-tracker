package filemanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

// Read the count in text file
func (fm FileManager) GetCount() (int64, error) {
	// Read file and get the count before it's updated
	text, err := os.ReadFile(fm.InputFilePath)

	if err != nil {
		return 0, errors.New("Failed to read file.")
	}

	count, err := strconv.ParseInt(string(text), 10, 0)

	if err != nil {
		return 0, errors.New("Failed to convert text to integer.")
	}

	return count, nil
}

// Update id count in text file and return new count as an id
func (fm FileManager) GenerateID() (int64, error) {
	// Read file and get the count before it's updated
	count, err := fm.GetCount()

	// Update file with new count
	updatedCountConverted := strconv.FormatInt(count+1, 10)
	err = os.WriteFile(fm.InputFilePath, []byte(updatedCountConverted), 0644)

	if err != nil {
		return 0, errors.New("Failed to write updated count to file.")
	}

	return count + 1, nil
}

// Read contents of JSON file and return it
func (fm FileManager) ReadJSONLines(task any) ([]string, error) {
	// Open file and read it's contents
	file, err := os.ReadFile(fm.OutputFilePath)

	if err != nil {
		return nil, err
	}

	// Unmarshall json file to read contents
	err = json.Unmarshal(file, task)

	// Parse contents for two fields: id and description

	return nil, nil
}

// Create JSON file and write data to it.
func (fm FileManager) CreateFile(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("Failed to create file.")
	}

	defer file.Close()

	// JSON package, NewEncoder() that converts data into JSON format
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("Failed to convert data to JSON.")
	}

	return nil
}

// // Update JSON file with task
func (fm *FileManager) UpdateFile(data any) error {
	// Read the file
	fileData, err := os.ReadFile(fm.OutputFilePath)
	if err != nil {
		return errors.New("Failed to read file.")
	}

	// Unmarshal existing data into a slice of maps
	tasks := make([]map[string]any, 0, 10)
	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &tasks)
		if err != nil {
			return errors.New("Failed to parse JSON file.")
		}
	}

	// Convert incoming data into a map via JSON round-trip
	// This avoids importing the Task struct entirely
	rawBytes, err := json.Marshal(data)
	if err != nil {
		return errors.New("Failed to marshal incoming data.")
	}

	var newEntry map[string]any
	if err = json.Unmarshal(rawBytes, &newEntry); err != nil {
		return errors.New("Failed to convert data to map.")
	}

	fmt.Println(newEntry)

	// Append new task, data, to map tasks
	tasks = append(tasks, newEntry)

	// Marshal the modified struct back into a JSON formatted byte slice
	updatedFileData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return errors.New("Failed to marshal updated data.")
	}

	// Write the new task to JSON file
	err = os.WriteFile(fm.OutputFilePath, updatedFileData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (fm FileManager) FileExists() bool {
	_, err := os.Stat(fm.OutputFilePath)
	if err == nil {
		return true // File exists so no error
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
