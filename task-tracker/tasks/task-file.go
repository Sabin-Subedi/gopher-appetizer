package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func getFilePath() string {
	cwd, error := os.Getwd()

	if error != nil {
		fmt.Println("Error getting file Path")
	}

	return filepath.Join(cwd, "tasks.json")
}

func readTaskFromJSONFile() []Task {
	tasks := []Task{}
	taskFilePath := getFilePath()
	_, error := os.Stat(taskFilePath)
	if error != nil {
		if os.IsNotExist(error) {
			// Create the file if it doesn't exist
			file, error := os.Create(taskFilePath)
			os.WriteFile(taskFilePath, []byte("[]"), os.ModeAppend.Perm())

			if error != nil {
				fmt.Println("Error creating tasks file: ", error)
			}

			// Close the file
			file.Close()
		} else {
			fmt.Println("Error checking if file exists: ", error)
		}
	}

	// Read tasks from file
	bytes, error := os.ReadFile(taskFilePath)

	if error != nil {
		fmt.Println("Error reading tasks from file: ", error)
	}

	// Unmarshal the JSON bytes into the tasks slice
	error = json.Unmarshal(bytes, &tasks)

	if error != nil {
		fmt.Println("Error unmarshalling tasks: ", error)
	}

	return tasks
}

func writeTaskToJSONFile(tasks []Task) {
	jsonData, err := json.Marshal(tasks)
	taskFilePath := getFilePath()

	if err != nil {
		fmt.Println("Error marshalling tasks: ", err)
	}

	// Write the JSON data to the file
	err = os.WriteFile(taskFilePath, jsonData, os.ModeAppend.Perm())

	if err != nil {
		fmt.Println("Error writing tasks to file: ", err)
	}
}
