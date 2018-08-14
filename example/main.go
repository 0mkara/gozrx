package main

import (
	"fmt"
	"log"

	"github.com/0mkara/gozrx"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("/Users/0mkar/Karma/Ethereum/rinkeby/geth.ipc")
	if err != nil {
		log.Fatalf("could not create ipc client: %v", err)
	}
	zerox, err := gozrx.NewZeroEx(conn)
	if err != nil {
		log.Fatalf("could not create 0x instance: %v", err)
	}
	tokenAddrs, err := zerox.TokenRegistry.GetTokenAddresses(nil)
	if err != nil {
		log.Fatalf("could not get token addresses: %v", err)
	}
	for _, addr := range tokenAddrs {
		fmt.Println("Token addresse:", addr.Hex())
	}
}
