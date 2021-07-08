package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// connect to websockets endpoint, MUST BE WEBSOCKETS TO TRIGGER THE CRASH
	client, err := ethclient.Dial("ws://localhost:8646")
	if err != nil {
		log.Fatal(err)
	}

	// retrieve the chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// generate a random private key to pass in to the transactor generation function
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// create transactor
	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	// transactor.Context = context.Background() <== NOTE: UNCOMMENT TO NOT CRASH

	// bind the smart contract
	deployedAt := common.HexToAddress("0x0000000000000000000000000000000000000000") // doesn't matter to trigger the crash
	transfer, err := NewTransfer(deployedAt, client)
	if err != nil {
		log.Fatal(err)
	}

	recipient := deployedAt // doesn't mater to trigger the crash

	// compose and sign a transaction to interact with the smart contract: attempt to send funds to recipient
	// NOTE: CRASH IS IN Send()
	tx, err := transfer.Send(transactor, recipient, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Transaction sent to the blockchain: %+v", tx)
}
