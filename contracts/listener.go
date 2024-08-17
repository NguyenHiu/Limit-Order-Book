package contracts

import (
	"context"
	"fmt"
	"log"

	"github.com/NguyenHiu/lob/contracts/generated/lOB"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type Listener struct {
	NoFulfillOrders int
}

func NewListner() *Listener {
	return &Listener{
		NoFulfillOrders: 0,
	}
}

func (l *Listener) StartListen(inst *lOB.LOB) {
	opt := &bind.WatchOpts{Context: context.Background()}
	go l.watchMatchAnOrder(inst, opt)
}

func (l *Listener) watchMatchAnOrder(inst *lOB.LOB, opt *bind.WatchOpts) {
	logs := make(chan *lOB.LOBMatchAnOrder)
	sub, err := inst.WatchMatchAnOrder(opt, logs)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case <-logs:
			fmt.Println("Matched an order")
			l.NoFulfillOrders += 1
		}
	}
}
