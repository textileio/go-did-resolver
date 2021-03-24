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
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/iancoleman/orderedmap"
	"github.com/multiformats/go-multibase"
	"github.com/ockam-network/did"

	contracts "github.com/textileio/go-did-resolver/ethr/contracts"
	resolver "github.com/textileio/go-did-resolver/resolver"
)

var (
	// ContractAddress is the address for the Ethereum DID Registry smart contract
	ContractAddress   = common.HexToAddress("0xdca7ef03e98e0dc2b855be647c39abe984fcf21b")
	identifierMatcher = regexp.MustCompile(`^(.*)?(0x[0-9a-fA-F]{40}|0x[0-9a-fA-F]{66})$`)
	attributeMatcher  = regexp.MustCompile(`^did\/(pub|svc)\/(\w+)(\/(\w+))?(\/(\w+))?$`)
	nullAddress       = "0x0000000000000000000000000000000000000000"
	trimmer           = regexp.MustCompile(`\0+$`)
)

var knownNetworks = map[string]*big.Int{
	"dev":     big.NewInt(1337),
	"mainnet": big.NewInt(1),
	"ropsten": big.NewInt(3),
	"rinkeby": big.NewInt(4),
	"goerli":  big.NewInt(5),
	"kovan":   big.NewInt(42),
}

var legacyAttrTypes = map[string]string{
	"sigAuth": "SignatureAuthentication2018",
	"veriKey": "VerificationKey2018",
	"enc":     "KeyAgreementKey2019",
}

var legacyAlgoMap = map[string]string{
	"Secp256k1VerificationKey2018":         "EcdsaSecp256k1VerificationKey2019",
	"Ed25519SignatureAuthentication2018":   "Ed25519VerificationKey2018",
	"Secp256k1SignatureAuthentication2018": "EcdsaSecp256k1VerificationKey2019",
	"RSAVerificationKey2018":               "RSAVerificationKey2018",
	"Ed25519VerificationKey2018":           "Ed25519VerificationKey2018",
	"X25519KeyAgreementKey2019":            "X25519KeyAgreementKey2019",
}

var VeriKey = [32]byte{118, 101, 114, 105, 75, 101, 121}
var SigAuth = [32]byte{115, 105, 103, 65, 117, 116, 104}

// Resolver defines a basic key resolver, conforming the to Resolver interface.
type Resolver struct {
	client  bind.ContractBackend
	chainID *big.Int
}

// New creates and returns a new key Resolver.
func New(client bind.ContractBackend, chainID *big.Int) *Resolver {
	return &Resolver{
		client,
		chainID,
	}
}

// Method returns the method that this resolver is capable of resolving.
func (r *Resolver) Method() string {
	return "ethr"
}

// Resolve is the primary resolution method for this resolver.
func (r *Resolver) Resolve(did string, parsed *did.DID, res resolver.Resolver) (*resolver.Document, error) {
	if parsed.Method != r.Method() {
		return nil, fmt.Errorf("unknown did method: '%s'", parsed.Method)
	}
	fullID := identifierMatcher.FindStringSubmatch(parsed.ID)
	if len(fullID) < 1 {
		return nil, fmt.Errorf("not a valid ethr did: %s", did)
	}
	networkID := "dev"
	var chainID *big.Int
	id := fullID[2]
	if fullID[1] != "" {
		if !strings.HasSuffix(fullID[1], ":") {
			return nil, fmt.Errorf("not a valid ethr did: %s", did)
		}
		networkID = fullID[1][:len(fullID[1])-1]
	}
	var ok bool
	// First check to see if this networkID matches a known plaintext network name
	chainID, ok = knownNetworks[networkID]
	if !ok {
		// If it doesn't, we assume it is a hex chainID code and work with that
		chainID = new(big.Int)
		chainID.SetString(networkID, 16)
	}
	if chainID.Cmp(r.chainID) != 0 {
		// TODO: Rather than error, we'll want to support multiple chains
		return nil, fmt.Errorf("non-matching chain ids")
	}

	controller, history, publicKey, err := changeLog(r.client, id)
	if err != nil {
		return nil, err
	}
	return wrapDocument(did, controller, publicKey, history, r.chainID)
}

