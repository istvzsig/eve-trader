package esi

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/istvzsig/eve-trader/internal/model"
)

func (c *Client) GetOrders(
	typeID int,
	regionID int,
) ([]model.MarketOrder, error) {

	url := fmt.Sprintf(
		"%s/markets/%d/orders/?type_id=%d",
		c.baseURL,
		regionID,
		typeID,
	)

	resp, err := c.httpClient.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 ||
		resp.StatusCode >= 300 {

		return nil,
			fmt.Errorf(
				"market error: %s",
				resp.Status,
			)
	}

	var orders []model.MarketOrder

	err = json.NewDecoder(
		resp.Body,
	).Decode(&orders)

	return orders, err
}

func (c *Client) GetRegionOrders(
	page int,
	regionID int,
) ([]model.MarketOrder, int, error) {

	url := fmt.Sprintf(
		"%s/markets/%d/orders/?page=%d",
		c.baseURL,
		regionID,
		page,
	)

	resp, err := c.httpClient.Get(url)

	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, 0, err
	}

	if resp.StatusCode < 200 ||
		resp.StatusCode >= 300 {

		return nil,
			0,
			fmt.Errorf(
				"esi market error: %s",
				string(body),
			)
	}

	totalPages := 1

	if value := resp.Header.Get("X-Pages"); value != "" {

		if n, err := strconv.Atoi(value); err == nil {

			totalPages = n
		}
	}

	var orders []model.MarketOrder

	err = json.Unmarshal(
		body,
		&orders,
	)

	return orders, totalPages, err
}
