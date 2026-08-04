package trader

import (
	"sort"

	"github.com/istvzsig/eve-trader/internal/model"
)

func (t *MarginTrader) FindCandidates(
	opts MarginOptions,
) ([]model.Candidate, error) {

	orders, err := FetchBestOrders(
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