var _ resolver.Resolver = (*Resolver)(nil)

// EthereumDIDRegistryDIDEventUnion represents any of the available DIDRegistry
// event types: *OwnerChanged, *DelegateChanged, *AttributeChange
type EthereumDIDRegistryDIDEventUnion struct {
	Identity       common.Address
	Owner          common.Address
	DelegateType   [32]byte
	Delegate       common.Address
	ValidTo        *big.Int
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
	Name           [32]byte
	Value          []byte
	Type           string
}

func wrapDocument(did string, controller *common.Address, publicKey *ecdsa.PublicKey, history []*EthereumDIDRegistryDIDEventUnion, chainID *big.Int) (*resolver.Document, error) {
	now := big.NewInt(time.Now().Unix())
	verificationMethod := []resolver.VerificationMethod{{
		ID:                  fmt.Sprintf("%s#controller", did),
		Type:                "EcdsaSecp256k1RecoveryMethod2020",
		Controller:          did,
		BlockchainAccountID: fmt.Sprintf("%s@eip155:%d", controller.Hex(), chainID),
	}}

	authentication := []string{fmt.Sprintf("%s#controller", did)}

	if publicKey != nil {
		// NOTE: This is the uncompressed key bytes
		publicKeyBytes := crypto.FromECDSAPub(publicKey)
		publicKeyMultibase, err := multibase.Encode(multibase.Base16, publicKeyBytes)
		if err != nil {
			return nil, err
		}
		verificationMethod = append(verificationMethod, resolver.VerificationMethod{
			ID:                 fmt.Sprintf("%s#controllerKey", did),
			Type:               "EcdsaSecp256k1VerificationKey2019",
			Controller:         did,
			PublicKeyMultibase: publicKeyMultibase,
		})

		authentication = append(authentication, fmt.Sprintf("%s#controllerKey", did))
	}

	deactivated := false
	auths := orderedmap.New()
	pks := orderedmap.New()
	services := orderedmap.New()
	delegateCount := 0
	serviceCount := 0

	for _, event := range history {
		validTo := event.ValidTo
		// fmt.Println(validTo.Uint64(), now.Uint64())
		if (validTo != nil) && validTo.Cmp(now) >= 0 {
			switch event.Type {
			case "DIDDelegateChanged":
				delegateCount++
				delegateType := trimmer.ReplaceAllString(string(event.DelegateType[:]), "")
				eventIndex := fmt.Sprintf("%s-%s-%s", event.Type, delegateType, event.Delegate.Hex())
				switch event.DelegateType {
				case SigAuth:
					auths.Set(eventIndex, fmt.Sprintf("%s#delegate-%d", did, delegateCount))
					fallthrough
				case VeriKey:
					pks.Set(eventIndex, resolver.VerificationMethod{
						ID:                  fmt.Sprintf("%s#delegate-%d", did, delegateCount),
						Type:                "EcdsaSecp256k1RecoveryMethod2020",
						Controller:          did,
						BlockchainAccountID: fmt.Sprintf("%s@eip155:%d", event.Delegate, chainID),
					})
				}
			case "DIDAttributeChanged":
				attributeName := trimmer.ReplaceAllString(string(event.Name[:]), "")
				eventIndex := fmt.Sprintf("%s-%s-%s", event.Type, attributeName, event.Value[:])
				match := attributeMatcher.FindStringSubmatch(attributeName)
				if len(match) > 0 {
					section := match[1]
					algorithm := match[2]
					kind := match[4]
					legacy, ok := legacyAttrTypes[match[4]]
					if ok {
						kind = legacy
					}
					encoding := match[6]
					keyOrService := string(event.Value[:])
					if strings.HasPrefix(keyOrService, "0x") {
						keyOrService = keyOrService[2:]
					}
					switch section {
					case "pub":
						delegateCount++
						pk := resolver.VerificationMethod{
							ID:         fmt.Sprintf("%s#delegate-%d", did, delegateCount),
							Controller: did,
							Type:       algorithm,
						}
						legacy, ok := legacyAlgoMap[algorithm+kind]
						if ok {
							pk.Type = legacy
						}
						switch encoding {
						case "", "hex":
							keyOrServiceBytes, err := hex.DecodeString(keyOrService)
							if err != nil {
								return nil, err
							}
							publicKeyMultibase, err := multibase.Encode(multibase.Base16, keyOrServiceBytes)
							if err != nil {
								return nil, err
							}
							pk.PublicKeyMultibase = publicKeyMultibase
						case "base64":
							keyOrServiceBytes, err := hex.DecodeString(keyOrService)
							if err != nil {
								return nil, err
							}
							publicKeyMultibase, err := multibase.Encode(multibase.Base64, keyOrServiceBytes)
							if err != nil {
								return nil, err
							}
							pk.PublicKeyMultibase = publicKeyMultibase
						case "base58":
							keyOrServiceBytes, err := hex.DecodeString(keyOrService)
							if err != nil {
								return nil, err
							}
							publicKeyMultibase, err := multibase.Encode(multibase.Base58BTC, keyOrServiceBytes)
							if err != nil {
								return nil, err
							}
							pk.PublicKeyMultibase = publicKeyMultibase
						case "pem":
							pk.PublicKeyPem = keyOrService
						default:
							// TODO: Should we have a default case?
						}
						pks.Set(eventIndex, pk)
						if kind == "sigAuth" {
							auths.Set(eventIndex, pk.ID)
						}
					case "svc":
						serviceCount++
						keyOrServiceBytes, err := hex.DecodeString(keyOrService)
						if err != nil {
							return nil, err
						}
						services.Set(eventIndex, resolver.ServiceEndpoint{
							ID:              fmt.Sprintf("%s#service-%d", did, serviceCount),
							Type:            algorithm,
							ServiceEndpoint: string(keyOrServiceBytes),
						})
					}
				}
			}
		} else {
			var eventIndex string
			if event.Type == "DIDDelegateChanged" {
				delegateCount++
				delegateType := trimmer.ReplaceAllString(string(event.DelegateType[:]), "")
				eventIndex = fmt.Sprintf("%s-%s-%s", event.Type, delegateType, event.Delegate.Hex())
			} else if event.Type == "DIDAttributeChanged" {
				attributeName := trimmer.ReplaceAllString(string(event.Name[:]), "")
				eventIndex = fmt.Sprintf("%s-%s-%s", event.Type, attributeName, event.Value)
				if strings.HasPrefix(string(event.Name[:]), "did/pub/") {
					delegateCount++
				} else if strings.HasPrefix(string(event.Name[:]), "did/svc/") {
					serviceCount++
				}
			}
			auths.Delete(eventIndex)
			pks.Delete(eventIndex)
			services.Delete(eventIndex)
			if event.Type == "DIDOwnerChanged" {
				if event.Owner.Hex() == (common.Address{}).Hex() {
					deactivated = true
					// TODO: Should we still build up the document?
					break
				}
			}
		}

	}

	var doc = &resolver.Document{
		ID: did,
		Context: []string{
			"https://w3id.org/did/v1",
			"https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/lds-ecdsa-secp256k1-recovery2020-0.0.jsonld",
		},
	}

	if deactivated {
		doc.Context = []string{
			"https://w3id.org/did/v1",
		}
		return doc, fmt.Errorf("deactivated")
	}
	// Sorted pks
	keys := pks.Keys()
	for _, k := range keys {
		v, _ := pks.Get(k)
		verificationMethod = append(verificationMethod, v.(resolver.VerificationMethod))
	}
	// Sorted auths
	keys = auths.Keys()
	for _, k := range keys {
		v, _ := auths.Get(k)
		authentication = append(authentication, v.(string))
	}
	doc.Authentication = authentication
	doc.VerificationMethod = verificationMethod

	// Sorted services
	keys = services.Keys()
	for _, k := range keys {
		v, _ := services.Get(k)
		doc.Service = append(doc.Service, v.(resolver.ServiceEndpoint))
	}

	return doc, nil
}

