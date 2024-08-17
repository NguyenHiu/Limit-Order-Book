package lob

type Order struct {
	Price  int
	Amount int
	Side   bool
	// Owner  string
}

func NewOrder(_price, _amount int, _side bool) *Order {
	return &Order{
		Price:  _price,
		Amount: _amount,
		Side:   _side,
	}
}

func (o *Order) IsBid() bool {
	return (o.Side == BID)
}
