package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
)

func ExportPriceCurve(priceCurve []*big.Int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file: %v", filename)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	if err := enc.Encode(&priceCurve); err != nil {
		return fmt.Errorf("cannot write data into file: %v", filename)
	}

	return nil
}
