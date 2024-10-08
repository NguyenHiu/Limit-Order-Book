package lob

import (
	"fmt"
)

type LOB struct {
	BidOrders []*Order
	AskOrders []*Order

	NoFulfill int
}

func NewLOB() *LOB {
	return &LOB{
		BidOrders: make([]*Order, 0),
		AskOrders: make([]*Order, 0),
		NoFulfill: 0,
	}
}

func (l *LOB) AddOrder(_order *Order) {
	if _order.IsBid() {
		l.BidOrders = addOrderAccording(_order, l.BidOrders)
	} else {
		l.AskOrders = addOrderAccording(_order, l.AskOrders)
	}
	l.Matching()
}

func addOrderAccording(_order *Order, _orders []*Order) []*Order {
	l := len(_orders)
	if l == 0 {
		_orders = append(_orders, _order)
	} else if l == 1 {
		if (_order.IsBid() && _order.Price > _orders[0].Price) ||
			(!_order.IsBid() && _order.Price < _orders[0].Price) {
			_orders = append([]*Order{_order}, _orders...)
		} else {
			_orders = append(_orders, _order)
		}
	} else {
		for i := 0; i < l; i++ {
			if (_order.IsBid() && _order.Price > _orders[i].Price) ||
				(!_order.IsBid() && _order.Price < _orders[i].Price) {
				_orders = append(_orders, nil)
				copy(_orders[i+1:], _orders[i:])
				_orders[i] = _order
				return _orders
			}
		}
		_orders = append(_orders, _order)
	}

	return _orders
}

func (l *LOB) Matching() {
	for canMatch(l.BidOrders, l.AskOrders) {
		minAmount := l.BidOrders[0].Amount
		if minAmount > l.AskOrders[0].Amount {
			minAmount = l.AskOrders[0].Amount
		}

		l.BidOrders[0].Amount -= minAmount
		l.AskOrders[0].Amount -= minAmount

		if l.BidOrders[0].Amount == 0 {
			l.NoFulfill += 1
			l.BidOrders = l.BidOrders[1:]
		}
		if l.AskOrders[0].Amount == 0 {
			l.NoFulfill += 1
			l.AskOrders = l.AskOrders[1:]
		}
	}
}

func canMatch(_bidOrders []*Order, _askOrders []*Order) bool {
	if len(_bidOrders) == 0 || len(_askOrders) == 0 {
		return false
	}

	return _bidOrders[0].Price >= _askOrders[0].Price
}

func (l *LOB) GetOrderBook() string {
	s := ""
	for _, o := range l.BidOrders {
		s += fmt.Sprintf("%v,%v,", o.Price, o.Amount)
	}
	for _, o := range l.AskOrders {
		s += fmt.Sprintf("%v,%v,", o.Price, o.Amount)
	}
	return s
}
