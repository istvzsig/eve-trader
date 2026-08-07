package esi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/istvzsig/eve-trader/internal/model"
)

func (c *Client) FindItemID(
	ctx context.Context,
	name string,
) (int, error) {

	// TODO:
	// ESI has no direct name lookup.
	// Usually you use:
	// /universe/ids/
	//
	// implement later

	return 0, fmt.Errorf("esi: FindItemID not implemented")
}

func (c *Client) GetItemName(
	ctx context.Context,
	typeID int,
) (string, error) {

	c.cacheMu.RLock()

	if name, ok := c.cache[typeID]; ok {
		c.cacheMu.RUnlock()
		return name, nil
	}

	c.cacheMu.RUnlock()

	u := fmt.Sprintf("%s/universe/types/%d/", c.baseURL, typeID)

	body, _, err := c.get(ctx, u)
	if err != nil {
		return "", err
	}

	var item model.TypeInfo
	if err := json.Unmarshal(body, &item); err != nil {
		return "", fmt.Errorf("esi: decoding item name: %w", err)
	}

	c.cacheMu.Lock()
	c.cache[typeID] = item.Name
	c.cacheMu.Unlock()

	return item.Name, nil
}
