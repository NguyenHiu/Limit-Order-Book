package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/NguyenHiu/lob/lob"
)

type OrderData struct {
	Price  int
	Amount int
	Side   bool
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

func SaveData(data []lob.Order, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file: %v", filename)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	if err := enc.Encode(&data); err != nil {
		return fmt.Errorf("cannot write data into file: %v", filename)
	}

	return nil
}
