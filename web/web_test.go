package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ockam-network/did"
	"github.com/stretchr/testify/suite"
	"github.com/textileio/go-did-resolver/resolver"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) SetupSuite() {
	Client = &MockClient{}
}

var (
	// GetFunc fetches the mock client's `Get` func
	GetFunc func(url string) (*http.Response, error)
	// Key is a randomly generated EcdsaSecp256k1RecoveryMethod2020
	Key = "f02d41c913f542ae671e6df17a456c7dd6a280479d491120629daf92130602cd135"
	// Ethereum identity
	Identity = "0x2Cc31912B2b0f3075A87b3640923D45A26cef3Ee"
)

// MockClient is the mock client
type MockClient struct {
	Getter
}

// Get is the mock client's `Get` func
func (m *MockClient) Get(url string) (*http.Response, error) {
	return GetFunc(url)
}

func (suite *TestSuite) TestUnknownMethod() {
	id := "did:borg:example.com"

	parsed, err := did.Parse(id)
	suite.NoError(err)
	r := New()
	_, err = r.Resolve(id, parsed, r)
	suite.EqualError(err, "unknown did method: 'borg'")
}

func (suite *TestSuite) TestFailOnGet() {
	id := "did:web:example.com"

	parsed, err := did.Parse(id)
	suite.NoError(err)
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 404,
		}, fmt.Errorf("expected error")
	}
	r := New()
	_, err = r.Resolve(id, parsed, r)
	suite.EqualError(err, "expected error")
}

func (suite *TestSuite) TestResolveDocument() {
	id := "did:web:example.com"
	expected := &resolver.Document{
		Context: []string{"https://www.w3.org/ns/did/v1"},
		ID:      id,
		VerificationMethod: []resolver.VerificationMethod{
			{
				ID:                  fmt.Sprintf("%s#owner", id),
				Type:                "EcdsaSecp256k1RecoveryMethod2020",
				Controller:          id,
				BlockchainAccountID: Identity,
			},
		},
		Authentication: []string{fmt.Sprintf("%s#owner", id)},
	}
	parsed, err := did.Parse(id)
	suite.NoError(err)
	res := New()
	data, err := json.Marshal(expected)
	suite.NoError(err)
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(data)),
		}, nil
	}
	observed, err := res.Resolve(id, parsed, res)
	suite.NoError(err)
	suite.Equal(expected, observed)
}

func (suite *TestSuite) TestResolveLongDid() {
	id := "did:web:example.com:user:alice"
	expected := &resolver.Document{
		Context: []string{"https://www.w3.org/ns/did/v1"},
		ID:      id,
		VerificationMethod: []resolver.VerificationMethod{
			{
				ID:                 fmt.Sprintf("%s#owner", id),
				Type:               "EcdsaSecp256k1RecoveryMethod2020",
				Controller:         id,
				PublicKeyMultibase: Key,
			},
		},
	}
	parsed, err := did.Parse(id)
	suite.NoError(err)
	res := New()
	data, err := json.Marshal(expected)
	suite.NoError(err)
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(data)),
		}, nil
	}
	observed, err := res.Resolve(id, parsed, res)
	suite.NoError(err)
	suite.Equal(expected, observed)
}

func (suite *TestSuite) TestFailOnStatus() {
	id := "did:web:example.com"
	parsed, err := did.Parse(id)
	suite.NoError(err)
	res := New()
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 404,
			Status:     "Not Found",
		}, nil
	}
	_, err = res.Resolve(id, parsed, res)
	suite.EqualError(err, "unable to resolve: 'Not Found'")
}

func (suite *TestSuite) TestFailOnBadResponse() {
	id := "did:web:example.com"
	parsed, err := did.Parse(id)
	suite.NoError(err)
	res := New()
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte("bad json"))),
		}, nil
	}
	_, err = res.Resolve(id, parsed, res)
	suite.EqualError(err, "invalid character 'b' looking for beginning of value")
}

func (suite *TestSuite) TestFailOnEmptyResponse() {
	id := "did:web:example.com"
	parsed, err := did.Parse(id)
	suite.NoError(err)
	res := New()
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
		}, nil
	}
	_, err = res.Resolve(id, parsed, res)
	suite.EqualError(err, "empty did document")
}

func (suite *TestSuite) TestFailOnMismatch() {
	id := "did:web:example.com"
	wrong := &resolver.Document{
		Context: []string{"https://www.w3.org/ns/did/v1"},
		ID:      "did:web:wrong.com",
		VerificationMethod: []resolver.VerificationMethod{
			{
				ID:                 fmt.Sprintf("%s#owner", id),
				Type:               "EcdsaSecp256k1RecoveryMethod2020",
				Controller:         id,
				PublicKeyMultibase: Key,
			},
		},
	}
	parsed, err := did.Parse(id)
	suite.NoError(err)
	res := New()
	data, err := json.Marshal(wrong)
	suite.NoError(err)
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(data)),
		}, nil
	}
	_, err = res.Resolve(id, parsed, res)
	suite.EqualError(err, "id does not match requested did")
}

func (suite *TestSuite) TestFailNoVerificationMethods() {
	id := "did:web:example.com"
	wrong := &resolver.Document{
		Context: []string{"https://www.w3.org/ns/did/v1"},
		ID:      id,
	}
	parsed, err := did.Parse(id)
	suite.NoError(err)
	res := New()
	data, err := json.Marshal(wrong)
	suite.NoError(err)
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(data)),
		}, nil
	}
	_, err = res.Resolve(id, parsed, res)
	suite.EqualError(err, "no verification methods")
}

func (suite *TestSuite) TestDocWithPortDid() {
	id := "did:web:localhost:8443"
	parsed, err := did.Parse(id)
	suite.NoError(err)
	res := New()
	suite.NoError(err)
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		suite.Equal("https://localhost:8443/.well-known/did.json", url)
		return &http.Response{
			StatusCode: 200,
		}, nil
	}
	res.Resolve(id, parsed, res)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
