package trader

import (
	"context"

	"github.com/istvzsig/eve-trader/internal/config"
	"github.com/istvzsig/eve-trader/internal/esi"
)

type BestOrder struct {
	Buy        float64
	Sell       float64
	BuyVolume  int
	SellVolume int
}

func FetchBestOrders(
	ctx context.Context,
	client esi.MarketClient,
) (map[int]*BestOrder, error) {

	byItem := make(map[int]*BestOrder)

	page := 1
	totalPages := 1

	for page <= totalPages {

		orders, pages, err := client.GetRegionOrders(ctx, page, config.JitaID)

		if err != nil {
			return nil, err
		}

		totalPages = pages

		for _, order := range orders {

			best := byItem[order.TypeID]

			if best == nil {
				best = &BestOrder{}
				byItem[order.TypeID] = best
			}

			if order.IsBuyOrder {

				if best.Buy == 0 ||
					order.Price > best.Buy {

					best.Buy = order.Price
					best.BuyVolume = order.VolumeRemain
				}

			} else {

				if best.Sell == 0 ||
					order.Price < best.Sell {

					best.Sell = order.Price
					best.SellVolume = order.VolumeRemain
				}
			}
		}

		page++
	}

	return byItem, nil
}
