package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		help()
		return
	}

	switch os.Args[1] {

	case "calculate":
		RunCalculator(os.Args[2:])

	case "market":
		RunMarket(os.Args[2:])

	default:
		help()
	}
}

func help() {

	fmt.Println(`
EVE Trader

Commands:

	calculate SELL BUY [VOLUME]

	Example:
		eve-trader calculate 417 282 3


	market TYPE_ID

	Example:
		eve-trader market 22291
`)
}
