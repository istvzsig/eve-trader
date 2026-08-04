package main

import (
	"fmt"
	"strconv"

	"github.com/istvzsig/eve-trader/internal/esi"
	"github.com/istvzsig/eve-trader/internal/parse"
)

type MarginOptions struct {
	Target    float64
	Low       float64
	High      float64
	MaxVolume int
}

func RunMarginTrader(
	esiClient esi.EsiClient,
	args []string,
) {

	target, err := parse.PercentArg(args[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	maxVolume := 1

	if len(args) == 2 {
		maxVolume, err = strconv.Atoi(args[1])

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	opts := MarginOptions{
		Target:    target,
		Low:       target - 5,
		High:      target + 5,
		MaxVolume: maxVolume,
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	orders, err := FetchBestOrders(esiClient)

	if err != nil {
		fmt.Println(err)
		return
	}

	candidates := BuildCandidates(
		orders,
		opts.Low,
		opts.High,
		opts.MaxVolume,
	)

	SortCandidates(candidates)

	PrintHeader(opts)

	PrintCandidates(
		esiClient,
		candidates,
	)
}
