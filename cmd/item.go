package main

import (
	"fmt"
	"strconv"

	"github.com/istvzsig/eve-trader/internal/esi"
)

func RunItem(c esi.Esi, args []string) {
	if len(args) < 1 || len(args) > 2 {
		fmt.Println("usage: item TYPE_NAME")
		return
	}

	typeID, _ := strconv.Atoi(args[2])
	typeName, _ := c.GetTypeName(typeID)

	fmt.Println("typeName", typeName)

	// typeID, err := strconv.Atoi(args[0])

	// if err != nil {
	// 	fmt.Println("invalid type id")
	// 	return
	// }

	// volume := 1

	// if len(args) == 2 {

	// 	volume, err = strconv.Atoi(args[1])

	// 	if err != nil || volume <= 0 {
	// 		fmt.Println("invalid quantity")
	// 		return
	// 	}
	// }

	// orders, err := esi.GetOrders(typeID)

	// if err != nil {
	// 	panic(err)
	// }

	// var buy float64
	// var sell float64

	// for _, o := range orders {

	// 	if o.IsBuyOrder {

	// 		if buy == 0 || o.Price > buy {
	// 			buy = o.Price
	// 		}

	// 	} else {

	// 		if sell == 0 || o.Price < sell {
	// 			sell = o.Price
	// 		}
	// 	}
	// }

	// result := market.Calculate(
	// 	buy,
	// 	sell,
	// 	volume,
	// )

	// fmt.Println("==============================")
	// fmt.Println("ITEM CHECK")
	// fmt.Println("==============================")

	// fmt.Println("Buy:", format.ISK(result.BuyPrice))
	// fmt.Println("Sell:", format.ISK(result.SellPrice))
	// fmt.Println("Volume:", result.Volume)
	// fmt.Printf("ROI: %.2f%%\n", result.ROI)
	// fmt.Println("Verdict:", result.Verdict)
	// fmt.Println("Total Profit:", format.ISK(result.TotalNetProfit))
}
