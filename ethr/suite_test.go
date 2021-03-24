package ethr

// Basic imports
import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/multiformats/go-multibase"
	"github.com/ockam-network/did"
	"github.com/stretchr/testify/suite"

	"github.com/textileio/go-did-resolver/ethr/contracts"
	resolver "github.com/textileio/go-did-resolver/resolver"
)

var (
	blockchain *backends.SimulatedBackend
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type TestSuite struct {
	suite.Suite
	Identity   *ecdsa.PrivateKey
	Controller *ecdsa.PrivateKey
	Delegate1  *ecdsa.PrivateKey
	Delegate2  *ecdsa.PrivateKey
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *TestSuite) SetupSuite() {
	// Setup simulated block chain
	key, err := crypto.GenerateKey()
	if err != nil {
		suite.NoErrorf(err, "Invalid secret key")
	}
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{
		Balance: big.NewInt(100000000000),
	}

	// Specify default keys to use
	suite.Identity, err = crypto.HexToECDSA("0000000000000000000000000000000000000000000000000000000000000001")
	if err != nil {
		suite.NoErrorf(err, "Invalid secret key")
	}

	// Specify default keys to use
	suite.Controller, err = crypto.HexToECDSA("0000000000000000000000000000000000000000000000000000000000000002")
	if err != nil {
		suite.NoErrorf(err, "Invalid secret key")
	}

	// Specify default keys to use
	suite.Delegate1, err = crypto.HexToECDSA("0000000000000000000000000000000000000000000000000000000000000003")
	if err != nil {
		suite.NoErrorf(err, "Invalid secret key")
	}

	suite.Delegate2, err = crypto.HexToECDSA("0000000000000000000000000000000000000000000000000000000000000004")
	if err != nil {
		suite.NoErrorf(err, "Invalid secret key")
	}

	publicKey := suite.Identity.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	identityAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// The account to use for actually calling the contract
	alloc[identityAddress] = core.GenesisAccount{
		Balance: big.NewInt(100000000000),
	}

	blockchain = backends.NewSimulatedBackend(alloc, uint64(4712388))

	// Deploy contract
	addr, _, _, err := contracts.DeployEthereumDIDRegistry(auth, blockchain)
	suite.NoErrorf(err, "Failed to deploy the contract")

	// Override contract address
	ContractAddress = addr

	// blockchain.AdjustTime(time.Duration(time.Now().UnixNano()))
	// blockchain.Commit()

	// Commit all pending transactions
	blockchain.Commit()

}

func (suite *TestSuite) TestInterpretIdentifierKey() {
	publicKey := suite.Identity.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	publicKeyHex := hexutil.Encode(crypto.CompressPubkey(publicKeyECDSA))
	expected := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	observed, key, err := InterpretIdentifier(publicKeyHex)

	suite.NoError(err)
	suite.Equalf(observed.Hex(), expected, "incorrect address encoding")
	suite.Truef(key.Equal(publicKeyECDSA), "invalid key")
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *TestSuite) TestResolvesDocument() {
	publicKey := suite.Identity.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	addr := crypto.PubkeyToAddress(*publicKeyECDSA)
	id := fmt.Sprintf("did:ethr:%s", addr.Hex())
	identity := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	parsed, err := did.Parse(id)
	suite.NoError(err)

	r := New(blockchain, blockchain.Blockchain().Config().ChainID)
	observed, err := r.Resolve(id, parsed, r)
	suite.NoError(err)

	expected := &resolver.Document{
		Context: []string{
			"https://w3id.org/did/v1",
			"https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/lds-ecdsa-secp256k1-recovery2020-0.0.jsonld",
		},
		ID: id,
		VerificationMethod: []resolver.VerificationMethod{{
			ID:                  fmt.Sprintf("%s#controller", id),
			Type:                "EcdsaSecp256k1RecoveryMethod2020",
			Controller:          id,
			BlockchainAccountID: identity + "@eip155:1337",
		}},
		Authentication: []string{fmt.Sprintf("%s#controller", id)},
	}
	suite.Equal(expected, observed)
}

func (suite *TestSuite) TestResolvesPublicKey() {
	publicKey := suite.Identity.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	publicKeyHex := hexutil.Encode(crypto.CompressPubkey(publicKeyECDSA))
	id := fmt.Sprintf("did:ethr:%s", publicKeyHex)
	identity := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	parsed, err := did.Parse(id)
	suite.NoError(err)

	r := New(blockchain, blockchain.Blockchain().Config().ChainID)
	observed, err := r.Resolve(id, parsed, r)
	suite.NoError(err)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyMultibase, err := multibase.Encode(multibase.Base16, publicKeyBytes)

	expected := &resolver.Document{
		Context: []string{
			"https://w3id.org/did/v1",
			"https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/lds-ecdsa-secp256k1-recovery2020-0.0.jsonld",
		},
		ID: id,
		VerificationMethod: []resolver.VerificationMethod{
			{
				ID:                  fmt.Sprintf("%s#controller", id),
				Type:                "EcdsaSecp256k1RecoveryMethod2020",
				Controller:          id,
				BlockchainAccountID: identity + "@eip155:1337",
			},
			{
				ID:                 fmt.Sprintf("%s#controllerKey", id),
				Type:               "EcdsaSecp256k1VerificationKey2019",
				Controller:         id,
				PublicKeyMultibase: publicKeyMultibase,
			},
		},
		Authentication: []string{
			fmt.Sprintf("%s#controller", id),
			fmt.Sprintf("%s#controllerKey", id),
		},
	}
	suite.Equal(observed, expected)
}

func (suite *TestSuite) TestRejectInvalid() {
	id := "did:ethr:2nQtiQG6Cgm1GYTBaaKAgr76uY7iSexUkqX"

	parsed, err := did.Parse(id)
	suite.NoError(err)

	r := New(blockchain, blockchain.Blockchain().Config().ChainID)
	_, err = r.Resolve(id, parsed, r)
	suite.EqualErrorf(err, fmt.Sprintf("not a valid ethr did: %s", id), "Error message not equal")
}

// TODO: Write a test that validates which network we're prepared to resolve on?

func (suite *TestSuite) TestController() {
	// Setup
	registry, err := contracts.NewEthereumDIDRegistry(ContractAddress, blockchain)
	if err != nil {
		suite.NoError(err)
	}
	publicKey := suite.Identity.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	controllerKey := suite.Controller.Public()
	controllerKeyECDSA, _ := controllerKey.(*ecdsa.PublicKey)
	controller := crypto.PubkeyToAddress(*controllerKeyECDSA)
	suite.NoError(err)

	// We need to setup some values for any calls that will update state
	nonce, err := blockchain.PendingNonceAt(context.Background(), fromAddress)
	suite.NoError(err)
	// This will generally return reasonable, fake suggestions, but the process
	// is the same for the "real world"
	gasPrice, err := blockchain.SuggestGasPrice(context.Background())
	suite.NoError(err)

	// We'll setup our new transactor to be our original identity
	auth := bind.NewKeyedTransactor(suite.Identity)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// This is the only real mutation we need to perform here
	_, err = registry.ChangeOwner(auth, fromAddress, controller)
	suite.NoError(err)

	// Commit all pending transactions
	blockchain.Commit()

	// Let's just do a low-level check to make sure this worked
	addr, err := registry.IdentityOwner(nil, fromAddress)
	suite.NoError(err)

	suite.Equal(controller.Hex(), addr.Hex(), "controller should be the new owner")

	r := New(blockchain, blockchain.Blockchain().Config().ChainID)

	suite.Run("ControllerChanged", func() {
		id := fmt.Sprintf("did:ethr:%s", fromAddress.Hex())
		parsed, err := did.Parse(id)
		suite.NoError(err)

		observed, err := r.Resolve(id, parsed, r)
		suite.NoError(err)

		expected := &resolver.Document{
			Context: []string{
				"https://w3id.org/did/v1",
				"https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/lds-ecdsa-secp256k1-recovery2020-0.0.jsonld",
			},
			ID: id,
			VerificationMethod: []resolver.VerificationMethod{
				{
					ID:                  fmt.Sprintf("%s#controller", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", controller.Hex()),
				},
			},
			Authentication: []string{fmt.Sprintf("%s#controller", id)},
		}
		suite.Equal(observed, expected)
	})

	suite.Run("InvalidatePublicKey", func() {
		publicKeyHex := hexutil.Encode(crypto.CompressPubkey(publicKeyECDSA))
		id := fmt.Sprintf("did:ethr:%s", publicKeyHex)
		parsed, err := did.Parse(id)
		suite.NoError(err)

		observed, err := r.Resolve(id, parsed, r)
		suite.NoError(err)

		expected := &resolver.Document{
			Context: []string{
				"https://w3id.org/did/v1",
				"https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/lds-ecdsa-secp256k1-recovery2020-0.0.jsonld",
			},
			ID: id,
			VerificationMethod: []resolver.VerificationMethod{
				{
					ID:                  fmt.Sprintf("%s#controller", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", controller.Hex()),
				},
			},
			Authentication: []string{fmt.Sprintf("%s#controller", id)},
		}
		suite.Equal(observed, expected)
	})
}

func (suite *TestSuite) TestDelegates() {
	// Setup
	registry, err := contracts.NewEthereumDIDRegistry(ContractAddress, blockchain)
	if err != nil {
		suite.NoError(err)
	}
	publicKey := suite.Identity.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// We need to setup some values for any calls that will update state
	nonce, err := blockchain.PendingNonceAt(context.Background(), fromAddress)
	suite.NoError(err)
	// This will generally return reasonable, fake suggestions, but the process
	// is the same for the "real world"
	gasPrice, err := blockchain.SuggestGasPrice(context.Background())
	suite.NoError(err)

	// We'll setup our new transactor to be our original identity
	auth := bind.NewKeyedTransactor(suite.Identity)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	r := New(blockchain, blockchain.Blockchain().Config().ChainID)

	delegate1Key := suite.Delegate1.Public()
	delegate1ECDSA, _ := delegate1Key.(*ecdsa.PublicKey)
	delegate1Address := crypto.PubkeyToAddress(*delegate1ECDSA)

	delegate2Key := suite.Delegate2.Public()
	delegate2ECDSA, _ := delegate2Key.(*ecdsa.PublicKey)
	delegate2Address := crypto.PubkeyToAddress(*delegate2ECDSA)

	suite.Run("WithSigningDelegate", func() {
		// TODO: I don't think we should have to ajust to unix time for these validTo values
		validity := big.NewInt(0).Add(big.NewInt(time.Now().Unix()), big.NewInt(1000))
		_, err = registry.AddDelegate(auth, fromAddress, VeriKey, delegate1Address, validity)
		suite.NoError(err)

		// Commit all pending transactions
		blockchain.Commit()

		id := fmt.Sprintf("did:ethr:dev:%s", fromAddress.Hex())
		parsed, err := did.Parse(id)
		suite.NoError(err)

		observed, err := r.Resolve(id, parsed, r)
		suite.NoError(err)

		expected := &resolver.Document{
			Context: []string{
				"https://w3id.org/did/v1",
				"https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/lds-ecdsa-secp256k1-recovery2020-0.0.jsonld",
			},
			ID: id,
			VerificationMethod: []resolver.VerificationMethod{
				{
					ID:                  fmt.Sprintf("%s#controller", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", fromAddress.Hex()),
				},
				{
					ID:                  fmt.Sprintf("%s#delegate-1", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", delegate1Address.Hex()),
				},
			},
			Authentication: []string{fmt.Sprintf("%s#controller", id)},
		}
		suite.Equal(expected, observed)
	})

	suite.Run("WithAuthDelegate", func() {
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
		validity := big.NewInt(0).Add(big.NewInt(time.Now().Unix()), big.NewInt(-30))
		_, err = registry.AddDelegate(auth, fromAddress, SigAuth, delegate2Address, validity)
		suite.NoError(err)

		// Commit all pending transactions
		blockchain.Commit()

		id := fmt.Sprintf("did:ethr:%s", fromAddress.Hex())
		parsed, err := did.Parse(id)
		suite.NoError(err)

		observed, err := r.Resolve(id, parsed, r)
		suite.NoError(err)

		expected := &resolver.Document{
			Context: []string{
				"https://w3id.org/did/v1",
				"https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/lds-ecdsa-secp256k1-recovery2020-0.0.jsonld",
			},
			ID: id,
			VerificationMethod: []resolver.VerificationMethod{
				{
					ID:                  fmt.Sprintf("%s#controller", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", fromAddress.Hex()),
				},
				{
					ID:                  fmt.Sprintf("%s#delegate-1", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", delegate1Address.Hex()),
				},
				{
					ID:                  fmt.Sprintf("%s#delegate-2", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", delegate2Address.Hex()),
				},
			},
			Authentication: []string{
				fmt.Sprintf("%s#controller", id),
				fmt.Sprintf("%s#delegate-2", id),
			},
		}
		suite.Equal(expected, observed)
	})

	suite.Run("ExpiresAutomatically", func() {
		time.Sleep(time.Second)
		id := fmt.Sprintf("did:ethr:dev:%s", fromAddress.Hex())
		parsed, err := did.Parse(id)
		suite.NoError(err)

		observed, err := r.Resolve(id, parsed, r)
		suite.NoError(err)

		expected := &resolver.Document{
			Context: []string{
				"https://w3id.org/did/v1",
				"https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/lds-ecdsa-secp256k1-recovery2020-0.0.jsonld",
			},
			ID: id,
			VerificationMethod: []resolver.VerificationMethod{
				{
					ID:                  fmt.Sprintf("%s#controller", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", fromAddress.Hex()),
				},
				{
					ID:                  fmt.Sprintf("%s#delegate-1", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", delegate1Address.Hex()),
				},
			},
			Authentication: []string{fmt.Sprintf("%s#controller", id)},
		}
		suite.Equal(expected, observed)
	})

	suite.Run("RevokesDelegate", func() {
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
		_, err = registry.RevokeDelegate(auth, fromAddress, VeriKey, delegate1Address)
		suite.NoError(err)

		// Commit all pending transactions
		blockchain.Commit()

		time.Sleep(time.Second * 2)

		id := fmt.Sprintf("did:ethr:dev:%s", fromAddress.Hex())
		parsed, err := did.Parse(id)
		suite.NoError(err)

		observed, err := r.Resolve(id, parsed, r)
		suite.NoError(err)

		expected := &resolver.Document{
			Context: []string{
				"https://w3id.org/did/v1",
				"https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/lds-ecdsa-secp256k1-recovery2020-0.0.jsonld",
			},
			ID: id,
			VerificationMethod: []resolver.VerificationMethod{
				{
					ID:                  fmt.Sprintf("%s#controller", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", fromAddress.Hex()),
				},
			},
			Authentication: []string{fmt.Sprintf("%s#controller", id)},
		}
		suite.Equal(expected, observed)
	})

	suite.Run("ReAddDelegate", func() {
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
		validity := big.NewInt(0).Add(big.NewInt(time.Now().Unix()), big.NewInt(86400))
		_, err = registry.AddDelegate(auth, fromAddress, SigAuth, delegate2Address, validity)
		suite.NoError(err)

		// Commit all pending transactions
		blockchain.Commit()

		time.Sleep(time.Second * 2)

		id := fmt.Sprintf("did:ethr:dev:%s", fromAddress.Hex())
		parsed, err := did.Parse(id)
		suite.NoError(err)

		observed, err := r.Resolve(id, parsed, r)
		suite.NoError(err)

		expected := &resolver.Document{
			Context: []string{
				"https://w3id.org/did/v1",
				"https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/lds-ecdsa-secp256k1-recovery2020-0.0.jsonld",
			},
			ID: id,
			VerificationMethod: []resolver.VerificationMethod{
				{
					ID:                  fmt.Sprintf("%s#controller", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", fromAddress.Hex()),
				},
				{
					ID:                  fmt.Sprintf("%s#delegate-4", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", delegate2Address.Hex()),
				},
			},

			Authentication: []string{
				fmt.Sprintf("%s#controller", id),
				fmt.Sprintf("%s#delegate-4", id),
			},
		}
		suite.Equal(expected, observed)
	})
}

func (suite *TestSuite) TestAttributes() {
	// Setup
	registry, err := contracts.NewEthereumDIDRegistry(ContractAddress, blockchain)
	if err != nil {
		suite.NoError(err)
	}
	publicKey := suite.Identity.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// We need to setup some values for any calls that will update state
	nonce, err := blockchain.PendingNonceAt(context.Background(), fromAddress)
	suite.NoError(err)
	// This will generally return reasonable, fake suggestions, but the process
	// is the same for the "real world"
	gasPrice, err := blockchain.SuggestGasPrice(context.Background())
	suite.NoError(err)

	// We'll setup our new transactor to be our original identity
	auth := bind.NewKeyedTransactor(suite.Identity)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	r := New(blockchain, blockchain.Blockchain().Config().ChainID)

	// delegate1Key := suite.Delegate1.Public()
	// delegate1ECDSA, _ := delegate1Key.(*ecdsa.PublicKey)
	// delegate1Address := crypto.PubkeyToAddress(*delegate1ECDSA)

	// delegate2Key := suite.Delegate2.Public()
	// delegate2ECDSA, _ := delegate2Key.(*ecdsa.PublicKey)
	// delegate2Address := crypto.PubkeyToAddress(*delegate2ECDSA)

	suite.Run("AddPublicKey", func() {
		// TODO: I don't think we should have to ajust to unix time for these validTo values
		validity := big.NewInt(0).Add(big.NewInt(time.Now().Unix()), big.NewInt(50))
		var name [32]byte
		copy(name[:], "did/pub/Secp256k1/veriKey")
		_, err = registry.SetAttribute(auth, fromAddress, name, []byte("0x02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71"), validity)
		suite.NoError(err)

		// Commit all pending transactions
		blockchain.Commit()

		id := fmt.Sprintf("did:ethr:dev:%s", fromAddress.Hex())
		parsed, err := did.Parse(id)
		suite.NoError(err)

		observed, err := r.Resolve(id, parsed, r)
		suite.NoError(err)

		expected := &resolver.Document{
			Context: []string{
				"https://w3id.org/did/v1",
				"https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/lds-ecdsa-secp256k1-recovery2020-0.0.jsonld",
			},
			ID: id,
			VerificationMethod: []resolver.VerificationMethod{
				{
					ID:                  fmt.Sprintf("%s#controller", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", fromAddress.Hex()),
				},
				{
					ID:                 fmt.Sprintf("%s#delegate-1", id),
					Type:               "EcdsaSecp256k1VerificationKey2019",
					Controller:         id,
					PublicKeyMultibase: "f02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71",
				},
			},
			Authentication: []string{fmt.Sprintf("%s#controller", id)},
		}
		suite.Equal(expected, observed)
	})

	suite.Run("ResolveEd25519VerificationKey", func() {
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
		// TODO: I don't think we should have to ajust to unix time for these validTo values
		validity := big.NewInt(0).Add(big.NewInt(time.Now().Unix()), big.NewInt(50))
		var name [32]byte
		copy(name[:], "did/pub/Ed25519/veriKey/base64")
		_, err = registry.SetAttribute(auth, fromAddress, name, []byte("0x02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71"), validity)
		suite.NoError(err)

		// Commit all pending transactions
		blockchain.Commit()

		id := fmt.Sprintf("did:ethr:dev:%s", fromAddress.Hex())
		parsed, err := did.Parse(id)
		suite.NoError(err)

		observed, err := r.Resolve(id, parsed, r)
		suite.NoError(err)

		expected := &resolver.Document{
			Context: []string{
				"https://w3id.org/did/v1",
				"https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/lds-ecdsa-secp256k1-recovery2020-0.0.jsonld",
			},
			ID: id,
			VerificationMethod: []resolver.VerificationMethod{
				{
					ID:                  fmt.Sprintf("%s#controller", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", fromAddress.Hex()),
				},
				{
					ID:                 fmt.Sprintf("%s#delegate-1", id),
					Type:               "EcdsaSecp256k1VerificationKey2019",
					Controller:         id,
					PublicKeyMultibase: "f02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71",
				},
				{
					ID:                 fmt.Sprintf("%s#delegate-2", id),
					Type:               "Ed25519VerificationKey2018",
					Controller:         id,
					PublicKeyMultibase: "mArl8MN52fwhM4wgBaO4pMFO6M7I11xFqMmPSnxRQk2tx",
				},
			},
			Authentication: []string{fmt.Sprintf("%s#controller", id)},
		}
		suite.Equal(expected, observed)
	})

	suite.Run("ResolveRSAVerificationKey2018", func() {
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
		// TODO: I don't think we should have to ajust to unix time for these validTo values
		validity := big.NewInt(0).Add(big.NewInt(time.Now().Unix()), big.NewInt(50))
		var name [32]byte
		copy(name[:], "did/pub/RSA/veriKey/pem")
		_, err = registry.SetAttribute(auth, fromAddress, name, []byte("-----BEGIN PUBLIC KEY...END PUBLIC KEY-----\r\n"), validity)
		suite.NoError(err)

		// Commit all pending transactions
		blockchain.Commit()

		id := fmt.Sprintf("did:ethr:dev:%s", fromAddress.Hex())
		parsed, err := did.Parse(id)
		suite.NoError(err)

		observed, err := r.Resolve(id, parsed, r)
		suite.NoError(err)

		expected := &resolver.Document{
			Context: []string{
				"https://w3id.org/did/v1",
				"https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/lds-ecdsa-secp256k1-recovery2020-0.0.jsonld",
			},
			ID: id,
			VerificationMethod: []resolver.VerificationMethod{
				{
					ID:                  fmt.Sprintf("%s#controller", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", fromAddress.Hex()),
				},
				{
					ID:                 fmt.Sprintf("%s#delegate-1", id),
					Type:               "EcdsaSecp256k1VerificationKey2019",
					Controller:         id,
					PublicKeyMultibase: "f02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71",
				},
				{
					ID:                 fmt.Sprintf("%s#delegate-2", id),
					Type:               "Ed25519VerificationKey2018",
					Controller:         id,
					PublicKeyMultibase: "mArl8MN52fwhM4wgBaO4pMFO6M7I11xFqMmPSnxRQk2tx",
				},
				{
					ID:           fmt.Sprintf("%s#delegate-3", id),
					Type:         "RSAVerificationKey2018",
					Controller:   id,
					PublicKeyPem: "-----BEGIN PUBLIC KEY...END PUBLIC KEY-----\r\n",
				},
			},
			Authentication: []string{fmt.Sprintf("%s#controller", id)},
		}
		suite.Equal(expected, observed)
	})

	// TODO: Write test for resolving X25519KeyAgreementKey2019

	suite.Run("Add Service Endpoint", func() {
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
		// TODO: I don't think we should have to ajust to unix time for these validTo values
		validity := big.NewInt(0).Add(big.NewInt(time.Now().Unix()), big.NewInt(50))
		var name [32]byte
		copy(name[:], "did/svc/HubService")
		encodedEndpoint := []byte(hex.EncodeToString([]byte("https://hub.textile.io")))
		_, err = registry.SetAttribute(auth, fromAddress, name, encodedEndpoint, validity)
		suite.NoError(err)

		// Commit all pending transactions
		blockchain.Commit()

		id := fmt.Sprintf("did:ethr:dev:%s", fromAddress.Hex())
		parsed, err := did.Parse(id)
		suite.NoError(err)

		observed, err := r.Resolve(id, parsed, r)
		suite.NoError(err)

		expected := &resolver.Document{
			Context: []string{
				"https://w3id.org/did/v1",
				"https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/lds-ecdsa-secp256k1-recovery2020-0.0.jsonld",
			},
			ID: id,
			VerificationMethod: []resolver.VerificationMethod{
				{
					ID:                  fmt.Sprintf("%s#controller", id),
					Type:                "EcdsaSecp256k1RecoveryMethod2020",
					Controller:          id,
					BlockchainAccountID: fmt.Sprintf("%s@eip155:1337", fromAddress.Hex()),
				},
				{
					ID:                 fmt.Sprintf("%s#delegate-1", id),
					Type:               "EcdsaSecp256k1VerificationKey2019",
					Controller:         id,
					PublicKeyMultibase: "f02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71",
				},
				{
					ID:                 fmt.Sprintf("%s#delegate-2", id),
					Type:               "Ed25519VerificationKey2018",
					Controller:         id,
					PublicKeyMultibase: "mArl8MN52fwhM4wgBaO4pMFO6M7I11xFqMmPSnxRQk2tx",
				},
				{
					ID:           fmt.Sprintf("%s#delegate-3", id),
					Type:         "RSAVerificationKey2018",
					Controller:   id,
					PublicKeyPem: "-----BEGIN PUBLIC KEY...END PUBLIC KEY-----\r\n",
				},
			},
			Authentication: []string{fmt.Sprintf("%s#controller", id)},
			Service: []resolver.ServiceEndpoint{
				{
					ID:              fmt.Sprintf("%s#service-1", id),
					Type:            "HubService",
					ServiceEndpoint: "https://hub.textile.io",
				},
			},
		}
		suite.Equal(expected, observed)
	})
}

// TODO: RevokeAttributes

// TODO: Multiple events in a single block

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
