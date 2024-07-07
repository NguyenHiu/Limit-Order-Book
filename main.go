package main

import (
	"fmt"
	"log"

	"github.com/NguyenHiu/lob/analysis"
	"github.com/NguyenHiu/lob/data"
	"github.com/NguyenHiu/lob/lob"
)

func main() {
	orders, err := data.LoadOrders("../lightning-exchange/data/aliceOrders.json")
	if err != nil {
		log.Fatal(err)
	}

	_a := analysis.NewAnalysis()
	_lob := lob.NewLOB(_a)
	for _, order := range orders {
		_lob.AddOrder(lob.NewOrder(order.Price, order.Amount, order.Side))
	}

	fmt.Printf("matched amount: %v\n", _a.MatchedAmount)
}
