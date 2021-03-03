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
	"github.com/textileio/go-did-resolver/resolver"

	// https://github.com/golang/go/issues/20504
	"github.com/jorrizza/ed2curve25519"
)

func encodeEd25519PublicKey(key []byte) (string, error) {
	// https://github.com/multiformats/multicodec/blob/master/table.csv#L84
	// {name:"x25519-pub", tag:"key", code:0xec, description:"Curve25519 public key"}
	prefix := varint.ToUvarint(uint64(codec.X25519Pub))
	bytes := append(prefix, key...)
	// The spec specifies base58 btc...
	// No error checking here, because we're using mbase consts directly.
	str, _ := mbase.Encode(mbase.Base58BTC, bytes)
	// Leaving possible error return gere in case we need it in the future
	return str, nil
}

// ExpandEd25519Key creates a did Document from an input ed15519 public key.
func ExpandEd25519Key(bytes []byte, fingerprint string) (*resolver.Document, error) {
	did := fmt.Sprintf("did:key:%s", fingerprint)
	keyID := fmt.Sprintf("%s#%s", did, fingerprint)
	// No error checking here, because we're using mbase consts directly.
	keyMultiBase, _ := mbase.Encode(mbase.Base58BTC, bytes)
	x25519Bytes := ed2curve25519.Ed25519PublicKeyToCurve25519(bytes)
	// No error checking here, because we're using mbase consts directly.
	x25519Encoded, _ := encodeEd25519PublicKey(x25519Bytes)
	x25519ID := fmt.Sprintf("%s#%s", did, x25519Encoded)
	// No error checking here, because we're using mbase consts directly.
	x25519MultiBase, _ := mbase.Encode(mbase.Base58BTC, x25519Bytes)
	doc := &resolver.Document{
		Context: []string{"https://w3id.org/did/v1"},
		ID:      did,
		Authentication: []resolver.VerificationMethod{
			{
				ID:                 keyID,
				Type:               "Ed25519VerificationKey2018",
				Controller:         did,
				PublicKeyMultibase: keyMultiBase,
			},
		},
		VerificationMethod: []resolver.VerificationMethod{
			{
				ID:                 x25519ID,
				Type:               "X25519KeyAgreementKey2019",
				Controller:         did,
				PublicKeyMultibase: x25519MultiBase,
			},
		},
	}
	// Leaving possible error return gere in case we need it in the future
	return doc, nil
}
