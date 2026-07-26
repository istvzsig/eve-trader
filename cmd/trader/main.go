package main

import (
	"fmt"

	"github.com/istvzsig/eve-trader/internal/esi"
	"github.com/istvzsig/eve-trader/internal/market"
)

func main() {

	typeID := 22291 // Ballistic Control System II

	orders, err := esi.GetOrders(typeID)

	if err != nil {
		panic(err)
	}

	var buy float64
	var sell float64

	for _, o := range orders {

		if o.IsBuyOrder {

			if buy == 0 || o.Price > buy {
				buy = o.Price
			}

		} else {

			if sell == 0 || o.Price < sell {
				sell = o.Price
			}
		}
	}

	result := market.Calculate(
		buy,
		sell,
		1,
	)

	fmt.Println("==============================")
	fmt.Println("EVE TRADER")
	fmt.Println("==============================")

	fmt.Printf(
		"Buy: %.2f ISK\n",
		result.BuyPrice,
	)

	fmt.Printf(
		"Sell: %.2f ISK\n",
		result.SellPrice,
	)

	fmt.Printf(
		"Gross: %.2f ISK\n",
		result.GrossProfit,
	)

	fmt.Printf(
		"Net: %.2f ISK\n",
		result.NetProfit,
	)

	fmt.Printf(
		"ROI: %.2f%%\n",
		result.ROI,
	)

}
