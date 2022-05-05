// Copyright 2020-2022 CYBERCRYPT

package keyserver

import (
	"testing"

	"context"
	"os"
)

var endpoint = "127.0.0.1:50051"
var certPath = ""
var kikID = ""

func TestMain(m *testing.M) {
	// Get test enpoint and cert path
	v, ok := os.LookupEnv("E2E_TEST_URL")
	if ok {
		endpoint = v
	}
	v, ok = os.LookupEnv("E2E_TEST_CERT")
	if ok {
		certPath = v
	}
	// Get kik ID of a key generated with script
	v, ok = os.LookupEnv("E2E_TEST_KIK_ID")
	if ok {
		kikID = v
	}

	os.Exit(m.Run())
}

func TestGetKeySet(t *testing.T) {
	ksClient, err := NewClient(context.Background(), endpoint, certPath)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	nonce := []byte("random")
	getKeySetResponse, err := ksClient.GetKeySet(kikID, nonce)
	if err != nil {
		t.Fatalf("Failed to retrieve a key set: %v", err)
	}

	if getKeySetResponse.Nonce == nil {
		t.Fatalf("Key set is not complete, nonce is missing: %v", err)
	}
	if getKeySetResponse.WrappedKeys == nil {
		t.Fatalf("Key set is not complete, kwp is missing: %v", err)
	}
}
