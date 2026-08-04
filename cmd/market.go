package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/istvzsig/eve-trader/internal/config"
	"github.com/istvzsig/eve-trader/internal/esi"
	"github.com/istvzsig/eve-trader/internal/format"
	"github.com/istvzsig/eve-trader/internal/market"
	"github.com/istvzsig/eve-trader/internal/model"
	"github.com/istvzsig/eve-trader/internal/parse"
)

func RunMarginTrader(c esi.Esi, args []string) {
	if len(args) < 1 || len(args) > 2 {
		fmt.Println("usage: eve-trader margin-trade PCT [MAX_QUANTITY]")
		return
	}

	target, err := parse.PercentArg(args[0])
	if err != nil {
		fmt.Println("invalid percent:", args[0])
		return
	}

	maxVolume := 1

	if len(args) == 2 {
		maxVolume, err = strconv.Atoi(args[1])
		if err != nil || maxVolume <= 0 {
			fmt.Println("invalid quantity:", args[1])
			return
		}
	}

	low := target - 5
	high := target + 5

	type best struct {
		buy        float64
		sell       float64
		buyVolume  int
		sellVolume int
	}

	byType := map[int]*best{}

	page := 1
	totalPages := 1

	for page <= totalPages {

		orders, pages, err := c.GetRegionOrders(page, config.JitaID)

		if err != nil {
			fmt.Println("error fetching orders:", err)
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
					b.buyVolume = o.VolumeRemain
				}

			} else {

				if b.sell == 0 || o.Price < b.sell {
					b.sell = o.Price
					b.sellVolume = o.VolumeRemain
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

		volume := min(
			maxVolume,
			b.buyVolume,
			b.sellVolume,
		)

		if volume <= 0 {
			continue
		}

		opp := market.Calculate(
			b.buy,
			b.sell,
			volume,
		)

		if opp.ROI >= low && opp.ROI <= high {
			candidates = append(
				candidates,
				model.Candidate{
					TypeID: typeID,
					Opp:    opp,
				},
			)
		}
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Opp.TotalNetProfit >
			candidates[j].Opp.TotalNetProfit
	})

	fmt.Println("==============================")
	fmt.Println("MARGIN TRADER")
	fmt.Printf(
		"Target ROI: %.2f%% (%.2f%% - %.2f%%)\n",
		target,
		low,
		high,
	)
	fmt.Printf("Max Quantity: %d\n", maxVolume)
	fmt.Println("==============================")

	for i := 0; i < len(candidates) && i < 10; i++ {

		cn := candidates[i]

		name, err := c.GetTypeName(cn.TypeID)

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
			"ROI: %.2f%%  %s\n",
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
