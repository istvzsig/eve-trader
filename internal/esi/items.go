package esi

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/istvzsig/eve-trader/internal/model"
)

func (c *Client) GetItemName(
	typeID int,
) (string, error) {

	c.cacheMu.RLock()

	if name, ok := c.cache[typeID]; ok {
		c.cacheMu.RUnlock()
		return name, nil
	}

	c.cacheMu.RUnlock()

	url := fmt.Sprintf(
		"%s/universe/types/%d/",
		c.baseURL,
		typeID,
	)

	resp, err := c.httpClient.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {

		body, _ := io.ReadAll(resp.Body)

		return "",
			fmt.Errorf(
				"esi type error: %s",
				string(body),
			)
	}

	var item model.TypeInfo

	err = json.NewDecoder(
		resp.Body,
	).Decode(&item)

	if err != nil {
		return "", err
	}

	c.cacheMu.Lock()

	c.cache[typeID] = item.Name

	c.cacheMu.Unlock()

	return item.Name, nil
}

func (c *Client) FindItemID(
	name string,
) (int, error) {

	// TODO:
	// ESI has no direct name lookup.
	// Usually you use:
	// /universe/ids/
	//
	// implement later

	return 0,
		fmt.Errorf(
			"FindItemID not implemented",
		)
}
