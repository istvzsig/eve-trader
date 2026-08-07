package main

import (
	"fmt"
	"os"

	"github.com/istvzsig/eve-trader/internal/config"
	"github.com/istvzsig/eve-trader/internal/esi"
	"github.com/istvzsig/eve-trader/internal/trader"
)

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	esiClient := esi.NewClient(config.BaseURL, config.Timeout)
	marginTrader := trader.NewMarginTrader(esiClient, esiClient)

	switch os.Args[1] {

	case "calculate":
		RunCalculator(os.Args[2:])

	case "margin-item":
		RunMarginItem(esiClient, os.Args[2:])

	case "margin-trade":
		RunMarginTrader(marginTrader, esiClient, os.Args[2:])

	case "wh-help":
		WormholeHelper()

	default:
		help()
	}
}

func help() {
	fmt.Println(`
EVE Trader Help:

Commands:
1. > calculate SELL BUY [VOLUME]
Example:
	eve-trader calculate 417 282 3
2. > margin-item ITEM_NAME
Example:
	eve-trader margin-item "Ballistic Control System II"
3. > margin-trade PERCENT [MAX_VOLUME]
Examples:
	eve-trader margin-trade 15%
4. > wh-help
Prints wormhole bookmarking checklist.
	`)
}
