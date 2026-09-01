package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/istvzsig/eve-trader/internal/config"
	"github.com/istvzsig/eve-trader/internal/esi"
	"github.com/istvzsig/eve-trader/internal/trader"
)

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "configuration error: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		help()
		return
	}

	esiClient := esi.NewClient(
		config.BaseURL,
		cfg.ClientID,
		cfg.ClientSecret,
		cfg.RedirectURI,
		config.Timeout,
	)

	marginTrader := trader.NewMarginTrader(esiClient, esiClient)

	switch os.Args[1] {
	case "auth":
		if err := runAuth(esiClient); err != nil {
			fmt.Fprintf(os.Stderr, "authentication failed: %v\n", err)
			os.Exit(1)
		}

	case "calculate":
		RunCalculator(os.Args[2:])

	case "item-price":
		RunItemPrice(esiClient, esiClient, os.Args[2:])

	case "margin-trade":
		RunMarginTrader(marginTrader, esiClient, os.Args[2:])

	case "wh-help":
		RunWormholeHelper()

	case "isk-challenge":
		RunISKChallenge()

	default:
		help()
	}
}

func runAuth(client *esi.Client) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if _, err := client.GetAccessToken(ctx); err != nil {
		return err
	}

	fmt.Println("EVE Online authentication successful.")
	return nil
}

func help() {
	fmt.Println(`
EVE Trader Help:

Commands:
1. > calculate SELL BUY [VOLUME]
Example:
	eve-trader calculate 41m 282k 3
2. > margin-item ITEM_NAME
Example:
	./eve-trader item-price "Ballistic Control System II"
With quantity:
	./eve-trader item-price "Ballistic Control System II" [QUANTITY]
3. > margin-trade PERCENT [MAX_VOLUME]
Examples:
	eve-trader margin-trade 15%
4. > wh-help
	Prints wormhole bookmarking checklist.
5. > isk-challenge TARGET CURRENT
Description: (TARGET - CURRENT)
Examples:
	eve-trader isk-challenge 420m 55k
6. > auth
	Authenticate with EVE Online.
	`)
}
