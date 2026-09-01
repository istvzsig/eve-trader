package model

// Contract is a subset of ESI's character contracts response - only
// the fields useful for listing/inspecting contracts. Full schema
// also has date_accepted, date_completed, buyout, collateral, etc.,
// added if/when a command actually needs them.
type Contract struct {
	ContractID  int64   `json:"contract_id"`
	Type        string  `json:"type"`
	Status      string  `json:"status"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Reward      float64 `json:"reward"`
	Volume      float64 `json:"volume"`
	DateIssued  string  `json:"date_issued"`
	DateExpired string  `json:"date_expired"`
}
