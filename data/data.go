package data

import (
	"encoding/json"
	"os"
)

type OrderData struct {
	Price  int  `json:"Price"`
	Amount int  `json:"Amount"`
	Side   bool `json:"Side"`
}

func LoadOrders(path string) ([]*OrderData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	orders := make([]*OrderData, 0)
	if err := json.Unmarshal(data, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}