// NewBoundContract binds a generic wrapper to an already deployed contract.
func NewBoundContract(address common.Address, backend bind.ContractBackend) (*abi.ABI, *bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(contracts.EthereumDIDRegistryABI))
	if err != nil {
		return nil, nil, err
	}
	return &parsed, bind.NewBoundContract(address, parsed, backend, backend, backend), nil
}

// InterpretIdentifier takes an input (hex) string and returns an address and/or public key string
func InterpretIdentifier(identifier string) (common.Address, *ecdsa.PublicKey, error) {
	if len(identifier) > 42 {
		if strings.HasPrefix(identifier, "0x") {
			identifier = identifier[2:]
		}
		publicKeyBytes, err := hex.DecodeString(identifier)
		if err != nil {
			return common.Address{}, nil, err
		}
		publicKey, err := crypto.DecompressPubkey(publicKeyBytes)
		if err != nil {
			return common.Address{}, nil, err
		}
		return crypto.PubkeyToAddress(*publicKey), publicKey, nil
	}
	return common.HexToAddress(identifier), nil, nil
}

func getLogs(client bind.ContractBackend, address common.Address, previousChange *big.Int) ([]*EthereumDIDRegistryDIDEventUnion, error) {
	// Leading bytes are used here, because BytesToHash trims from the left
	leading := common.Hex2Bytes("0x000000000000000000000000")
	topic := common.BytesToHash(append(leading, address.Bytes()...))
	query := ethereum.FilterQuery{
		FromBlock: previousChange,
		ToBlock:   previousChange,
		Addresses: []common.Address{
			ContractAddress,
		},
	}
	// The first topic has to be nil to allow all event types through
	// The second topic is set to the input address to get all events for a given
	// address, otherwise, we'd get all updates for all addresses
	query.Topics = append(query.Topics, nil, []common.Hash{topic})
	// We use the lower-level FilterLogs here rather than the auto-built
	// contract log filters because we want all the events, not just a single
	// type.
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	// Create a new filterer that is bound to the given address
	abi, bound, err := NewBoundContract(address, client)
	if err != nil {
		return nil, err
	}

	// We keep all event types together as generic interface{} to enable in-order
	// parsing later
	events := make([]*EthereumDIDRegistryDIDEventUnion, len(logs))

	// Parse the events in reverse order, so that we can build up the doc
	for i := len(logs) - 1; i >= 0; i-- {
		log := logs[i]
		evt, err := abi.EventByID(log.Topics[0])
		if err != nil {
			return nil, err
		}
		e := new(EthereumDIDRegistryDIDEventUnion)
		e.Type = evt.Name
		if err = bound.UnpackLog(e, evt.Name, log); err != nil {
			return nil, err
		}
		e.Raw = log
		events[i] = e
	}
	return events, nil
}

