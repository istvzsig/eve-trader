package esi

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

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

type categoryResponse struct {
	Name   string `json:"name"`
	Groups []int  `json:"groups"`
}

type groupResponse struct {
	Types []int `json:"types"`
}

// FindImplants discovers all implant types from the ESI Implant category,
// resolves their names, and returns them as TypeInfo values.
//
// The implant category contains multiple groups. Each group's type IDs are
// resolved separately in batches to avoid exceeding ESI's request limit.
func (c *Client) FindImplants(
	ctx context.Context,
) ([]model.TypeInfo, error) {
	const implantCategoryID = 20
	const nameBatchSize = 1000

	// TODO: Confirm whether ESI's maximum IDs per request changes in the future.

	// Get the Implant category.
	u := fmt.Sprintf(
		"%s/universe/categories/%d/",
		c.baseURL,
		implantCategoryID,
	)

	body, _, err := c.get(ctx, u)
	if err != nil {
		return nil, fmt.Errorf(
			"esi: getting implant category: %w",
			err,
		)
	}

	var category categoryResponse
	if err := json.Unmarshal(body, &category); err != nil {
		return nil, fmt.Errorf(
			"esi: decoding implant category: %w",
			err,
		)
	}

	implants := make([]model.TypeInfo, 0)

	// Get all implant type IDs from each implant group.
	for _, groupID := range category.Groups {
		u = fmt.Sprintf(
			"%s/universe/groups/%d/",
			c.baseURL,
			groupID,
		)

		body, _, err := c.get(ctx, u)
		if err != nil {
			return nil, fmt.Errorf(
				"esi: getting implant group %d: %w",
				groupID,
				err,
			)
		}

		var group groupResponse
		if err := json.Unmarshal(body, &group); err != nil {
			return nil, fmt.Errorf(
				"esi: decoding implant group %d: %w",
				groupID,
				err,
			)
		}

		// Resolve names in batches because a group can contain more
		// type IDs than ESI accepts in a single request.
		for start := 0; start < len(group.Types); start += nameBatchSize {
			end := start + nameBatchSize

			if end > len(group.Types) {
				end = len(group.Types)
			}

			batch := group.Types[start:end]

			names, err := c.ResolveNames(ctx, batch)
			if err != nil {
				return nil, fmt.Errorf(
					"esi: resolving names for implant group %d: %w",
					groupID,
					err,
				)
			}

			for _, typeID := range batch {
				name, ok := names[typeID]
				if !ok {
					continue
				}

				implants = append(implants, model.TypeInfo{
					TypeID: strconv.Itoa(typeID),
					Name:   name,
				})
			}
		}
	}

	return implants, nil
}
