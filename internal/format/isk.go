package format

import "fmt"

func ISK(value float64) string {

	if value >= 1_000_000_000 {
		return fmt.Sprintf("%.2fB ISK", value/1_000_000_000)
	}

	if value >= 1_000_000 {
		return fmt.Sprintf("%.2fM ISK", value/1_000_000)
	}

	if value >= 1_000 {
		return fmt.Sprintf("%.2fK ISK", value/1_000)
	}

	return fmt.Sprintf("%.0f ISK", value)
}
