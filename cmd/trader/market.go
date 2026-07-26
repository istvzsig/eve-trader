package main

import (
	"fmt"
	"strconv"

	"github.com/istvzsig/eve-trader/internal/esi"
	"github.com/istvzsig/eve-trader/internal/format"
	"github.com/istvzsig/eve-trader/internal/market"
)

func RunMarket(args []string) {

	if len(args) < 1 {
		fmt.Println("usage: market TYPE_ID")
		return
	}

	typeID, _ := strconv.Atoi(args[0])

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
	fmt.Println("MARKET CHECK")
	fmt.Println("==============================")

	fmt.Println("Buy:", format.ISK(result.BuyPrice))
	fmt.Println("Sell:", format.ISK(result.SellPrice))
	fmt.Println("ROI:", result.ROI)
	fmt.Println("Verdict:", result.Verdict)

}
