package adapters

type OptionListing struct {
	OptionType string   `json:"option_type"`
	Item       *Listing `json:"item"`
}
