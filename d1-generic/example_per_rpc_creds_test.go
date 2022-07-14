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

package client_test

import (
	"context"
	"fmt"
	"log"

	client "github.com/cybercryptio/d1-client-go/d1-generic"
	pbgeneric "github.com/cybercryptio/d1-client-go/d1-generic/protobuf/generic"
)

func Example_withPerRPCCredentials() {
	// Create a new D1 Generic client providing the hostname, and optionally, the client connection level and per RPC credentials.
	client, err := client.NewGenericClient(endpoint,
		client.WithTransportCredentials(creds),
		client.WithPerRPCCredentials(
			client.NewStandalonePerRPCToken(endpoint, uid, password, creds),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Encrypt sensitive data.
	encryptResponse, err := client.Generic.Encrypt(
		ctx,
		&pbgeneric.EncryptRequest{
			Plaintext:      []byte("secret data"),
			AssociatedData: []byte("metadata"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt the response.
	decryptResponse, err := client.Generic.Decrypt(
		ctx,
		&pbgeneric.DecryptRequest{
			ObjectId:       encryptResponse.ObjectId,
			Ciphertext:     encryptResponse.Ciphertext,
			AssociatedData: encryptResponse.AssociatedData,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("plaintext:%q associated_data:%q",
		decryptResponse.Plaintext,
		decryptResponse.AssociatedData,
	)
	// Output: plaintext:"secret data" associated_data:"metadata"
}
