package esi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/istvzsig/eve-trader/internal/model"
)

const (
	BaseURL = "https://esi.evetech.net/latest"
	JitaID  = 10000002
)

func GetOrders(typeID int) ([]model.MarketOrder, error) {

	url := fmt.Sprintf(
		"%s/markets/%d/orders/?type_id=%d",
		BaseURL,
		JitaID,
		typeID,
	)

	fmt.Println("esi_url", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var orders []model.MarketOrder

	err = json.NewDecoder(resp.Body).Decode(&orders)

	return orders, err
}
