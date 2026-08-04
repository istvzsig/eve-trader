package trader

import (
	"github.com/istvzsig/eve-trader/internal/config"
	"github.com/istvzsig/eve-trader/internal/esi"
	"github.com/istvzsig/eve-trader/internal/market"
	"github.com/istvzsig/eve-trader/internal/model"
	"github.com/istvzsig/eve-trader/internal/util"
)

type BestOrder struct {
	Buy        float64
	Sell       float64
	BuyVolume  int
	SellVolume int
}

func BuildCandidates(
	orders map[int]*BestOrder,
	low float64,
	high float64,
	maxVolume int,
) []model.Candidate {

	candidates := make([]model.Candidate, 0)

	for itemID, order := range orders {

		if order.Buy == 0 || order.Sell == 0 {
			continue
		}

		volume := util.MinValue(
			maxVolume,
			order.BuyVolume,
			order.SellVolume,
		)

		if volume <= 0 {
			continue
		}

		opportunity := market.Calculate(
			order.Buy,
			order.Sell,
			volume,
		)

		if opportunity.ROI < low ||
			opportunity.ROI > high {
			continue
		}

		candidates = append(
			candidates,
			model.Candidate{
				TypeID: itemID,
				Opp:    opportunity,
			},
		)
	}

	return candidates
}

func FetchBestOrders(
	client esi.MarketClient,
) (map[int]*BestOrder, error) {

	byItem := make(map[int]*BestOrder)

	page := 1
	totalPages := 1

	for page <= totalPages {

		orders, pages, err := client.GetRegionOrders(
			page,
			config.JitaID,
		)

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
