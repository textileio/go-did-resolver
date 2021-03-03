// Package threeid provides tools for resolving the did:3 method format for
// the ceramic network:
// https://github.com/ceramicnetwork/CIP/blob/main/CIPs/CIP-79/CIP-79.md
// Copyright 2021 Textile
// Copyright 2021 Ceramic Network
package threeid

import (
	"fmt"
	"strings"

	cid "github.com/ipfs/go-cid"
	mbase "github.com/multiformats/go-multibase"
	varint "github.com/multiformats/go-varint"
)

const docIDCodec = 206

// DocIdentifier defines any cid-based document with a type and bytes representation.
type DocIdentifier interface {
	fmt.Stringer
	Type() uint64
	CID() cid.Cid
	Bytes() []byte
}

// DocID is a document identifier, no commit information included.
// Encoded as '<multibase-prefix><multicodec-docid><doctype><genesis-cid-bytes>'
type DocID struct {
	// DocType is the document type. Current supported types include "tile": 0,
	// and "caip10-link": 1
	docType uint64
	// Genesis CID
	cid cid.Cid
}

// Parse DocID from bytes representation.
func fromBytes(bytes []byte) (*DocID, error) {
	codec, n1, err := varint.FromUvarint(bytes)
	if err != nil {
		return nil, err
	}
	if codec != docIDCodec {
		return nil, fmt.Errorf("invalid docid, does not include docid codec")
	}
	t, n2, err := varint.FromUvarint(bytes[n1:])
	if err != nil {
		return nil, err
	}
	c, err := cid.Cast(bytes[n1+n2:])
	if err != nil {
		return nil, err
	}
	return &DocID{t, c}, nil
}

// Parse DocID from string representation.
// str is the string representation of DocID, be it base36-encoded string or URL.
func fromString(str string) (*DocID, error) {
	protocolFree := strings.Replace(strings.Replace(str, "ceramic://", "", -1), "/ceramic/", "", -1)
	var commitFree string
	if strings.Contains(protocolFree, "commit") {
		commitFree = strings.Split(protocolFree, "?")[0]
	} else {
		commitFree = protocolFree
	}
	_, bytes, err := mbase.Decode(commitFree)
	if err != nil {
		return nil, err
	}
	return fromBytes(bytes)
}

// Type returns the document type.
func (doc *DocID) Type() uint64 {
	return doc.docType
}

// CID returns the document cid. The genesis cid is undefined (i.e., cid.Undef).
func (doc *DocID) CID() cid.Cid {
	return doc.cid
}

// AtCommit returns a CommitID based on the DocID
func (doc *DocID) AtCommit(commit cid.Cid) (*CommitID, error) {
	return &CommitID{DocID{doc.Type(), doc.CID()}, commit}, nil
}

// Bytes returns the bytes representation for a doc id.
func (doc *DocID) Bytes() []byte {
	codec := varint.ToUvarint(docIDCodec)
	docType := varint.ToUvarint(doc.Type())
	bytes := append(codec, append(docType, doc.CID().Bytes()...)...)
	return bytes
}

func (doc *DocID) String() string {
	bytes := doc.Bytes()
	encoded, _ := mbase.Encode(mbase.Base36, bytes)
	return encoded
}

var _ DocIdentifier = (*DocID)(nil)

// CommitID is a commit identifier, includes doctype, genesis cid, commit cid.
type CommitID struct {
	DocID
	// cid.Undef ≝ genesis commit
	commit cid.Cid
}

// Commit returns the document commit cid.
func (doc *CommitID) Commit() cid.Cid {
	if doc.commit.Equals(cid.Undef) {
		return doc.CID()
	}
	return doc.commit
}

// Bytes returns the bytes representation for a doc id.
func (doc *CommitID) Bytes() []byte {
	codec := varint.ToUvarint(docIDCodec)
	docType := varint.ToUvarint(doc.Type())
	bytes := append(codec, append(docType, doc.CID().Bytes()...)...)
	var commit []byte
	if doc.commit.Equals(cid.Undef) {
		commit = []byte{0}
	} else {
		commit = doc.Commit().Bytes()
	}
	return append(bytes, commit...)
}

func (doc *CommitID) String() string {
	bytes := doc.Bytes()
	encoded, _ := mbase.Encode(mbase.Base36, bytes)
	return encoded
}

var _ DocIdentifier = (*CommitID)(nil)
