package esi

import (
	"context"

	"github.com/istvzsig/eve-trader/internal/model"
)

type ItemResolver interface {
	FindItemID(ctx context.Context, name string) (int, error)
	ResolveNames(ctx context.Context, ids []int) (map[int]string, error)
	FindImplants(ctx context.Context) ([]model.TypeInfo, error)
}

type MarketClient interface {
	GetRegionOrders(ctx context.Context, page, regionID int) ([]model.MarketOrder, int, error)
	GetOrders(ctx context.Context, itemID, regionID int) ([]model.MarketOrder, error)
}
