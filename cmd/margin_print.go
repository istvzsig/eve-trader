package main

import (
	"fmt"

	"github.com/istvzsig/eve-trader/internal/esi"
	"github.com/istvzsig/eve-trader/internal/format"
	"github.com/istvzsig/eve-trader/internal/model"
)

func PrintHeader(
	opts MarginOptions,
) {
	fmt.Println("==============================")
	fmt.Println("MARGIN TRADER")

	fmt.Printf(
		"Target ROI: %.2f%% (%.2f%% - %.2f%%)\n",
		opts.Target,
		opts.Low,
		opts.High,
	)

	fmt.Printf(
		"Max Quantity: %d\n",
		opts.MaxVolume,
	)

	fmt.Println("==============================")
}

func PrintCandidates(
	items esi.ItemResolver,
	candidates []model.Candidate,
) {

	for i := 0; i < len(candidates) && i < 10; i++ {

		cn := candidates[i]

		name, err := items.GetItemName(
			cn.TypeID,
		)

		if err != nil {
			name = "(unknown)"
		}

		fmt.Printf(
			"\n[%d] %s\n",
			i+1,
			name,
		)

		fmt.Printf(
			"Buy: %s  Sell: %s\n",
			format.ISK(cn.Opp.BuyPrice),
			format.ISK(cn.Opp.SellPrice),
		)

		fmt.Printf(
			"Volume: %d\n",
			cn.Opp.Volume,
		)

		fmt.Printf(
			"ROI: %.2f%% %s\n",
			cn.Opp.ROI,
			cn.Opp.Verdict,
		)

		fmt.Printf(
			"Profit/unit: %s\n",
			format.ISK(cn.Opp.NetProfit),
		)

		fmt.Printf(
			"Total Profit: %s\n",
			format.ISK(cn.Opp.TotalNetProfit),
		)
	}
}
