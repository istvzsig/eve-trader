package esi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/istvzsig/eve-trader/internal/model"
)

type inventoryType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type idsResponse struct {
	InventoryTypes []inventoryType `json:"inventory_types"`
}

func (c *Client) FindItemID(
	ctx context.Context,
	name string,
) (int, error) {

	u := fmt.Sprintf("%s/universe/ids/", c.baseURL)

	payload, err := json.Marshal([]string{name})
	if err != nil {
		return 0, fmt.Errorf("esi: encoding item name: %w", err)
	}

	body, _, err := c.post(ctx, u, payload)
	if err != nil {
		return 0, fmt.Errorf("esi: finding item %q: %w", name, err)
	}

	var result idsResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, fmt.Errorf("esi: decoding item IDs: %w", err)
	}

	for _, item := range result.InventoryTypes {
		if item.Name == name {
			return item.ID, nil
		}
	}

	return 0, fmt.Errorf("esi: item not found: %q", name)
}

func (c *Client) ResolveNames(
	ctx context.Context,
	ids []int,
) (map[int]string, error) {

	if len(ids) == 0 {
		return map[int]string{}, nil
	}

	u := fmt.Sprintf("%s/universe/names/", c.baseURL)

	payload, err := json.Marshal(ids)
	if err != nil {
		return nil, fmt.Errorf("esi: encoding IDs: %w", err)
	}

	body, _, err := c.post(ctx, u, payload)
	if err != nil {
		return nil, fmt.Errorf("esi: resolving names: %w", err)
	}

	var result []model.NameInfo

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("esi: decoding names: %w", err)
	}

	names := make(map[int]string, len(result))

	for _, item := range result {
		if item.Category != "inventory_type" {
			continue
		}

		names[item.ID] = item.Name
	}

	return names, nil
}
