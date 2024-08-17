# Limit Order Book 

This project simulates an exchange based on the Limit Order Book (LOB) in both off-chain and on-chain.

These models are used to analyze attributes of exchange based on LOB, such as *the number of matched orders*, *the number of orders per second*, *total gas fee*, *total profit*,... Currently, I only get *the number of fulfilled orders* (`NoFulfillOrders`), you can add what you want in the `lob` module (offchain) or `listener` module (onchain).

## How to run?
```
$ go run . <filename> <is-offchain>
```
- `filename` is the path to order file 
- `is-offchain` to determine the lob you wanna use to run

e.g.
```
$ go run . ./data/sample_orders.json false
```

## Missing dependencies?
### 1. Re-generate LOB onchain contract
You can use the `contracts/generate_docker.sh` to generate the `LOB.go` file automatically from `LOB.sol` contract. 

*Note that: you must have the [`ethereum/solc:0.8.24`](https://docs.soliditylang.org/en/latest/installing-solidity.html#docker) docker installed in your machine*

### 2. Openzeppelin contracts
Don't forget the openzeppelin contracts in `contracts` folder, you can pull the repo manually from [`github.com/openzeppelin-contracts`](https://github.com/OpenZeppelin/openzeppelin-contracts).