package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/istvzsig/eve-trader/internal/format"
	"github.com/istvzsig/eve-trader/internal/parse"
)

func RunISKChallenge() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s isk-challenge <target> <current>\n", os.Args[0])
		return
	}

	t, _ := strconv.ParseFloat(os.Args[2], 64)
	target := parse.ISKMillion(t)

	c, _ := strconv.ParseFloat(os.Args[3], 64)
	current := parse.ISKMillion(c)

	remaining := target - current
	if remaining < 0 {
		remaining = 0
	}

	progress := float64(current) / float64(target) * 100

	fmt.Println("🚀 EVE ISK CHALLENGE")
	fmt.Println("===================================")
	fmt.Printf("Target:      %s\n", format.ISK(target))
	fmt.Printf("Exploration: %s\n", format.ISK(current))
	fmt.Printf("Remaining:   %s\n", format.ISK(remaining))
	fmt.Printf("Progress:    %.1f%%\n", progress)
}
