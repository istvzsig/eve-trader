package main

import (
	"sort"

	"github.com/istvzsig/eve-trader/internal/market"
	"github.com/istvzsig/eve-trader/internal/model"
)

func BuildCandidates(
	byType map[int]*BestOrder,
	low float64,
	high float64,
	maxVolume int,
) []model.Candidate {

	candidates := make([]model.Candidate, 0)

	for typeID, b := range byType {

		if b.Buy == 0 || b.Sell == 0 {
			continue
		}

		if b.Sell <= b.Buy {
			continue
		}

		volume := min(
			maxVolume,
			b.BuyVolume,
			b.SellVolume,
		)

		if volume <= 0 {
			continue
		}

		opp := market.Calculate(
			b.Buy,
			b.Sell,
			volume,
		)

		if opp.ROI < low || opp.ROI > high {
			continue
		}

		candidates = append(
			candidates,
			model.Candidate{
				TypeID: typeID,
				Opp:    opp,
			},
		)
	}

	return candidates
}

func SortCandidates(
	candidates []model.Candidate,
) {

	sort.Slice(
		candidates,
		func(i, j int) bool {

			return candidates[i].Opp.TotalNetProfit >
				candidates[j].Opp.TotalNetProfit
		},
	)
}
