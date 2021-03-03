// Package keys provides tools for resolving the w3c did:key format, for
// static cryptographic keys: https://w3c-ccg.github.io/did-method-key/#format
// Copyright 2021 Textile
// Copyright 2021 Ceramic Network
package keys

import (
	"fmt"

	mbase "github.com/multiformats/go-multibase"
	codec "github.com/multiformats/go-multicodec"
	varint "github.com/multiformats/go-varint"
	"github.com/ockam-network/did"
	"github.com/textileio/go-did-resolver/resolver"
)

// Resolver defines a basic key resolver, conforming the to Resolver interface.
type Resolver struct{}

// New creates and returns a new key Resolver.
func New() *Resolver {
	return &Resolver{}
}

// Method returns the method that this resolver is capable of resolving.
func (r *Resolver) Method() string {
	return "key"
}

// Resolve is the primary resolution method for this resolver.
func (r *Resolver) Resolve(did string, parsed *did.DID, res resolver.Resolver) (*resolver.Document, error) {
	if parsed.Method != r.Method() {
		return nil, fmt.Errorf("unknown did method: '%s'", parsed.Method)
	}
	_, bytes, err := mbase.Decode(parsed.ID)
	if err != nil {
		return nil, err
	}
	keyType, n, err := varint.FromUvarint(bytes)
	if err != nil {
		return nil, err
	}
	if n != 2 {
		return nil, fmt.Errorf("error parsing varint")
	}
	switch keyType {
	case uint64(codec.Ed25519Pub):
		return ExpandEd25519Key(bytes[n:], parsed.ID)
	case uint64(codec.Secp256k1Pub):
		return ExpandSecp256k1Key(bytes[n:], parsed.ID)
	default:
		return nil, fmt.Errorf("unknown key type: '%s'", codec.Code(keyType).String())
	}
}

var _ resolver.Resolver = (*Resolver)(nil)
