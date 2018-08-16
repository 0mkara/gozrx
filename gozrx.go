package gozrx

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ZeroEx instance
type ZeroEx struct {
	*TokenRegistry
	*Exchange
}

// NewZeroEx will create a new ZeroEx instance
func NewZeroEx(ec *ethclient.Client) (*ZeroEx, error) {
	tknrg, err := NewTokenRegistry(common.HexToAddress("0x4E9Aad8184DE8833365fEA970Cd9149372FDF1e6"), ec)
	if err != nil {
		log.Fatalf("could not create TokenRegistry: %v", err)
	}
	exchng, err := NewExchange(common.HexToAddress("0x1D16EF40FAC01cEC8aDAC2AC49427b9384192C05"), ec)
	if err != nil {
		log.Fatalf("could not create Exchange: %v", err)
	}
	return &ZeroEx{TokenRegistry: tknrg, Exchange: exchng}, nil
}
