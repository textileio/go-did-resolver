// Package threeid provides tools for resolving the did:3 method format for
// the ceramic network:
// https://github.com/ceramicnetwork/CIP/blob/main/CIPs/CIP-79/CIP-79.md
// Copyright 2021 Textile
// Copyright 2021 Ceramic Network
package threeid

import (
	"fmt"
	"testing"
)

// TODO: More tests!

func TestBasicClient(t *testing.T) {
	client := HTTPClient{
		DefaultHost + DefaultAPIPath,
	}
	docID, err := fromString("ceramic://kjzl6cwe1jw14aa2ugzr1zumd8f2wtyl986skkv8kimbhxf39ajwek6gqrnxvfa")
	if err != nil {
		t.Error(err)
	}
	resp, err := client.Load(docID)
	if err != nil {
		t.Error(err)
	}
	if docID.String() != resp.DocID {
		t.Error("expected doc ids to match")
	}
	content := fmt.Sprintf("%s", *resp.State.Content)
	expected := "{\"Foo\":\"Bar\"}"
	if content != expected {
		t.Error("expected content to match")
	}
}
