package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/ockam-network/did"
	"github.com/textileio/go-did-resolver/resolver"
)

func init() {
	Client = &MockClient{}
}

var (
	// GetFunc fetches the mock client's `Get` func
	GetFunc func(url string) (*http.Response, error)
	// Key is a randomly generated Secp256k1VerificationKey2018
	Key = "f02d41c913f542ae671e6df17a456c7dd6a280479d491120629daf92130602cd135"
)

// MockClient is the mock client
type MockClient struct {
	Getter
}

// Get is the mock client's `Get` func
func (m *MockClient) Get(url string) (*http.Response, error) {
	return GetFunc(url)
}

func TestUnknownMethod(t *testing.T) {
	id := "did:borg:example.com"

	parsed, err := did.Parse(id)
	if err != nil {
		t.Fatal(err)
	}
	r := New()
	_, err = r.Resolve(id, parsed, r)
	if err == nil || err.Error() != "unknown did method: 'borg'" {
		t.Error("expected Resolve to return error")
	}
}

func TestFailOnGet(t *testing.T) {
	id := "did:web:example.com"

	parsed, err := did.Parse(id)
	if err != nil {
		t.Fatal(err)
	}
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 404,
		}, fmt.Errorf("expected error")
	}
	r := New()
	_, err = r.Resolve(id, parsed, r)
	if err == nil || err.Error() != "expected error" {
		t.Error("expected Resolve to return error")
	}
}

func TestResolveDocument(t *testing.T) {
	id := "did:web:example.com"
	expected := &resolver.Document{
		Context: []string{"https://w3id.org/did/v1"},
		ID:      id,
		VerificationMethod: []resolver.VerificationMethod{
			{
				ID:                 fmt.Sprintf("%s#owner", id),
				Type:               "Secp256k1VerificationKey2018",
				Controller:         id,
				PublicKeyMultibase: Key,
			},
		},
		Authentication: []resolver.VerificationMethod{
			{
				Type:      "Secp256k1SignatureAuthentication2018",
				PublicKey: fmt.Sprintf("%s#owner", id),
			},
		},
	}
	parsed, err := did.Parse(id)
	if err != nil {
		t.Fatal(err)
	}
	res := New()
	data, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(data)),
		}, nil
	}
	observed, err := res.Resolve(id, parsed, res)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(observed, expected) {
		t.Errorf("expected document to match: %s", observed.ID)
	}
}

func TestResolveLongDid(t *testing.T) {
	id := "did:web:example.com:user:alice"
	expected := &resolver.Document{
		Context: []string{"https://w3id.org/did/v1"},
		ID:      id,
		VerificationMethod: []resolver.VerificationMethod{
			{
				ID:                 fmt.Sprintf("%s#owner", id),
				Type:               "Secp256k1VerificationKey2018",
				Controller:         id,
				PublicKeyMultibase: Key,
			},
		},
		Authentication: []resolver.VerificationMethod{
			{
				Type:      "Secp256k1SignatureAuthentication2018",
				PublicKey: fmt.Sprintf("%s#owner", id),
			},
		},
	}
	parsed, err := did.Parse(id)
	if err != nil {
		t.Fatal(err)
	}
	res := New()
	data, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(data)),
		}, nil
	}
	observed, err := res.Resolve(id, parsed, res)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(observed, expected) {
		t.Errorf("expected document to match: %s", observed.ID)
	}
}

func TestFailOnStatus(t *testing.T) {
	id := "did:web:example.com"
	parsed, err := did.Parse(id)
	if err != nil {
		t.Fatal(err)
	}
	res := New()
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 404,
			Status:     "Not Found",
		}, nil
	}
	_, err = res.Resolve(id, parsed, res)
	if err == nil || err.Error() != "unable to resolve: 'Not Found'" {
		t.Error("expected Resolve to return error")
	}
}

func TestFailOnBadResponse(t *testing.T) {
	id := "did:web:example.com"
	parsed, err := did.Parse(id)
	if err != nil {
		t.Fatal(err)
	}
	res := New()
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte("bad json"))),
		}, nil
	}
	_, err = res.Resolve(id, parsed, res)
	if err == nil || !strings.Contains(err.Error(), "invalid character") {
		t.Error("expected Resolve to return error")
	}
}

func TestFailOnEmptyResponse(t *testing.T) {
	id := "did:web:example.com"
	parsed, err := did.Parse(id)
	if err != nil {
		t.Fatal(err)
	}
	res := New()
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
		}, nil
	}
	_, err = res.Resolve(id, parsed, res)
	if err == nil || err.Error() != "empty did document" {
		t.Error("expected Resolve to return error")
	}
}

func TestFailOnMismatch(t *testing.T) {
	id := "did:web:example.com"
	wrong := &resolver.Document{
		Context: []string{"https://w3id.org/did/v1"},
		ID:      "did:web:wrong.com",
		VerificationMethod: []resolver.VerificationMethod{
			{
				ID:                 fmt.Sprintf("%s#owner", id),
				Type:               "Secp256k1VerificationKey2018",
				Controller:         id,
				PublicKeyMultibase: Key,
			},
		},
		Authentication: []resolver.VerificationMethod{
			{
				Type:      "Secp256k1SignatureAuthentication2018",
				PublicKey: fmt.Sprintf("%s#owner", id),
			},
		},
	}
	parsed, err := did.Parse(id)
	if err != nil {
		t.Fatal(err)
	}
	res := New()
	data, err := json.Marshal(wrong)
	if err != nil {
		t.Fatal(err)
	}
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(data)),
		}, nil
	}
	_, err = res.Resolve(id, parsed, res)
	if err == nil || err.Error() != "id does not match requested did" {
		t.Error("expected Resolve to return error")
	}
}

func TestFailNoVerificationMethods(t *testing.T) {
	id := "did:web:example.com"
	wrong := &resolver.Document{
		Context: []string{"https://w3id.org/did/v1"},
		ID:      id,
		Authentication: []resolver.VerificationMethod{
			{
				Type:      "Secp256k1SignatureAuthentication2018",
				PublicKey: fmt.Sprintf("%s#owner", id),
			},
		},
	}
	parsed, err := did.Parse(id)
	if err != nil {
		t.Fatal(err)
	}
	res := New()
	data, err := json.Marshal(wrong)
	if err != nil {
		t.Fatal(err)
	}
	// Mock the Get function for our test
	GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(data)),
		}, nil
	}
	_, err = res.Resolve(id, parsed, res)
	if err == nil || err.Error() != "no verification methods" {
		t.Error("expected Resolve to return error")
	}
}
