// Package threeid provides tools for resolving the did:3 method format for
// the ceramic network:
// https://github.com/ceramicnetwork/CIP/blob/main/CIPs/CIP-79/CIP-79.md
// Copyright 2021 Textile
// Copyright 2021 Ceramic Network
package threeid

import (
	"testing"

	"github.com/ipfs/go-cid"
	mbase "github.com/multiformats/go-multibase"
)

var (
	CIDString              = "bagcqcerakszw2vsovxznyp5gfnpdj4cqm2xiv76yd24wkjewhhykovorwo6a"
	DocIDString            = "kjzl6cwe1jw147dvq16zluojmraqvwdmbh61dx9e0c59i344lcrsgqfohexp60s"
	DocIDURL               = "ceramic://kjzl6cwe1jw147dvq16zluojmraqvwdmbh61dx9e0c59i344lcrsgqfohexp60s"
	DocIDLegacy            = "/ceramic/kjzl6cwe1jw147dvq16zluojmraqvwdmbh61dx9e0c59i344lcrsgqfohexp60s"
	CommitCIDString        = "bagjqcgzaday6dzalvmy5ady2m5a5legq5zrbsnlxfc2bfxej532ds7htpova"
	DocIDWithCommit        = "k1dpgaqe3i64kjqcp801r3sn7ysi5i0k7nxvs7j351s7kewfzr3l7mdxnj7szwo4kr9mn2qki5nnj0cv836ythy1t1gya9s25cn1nexst3jxi5o3h6qprfyju"
	DocIDWith0Commit       = "k3y52l7qbv1frxwipl4hp7e6jlu4f6u8upm2xv0irmedfkm5cnutmezzi3u7mytj4"
	DocIDWithCommitLegacy  = "/ceramic/kjzl6cwe1jw147dvq16zluojmraqvwdmbh61dx9e0c59i344lcrsgqfohexp60s?commit=bagjqcgzaday6dzalvmy5ady2m5a5legq5zrbsnlxfc2bfxej532ds7htpova"
	DocIDWith0CommitLegacy = "/ceramic/kjzl6cwe1jw147dvq16zluojmraqvwdmbh61dx9e0c59i344lcrsgqfohexp60s?commit=0"
)

// DocID

func TestCreateByParts(t *testing.T) {
	d := uint64(0)
	c, _ := cid.Parse(CIDString)
	docID := DocID{d, c}
	if docID.CID().String() != CIDString {
		t.Error("expected cid strings to match")
	}
	if docID.Type() != d {
		t.Error("expected doc type to match")
	}
	observed := docID.String()
	if observed != DocIDString {
		t.Error("expected doc id string to match")
	}
}

func TestFromBytes(t *testing.T) {
	_, bytes, _ := mbase.Decode(DocIDString)
	docID, err := fromBytes(bytes)
	if err != nil {
		t.Error(err)
	}
	if docID.CID().String() != CIDString {
		t.Error("expected cid strings to match")
	}
	if docID.Type() != uint64(0) {
		t.Error("expected doc type to match")
	}
	observed := docID.String()
	if observed != DocIDString {
		t.Error("expected doc id string to match")
	}
}

func TestFromInvalidBytes(t *testing.T) {
	c, err := cid.Parse(CIDString)
	if err != nil {
		t.Error(err)
	}
	_, err = fromBytes(c.Bytes())
	if err == nil {
		t.Error("expected fromBytes to return error")
	}
}

func TestFromBytesWithCommit(t *testing.T) {
	_, bytes, _ := mbase.Decode(DocIDWithCommit)
	_, err := fromBytes(bytes)
	if err == nil {
		t.Error("expected fromBytes to return error")
	}
}

func TestFromBytesWith0Commit(t *testing.T) {
	_, bytes, _ := mbase.Decode(DocIDWith0Commit)
	_, err := fromBytes(bytes)
	if err == nil {
		t.Error("expected fromBytes to return error")
	}
}

func TestFromBytesRoundTrip(t *testing.T) {
	c, err := cid.Parse(CIDString)
	if err != nil {
		t.Error(err)
	}
	doc1 := DocID{0, c} // tile
	doc2, err := fromBytes(doc1.Bytes())
	if err != nil {
		t.Error(err)
	}
	observed := doc2.String()
	if observed != DocIDString {
		t.Error("expected doc id string to match")
	}
}

func TestFromString(t *testing.T) {
	docID, err := fromString(DocIDString)
	if err != nil {
		t.Error(err)
	}
	if docID.CID().String() != CIDString {
		t.Error("expected cid strings to match")
	}
	if docID.Type() != uint64(0) {
		t.Error("expected doc type to match")
	}
	observed := docID.String()
	if observed != DocIDString {
		t.Error("expected doc id string to match")
	}
}

func TestFromStringWithCommit(t *testing.T) {
	_, err := fromString(DocIDWithCommit)
	if err == nil {
		t.Error("expected fromString to return error")
	}
}

func TestFromStringWith0Commit(t *testing.T) {
	_, err := fromString(DocIDWith0Commit)
	if err == nil {
		t.Error("expected fromBytes to return error")
	}
}

