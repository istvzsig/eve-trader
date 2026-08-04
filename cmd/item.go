package main

import (
	"fmt"

	"github.com/istvzsig/eve-trader/internal/esi"
)

func RunItem(
	client esi.EsiClient,
	args []string,
) {

	if len(args) != 1 {

		fmt.Println(
			`usage: eve-trader item "ITEM_NAME"`,
		)

		return
	}

	typeID, err := client.FindItemID(args[0])

	if err != nil {

		fmt.Println(
			"item not found:",
			args[0],
		)

		return
	}

	name, err := client.GetItemName(typeID)

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println("TypeID:", typeID)
	fmt.Println("Name:", name)
}
