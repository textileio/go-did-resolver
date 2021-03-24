// Package threeid provides tools for resolving the did:3 method format for
// the ceramic network:
// https://github.com/ceramicnetwork/CIP/blob/main/CIPs/CIP-79/CIP-79.md
// Copyright 2021 Textile
// Copyright 2021 Ceramic Network
package threeid

import (
	"encoding/json"
	"fmt"
	"strings"

	cid "github.com/ipfs/go-cid"
	multibase "github.com/multiformats/go-multibase"
	codec "github.com/multiformats/go-multicodec"
	varint "github.com/multiformats/go-varint"
	did "github.com/ockam-network/did"
	resolver "github.com/textileio/go-did-resolver/resolver"
)

// Resolver defines a basic key resolver, conforming the to Resolver interface.
type Resolver struct {
	client Client
}

// New creates and returns a new key Resolver.
func New() *Resolver {
	return &Resolver{}
}

// Method returns the method that this resolver is capable of resolving.
func (r *Resolver) Method() string {
	return "3"
}

// Resolve is the primary resolution method for this resolver.
func (r *Resolver) Resolve(did string, parsed *did.DID, res resolver.Resolver) (*resolver.Document, error) {
	if parsed.Method != r.Method() {
		return nil, fmt.Errorf("unknown did method: '%s'", parsed.Method)
	}
	var c = cid.Undef
	var err error
	version := getVersion(parsed.Query)
	if version != "" {
		c, err = cid.Parse(version)
	}
	if err != nil {
		return nil, err
	}
	return resolve(r.client, parsed.ID, c)
}

func resolve(client Client, id string, commit cid.Cid) (*resolver.Document, error) {
	docID, err := fromString(id)
	if err != nil {
		return nil, err
	}
	commitID, err := docID.AtCommit(commit)
	if err != nil {
		return nil, err
	}
	state, err := client.LoadDocument(commitID)
	if err != nil {
		return nil, err
	}

	did := fmt.Sprintf("did:3:%s", id)
	var doc = &resolver.Document{
		ID:      did,
		Context: []string{"https://w3id.org/did/v1"},
	}

	type Content struct {
		PublicKeys map[string]string `json:"publicKeys,omitempty"`
	}

	var content Content
	err = json.Unmarshal(*state.Content, &content)
	if err != nil {
		return nil, err
	}

	// Loop through content.publicKeys to sort things out
	for keyName, keyValue := range content.PublicKeys {
		_, keyBuf, err := multibase.Decode(keyValue)
		if err != nil {
			return nil, err
		}
		keyType, n, err := varint.FromUvarint(keyBuf)
		if err != nil {
			return nil, err
		}
		if n != 2 {
			return nil, fmt.Errorf("error parsing varint")
		}
		publicKeyBase58, err := multibase.Encode(multibase.Base58BTC, keyBuf[2:])
		if err != nil {
			return nil, err
		}
		keyID := fmt.Sprintf("%s#%s", did, keyName)
		switch keyType {
		case uint64(codec.Secp256k1Pub):
			doc.VerificationMethod = append(doc.VerificationMethod, resolver.VerificationMethod{
				ID:                 keyID,
				Type:               "Secp256k1VerificationKey2018",
				Controller:         did,
				PublicKeyMultibase: publicKeyBase58,
			})
			doc.Authentication = append(doc.Authentication, keyID)
		case uint64(codec.X25519Pub):
			// Old key format, likely not needed in the future
			doc.VerificationMethod = append(doc.VerificationMethod, resolver.VerificationMethod{
				ID:                 keyID,
				Type:               "Curve25519EncryptionPublicKey",
				Controller:         did,
				PublicKeyMultibase: publicKeyBase58,
			})
			// New keyAgreement format for x25519 keys
			doc.KeyAgreement = append(doc.KeyAgreement, resolver.VerificationMethod{
				ID:                 keyID,
				Type:               "X25519KeyAgreementKey2019",
				Controller:         did,
				PublicKeyMultibase: publicKeyBase58,
			})
		}
	}
	return doc, nil
}

// getVersion gets the identifier of the version of the did document that was
// requested. This will correspond to a specific 'commit' of the ceramic
// document.
func getVersion(query string) string {
	parts := strings.Split(query, "&")
	for _, part := range parts {
		if strings.Contains(part, "versionId") {
			return strings.Split(part, "=")[1]
		}
	}
	return ""
}

var _ resolver.Resolver = (*Resolver)(nil)
