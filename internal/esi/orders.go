package esi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/istvzsig/eve-trader/internal/model"
)

func (c *Client) GetOrders(ctx context.Context, typeID, regionID int) ([]model.MarketOrder, error) {
	u := fmt.Sprintf("%s/markets/%d/orders/", c.baseURL, regionID)
	q := url.Values{"type_id": {strconv.Itoa(typeID)}}

	body, _, err := c.get(ctx, u+"?"+q.Encode())
	if err != nil {
		return nil, err
	}

	var orders []model.MarketOrder
	if err := json.Unmarshal(body, &orders); err != nil {
		return nil, fmt.Errorf("esi: decoding orders: %w", err)
	}
	return orders, nil
}

func (c *Client) GetRegionOrders(ctx context.Context, page, regionID int) ([]model.MarketOrder, int, error) {
	u := fmt.Sprintf("%s/markets/%d/orders/", c.baseURL, regionID)
	q := url.Values{"page": {strconv.Itoa(page)}}

	body, headers, err := c.get(ctx, u+"?"+q.Encode())
	if err != nil {
		return nil, 0, err
	}

	totalPages := 1
	if v := headers.Get("X-Pages"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			totalPages = n
		} else {
			// don't fail the request over a malformed pagination header -
			// but don't stay silent either
			log.Printf("esi: could not parse X-Pages header %q: %v", v, err)
		}
	}

	var orders []model.MarketOrder
	if err := json.Unmarshal(body, &orders); err != nil {
		return nil, 0, fmt.Errorf("esi: decoding region orders: %w", err)
	}
	return orders, totalPages, nil
}
