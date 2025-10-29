package adapters

type ApiProficiency struct {
	OptionSetType string           `json:"option_set_type"`
	Options       []*OptionListing `json:"options"`
}
