package model

type Opportunity struct {
	BuyPrice  float64
	SellPrice float64
	Volume    int

	GrossProfit      float64 // per unit
	NetProfit        float64 // per unit
	TotalGrossProfit float64
	TotalNetProfit   float64

	GrossMargin float64
	ROI         float64

	Verdict string
}
