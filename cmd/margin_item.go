package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/istvzsig/eve-trader/internal/esi"
)

func RunItemPrice(ir esi.ItemResolver, args []string) {
	if len(args) != 2 {
		fmt.Println(`usage: ./eve-trader item-price "ITEM_NAME" [QUANTITY]`)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	typeID, err := ir.FindItemID(ctx, args[0])
	if err != nil {
		fmt.Println("item not found:", args[0])
		return
	}

	quantity, err := strconv.Atoi(args[1])
	if err != nil || quantity <= 0 {
		fmt.Println("invalid quantity:", args[1])
		return
	}

	fmt.Println("TypeID:", typeID)
	fmt.Println("Name:", args[0])
	fmt.Println("Quantity:", quantity)
}
