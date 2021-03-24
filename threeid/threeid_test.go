// Package threeid provides tools for resolving the did:3 method format for
// the ceramic network:
// https://github.com/ceramicnetwork/CIP/blob/main/CIPs/CIP-79/CIP-79.md
// Copyright 2021 Textile
// Copyright 2021 Ceramic Network
package threeid

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	did "github.com/ockam-network/did"
	"github.com/stretchr/testify/suite"
	"github.com/textileio/go-did-resolver/resolver"
)

type TestSuite struct {
	suite.Suite
}

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

func (suite *TestSuite) TestInterpretIdentifierKey() {
	res := Resolver{&mockClient{}}
	suite.Equal(res.Method(), "3", "expected method to be '3'")
}

func (suite *TestSuite) TestResolveMockDocument() {
	parsed, err := did.Parse(Fake3ID)
	suite.NoError(err)
	res := &Resolver{&mockClient{}}
	observed, err := res.Resolve(Fake3ID, parsed, res)
	suite.NoError(err)

	byteValue, err := ioutil.ReadFile("testdata/threeid.json")
	suite.NoError(err)
	var expected resolver.Document
	err = json.Unmarshal(byteValue, &expected)
	suite.NoError(err)

	suite.Equal(&expected, observed)
}

func (suite *TestSuite) TestResolveMockDocumentWithService() {
	parsed, err := did.Parse(Fake3ID)
	suite.NoError(err)
	res := &Resolver{&mockClientWithIDX{}}
	observed, err := res.Resolve(Fake3ID, parsed, res)
	suite.NoError(err)

	byteValue, err := ioutil.ReadFile("testdata/threeid.json")
	suite.NoError(err)
	var expected resolver.Document
	err = json.Unmarshal(byteValue, &expected)
	suite.NoError(err)

	suite.Equal(&expected, observed)
}

func (suite *TestSuite) DontTestResolveRealDocument() {
	// TODO: This test requires a locally running ceramic peer.
	// And it also requires this peer to have added the following hex-encoded
	// private key seed: "70FC119BB28CAA736BFCF91A6B73ED544E6745E37484EFA535BE357269ED9B02"
	parsed, err := did.Parse("did:3:kjzl6cwe1jw148t9pgxvoty45b02rztk3f9zaf45d5yxwarikwgrds8rcke8uuv")
	suite.NoError(err)
	res := &Resolver{&HTTPClient{
		APIURL: DefaultHost + DefaultAPIPath,
	}}
	observed, err := res.Resolve(Fake3ID, parsed, res)
	suite.NoError(err)

	byteValue, err := ioutil.ReadFile("testdata/live.json")
	suite.NoError(err)
	var expected resolver.Document
	err = json.Unmarshal(byteValue, &expected)
	suite.NoError(err)

	suite.Equal(&expected, observed)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
