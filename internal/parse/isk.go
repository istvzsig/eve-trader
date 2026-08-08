package parse

import (
	"fmt"
	"strconv"
	"strings"
)

func ISK(value string) (float64, error) {
	value = strings.TrimSpace(strings.ToLower(value))

	multiplier := 1.0

	switch {
	case strings.HasSuffix(value, "k"):
		multiplier = 1_000
		value = strings.TrimSuffix(value, "k")
	case strings.HasSuffix(value, "m"):
		multiplier = 1_000_000
		value = strings.TrimSuffix(value, "m")
	case strings.HasSuffix(value, "b"):
		multiplier = 1_000_000_000
		value = strings.TrimSuffix(value, "b")
	}

	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid ISK: %q", value)
	}

	return number * multiplier, nil
}
