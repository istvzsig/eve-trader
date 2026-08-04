package main

import (
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
	c esi.MarketClient,
) (map[int]*BestOrder, error) {

	byType := map[int]*BestOrder{}

	page := 1
	totalPages := 1

	for page <= totalPages {

		orders, pages, err := c.GetRegionOrders(
			page,
			config.JitaID,
		)

		if err != nil {
			return nil, err
		}

		totalPages = pages

		for _, o := range orders {

			b := byType[o.TypeID]

			if b == nil {
				b = &BestOrder{}
				byType[o.TypeID] = b
			}

			if o.IsBuyOrder {

				if b.Buy == 0 || o.Price > b.Buy {
					b.Buy = o.Price
					b.BuyVolume = o.VolumeRemain
				}

			} else {

				if b.Sell == 0 || o.Price < b.Sell {
					b.Sell = o.Price
					b.SellVolume = o.VolumeRemain
				}
			}
		}

		page++
	}

	return byType, nil
}
