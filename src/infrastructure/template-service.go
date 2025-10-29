package infrastructure

import (
	"errors"
	"os"
	"text/template"
)

type TemplateService struct {
	filepath string
}

// Create a new instance of TemplateService.
// The template parameter specifies the name of the template file.
// This should be a valid path to a text file.
// An error is returned if the file does not exist or is a directory.
func NewTemplateService(template string) (*TemplateService, error) {
	filepath := "./templates/" + template

	file, err := os.Stat(filepath)
	if err != nil || file.IsDir() {
		return nil, errors.New("template file '" + template + "' does not exist or points to a directory")
	}

	return &TemplateService{filepath: filepath}, nil
}

// Read and return the content of the template file.
// An error is returned if the file cannot be read.
func (service *TemplateService) GetTemplateContent() (string, error) {
	content, err := os.ReadFile(service.filepath)
	if err != nil {
		return "", errors.New("error reading template file: " + err.Error())
	}

	return string(content), nil
}

// GetParsable applies the provided data to the template content.
// It uses a given data structure to replace placeholders in the template.
func (service *TemplateService) GetParsable() (*template.Template, error) {
	parsable, err := template.ParseFiles(service.filepath)
	if err != nil {
		return nil, err
	}

	return parsable, nil
}
