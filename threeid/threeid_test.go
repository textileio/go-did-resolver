// Package threeid provides tools for resolving the did:3 method format for
// the ceramic network:
// https://github.com/ceramicnetwork/CIP/blob/main/CIPs/CIP-79/CIP-79.md
// Copyright 2021 Textile
// Copyright 2021 Ceramic Network
package threeid

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	did "github.com/ockam-network/did"
	"github.com/textileio/go-did-resolver/resolver"
)

var (
	Fake3ID       = "did:3:k2t6wyfsu4pg0t2n4j8ms3s33xsgqjhtto04mvq8w5a2v5xo48idyz38l7ydki"
	FakeLegacy3ID = "did:3:bafyreiffkeeq4wq2htejqla2is5ognligi4lvjhwrpqpl2kazjdoecmugi"
)

type mockClient struct{}

func (client *mockClient) LoadDocument(docID DocIdentifier) (*DocState, error) {
	// Mimic how IDX sets the 3ID DID document
	str := `{
	"publicKeys": {
		"8VBUiUTCeHbTeTV": "zQ3shwsCgFanBax6UiaLu1oGvM7vhuqoW88VBUiUTCeHbTeTV",
		"6F7kmzdodfvUCo9": "z6LSfQabSbJzX8WAm1qdQcHCHTzVv8a2u6F7kmzdodfvUCo9"
  }
}`
	var content json.RawMessage
	doc := &DocState{
		Content: &content,
	}
	err := json.Unmarshal([]byte(str), &content)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

type mockClientWithIDX struct{}

func (client *mockClientWithIDX) LoadDocument(docID DocIdentifier) (*DocState, error) {
	// Mimic how IDX sets the 3ID DID document
	str := `{
	"publicKeys": {
		"8VBUiUTCeHbTeTV": "zQ3shwsCgFanBax6UiaLu1oGvM7vhuqoW88VBUiUTCeHbTeTV",
		"6F7kmzdodfvUCo9": "z6LSfQabSbJzX8WAm1qdQcHCHTzVv8a2u6F7kmzdodfvUCo9"
  },
	"idx": "ceramic://rootId"
}`
	var content json.RawMessage
	doc := &DocState{
		Content: &content,
	}
	err := json.Unmarshal([]byte(str), &content)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func TestBasicResolver(t *testing.T) {
	res := Resolver{&mockClient{}}
	if res.Method() != "3" {
		t.Error("expected method to be '3'")
	}
}

func TestResolveMockDocument(t *testing.T) {
	parsed, err := did.Parse(Fake3ID)
	if err != nil {
		t.Error(err)
	}
	res := &Resolver{&mockClient{}}
	observed, err := res.Resolve(Fake3ID, parsed, res)
	if err != nil {
		t.Error(err)
	}

	byteValue, err := ioutil.ReadFile("testdata/threeid.json")
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

func TestResolveMockDocumentWithService(t *testing.T) {
	parsed, err := did.Parse(Fake3ID)
	if err != nil {
		t.Error(err)
	}
	res := &Resolver{&mockClientWithIDX{}}
	observed, err := res.Resolve(Fake3ID, parsed, res)
	if err != nil {
		t.Error(err)
	}

	byteValue, err := ioutil.ReadFile("testdata/threeid.json")
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

func TestResolveRealDocument(t *testing.T) {
	// TODO: This test requires a locally running ceramic peer.
	// And it also requires this peer to have added the following hex-encoded
	// private key seed: "70FC119BB28CAA736BFCF91A6B73ED544E6745E37484EFA535BE357269ED9B02"
	parsed, err := did.Parse("did:3:kjzl6cwe1jw148t9pgxvoty45b02rztk3f9zaf45d5yxwarikwgrds8rcke8uuv")
	if err != nil {
		t.Error(err)
	}
	res := &Resolver{&HTTPClient{
		APIURL: DefaultHost + DefaultAPIPath,
	}}
	observed, err := res.Resolve(Fake3ID, parsed, res)
	if err != nil {
		t.Error(err)
	}

	byteValue, err := ioutil.ReadFile("testdata/live.json")
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
