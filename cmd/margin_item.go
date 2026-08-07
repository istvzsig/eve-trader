package main

import (
	"context"
	"fmt"
	"time"

	"github.com/istvzsig/eve-trader/internal/esi"
)

func RunMarginItem(ir esi.ItemResolver, args []string) {

	if len(args) != 1 {
		fmt.Println(`usage: eve-trader margin-item "ITEM_NAME"`)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	typeID, err := ir.FindItemID(ctx, args[0])
	if err != nil {
		fmt.Println("item not found:", args[0])
		return
	}

	name, err := ir.GetItemName(ctx, typeID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("TypeID:", typeID)
	fmt.Println("Name:", name)
}
