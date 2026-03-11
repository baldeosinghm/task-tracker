package filemanager

import (
	"encoding/json"
	"errors"
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

// Update the count in text file
func (fm FileManager) UpdateCount() (int64, error) {
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

// Save the task's fields to JSON file
func (fm FileManager) SaveTask(data any) error {
	// Step 1: Create file by using os.Create and pass a file path as the parameter
	file, err := os.Create(fm.OutputFilePath)

	// Step 2: Check if any errors occurred during file creation due to file path
	if err != nil {
		return err
	}

	// Step 3: Defer closing the file. You should always defer closing a file right
	// after it has been created.
	defer file.Close()

	// Step 4: Write to file using json's NewEncoder() package
	encoder := json.NewEncoder(file) // Give location of the file to write to
	err = encoder.Encode(data)       // Give data to put in JSON file

	// Check for errors
	if err != nil {
		return err
	}

	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
