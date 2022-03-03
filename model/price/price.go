package price

type Price struct {
	Symbol        string  `json:"symbol,omitempty"`
	Price         float64 `json:"price,omitempty"`
	OriginalPrice float64 `json:"original_price,omitempty"`
	Volume24      float64 `json:"volume_24,omitempty"`
	Change24H     float64 `json:"change_24,omitempty"`
	Logo          string  `json:"logo"`
}
type Prices []Price

func (p Price) String() string {
	return "price test"
}
