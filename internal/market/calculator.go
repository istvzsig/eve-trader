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
	net := sell*(1-BrokerFee-SalesTax) - buy

	totalGross := gross * float64(volume)
	totalNet := net * float64(volume)

	grossMargin := (gross / buy) * 100

	roi := (net / buy) * 100

	return model.Opportunity{
		BuyPrice:  buy,
		SellPrice: sell,
		Volume:    volume,

		GrossProfit:      gross,
		NetProfit:        net,
		TotalGrossProfit: totalGross,
		TotalNetProfit:   totalNet,

		GrossMargin: grossMargin,
		ROI:         roi,

		Verdict: verdict(roi),
	}
}

func verdict(roi float64) string {
	switch {
	case roi >= 20:
		return "✅ GOOD"
	case roi >= 10:
		return "👍 ACCEPTABLE"
	case roi >= 5:
		return "🟡 WEAK"
	default:
		return "⏭️ SKIP"
	}
}
