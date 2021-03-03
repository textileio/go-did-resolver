// Package resolver implements did resolution functions resolve a did into a
// did document. It aims to be did spec v1.0 compliant.
// See https://w3c.github.io/did-core/#did-resolution
// Copyright 2021 Textile
// Copyright 2018 ConsenSys AG
package resolver

import (
	"fmt"

	parse "github.com/ockam-network/did"
)

// Document is a did document that describes a did subject.
// See https://www.w3.org/TR/did-core/#dfn-did-documents.
// We include some non-standard keys here for compatability reasons
// See https://w3c.github.io/did-spec-registries/#publickey
type Document struct {
	Context            []string             `json:"@context"` // https://w3id.org/did/v1
	ID                 string               `json:"id"`
	Controller         []string             `json:"controller,omitempty"`
	VerificationMethod []VerificationMethod `json:"verificationMethod,omitempty"`
	Authentication     []VerificationMethod `json:"authentication,omitempty"`
	KeyAgreement       []VerificationMethod `json:"keyAgreement,omitempty"`
	Service            []ServiceEndpoint    `json:"service,omitempty"`
}

// ServiceEndpoint descrives a network address, such as an http url, at which services operate on behalf of a did subject.
// See https://www.w3.org/TR/did-core/#dfn-service-endpoints
type ServiceEndpoint struct {
	ID              string `json:"id"`
	Type            string `json:"type,omitempty"`
	ServiceEndpoint string `json:"serviceEndpoint,omitempty"`
	Description     string `json:"description,omitempty"`
}

// VerificationMethod describes how to authenticate or authorize interactions with a did subject.
// See https://www.w3.org/TR/did-core/#dfn-verification-method.
type VerificationMethod struct {
	ID                 string `json:"id,omitempty"`
	Type               string `json:"type,omitempty"`
	Controller         string `json:"controller,omitempty"`
	PublicKeyMultibase string `json:"publicKeyMultibase,omitempty"`
	PublicKey          string `json:"publicKey,omitempty"`
}

// VerificationString describes how to authenticate or authorize interactions with a did subject.
// See https://www.w3.org/TR/did-core/#dfn-verification-method.
// See https://www.w3.org/TR/did-core/#did-url-syntax
type VerificationString string

// ResolutionOptions defines a metadata structure containing properties that may be used to control did resolution.
// See https://w3c.github.io/did-core/#did-resolution-options
// See https://w3c.github.io/did-spec-registries/#did-resolution-input-metadata
type ResolutionOptions struct {
	// Accept indicates the media type of the caller's preferred representation of the did document.
	// This property must not be used with the resolve function.
	Accept string `json:"accept,omitempty"`
}

// ResolutionMetadata defined a metadata structure consisting of values relating to the results of the did resolution process.
// See https://w3c.github.io/did-core/#did-resolution-metadata
// See https://w3c.github.io/did-spec-registries/#did-resolution-metadata
// ContentType indicates the media type of the returned did document stream.
// This property must not be present if the resolve function was called.
// The error code from the resolution process. This property is required when there is an error in the resolution process.
// Curret possible values include: "invalid-did", "invalid-did-url", "not-found", "deactivated", "unsupported-did-method", "representation-not-supported"
type ResolutionMetadata struct {
	ContentType string `json:"content-type,omitempty"`
	Error       string `json:"error,omitempty"`
}

// DocumentMetadata defines a metadata structure consisting of values relating the resolved did document.
// See https://w3c.github.io/did-core/#did-document-metadata
// See https://w3c.github.io/did-spec-registries/#did-document-metadata
// Created is a string formatted as an XML Datetime normalized to UTC 00:00:00 and without sub-second decimal precision.
// Updated is a string formatted as an XML Datetime normalized to UTC 00:00:00 and without sub-second decimal precision.
// Deactivated is a boolean value indicating if the associated did document has been deactivated.
// VersionID indicates the version of the last update operation for the document version which was resolved.
type DocumentMetadata struct {
	Created     string `json:"created,omitempty"`
	Updated     string `json:"updated,omitempty"`
	Deactivated bool   `json:"deactivated,omitempty"`
	VersionID   string `json:"versionId,omitempty"`
}

// Resolver defines a basic did resolver implementation.
type Resolver interface {
	// Method returns the method that this resolver is capable of resolving.
	Method() string
	// Resolve is the primary resolution method for the Resolver.
	Resolve(did string, parsed *parse.DID, resolver Resolver) (*Document, error)
}

// wrappedResolve is a simple function type for wrapping a did Resolver.
type wrappedResolve func() (*Document, error)

// Cache is a simple did document cache.
type Cache interface {
	// Get checks the cache for a given did, and if found, returns the previous result.
	Get(parsed *parse.DID, resolve wrappedResolve) (*Document, error)
}

// memoryCache is a simple in-memory did document cache.
type memoryCache map[string]*Document

func (c memoryCache) Get(parsed *parse.DID, resolve wrappedResolve) (*Document, error) {
	if len(parsed.Params) > 0 {
		for _, p := range parsed.Params {
			if p.Name == "no-cache" && p.Value == "true" {
				return resolve()
			}
		}
	}
	cached, ok := c[parsed.String()]
	if ok {
		return cached, nil
	}
	doc, err := resolve()
	if err != nil {
		return nil, err
	}
	c[parsed.String()] = doc
	return doc, nil
}

// Registry is the singleton Resolver registery for resolving dids.
type Registry struct {
	registry map[string]Resolver
	cache    Cache
}

// New creates and returns a new resolver Registry.
func New(resolvers []Resolver, useCache bool) Registry {
	var cache Cache
	if useCache {
		cache = memoryCache{}
	}
	registry := make(map[string]Resolver)
	for _, resolver := range resolvers {
		registry[resolver.Method()] = resolver
	}
	return Registry{
		registry,
		cache,
	}
}

// Parse parses a did url into a did struct.
func (r Registry) Parse(did string) (*parse.DID, error) {
	return parse.Parse(did)
}

// Resolve resolves a did url using a Resolver from its internal registry.
// Currently, resolutionOptions are ignored, as the defaults to not pertain to the resolve function.
// See https://w3c.github.io/did-core/#did-resolution-options for details.
func (r Registry) Resolve(did string, resolutionOptions *ResolutionOptions) (ResolutionMetadata, *Document, DocumentMetadata, error) {
	parsed, err := r.Parse(did)
	if err != nil {
		return ResolutionMetadata{Error: "invalid-did-url"}, nil, DocumentMetadata{}, err
	}
	resolver, ok := r.registry[parsed.Method]
	var doc *Document
	if ok && resolver != nil {
		if r.cache != nil {
			doc, err = r.cache.Get(parsed, func() (*Document, error) {
				return resolver.Resolve(parsed.String(), parsed, resolver)
			})
		} else {
			doc, err = resolver.Resolve(parsed.String(), parsed, resolver)
		}
		if err != nil {
			return ResolutionMetadata{Error: "invalid-did"}, nil, DocumentMetadata{}, err
		}
		if doc == nil {
			return ResolutionMetadata{Error: "not-found"}, nil, DocumentMetadata{}, fmt.Errorf("resolver returned nil for: '%s'", parsed.String())
		}
		return ResolutionMetadata{}, doc, DocumentMetadata{}, nil
	}
	return ResolutionMetadata{Error: "invalid-did-url"}, nil, DocumentMetadata{}, fmt.Errorf("unknown did method: '%s'", parsed.Method)
}
