package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/NguyenHiu/lob/contracts"
	"github.com/NguyenHiu/lob/contracts/generated/lOB"
	"github.com/NguyenHiu/lob/data"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
)

func main() {
	deployerPrvkeyHex := "abf82ff96b463e9d82b83cb9bb450fe87e6166d4db6d7021d0c71d7e960d5abe"
	_, inst := contracts.DeployContracts(deployerPrvkeyHex)

	go contracts.ListenContract(inst)

	traderPrvkeyHex := "dcb7118c9946a39cd40b661e0d368e4afcc3cc48d21aa750d8164ca2e44961c4"
	prvkey, _ := crypto.HexToECDSA(traderPrvkeyHex)
	_addr := crypto.PubkeyToAddress(prvkey.PublicKey)
	auth, _ := bind.NewKeyedTransactorWithChainID(prvkey, big.NewInt(1338))
	_client, _ := ethclient.Dial("ws://127.0.0.1:8546")

	orders, _ := data.LoadOrders("../lightning-exchange/data/orders.json")

	for _, order := range orders {
		SetupTxSender(auth, _client, _addr)
		id, _ := uuid.NewRandom()
		inst.AddOrder(auth, lOB.LOBOrder{
			OrderID: id,
			Price:   big.NewInt(int64(order.Price)),
			Amount:  big.NewInt(int64(order.Amount)),
			Side:    order.Side,
		})
	}

	SetupTxSender(auth, _client, _addr)
	inst.AddOrder(auth, lOB.LOBOrder{
		OrderID: [16]byte{},
		Price:   new(big.Int),
		Amount:  new(big.Int),
		Side:    true,
	})

	fmt.Println("Done, calculating fee...")

	_prvkey, _ := crypto.HexToECDSA(traderPrvkeyHex)
	totalGas := CalculateTotalUsedGas(crypto.PubkeyToAddress(_prvkey.PublicKey))

	fmt.Println("total gas: ", totalGas)

}

func CalculateTotalUsedGas(addr common.Address) int {
	totalGas := 0
	_client, _ := ethclient.Dial("ws://127.0.0.1:8546")

	for i := new(big.Int); ; i.Add(i, big.NewInt(1)) {
		block, err := _client.BlockByNumber(context.Background(), i)
		if err != nil {
			fmt.Println("err: ", err)
			return totalGas
		}
		for _, tx := range block.Transactions() {
			if from, err := types.Sender(types.NewLondonSigner(big.NewInt(1338)), tx); err == nil {
				if from.Cmp(addr) == 0 {
					receipt, err := _client.TransactionReceipt(context.Background(), tx.Hash())
					if err != nil {
						log.Fatal(err)
					}
					totalGas += int(receipt.GasUsed)
				}
			}
		}

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
	auth.GasLimit = uint64(6_000_000)
}
