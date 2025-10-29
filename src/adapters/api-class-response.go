package adapters

type ClassResponse struct {
	Count   int        `json:"count"`
	Results []*Listing `json:"results"`
}

// List provides a list of class listings from the response.
// It will be returned as a slice of strings.
func (response *ClassResponse) List() []Listing {
	var listings []Listing

	for _, class := range response.Results {
		listings = append(listings, *class)
	}

	return listings
}
