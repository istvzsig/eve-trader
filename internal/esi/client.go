package esi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/istvzsig/eve-trader/internal/model"
)

const (
	BaseURL = "https://esi.evetech.net/latest"
	JitaID  = 10000002
)

var httpClient = &http.Client{
	Timeout: 20 * time.Second,
}

var (
	typeNameCache = map[int]string{}
	typeNameMu    = sync.Mutex{}
)

func GetTypeName(typeID int) (string, error) {
	typeNameMu.Lock()
	if n, ok := typeNameCache[typeID]; ok {
		typeNameMu.Unlock()
		return n, nil
	}
	typeNameMu.Unlock()

	url := fmt.Sprintf("%s/universe/types/%d/", BaseURL, typeID)

	resp, err := httpClient.Get(url)
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

	typeNameMu.Lock()
	typeNameCache[typeID] = ti.Name
	typeNameMu.Unlock()

	return ti.Name, nil
}

func GetOrders(typeID int) ([]model.MarketOrder, error) {
	url := fmt.Sprintf(
		"%s/markets/%d/orders/?type_id=%d",
		BaseURL,
		JitaID,
		typeID,
	)

	resp, err := httpClient.Get(url)
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

func GetRegionOrders(regionID int, page int) ([]model.MarketOrder, int, error) {
	url := fmt.Sprintf(
		"%s/markets/%d/orders/?page=%d",
		BaseURL,
		regionID,
		page,
	)

	resp, err := httpClient.Get(url)
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

	// ESI returns total pages in X-Pages header
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
