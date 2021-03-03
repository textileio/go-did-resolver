// Package threeid provides tools for resolving the did:3 method format for
// the ceramic network:
// https://github.com/ceramicnetwork/CIP/blob/main/CIPs/CIP-79/CIP-79.md
// Copyright 2021 Textile
// Copyright 2021 Ceramic Network
package threeid

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	// DefaultAPIPath is the default version path for the Ceramic api
	DefaultAPIPath = "/api/v0"
	// DefaultHost is the default host path/url for the Ceramic api
	DefaultHost = "http://localhost:7007"
)

// Client is a basic client interface for interacting with the Ceramic network.
type Client interface {
	LoadDocument(docID DocIdentifier) (*DocState, error)
}

// DocMetadata represents metadata about the primary ceramic document.
type DocMetadata struct {
	Controllers []string `json:"controllers,omitempty"`
	Schema      string   `json:"schema,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

// DocNext represents information about the next iteration.
type DocNext struct {
	Content     *json.RawMessage `json:"content,omitempty"`
	Controllers []string         `json:"contollers,omitempty"`
	Metadata    DocMetadata      `json:"metadata,omitempty"`
}

// LogEntry represents an entry in a document's update history.
type LogEntry struct {
	// TODO: We actually want this as a cid.CID at some point
	// CID  cid.Cid
	CID  string `json:"cid"`
	Type uint   `json:"type"`
}

// DocResponse represents the primary document query response.
type DocResponse struct {
	DocID string   `json:"docId"`
	State DocState `json:"state"`
}

// DocState represents the primary ceramic document.
type DocState struct {
	Doctype            string           `json:"doctype"`
	Content            *json.RawMessage `json:"content,omitempty"`
	Next               DocNext          `json:"next,omitempty"`
	Metadata           DocMetadata      `json:"metadata,omitempty"`
	Signature          uint64           `json:"signature,omitempty"`          // 0=GENESIS, 1=PARTIAL, 2=SIGNED
	AnchorStatus       string           `json:"anchorStatus,omitempty"`       // NOT_REQUESTED, PENDING, PROCESSING, ANCHORED, FAILED
	AnchorScheduledFor uint64           `json:"anchorScheduledFor,omitempty"` // only present when anchor status is pending
	AnchorProof        AnchorProof      `json:"anchorProof,omitempty"`        // The anchor proof of the latest anchor, only present when anchor status is anchored
	Log                []LogEntry       `json:"log,omitempty"`
}

// AnchorProof represents metadata about the on-chain anchor proof.
type AnchorProof struct {
	ChainID        string `json:"chainId,omitempty"`
	BlockNumber    uint64 `json:"blockNumber,omitempty"`
	BlockTimestamp uint64 `json:"blockTimestamp,omitempty"`
	TxHash         string `json:"txHash,omitempty"`
	Root           string `json:"root,omitempty"`
}

// HTTPClient interfaces with the ceramic network via local or remote ceramic daemon.
type HTTPClient struct {
	APIURL string
}

// Load fetches the remote Ceramic document and returns it.
func (client *HTTPClient) Load(docID DocIdentifier) (*DocResponse, error) {
	var c = &http.Client{Timeout: 10 * time.Second}
	resp, err := c.Get(client.APIURL + "/documents/" + docID.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var state DocResponse
		json.NewDecoder(resp.Body).Decode(&state)
		return &state, nil

	}
	// TODO: Should return a nicer error here
	return nil, fmt.Errorf(resp.Status)
}

// LoadDocument loads a ceramic document using a local or remote ceramic daemon.
func (client *HTTPClient) LoadDocument(docID DocIdentifier) (*DocState, error) {
	resp, err := client.Load(docID)
	if err != nil {
		return nil, err
	}
	return &resp.State, nil
}

var _ Client = (*HTTPClient)(nil)
