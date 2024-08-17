package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/NguyenHiu/lob/contracts"
	"github.com/NguyenHiu/lob/data"
	"github.com/NguyenHiu/lob/lob"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
)

func offchain(orders []*data.OrderData) {
	localOrderBook := lob.NewLOB()

	for _, order := range orders {
		localOrderBook.AddOrder(&lob.Order{
			Price:  order.Price,
			Amount: order.Amount,
			Side:   order.Side,
		})
	}

	fmt.Println("Done, the number of fulfilled orders: ", localOrderBook.NoFulfill)
}

func onchain(orders []*data.OrderData) {
	deployerPrvkeyHex := "abf82ff96b463e9d82b83cb9bb450fe87e6166d4db6d7021d0c71d7e960d5abe"
	_, inst := contracts.DeployContracts(deployerPrvkeyHex)

	_listenerInstance := contracts.NewListner()
	go _listenerInstance.StartListen(inst)

	traderPrvkeyHex := "dcb7118c9946a39cd40b661e0d368e4afcc3cc48d21aa750d8164ca2e44961c4"
	prvkey, _ := crypto.HexToECDSA(traderPrvkeyHex)
	_addr := crypto.PubkeyToAddress(prvkey.PublicKey)
	auth, _ := bind.NewKeyedTransactorWithChainID(prvkey, big.NewInt(1337))
	_client, _ := ethclient.Dial("ws://127.0.0.1:8545")

	for _, order := range orders {
		SetupTxSender(auth, _client, _addr)
		id, _ := uuid.NewRandom()
		inst.AddOrder(auth, id,
			big.NewInt(int64(order.Price)),
			big.NewInt(int64(order.Amount)),
			order.Side,
		)
	}

	fmt.Println("Done, the number of fulfilled orders: ", _listenerInstance.NoFulfillOrders)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <filename> <is-offchain>\ne.g. go run . data/orders.json true")
	}

	_FILENAME_ := os.Args[1]
	orders, err := data.LoadOrders(_FILENAME_)
	if err != nil {
		fmt.Println("Cannot read orders from file")
		fmt.Println("Usage: go run . <filename> <is-offchain>\ne.g. go run . data/orders.json true")
		log.Fatal(err)
	}
	_RUN_OFFCHAIN_ := os.Args[2]

	if _RUN_OFFCHAIN_ == "true" {
		offchain(orders)
	} else if _RUN_OFFCHAIN_ == "false" {
		onchain(orders)
	} else {
		fmt.Println("Usage: go run . <filename> <is-offchain>\ne.g. go run . data/orders.json true")
	}
}

func SetupTxSender(auth *bind.TransactOpts, _client *ethclient.Client, _addr common.Address) {
	nonce, err := _client.PendingNonceAt(context.Background(), _addr)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := _client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(30_000_000)
}
