package model

type MarketOrder struct {
	OrderID      int64   `json:"order_id"`
	TypeID       int     `json:"type_id"`
	LocationID   int64   `json:"location_id"`
	Price        float64 `json:"price"`
	VolumeRemain int     `json:"volume_remain"`
	IsBuyOrder   bool    `json:"is_buy_order"`
}

type Opportunity struct {
	TypeID int

	BuyPrice  float64
	SellPrice float64
	Volume    int

	GrossProfit float64
	NetProfit   float64

	GrossMargin float64
	ROI         float64

	Verdict string
}
