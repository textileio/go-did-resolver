// Package resolver provides tools for resolving did documents.
// Copyright 2021 Textile
// Copyright 2018 ConsenSys AG

package resolver

import (
	"fmt"
	"reflect"
	"testing"

	did "github.com/ockam-network/did"
)

// TestUnhandledMethods calls resolve with an unregistered resolver, expecting
// an error to be thrown.
func TestUnknownError(t *testing.T) {
	r := New([]Resolver{}, false)
	_, _, _, err := r.Resolve("did:borg:2nQtiQG6Cgm1GYTBaaKAgr76uY7iSexUkqX", nil)
	if err == nil || err.Error() != "unknown did method: 'borg'" {
		t.Error("expected Resolve to return error")
	}
}

// TestParseError calls resolve with an invalid did url, expecting an error
// to be thrown.
func TestParseError(t *testing.T) {
	r := New([]Resolver{}, false)
	_, _, _, err := r.Resolve("did:borg:", nil)
	if err == nil || err.Error() != "idstring must be atleast one char long" {
		t.Error("expected Resolve to return error")
	}
}

type basicResolver struct{}

func (r basicResolver) Resolve(did string, parsed *did.DID, resolver Resolver) (*Document, error) {
	doc := &Document{
		Context:    []string{"https://w3id.org/did/v1"},
		ID:         did,
		Controller: []string{"1234"},
	}
	return doc, nil
}

func (r basicResolver) Method() string {
	return "basic"
}

// TestBasicResolve calls resolve and expects the result to include an basic
// id and to equal a basic document.
func TestBasicResolve(t *testing.T) {
	r := New([]Resolver{
		basicResolver{},
	}, false)
	expected := &Document{
		Context:    []string{"https://w3id.org/did/v1"},
		ID:         "did:basic:123456789",
		Controller: []string{"1234"},
	}
	_, observed, _, err := r.Resolve("did:basic:123456789", nil)
	if !reflect.DeepEqual(observed, expected) {
		t.Errorf("expected document to match: %s", expected.ID)
	}
	if err != nil {
		t.Error("unexpected error")
	}
}

type nilResolver struct{}

func (r nilResolver) Resolve(did string, parsed *did.DID, resolver Resolver) (*Document, error) {
	return nil, nil
}

func (r nilResolver) Method() string {
	return "nil"
}

// TestNilError calls resolve with did that should return nil, expecting an
// error to be thrown.
func TestNilError(t *testing.T) {
	r := New([]Resolver{
		nilResolver{},
	}, false)
	_, _, _, err := r.Resolve("did:nil:asdfghjk", nil)
	if err == nil || err.Error() != "resolver returned nil for: 'did:nil:asdfghjk'" {
		t.Error("expected Resolve to return error")
	}
}

type countingResolver struct {
	Count int
}

func (r *countingResolver) Resolve(did string, parsed *did.DID, resolver Resolver) (*Document, error) {
	r.Count = r.Count + 1
	doc := &Document{
		Context:    []string{"https://w3id.org/did/v1"},
		ID:         fmt.Sprintf("did:%s:%s", parsed.Method, parsed.ID),
		Controller: []string{"1234"},
	}
	return doc, nil
}

func (r *countingResolver) Method() string {
	return "mock"
}

// TestCacheFalse calls resolve multiple times with caching disabled, and
// expects Resolve to be called each time.
func TestCacheFalse(t *testing.T) {
	resolver := &countingResolver{}
	r := New([]Resolver{
		resolver,
	}, false)
	expected := &Document{
		Context:    []string{"https://w3id.org/did/v1"},
		ID:         "did:mock:123456789",
		Controller: []string{"1234"},
	}
	// First time
	_, observed, _, err := r.Resolve("did:mock:123456789", nil)
	if !reflect.DeepEqual(observed, expected) {
		t.Errorf("expected document to match: %s", expected.ID)
	}
	if err != nil {
		t.Error("unexpected error")
	}

	// Second time
	_, observed, _, err = r.Resolve("did:mock:123456789", nil)
	if !reflect.DeepEqual(observed, expected) {
		t.Errorf("expected document to match: %s", expected.ID)
	}
	if err != nil {
		t.Error("unexpected error")
	}

	// Should be called twice
	if resolver.Count != 2 {
		t.Errorf("expected Resolve to be called 2 times, got %d", resolver.Count)
	}
}

