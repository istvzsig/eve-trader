package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/istvzsig/eve-trader/internal/esi"
	"github.com/istvzsig/eve-trader/internal/format"
	"github.com/istvzsig/eve-trader/internal/market"
	"github.com/istvzsig/eve-trader/internal/model"
	"github.com/istvzsig/eve-trader/internal/parse"
)

func RunMarginTrader(args []string) {
	if len(args) != 1 {
		fmt.Println("usage: eve-trader margin-trade PCT   (e.g. 15 or 20%)")
		return
	}

	target, err := parse.PercentArg(args[0])
	if err != nil {
		fmt.Println("invalid percent:", args[0])
		return
	}

	low := target - 5
	high := target + 5

	type best struct {
		buy  float64 // best bid (max of buy orders)
		sell float64 // best ask (min of sell orders)
	}

	byType := map[int]*best{}

	regionID := esi.JitaID
	page := 1
	totalPages := 1

	for page <= totalPages {
		orders, pages, err := esi.GetRegionOrders(regionID, page)
		if err != nil {
			fmt.Println("error fetching region orders:", err)
			return
		}
		totalPages = pages

		for _, o := range orders {
			b := byType[o.TypeID]
			if b == nil {
				b = &best{}
				byType[o.TypeID] = b
			}

			if o.IsBuyOrder {
				if b.buy == 0 || o.Price > b.buy {
					b.buy = o.Price
				}
			} else {
				if b.sell == 0 || o.Price < b.sell {
					b.sell = o.Price
				}
			}
		}

		page++
	}

	candidates := make([]model.Candidate, 0)

	for typeID, b := range byType {
		if b.buy == 0 || b.sell == 0 {
			continue
		}

		opp := market.Calculate(b.buy, b.sell, 1)
		if opp.ROI >= low && opp.ROI <= high {
			candidates = append(candidates, model.Candidate{TypeID: typeID, Opp: opp})
		}
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Opp.ROI > candidates[j].Opp.ROI
	})

	fmt.Println("==============================")
	fmt.Println("MARGIN TRADER (ROI within target +/- 5%)")
	fmt.Printf("Target: %.2f%%  Range: %.2f%% - %.2f%%\n", target, low, high)
	fmt.Println("==============================")

	for i := 0; i < len(candidates) && i < 10; i++ {
		c := candidates[i]
		fmt.Printf(
			"[%d] TypeID=%d Buy=%s Sell=%s ROI=%.2f%% Verdict=%s Net=%s\n",
			i+1,
			c.TypeID,
			format.ISK(c.Opp.BuyPrice),
			format.ISK(c.Opp.SellPrice),
			c.Opp.ROI,
			c.Opp.Verdict,
			format.ISK(c.Opp.NetProfit),
		)
	}
}

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
