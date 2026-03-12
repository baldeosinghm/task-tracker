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

// Update the id count and return it as an id
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

// Save the task's fields to JSON file
func (fm FileManager) SaveTask(data any) error {
	// If there are zero tasks, create a file to store tasks using os.Create
	// and pass a file path as the parameter
	file, err := os.Create(fm.OutputFilePath)

	// Check if any errors occurred during file creation due to file path
	if err != nil {
		return err
	}

	// Defer closing the file. You should always defer closing a file right
	// after it has been created.
	defer file.Close()

	// If JSON file is empty, add new task without reading file
	if FileExists(fm.OutputFilePath)


	if isEmpty == true {
		// Write to file using json's NewEncoder() package
		encoder := json.NewEncoder(file) // Give location of the file to write to
		err = encoder.Encode(data)       // Give data to put in JSON file

		// Check for errors
		if err != nil {
			return err
		}

		return nil
	}

	// If JSON file isn't empty, read it
	fileData, err := os.ReadFile(fm.OutputFilePath)

	// Unmarshal data into map
	var tasks map[string]any
	json.Unmarshal(fileData, &tasks)
	if err != nil {
		return errors.New("Failed to parsed JSON-encoded data.")
	}

	// Marshal the modified struct back into a JSON formatted byte slice
	updatedFileData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return errors.New("Failed to write data to JSON file.")
	}

	// Write the new task to JSON file
	err = os.WriteFile(fm.OutputFilePath, updatedFileData, 0644)
	if err != nil {
		return err
	}

	return nil
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

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true // File exists so no error
	}
	if errors.Is(err, os.ErrNotExist) {
		return False
	}
	return false
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