// TestCacheTrue calls resolve multiple times with caching enabled, and
// expects Resolve to be called once, falling back on cache the second time.
func TestCacheTrue(t *testing.T) {
	resolver := &countingResolver{}
	r := New([]Resolver{
		resolver,
	}, true)
	expected := &Document{
		Context:    []string{"https://w3id.org/did/v1"},
		ID:         "did:mock:123456789",
		Controller: []string{"1234"},
	}
	// First time
	_, observed, _, err := r.Resolve("did:mock:123456789", nil)
	if !reflect.DeepEqual(observed, expected) {
		t.Errorf("expected document to match: %s", expected.ID)
	}
	if err != nil {
		t.Error("unexpected error")
	}

	// Second time
	_, observed, _, err = r.Resolve("did:mock:123456789", nil)
	if !reflect.DeepEqual(observed, expected) {
		t.Errorf("expected document to match: %s", expected.ID)
	}
	if err != nil {
		t.Error("unexpected error")
	}

	// Should be called only once
	if resolver.Count != 1 {
		t.Errorf("expected Resolve to be called 1 times, got %d", resolver.Count)
	}
}

// TestRespectCache calls resolve multiple times with caching enabled, but
// with no-cache=true on the second call, expecting Resolve to be called each
// time.
func TestRespectCache(t *testing.T) {
	resolver := &countingResolver{}
	r := New([]Resolver{
		resolver,
	}, true)
	expected := &Document{
		Context:    []string{"https://w3id.org/did/v1"},
		ID:         "did:mock:123456789",
		Controller: []string{"1234"},
	}
	// First time
	_, observed, _, err := r.Resolve("did:mock:123456789", nil)
	if !reflect.DeepEqual(observed, expected) {
		t.Errorf("expected document to match: %s", expected.ID)
	}
	if err != nil {
		t.Error("unexpected error")
	}

	// Second time
	_, observed, _, err = r.Resolve("did:mock:123456789;no-cache=true", nil)
	if !reflect.DeepEqual(observed, expected) {
		t.Errorf("expected document to match: %s", expected.ID)
	}
	if err != nil {
		t.Error("unexpected error")
	}

	// Should be called twice
	if resolver.Count != 2 {
		t.Errorf("expected Resolve to be called 2 times, got %d", resolver.Count)
	}
}

type errorResolver struct{}

func (r errorResolver) Resolve(did string, parsed *did.DID, resolver Resolver) (*Document, error) {
	return nil, fmt.Errorf("error resolving: '%s'", did)
}

func (r errorResolver) Method() string {
	return "error"
}

// TestResolverError calls Resolve on a Resolver that returns an error,
// expecting the error to be propagated back to the Registry.
func TestResolverError(t *testing.T) {
	r := New([]Resolver{
		errorResolver{},
	}, false)
	_, _, _, err := r.Resolve("did:error:asdfghjk", nil)
	if err == nil || err.Error() != "error resolving: 'did:error:asdfghjk'" {
		t.Error("expected Resolve to return error")
	}
}

func TestResolverErrorWithCache(t *testing.T) {
	r := New([]Resolver{
		errorResolver{},
	}, true)
	_, _, _, err := r.Resolve("did:error:asdfghjk", nil)
	if err == nil || err.Error() != "error resolving: 'did:error:asdfghjk'" {
		t.Error("expected Resolve to return error")
	}
	_, _, _, err = r.Resolve("did:error:asdfghjk", nil)
	if err == nil || err.Error() != "error resolving: 'did:error:asdfghjk'" {
		t.Error("expected Resolve to return error")
	}
	// TODO: We SHOULD actually check to make sure the cache wasn't updated.
}
