package market

import "github.com/istvzsig/eve-trader/internal/model"

const (
	BrokerFee = 0.03
	SalesTax  = 0.075
)

func Calculate(
	buy float64,
	sell float64,
	volume int,
) model.Opportunity {

	gross := sell - buy

	fees := sell * (BrokerFee + SalesTax)

	net := sell - fees - buy

	return model.Opportunity{
		BuyPrice:  buy,
		SellPrice: sell,
		Volume:    volume,

		GrossProfit: gross,
		NetProfit:   net,

		ROI: net / buy * 100,
	}
}