func changeLog(client bind.ContractBackend, identifier string) (*common.Address, []*EthereumDIDRegistryDIDEventUnion, *ecdsa.PublicKey, error) {
	contract, err := contracts.NewEthereumDIDRegistry(ContractAddress, client)
	if err != nil {
		return nil, nil, nil, err
	}
	var history []*EthereumDIDRegistryDIDEventUnion
	address, publicKey, err := InterpretIdentifier(identifier)
	previousChange, err := contract.Changed(nil, address)
	if err != nil {
		return nil, nil, nil, err
	}
	var controller = address
	if previousChange != nil {
		controllerRecord, err := contract.IdentityOwner(nil, address)
		if err != nil {
			return nil, nil, nil, err
		}
		if strings.ToLower(controllerRecord.Hex()) != strings.ToLower(controller.Hex()) {
			publicKey = nil
		}
		controller = controllerRecord
	}
	type info struct {
		// TODO: Do we want to add more keys here?
		PreviousChange *big.Int
	}
	for previousChange != nil {
		blockNumber := previousChange
		logs, err := getLogs(client, address, previousChange)
		if err != nil {
			return nil, nil, nil, err
		}
		previousChange = nil
		for _, event := range logs {
			history = append([]*EthereumDIDRegistryDIDEventUnion{event}, history...)
			if event.PreviousChange.Cmp(blockNumber) < 0 {
				previousChange = event.PreviousChange
			}
		}
	}
	return &controller, history, publicKey, nil
}
