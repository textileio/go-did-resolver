// Package keys provides tools for resolving the w3c did:key format, for
// static cryptographic keys: https://w3c-ccg.github.io/did-method-key/#format
// Copyright 2021 Textile
// Copyright 2021 Ceramic Network
package keys

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	mbase "github.com/multiformats/go-multibase"
	codec "github.com/multiformats/go-multicodec"
	varint "github.com/multiformats/go-varint"
	did "github.com/ockam-network/did"
	"github.com/textileio/go-did-resolver/resolver"
)

// TestExpandEd25519Key calls ExpandEd25519Key directly, and checks that the
// document is correctly derived from the input key.
func TestExpandEd25519Key(t *testing.T) {
	id := "z6MktvqCyLxTsXUH1tUZncNdVeEZ7hNh7npPRbUU27GTrYb8"
	encoding, multicodecPublicKey, err := mbase.Decode(id)
	if encoding != mbase.Base58BTC {
		t.Errorf("expected base58 btc, got: %s", mbase.EncodingToStr[encoding])
	}
	keyType, n, err := varint.FromUvarint(multicodecPublicKey)
	if n != 2 {
		t.Error("error parsing varint")
	}
	if keyType != uint64(codec.Ed25519Pub) {
		t.Errorf("expected ed25519 public key, got: %s", codec.Code(keyType).String())
	}
	publicKeyBytes := multicodecPublicKey[n:]

	byteValue, err := ioutil.ReadFile("testdata/ed25519.json")
	if err != nil {
		t.Error(err)
	}
	var expected resolver.Document
	err = json.Unmarshal(byteValue, &expected)
	if err != nil {
		t.Error(err)
	}

	observed, err := ExpandEd25519Key(publicKeyBytes, id)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(observed, &expected) {
		t.Errorf("expected document to match: %s", expected.ID)
	}
}

// TestExpandSecp256k1Key calls ExpandSecp256k1Key directly, and checks that
// the document is correctly derived from the input key.
func TestExpandSecp256k1Key(t *testing.T) {
	id := "zQ3shbgnTGcgBpXPdBjDur3ATMDWhS7aPs6FRFkWR19Lb9Zwz"
	encoding, multicodecPublicKey, err := mbase.Decode(id)
	if encoding != mbase.Base58BTC {
		t.Errorf("expected base58 btc, got: %s", mbase.EncodingToStr[encoding])
	}
	keyType, n, err := varint.FromUvarint(multicodecPublicKey)
	if n != 2 {
		t.Error("error parsing varint")
	}
	if keyType != uint64(codec.Secp256k1Pub) {
		t.Errorf("expected secp256k1 public key, got: %s", codec.Code(keyType).String())
	}
	publicKeyBytes := multicodecPublicKey[n:]

	byteValue, err := ioutil.ReadFile("testdata/secp256k1.json")
	if err != nil {
		t.Error(err)
	}
	var expected resolver.Document
	err = json.Unmarshal(byteValue, &expected)
	if err != nil {
		t.Error(err)
	}

	observed, err := ExpandSecp256k1Key(publicKeyBytes, id)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(observed, &expected) {
		t.Errorf("expected document to match: %s", expected.ID)
	}
}

func TestEd25519KeyResolver(t *testing.T) {
	id := "did:key:z6MktvqCyLxTsXUH1tUZncNdVeEZ7hNh7npPRbUU27GTrYb8"
	// TODO: Carson's local random did:
	// id := "did:key:z6Mkg9cRA65MvVbNdZCMpiqiFiWD8guY9oRGHeeC3J1qsCk5"

	parsed, err := did.Parse(id)
	if err != nil {
		t.Error(err)
	}
	r := New()
	observed, err := r.Resolve(id, parsed, r)
	if err != nil {
		t.Error(err)
	}
	byteValue, err := ioutil.ReadFile("testdata/ed25519.json")
	if err != nil {
		t.Error(err)
	}
	var expected resolver.Document
	err = json.Unmarshal(byteValue, &expected)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(observed, &expected) {
		t.Errorf("expected document to match: %s", expected.ID)
	}
}

func TestSecp256k1KeyResolver(t *testing.T) {
	id := "did:key:zQ3shbgnTGcgBpXPdBjDur3ATMDWhS7aPs6FRFkWR19Lb9Zwz"

	parsed, err := did.Parse(id)
	if err != nil {
		t.Error(err)
	}
	r := New()
	observed, err := r.Resolve(id, parsed, r)
	if err != nil {
		t.Error(err)
	}
	byteValue, err := ioutil.ReadFile("testdata/secp256k1.json")
	if err != nil {
		t.Error(err)
	}
	var expected resolver.Document
	err = json.Unmarshal(byteValue, &expected)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(observed, &expected) {
		t.Errorf("expected document to match: %s", expected.ID)
	}
}

func TestUnknownMethod(t *testing.T) {
	id := "did:borg:zQ3shbgnTGcgBpXPdBjDur3ATMDWhS7aPs6FRFkWR19Lb9Zwz"

	parsed, err := did.Parse(id)
	if err != nil {
		t.Error(err)
	}
	r := New()
	_, err = r.Resolve(id, parsed, r)
	if err == nil || err.Error() != "unknown did method: 'borg'" {
		t.Error("expected Resolve to return error")
	}
}

func TestVarintParsingError(t *testing.T) {
	id := "did:key:zhash"

	parsed, err := did.Parse(id)
	if err != nil {
		t.Error(err)
	}
	r := New()
	_, err = r.Resolve(id, parsed, r)
	if err == nil || err.Error() != "error parsing varint" {
		t.Error("expected Resolve to return error")
	}
}

func TestUnknownKeyType(t *testing.T) {
	id := "did:key:z6LSnkZe3JZPCo88XsdQVJi8j1TomzX2yRW7ZnvvWhmrSdmG"

	parsed, err := did.Parse(id)
	if err != nil {
		t.Error(err)
	}
	r := New()
	_, err = r.Resolve(id, parsed, r)
	// We don't support "x25519-pub" keys directly
	if err == nil || err.Error() != "unknown key type: 'x25519-pub'" {
		t.Error("expected Resolve to return error")
	}
}

func TestInvalidMultibase(t *testing.T) {
	id := "did:key:noop"

	parsed, err := did.Parse(id)
	if err != nil {
		t.Error(err)
	}
	r := New()
	_, err = r.Resolve(id, parsed, r)
	// We don't support "x25519-pub" keys directly
	if err == nil || err.Error() != "selected encoding not supported" {
		t.Error("expected Resolve to return error")
	}
}

func TestInvalidVarint(t *testing.T) {
	id := "did:key:zosEoy933LkHyyBbwyawsww567i"

	parsed, err := did.Parse(id)
	if err != nil {
		t.Error(err)
	}
	r := New()
	_, err = r.Resolve(id, parsed, r)
	// We don't support "x25519-pub" keys directly
	if err == nil || err.Error() != "varints larger than uint63 not supported" {
		t.Error("expected Resolve to return error")
	}
}
