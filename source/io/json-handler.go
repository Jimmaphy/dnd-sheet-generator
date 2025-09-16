package io

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type JsonHandler struct {
	StoragePath string
}

func NewJsonHandler(storagePath string) (*JsonHandler, error) {
	if err := os.MkdirAll(storagePath, 0755); err != nil {
		return nil, err
	}
	return &JsonHandler{StoragePath: storagePath}, nil
}

func (handler *JsonHandler) List() ([]string, error) {
	entries, err := os.ReadDir(handler.StoragePath)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
			name := strings.TrimSuffix(entry.Name(), ".json")
			files = append(files, name)
		}
	}

	return files, nil
}

func (handler *JsonHandler) Store(obj any, filename string) error {
	filename += ".json"

	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	path := "./" + filepath.Join(handler.StoragePath, filename)
	return os.WriteFile(path, data, 0644)
}

func (handler *JsonHandler) Read(filename string, obj any) error {
	filename += ".json"
	path := filepath.Join(handler.StoragePath, filename)

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, obj)
}
