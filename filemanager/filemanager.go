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

// Update JSON file with task
func (fm *FileManager) UpdateFile(data any) error {
	// Read the file
	tasks, err := fm.Parser()

	// Convert incoming data into a map via JSON round-trip
	// This avoids import cycle from importing the Task struct entirely
	rawBytes, err := json.Marshal(data)
	if err != nil {
		return errors.New("Failed to marshal incoming data.")
	}

	var newEntry map[string]any
	if err = json.Unmarshal(rawBytes, &newEntry); err != nil {
		return errors.New("Failed to convert data to map.")
	}

	// Append new task, data, to map tasks
	tasks = append(tasks, newEntry)
	err = fm.WriteResult(tasks)
	if err != nil {
		return err
	}
	return nil
}

func (fm FileManager) WriteResult(tasks []map[string]any) error {
	// Marshal the modified struct back into a JSON formatted byte slice
	updatedFileData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return errors.New("Failed to marshal updated data.")
	}

	// NOTE: WriteFile will overwrite the JSON file
	if err = os.WriteFile(fm.OutputFilePath, updatedFileData, 0644); err != nil {
		return errors.New("Failed to write to JSON file.")
	}

	return nil
}

// Reads json file, parses it, and returns all tasks as a slice of maps
func (fm FileManager) Parser() (tasks []map[string]any, err error) {
	// Read json file
	data, err := os.ReadFile(fm.OutputFilePath)
	if err != nil {
		return nil, errors.New("Failed to read file.")
	}

	// Unmarshall json file
	var items []map[string]any
	err = json.Unmarshal(data, &items)
	if err != nil {
		return nil, errors.New("Failed to parse JSON file.")
	}

	return items, nil
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
