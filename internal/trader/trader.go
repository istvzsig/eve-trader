package trader

import (
	"github.com/istvzsig/eve-trader/internal/esi"
)

type MarginTrader struct {
	market esi.MarketClient
	items  esi.ItemResolver
}

func NewMarginTrader(
	market esi.MarketClient,
	items esi.ItemResolver,
) *MarginTrader {

	return &MarginTrader{
		market: market,
		items:  items,
	}
}
