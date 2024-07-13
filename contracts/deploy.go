package contracts

import (
	"math/big"

	"github.com/NguyenHiu/lob/contracts/generated/lOB"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployContracts(privateKeyHex string) (common.Address, *lOB.LOB) {
	prvkey, _ := crypto.HexToECDSA(privateKeyHex)

	chainID := big.NewInt(1338)
	auth, _ := bind.NewKeyedTransactorWithChainID(prvkey, chainID)

	_client, _ := ethclient.Dial("ws://127.0.0.1:8546")

	contractAddr, _, contractInst, _ := lOB.DeployLOB(auth, _client)

	return contractAddr, contractInst
}
