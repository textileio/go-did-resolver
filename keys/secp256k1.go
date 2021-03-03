// Package keys provides tools for resolving the w3c did:key format, for
// static cryptographic keys: https://w3c-ccg.github.io/did-method-key/#format
// Copyright 2021 Textile
// Copyright 2021 Ceramic Network
package keys

import (
	"fmt"

	mbase "github.com/multiformats/go-multibase"
	"github.com/textileio/go-did-resolver/resolver"
)

// ExpandSecp256k1Key creates a did Document from an input secp256k1 public key.
func ExpandSecp256k1Key(bytes []byte, fingerprint string) (*resolver.Document, error) {
	did := fmt.Sprintf("did:key:%s", fingerprint)
	keyID := fmt.Sprintf("%s#%s", did, fingerprint)
	// No error checking here, because we're using mbase consts directly.
	keyMultiBase, _ := mbase.Encode(mbase.Base16, bytes)
	doc := &resolver.Document{
		Context: []string{"https://w3id.org/did/v1"},
		ID:      did,
		VerificationMethod: []resolver.VerificationMethod{
			{
				ID:                 keyID,
				Type:               "Secp256k1VerificationKey2018",
				Controller:         did,
				PublicKeyMultibase: keyMultiBase,
			},
		},
	}
	// Leaving possible error return gere in case we need it in the future
	return doc, nil
}
