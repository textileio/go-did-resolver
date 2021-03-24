// Package web provides tools for resolving dids formed by taking an http(s)
// domain and producing a well-known URI to access the did document.
// See https://tools.ietf.org/html/rfc5785 for details.
// Copyright 2021 Textile
// Copyright 2021 Decentralized Identity Foundation
package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ockam-network/did"
	"github.com/textileio/go-did-resolver/resolver"
)

// DocPath is the default path at which to look for the did.json document
const DocPath = "/.well-known/did.json"

// Getter is a simple interface for Getting an *http.Response.
type Getter interface {
	Get(url string) (resp *http.Response, err error)
}

// Client is the default client used to Get *http.Responses.
var Client Getter

// The init function sets the Client var to instance of http.Client.
// https://www.thegreatcodeadventure.com/mocking-http-requests-in-golang/
func init() {
	Client = &http.Client{}
}

// Resolver takes an http(s) domain and resolves a well-known URI-based did
// document.
type Resolver struct{}

// New creates and returns a new Resolver.
func New() *Resolver {
	return &Resolver{}
}

// Method returns the method that this resolver is capable of resolving.
func (r *Resolver) Method() string {
	return "web"
}

// Resolve is the primary resolution method for this resolver.
// For a did did:web:example.com, the resolver will attempt to access the
// document at https://example.com/.well-known/did.json
func (r *Resolver) Resolve(did string, parsed *did.DID, res resolver.Resolver) (*resolver.Document, error) {
	if parsed.Method != r.Method() {
		return nil, fmt.Errorf("unknown did method: '%s'", parsed.Method)
	}
	path := parsed.ID + DocPath
	parts := parsed.IDStrings
	if len(parts) > 1 {
		_, err := strconv.Atoi(parts[1])
		if err == nil { // Probably a port number
			path = strings.Join(append([]string{parts[0] + ":" + parts[1]}, parts[2:]...), "/") + DocPath
		} else {
			path = strings.Join(parts, "/") + DocPath
		}

	}
	url := "https://" + path
	resp, err := Client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to resolve: '%s'", resp.Status)
	}
	if resp.Body == nil {
		return nil, fmt.Errorf("empty did document")
	}
	defer resp.Body.Close()

	var state resolver.Document
	err = json.NewDecoder(resp.Body).Decode(&state)
	if err != nil {
		return nil, err
	}
	// Ignores query params and things in the did url
	if state.ID != fmt.Sprintf("did:web:%s", parsed.ID) {
		return nil, fmt.Errorf("id does not match requested did")
	}
	if len(state.VerificationMethod) < 1 {
		return nil, fmt.Errorf("no verification methods")
	}
	return &state, nil
}

var _ resolver.Resolver = (*Resolver)(nil)
