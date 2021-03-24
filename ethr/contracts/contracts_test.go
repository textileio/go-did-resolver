package contracts

import (
	"crypto/ecdsa"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestDeployDIDRegistryAndQuery(t *testing.T) {

	// Setup simulated block chain
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, uint64(2000000))

	//Deploy contract
	address, _, contract, err := DeployEthereumDIDRegistry(auth, blockchain)
	// commit all pending transactions
	blockchain.Commit()

	if err != nil {
		t.Fatalf("Failed to deploy the contract: %v", err)
	}

	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address. Received empty address byte array instead")
	}

	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("PublicKey is not of type *ecdsa.PublicKey")
	}

	addr, err := contract.IdentityOwner(nil, crypto.PubkeyToAddress(*publicKeyECDSA))
	if err != nil {
		t.Error(err)
	}

	if !common.IsHexAddress(addr.Hex()) {
		t.Fatalf("Not a valid Ethereum address: %s", addr.Hex())
	}
}

func TestLiveAndQuery(t *testing.T) {
	conn, err := ethclient.Dial("https://rinkeby.infura.io/v3/ca30ba98858d47b4adfa22256f13e6f6")
	if err != nil {
		t.Error(err)
	}

	contractAddr := "0xdCa7EF03e98e0DC2B855bE647C39ABe984fcF21B"

	contract, err := NewEthereumDIDRegistry(common.HexToAddress(contractAddr), conn)
	if err != nil {
		t.Error(err)
	}

	carsonAddr := "0x53e448B7A37bE85f7D6A63d83eEcBb3a5c350471"
	opts := &bind.CallOpts{
		From: common.HexToAddress(carsonAddr),
	}

	addr, err := contract.IdentityOwner(opts, common.HexToAddress(carsonAddr))
	if err != nil {
		t.Error(err)
	}
	if !common.IsHexAddress(addr.Hex()) {
		t.Fatalf("Not a valid Ethereum address: %s", addr.Hex())
	}

	if addr.Hex() != carsonAddr {
		t.Fatalf("Unexpected owner address: %s", addr.Hex())
	}
}
