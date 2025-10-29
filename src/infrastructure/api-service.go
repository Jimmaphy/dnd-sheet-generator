package infrastructure

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type APIService struct {
	baseURL string
}

// While creating a new instance of the APIService, the base URL for the API is specified.
// Allowing for reduced input in later calls.
func NewAPIService(baseURL string) *APIService {
	return &APIService{
		baseURL: baseURL,
	}
}

// GetData will perform a GET request to the specified endpoint of the API.
// The endpoint parameter should be the path after the base URL.
// The index parameter is optional and can be used to specify a particular resource.
// There can be more than one index, whoch will be concatenated with slashes.
// The function returns the response body as object.
// It is assumed that the response body is in JSON format and can be unmarshaled into the provided structure.
func (service *APIService) GetData(structure any, endpoint string, index ...string) error {
	url := service.baseURL + "/" + endpoint

	for _, indexName := range index {
		url += "/" + indexName
	}

	response, err := http.Get(url)
	if err != nil {
		return errors.New("error while making GET request to: " + url)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.New("error while reading response body from: " + url)
	}

	err = json.Unmarshal(body, structure)
	if err != nil {
		return errors.New("error while unmarshaling JSON data from: " + url)
	}

	return nil
}
