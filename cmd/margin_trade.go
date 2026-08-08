package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/istvzsig/eve-trader/internal/esi"
	"github.com/istvzsig/eve-trader/internal/format"
	"github.com/istvzsig/eve-trader/internal/model"
	"github.com/istvzsig/eve-trader/internal/parse"
	"github.com/istvzsig/eve-trader/internal/trader"
)

func RunMarginTrader(
	marginTrader *trader.MarginTrader,
	itemResolver esi.ItemResolver,
	args []string,
) {

	if len(args) < 1 || len(args) > 2 {
		fmt.Println(
			"usage: eve-trader margin-trade ROI [MAX_VOLUME]",
		)
		return
	}

	target, err := parse.PercentArg(args[0])
	if err != nil {
		fmt.Println(
			"invalid percent:",
			args[0],
		)
		return
	}

	maxVolume := 1

	if len(args) == 2 {
		maxVolume, err = strconv.Atoi(args[1])

		if err != nil || maxVolume <= 0 {
			fmt.Println(
				"invalid volume:",
				args[1],
			)
			return
		}
	}

	low := target - 5
	high := target + 5

	opts := trader.MarginOptions{
		Target:    target,
		Low:       low,
		High:      high,
		MaxVolume: maxVolume,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	candidates, err := marginTrader.FindCandidates(ctx, opts)

	if err != nil {
		fmt.Println(
			"unable to find candidates:",
			err,
		)
		return
	}

	PrintHeader(opts)

	PrintCandidates(ctx, itemResolver, candidates)
}

func PrintHeader(
	opts trader.MarginOptions,
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
	ctx context.Context,
	items esi.ItemResolver,
	candidates []model.Candidate,
) {

	for i := 0; i < len(candidates) && i < 10; i++ {

		cn := candidates[i]

		name, err := items.GetItemName(ctx, cn.TypeID)

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
