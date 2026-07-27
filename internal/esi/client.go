package esi

import (
	"encoding/json"
	"fmt"
	"net/http"
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
