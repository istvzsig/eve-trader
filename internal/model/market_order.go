package model

type MarketOrder struct {
	OrderID      int64   `json:"order_id"`
	TypeID       int     `json:"type_id"`
	LocationID   int64   `json:"location_id"`
	Price        float64 `json:"price"`
	VolumeRemain int     `json:"volume_remain"`
	IsBuyOrder   bool    `json:"is_buy_order"`
}
