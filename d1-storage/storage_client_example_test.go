// Copyright 2022 CYBERCRYPT
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc/metadata"

	pbauthn "github.com/cybercryptio/d1-client-go/d1-generic/protobuf/authn"
	pbstorage "github.com/cybercryptio/d1-client-go/d1-storage/protobuf/storage"
)

var endpoint = os.Getenv("D1_ENDPOINT")
var uid = os.Getenv("D1_UID")
var password = os.Getenv("D1_PASS")
var certPath = os.Getenv("D1_CERT")

func ExampleNewStorageClient() {
	ctx := context.Background()

	// Create a new D1 Storage client, providing the hostname and a root CA certificate.
	client, err := NewStorageClient(endpoint, certPath)
	if err != nil {
		log.Fatal(err)
	}

	// Login the user with their credentials.
	loginResponse, err := client.Authn.LoginUser(
		ctx,
		&pbauthn.LoginUserRequest{
			UserId:   uid,
			Password: password,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Set access token for future calls
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "bearer "+loginResponse.AccessToken)

	// Store sensitive data in encrypted form.
	storeResponse, err := client.Storage.Store(
		ctx,
		&pbstorage.StoreRequest{
			Plaintext:      []byte("secret data"),
			AssociatedData: []byte("metadata"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the stored data.
	retrieveResponse, err := client.Storage.Retrieve(
		ctx,
		&pbstorage.RetrieveRequest{
			ObjectId: storeResponse.ObjectId,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("plaintext:%q associated_data:%q",
		retrieveResponse.Plaintext,
		retrieveResponse.AssociatedData,
	)
	// Output: plaintext:"secret data" associated_data:"metadata"
}

func ExampleNewStorageClientWR() {
	ctx := context.Background()

	// Create a new D1 Storage client, providing the hostname, a root CA certificate, and
	// user credentials.
	client, err := NewStorageClientWR(endpoint, certPath, uid, password)
	if err != nil {
		log.Fatal(err)
	}

	// Store sensitive data in encrypted form.
	storeResponse, err := client.Storage.Store(
		ctx,
		&pbstorage.StoreRequest{
			Plaintext:      []byte("secret data"),
			AssociatedData: []byte("metadata"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the stored data.
	retrieveResponse, err := client.Storage.Retrieve(
		ctx,
		&pbstorage.RetrieveRequest{
			ObjectId: storeResponse.ObjectId,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("plaintext:%q associated_data:%q",
		retrieveResponse.Plaintext,
		retrieveResponse.AssociatedData,
	)
	// Output: plaintext:"secret data" associated_data:"metadata"
}
