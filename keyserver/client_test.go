// Copyright 2020-2022 CYBERCRYPT

package keyserver

import (
	"testing"

	"bytes"
	"context"
	"os"

	"github.com/gofrs/uuid"

	"github.com/cybercryptio/k1/service"
)

var endpoint = "localhost:50051"
var certPath = ""
var kikID = uuid.Nil
var kik = ""

func TestMain(m *testing.M) {
	// Get test enpoint and cert path
	if v, ok := os.LookupEnv("E2E_TEST_URL"); ok {
		endpoint = v
	}
	if v, ok := os.LookupEnv("E2E_TEST_CERT"); ok {
		certPath = v
	}
	// Get kik ID of a key generated with script
	if v, ok := os.LookupEnv("E2E_TEST_KIK_ID"); ok {
		kikID = uuid.Must(uuid.FromString(v))
	}
	if v, ok := os.LookupEnv("E2E_TEST_KIK"); ok {
		kik = v
	}

	os.Exit(m.Run())
}

func TestGetKeySet(t *testing.T) {
	ksClient, err := NewClient(context.Background(), endpoint, certPath, kik, kikID)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	keys, err := ksClient.GetKeys()
	if err != nil {
		t.Fatalf("Failed to retrieve a key set: %v", err)
	}

	if len(keys.KEK) != service.KeySize {
		t.Fatalf("Key is the wrong size")
	}
	if len(keys.AEK) != service.KeySize {
		t.Fatalf("Key is the wrong size")
	}
	if len(keys.TEK) != service.KeySize {
		t.Fatalf("Key is the wrong size")
	}
	if len(keys.IEK) != service.KeySize {
		t.Fatalf("Key is the wrong size")
	}

	if bytes.Equal(keys.KEK, make([]byte, service.KeySize)) {
		t.Fatalf("Key is not initialized")
	}
	if bytes.Equal(keys.AEK, make([]byte, service.KeySize)) {
		t.Fatalf("Key is not initialized")
	}
	if bytes.Equal(keys.TEK, make([]byte, service.KeySize)) {
		t.Fatalf("Key is not initialized")
	}
	if bytes.Equal(keys.IEK, make([]byte, service.KeySize)) {
		t.Fatalf("Key is not initialized")
	}
}
