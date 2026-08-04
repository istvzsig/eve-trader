package main

import (
	"fmt"
	"strconv"

	"github.com/istvzsig/eve-trader/internal/esi"
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

	candidates, err := marginTrader.FindCandidates(opts)

	if err != nil {
		fmt.Println(
			"unable to find candidates:",
			err,
		)
		return
	}

	PrintHeader(opts)

	PrintCandidates(
		itemResolver,
		candidates,
	)
}
