package main

import (
	"fmt"
	"os"

	"github.com/istvzsig/eve-trader/internal/config"
	"github.com/istvzsig/eve-trader/internal/esi"
)

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	esiClient := esi.NewClient(config.BaseURL, config.Timeout)

	switch os.Args[1] {

	case "calculate":
		RunCalculator(os.Args[2:])

	case "item":
		RunItem(esiClient, os.Args[2:])

	case "margin-trade":
		RunMarginTrader(esiClient, os.Args[2:])

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
2. > item ITEM_NAME
Example:
	eve-trader item "Ballistic Control System II"
3. > margin-trade PERCENT
Examples:
	eve-trader margin-trade 15%
	`)
}
