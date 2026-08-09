package market

import (
	"fmt"

	"github.com/istvzsig/eve-trader/internal/model"
)

func BestPrices(
	orders []model.MarketOrder,
) (buy float64, sell float64, err error) {

	for _, order := range orders {
		if order.IsBuyOrder {
			if order.Price > buy {
				buy = order.Price
			}
			continue
		}

		if sell == 0 || order.Price < sell {
			sell = order.Price
		}
	}

	if buy == 0 {
		return 0, 0, fmt.Errorf("no buy orders found")
	}

	if sell == 0 {
		return 0, 0, fmt.Errorf("no sell orders found")
	}

	return buy, sell, nil
}