func TestFromStringURL(t *testing.T) {
	docID, err := fromString(DocIDURL)
	if err != nil {
		t.Error(err)
	}
	if docID.CID().String() != CIDString {
		t.Error("expected cid strings to match")
	}
	if docID.Type() != uint64(0) {
		t.Error("expected doc type to match")
	}
	observed := docID.String()
	if observed != DocIDString {
		t.Error("expected doc id string to match")
	}
}

func TestFromStringLegacy(t *testing.T) {
	docID, err := fromString(DocIDLegacy)
	if err != nil {
		t.Error(err)
	}
	if docID.CID().String() != CIDString {
		t.Error("expected cid strings to match")
	}
	if docID.Type() != uint64(0) {
		t.Error("expected doc type to match")
	}
	observed := docID.String()
	if observed != DocIDString {
		t.Error("expected doc id string to match")
	}
}

func TestFromStringLegacyWithCommit(t *testing.T) {
	docID, err := fromString(DocIDWithCommitLegacy)
	if err != nil {
		t.Error(err)
	}
	if docID.CID().String() != CIDString {
		t.Error("expected cid strings to match")
	}
	if docID.Type() != uint64(0) {
		t.Error("expected doc type to match")
	}
	observed := docID.String()
	if observed != DocIDString {
		t.Error("expected doc id string to match")
	}
}

func TestFromStringLegacyWith0Commit(t *testing.T) {
	docID, err := fromString(DocIDWith0CommitLegacy)
	if err != nil {
		t.Error(err)
	}
	if docID.CID().String() != CIDString {
		t.Error("expected cid strings to match")
	}
	if docID.Type() != uint64(0) {
		t.Error("expected doc type to match")
	}
	observed := docID.String()
	if observed != DocIDString {
		t.Error("expected doc id string to match")
	}
}

func TestFromStringRoundTrip(t *testing.T) {
	c, err := cid.Parse(CIDString)
	if err != nil {
		t.Error(err)
	}
	doc1 := DocID{0, c} // tile
	doc2, err := fromString(doc1.String())
	if err != nil {
		t.Error(err)
	}
	observed := doc2.String()
	if observed != DocIDString {
		t.Error("expected doc id string to match")
	}
}

func TestAtCommit(t *testing.T) {
	baseCID, err := cid.Parse(CIDString)
	if err != nil {
		t.Error(err)
	}
	commitCID, err := cid.Parse(CommitCIDString)
	if err != nil {
		t.Error(err)
	}
	docID := DocID{0, baseCID} // tile

	// to genesis commit
	commitID, err := docID.AtCommit(cid.Undef)
	if err != nil {
		t.Error(err)
	}
	if !commitID.Commit().Equals(baseCID) {
		t.Error("expected cids to match")
	}

	// to latest commit
	commitID, err = docID.AtCommit(commitCID)
	if err != nil {
		t.Error(err)
	}
	if !commitID.Commit().Equals(commitCID) {
		t.Error("expected cids to match")
	}
}

// CommitID

func TestCreateCommitByParts(t *testing.T) {
	d := uint64(0)
	c, _ := cid.Parse(CIDString)
	docID := CommitID{DocID{d, c}, cid.Undef}
	if docID.CID().String() != CIDString {
		t.Error("expected cid strings to match")
	}
	if docID.Type() != d {
		t.Error("expected doc type to match")
	}
	observed := docID.String()
	if observed != DocIDWith0Commit {
		t.Error("expected doc id string to match")
	}
}

func TestCommitAtCommit(t *testing.T) {
	baseCID, err := cid.Parse(CIDString)
	if err != nil {
		t.Error(err)
	}
	commitCID, err := cid.Parse(CommitCIDString)
	if err != nil {
		t.Error(err)
	}
	docID := DocID{0, baseCID} // tile

	// to genesis commit
	commitID, err := docID.AtCommit(cid.Undef)
	if err != nil {
		t.Error(err)
	}
	if !commitID.Commit().Equals(baseCID) {
		t.Error("expected cids to match")
	}

	// to latest commit
	commitID, err = docID.AtCommit(commitCID)
	if err != nil {
		t.Error(err)
	}
	if !commitID.Commit().Equals(commitCID) {
		t.Error("expected cids to match")
	}
}

func TestCommitToString(t *testing.T) {
	c, err := cid.Parse(CIDString)
	if err != nil {
		t.Error(err)
	}
	doc := DocID{0, c} // tile
	commit, err := doc.AtCommit(cid.Undef)
	if err != nil {
		t.Error(err)
	}
	observed := commit.String()
	if observed != DocIDWith0Commit {
		t.Error("expected doc id string to match")
	}
	commitCID, err := cid.Parse(CommitCIDString)
	if err != nil {
		t.Error(err)
	}
	commit, err = doc.AtCommit(commitCID)
	if err != nil {
		t.Error(err)
	}
	observed = commit.String()
	if observed != DocIDWithCommit {
		t.Error("expected doc id string to match")
	}
}
