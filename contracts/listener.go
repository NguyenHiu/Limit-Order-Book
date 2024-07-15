package contracts

import (
	"context"
	"fmt"
	"log"

	"github.com/NguyenHiu/lob/contracts/generated/lOB"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func ListenContract(inst *lOB.LOB) {
	opt := &bind.WatchOpts{Context: context.Background()}
	go watchMatchedEvent(inst, opt)
	go watchMatchTimestampEvent(inst, opt)
}

func watchMatchedEvent(inst *lOB.LOB, opt *bind.WatchOpts) {
	logs := make(chan *lOB.LOBMyMatchEvent)
	sub, err := inst.WatchMyMatchEvent(opt, logs)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLogs := <-logs:
			fmt.Println("Matched, amount: ", vLogs.Arg1)
		}
	}
}

func watchMatchTimestampEvent(inst *lOB.LOB, opt *bind.WatchOpts) {
	logs := make(chan *lOB.LOBMatchTimestamp)
	sub, err := inst.WatchMatchTimestamp(opt, logs)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLogs := <-logs:
			fmt.Println("Time: ", vLogs.Arg0)
		}
	}
}
