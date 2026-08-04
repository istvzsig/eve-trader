package esi

import "github.com/istvzsig/eve-trader/internal/model"

type ItemResolver interface {
	FindItemID(name string) (int, error)
	GetItemName(id int) (string, error)
}

type MarketClient interface {
	GetRegionOrders(page, regionID int) ([]model.MarketOrder, int, error)
	GetOrders(itemID, regionID int) ([]model.MarketOrder, error)
}

type EsiClient interface {
	ItemResolver
	MarketClient
}
