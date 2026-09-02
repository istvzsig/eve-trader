package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/istvzsig/eve-trader/internal/config"
	"github.com/istvzsig/eve-trader/internal/esi"
	"github.com/istvzsig/eve-trader/internal/format"
	"github.com/istvzsig/eve-trader/internal/market"
)

type ImplantPrice struct {
	TypeID string
	Name   string
	Price  string
}

func RunImplantsPrice(
	esiClient esi.ItemResolver,
	marketClient esi.MarketClient,
) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)
	defer cancel()

	implants, err := esiClient.FindImplants(ctx)
	if err != nil {
		fmt.Println("find implants:", err)
		return
	}

	if len(implants) == 0 {
		fmt.Println("no implants found")
		return
	}

	results := make([]ImplantPrice, 0, len(implants))

	for _, implant := range implants {
		typeID, err := strconv.Atoi(implant.TypeID)
		if err != nil {
			fmt.Printf(
				"invalid type ID %q for %s: %v\n",
				implant.TypeID,
				implant.Name,
				err,
			)
			continue
		}

		orders, err := marketClient.GetOrders(
			ctx,
			typeID,
			config.JitaID,
		)
		if err != nil {
			fmt.Printf(
				"get orders for %s: %v\n",
				implant.Name,
				err,
			)
			continue
		}

		_, price, err := market.BestPrices(orders)
		if err != nil {
			fmt.Printf(
				"get best price for %s: %v\n",
				implant.Name,
				err,
			)
			continue
		}

		results = append(results, ImplantPrice{
			TypeID: implant.TypeID,
			Name:   implant.Name,
			Price:  format.ISK(price),
		})
	}

	for _, result := range results {
		fmt.Printf(
			"%s : %s\n",
			result.Name,
			result.Price,
		)
	}
}
