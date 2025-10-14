package services

import (
	"encoding/json"
	"errors"
	"os"
)

type JSONService struct {
	filePath string
}

// Create a new instance of JSONService with the specified folder and filename.
// The folder is where the JSON file will be stored, and the filename is the name of the JSON file.
// Before creating the service, ensure that the folder exists or create it if it doesn't.
// The extension and path seperators are handled internally.
func NewJSONService(folder, filename string) (*JSONService, error) {
	folderPath := "storage/" + folder
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return nil, errors.New("error while loading storage directory: " + folder)
	}

	return &JSONService{
		filePath: folderPath + "/" + filename + ".json",
	}, nil
}

// Save will save the provided data as a JSON file
// Any file with the same name will be overwritten.
// The data parameter should be a struct that can be marshaled into JSON format.
func (service *JSONService) Save(data any) error {
	file, err := os.Create(service.filePath)
	if err != nil {
		return errors.New("error while creating JSON file: " + service.filePath)
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(data)
	if err != nil {
		return errors.New("error while encoding JSON data: " + service.filePath)
	}

	return nil
}
