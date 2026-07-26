package main

import (
	"fmt"
	"strconv"

	"github.com/istvzsig/eve-trader/internal/format"
	"github.com/istvzsig/eve-trader/internal/market"
)

func ParseISK(value float64) float64 {
	return value * 1000
}

func RunCalculator(args []string) {

	if len(args) < 2 {
		fmt.Println("usage: calculate SELL BUY [VOLUME]")
		return
	}

	sellInput, _ := strconv.ParseFloat(args[0], 64)
	buyInput, _ := strconv.ParseFloat(args[1], 64)

	sell := ParseISK(sellInput)
	buy := ParseISK(buyInput)

	volume := 1

	if len(args) >= 3 {
		volume, _ = strconv.Atoi(args[2])
	}

	result := market.Calculate(
		buy,
		sell,
		volume,
	)

	fmt.Println()
	fmt.Println("===========================================")

	fmt.Printf(
		"Sell Price:      %s\n",
		format.ISK(sell),
	)

	fmt.Printf(
		"Buy Price:       %s\n",
		format.ISK(buy),
	)

	fmt.Printf(
		"Volume:          %d\n",
		volume,
	)

	fmt.Println("-------------------------------------------")

	fmt.Printf(
		"Gross Profit:    %s/unit\n",
		format.ISK(result.GrossProfit),
	)

	fmt.Printf(
		"Net Profit:      %s/unit\n",
		format.ISK(result.NetProfit),
	)

	fmt.Printf(
		"Gross Margin:    %.2f%%\n",
		result.GrossMargin,
	)

	fmt.Printf(
		"Net ROI:         %.2f%%\n",
		result.ROI,
	)

	fmt.Println("-------------------------------------------")

	fmt.Printf(
		"Total Gross:     %s\n",
		format.ISK(result.GrossProfit*float64(volume)),
	)

	fmt.Printf(
		"Total Net:       %s\n",
		format.ISK(result.NetProfit*float64(volume)),
	)

	fmt.Println("-------------------------------------------")

	fmt.Printf(
		"Verdict:         %s\n",
		result.Verdict,
	)

	fmt.Println("===========================================")
}
