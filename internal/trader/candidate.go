package trader

import (
	"context"
	"sort"

	"github.com/istvzsig/eve-trader/internal/market"
	"github.com/istvzsig/eve-trader/internal/model"
	"github.com/istvzsig/eve-trader/internal/util"
)

func (t *MarginTrader) FindCandidates(
	ctx context.Context,
	opts MarginOptions,
) ([]model.Candidate, error) {

	orders, err := FetchBestOrders(
		ctx,
		t.market,
	)

	if err != nil {
		return nil, err
	}

	candidates := BuildCandidates(
		orders,
		opts.Low,
		opts.High,
		opts.MaxVolume,
	)

	sortCandidates(candidates)

	return candidates, nil
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

func sortCandidates(
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
