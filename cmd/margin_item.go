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

func RunItemPrice(
	itemResolver esi.ItemResolver,
	marketClient esi.MarketClient,
	args []string,
) {
	if len(args) != 2 {
		fmt.Println(`usage: ./eve-trader item-price "ITEM_NAME" [QUANTITY]`)
		return
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	typeID, err := itemResolver.FindItemID(ctx, args[0])
	if err != nil {
		fmt.Println("unable to find item:", err)
		return
	}

	quantity, err := strconv.Atoi(args[1])
	if err != nil || quantity <= 0 {
		fmt.Println("invalid quantity:", args[1])
		return
	}

	orders, err := marketClient.GetOrders(
		ctx,
		typeID,
		config.JitaID, // The Forge
	)
	if err != nil {
		fmt.Println("unable to get market orders:", err)
		return
	}

	buy, sell, err := market.BestPrices(orders)
	if err != nil {
		fmt.Println("unable to determine market prices:", err)
		return
	}

	opportunity := market.Calculate(
		buy,
		sell,
		quantity,
	)

	fmt.Println()
	fmt.Println("===========================================")
	fmt.Println("ITEM TRADER")
	fmt.Println("===========================================")

	fmt.Printf("TypeID:          %d\n", typeID)
	fmt.Printf("Name:            %s\n", args[0])
	fmt.Printf("Quantity:        %d\n", quantity)

	fmt.Println("-------------------------------------------")

	fmt.Printf(
		"Buy Price:       %s (%s)\n",
		format.ISK(buy),
		format.ISK(buy*float64(quantity)),
	)

	fmt.Printf(
		"Sell Price:      %s (%s)\n",
		format.ISK(sell),
		format.ISK(sell*float64(quantity)),
	)

	fmt.Println("-------------------------------------------")

	fmt.Printf(
		"Gross Profit:    %s/unit\n",
		format.ISK(opportunity.GrossProfit),
	)

	fmt.Printf(
		"Net Profit:      %s/unit\n",
		format.ISK(opportunity.NetProfit),
	)

	fmt.Printf(
		"Gross Margin:    %.2f%%\n",
		opportunity.GrossMargin,
	)

	fmt.Printf(
		"Net ROI:         %.2f%%\n",
		opportunity.ROI,
	)

	fmt.Println("-------------------------------------------")

	fmt.Printf(
		"Total Gross:     %s\n",
		format.ISK(opportunity.TotalGrossProfit),
	)

	fmt.Printf(
		"Total Net:       %s\n",
		format.ISK(opportunity.TotalNetProfit),
	)

	fmt.Println("-------------------------------------------")

	fmt.Printf(
		"Verdict:         %s\n",
		opportunity.Verdict,
	)

	fmt.Println("===========================================")
}
