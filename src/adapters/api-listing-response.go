package adapters

import "strings"

type ListingResponse struct {
	Index     string     `json:"index"`
	Name      string     `json:"name"`
	Equipment []*Listing `json:"equipment"`
	Url       string     `json:"url"`
	UpdatedAt string     `json:"updated_at"`
}

// List provides a list of item listings from the response.
// It will be returned as a slice of strings.
// Each string represents the name of an item listing.
// Magic items and shield are excluded from the list.
func (response *ListingResponse) List() []Listing {
	var listings []Listing

	for _, item := range response.Equipment {
		if strings.Contains(item.Name, "Shield") {
			continue
		}

		if !strings.Contains(item.Url, "equipment") {
			continue
		}

		listings = append(listings, *item)
	}

	return listings
}
