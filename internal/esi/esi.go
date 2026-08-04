package esi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"

	"github.com/istvzsig/eve-trader/internal/config"
	"github.com/istvzsig/eve-trader/internal/model"
)

// Esi interface defines the contract for ESI API interactions
type Esi interface {
	GetTypeName(typeID int) (string, error)
	GetOrders(typeID, regionID int) ([]model.MarketOrder, error)
	GetRegionOrders(page, regionID int) ([]model.MarketOrder, int, error)
}

// Client handles all ESI API interactions
type Client struct {
	baseURL    string
	httpClient *http.Client

	cache   map[int]string
	cacheMu sync.RWMutex
}

var _ Esi = (*Client)(nil)

// NewClient creates a new ESI client
func NewClient(baseURL string) Esi {
	return &Client{
		baseURL: config.BaseURL,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
		cache: make(map[int]string),
	}
}

// GetTypeName retrieves the name for a given type ID
func (c *Client) GetTypeName(typeID int) (string, error) {
	c.cacheMu.RLock()
	if n, ok := c.cache[typeID]; ok {
		c.cacheMu.RUnlock()
		return n, nil
	}
	c.cacheMu.RUnlock()

	url := fmt.Sprintf("%s/universe/types/%d/", c.baseURL, typeID)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("esi type name error: status=%d body=%s", resp.StatusCode, string(body))
	}

	var ti model.TypeInfo
	if err := json.NewDecoder(resp.Body).Decode(&ti); err != nil {
		return "", err
	}

	c.cacheMu.Lock()
	c.cache[typeID] = ti.Name
	c.cacheMu.Unlock()

	return ti.Name, nil
}

// GetOrders retrieves market orders for a type in Jita
func (c *Client) GetOrders(typeID, regionID int) ([]model.MarketOrder, error) {
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

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("esi markets orders: status %s", resp.Status)
	}

	defer resp.Body.Close()

	var orders []model.MarketOrder
	err = json.NewDecoder(resp.Body).Decode(&orders)
	return orders, err
}

// GetRegionOrders retrieves paginated market orders for a region
func (c *Client) GetRegionOrders(page, regionID int) ([]model.MarketOrder, int, error) {
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

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, 0, fmt.Errorf("esi error: status=%d body=%s", resp.StatusCode, string(body))
	}

	totalPages := 1
	if v := resp.Header.Get("X-Pages"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			totalPages = n
		}
	}

	var orders []model.MarketOrder
	if err := json.Unmarshal(body, &orders); err != nil {
		return nil, 0, err
	}

	return orders, totalPages, nil
}
