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

package d1

import (
	"context"
	"log"

	gpb "github.com/cybercryptio/d1-service-generic/protobuf"
	spb "github.com/cybercryptio/d1-service-storage/protobuf"
	"google.golang.org/grpc/metadata"
)

func ExampleNewGenericClient() {
	ctx := context.Background()

	// Create a new D1 Generic client, providing the hostname and a root CA certificate.
	client, err := NewGenericClient("localhost:9000", "./ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	// Login the user with their credentials.
	loginResponse, err := client.LoginUser(
		ctx,
		&gpb.LoginUserRequest{
			UserId:   uid,
			Password: password,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Set access token for future calls
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "bearer "+loginResponse.AccessToken)

	// Encrypt sensitive data.
	encryptResponse, err := client.Encrypt(
		ctx,
		&gpb.EncryptRequest{
			Plaintext:      []byte("secret data"),
			AssociatedData: []byte("metadata"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt the response.
	decryptResponse, err := client.Decrypt(
		ctx,
		&gpb.DecryptRequest{
			ObjectId:       encryptResponse.ObjectId,
			Ciphertext:     encryptResponse.Ciphertext,
			AssociatedData: encryptResponse.AssociatedData,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", decryptResponse)
}

func ExampleNewGenericClientWR() {
	ctx := context.Background()

	// Create a new D1 Generic client, providing the hostname, a root CA certificate, and user
	// credentials.
	client, err := NewGenericClientWR("localhost:9000", "./ca.crt", "user id", "password")
	if err != nil {
		log.Fatal(err)
	}

	// Encrypt sensitive data.
	encryptResponse, err := client.Encrypt(
		ctx,
		&gpb.EncryptRequest{
			Plaintext:      []byte("secret data"),
			AssociatedData: []byte("metadata"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt the response.
	decryptResponse, err := client.Decrypt(
		ctx,
		&gpb.DecryptRequest{
			ObjectId:       encryptResponse.ObjectId,
			Ciphertext:     encryptResponse.Ciphertext,
			AssociatedData: encryptResponse.AssociatedData,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", decryptResponse)
}

func ExampleNewStorageClient() {
	ctx := context.Background()

	// Create a new D1 Storage client, providing the hostname and a root CA certificate.
	client, err := NewStorageClient("localhost:9000", "./ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	// Login the user with their credentials.
	loginResponse, err := client.LoginUser(
		ctx,
		&gpb.LoginUserRequest{
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
	storeResponse, err := client.Store(
		ctx,
		&spb.StoreRequest{
			Plaintext:      []byte("secret data"),
			AssociatedData: []byte("metadata"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the stored data.
	retrieveResponse, err := client.Retrieve(
		ctx,
		&spb.RetrieveRequest{
			ObjectId: storeResponse.ObjectId,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", retrieveResponse)
}

func ExampleNewStorageClientWR() {
	ctx := context.Background()

	// Create a new D1 Storage client, providing the hostname, a root CA certificate, and
	// user credentials.
	client, err := NewStorageClientWR("localhost:9000", "./ca.crt", "user id", "password")
	if err != nil {
		log.Fatal(err)
	}

	// Store sensitive data in encrypted form.
	storeResponse, err := client.Store(
		ctx,
		&spb.StoreRequest{
			Plaintext:      []byte("secret data"),
			AssociatedData: []byte("metadata"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the stored data.
	retrieveResponse, err := client.Retrieve(
		ctx,
		&spb.RetrieveRequest{
			ObjectId: storeResponse.ObjectId,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", retrieveResponse)
}
