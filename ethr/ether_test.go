// Package ether provides tools for resolving the did:ether format, for
// resolving ethereum addresses as did documents.
// This resolver takes an ethereum address, checks for the current controller,
// looks at contract events, and builds a simple did document.
// Copyright 2021 Textile
// Copyright 2018 Consensys AG

// abigen --abi contracts/ethr-did-registry.json --pkg contracts --out contracts/ethr-did-registry.go
// abigen --sol contracts/ethr-did-registry.sol --pkg contracts --out contracts/ethr-did-registry.go
package ethr

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ockam-network/did"
)

func TestResolvePublicKey(t *testing.T) {
	conn, err := ethclient.Dial("https://rinkeby.infura.io/v3/ca30ba98858d47b4adfa22256f13e6f6")
	if err != nil {
		log.Fatalf("Whoops something went wrong: %s", err)
	}

	pubKey := "0x0279be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"
	pubID := fmt.Sprintf("did:ethr:%s", pubKey)

	parsed, err := did.Parse(pubID)
	if err != nil {
		t.Error(err)
	}

	chainID, err := conn.ChainID(context.Background())
	client := New(conn, chainID)
	doc, err := client.Resolve(pubID, parsed, client)
	if err != nil {
		t.Error(err)
	}
	j, err := json.Marshal(doc)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(j))
}
