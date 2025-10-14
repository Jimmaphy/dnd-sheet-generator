package services

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

type JSONService struct {
	folder string
}

// A new instance of JSONService is created with for the specified folder.
// The folder will be placed inside the storage directory, creating it if it doesn't exist.
func NewJSONService(folder string) (*JSONService, error) {
	folderPath := "storage/" + folder
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return nil, errors.New("error while loading storage directory: " + folder)
	}

	return &JSONService{
		folder: folderPath,
	}, nil
}

// Save will create or overwrite a JSON file with the specified name in the service's folder.
// The exetension ".json" will be automatically added to the file name.
// Any file with the same name will be overwritten.
// The data parameter should be a struct that can be marshaled into JSON format.
func (service *JSONService) Save(fileName string, data any) error {
	filePath := service.folder + "/" + fileName + ".json"

	file, err := os.Create(filePath)
	if err != nil {
		return errors.New("error while creating JSON file: " + filePath)
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(data)
	if err != nil {
		return errors.New("error while encoding JSON data: " + filePath)
	}

	return nil
}

// List will read the folder containing JSON files and return a list of file names without the ".json" extension.
// If the folder does not exist or cannot be read, an error will be returned.
func (service *JSONService) List() ([]string, error) {
	entries, err := os.ReadDir(service.folder)
	if err != nil {
		return nil, errors.New("error while reading JSON directory: " + service.folder)
	}

	var characterNames []string = []string{}
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".json") {
			name := strings.TrimSuffix(entry.Name(), ".json")
			characterNames = append(characterNames, name)
		}
	}

	return characterNames, nil
}

// Delete will remove the JSON file with the specified name from the service's folder.
// The exetension ".json" will be automatically added to the file name.
// If the file does not exist, this function does nothing.
func (service *JSONService) Delete(fileName string) error {
	filePath := service.folder + "/" + fileName + ".json"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil
	}

	return os.Remove(filePath)
}
