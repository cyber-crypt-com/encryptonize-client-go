// Copyright 2020-2022 CYBERCRYPT
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encryptonize

import (
	"context"
	"log"
)

func ExampleNewCoreClient() {
	// Create a new Encryptonize Core client, providing the hostname and a root CA certificate.
	client, err := NewCoreClient(context.Background(), "localhost:9000", "./ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	// Login the user with their credentials.
	err = client.LoginUser("user id", "password")
	if err != nil {
		log.Fatal(err)
	}

	// Encrypt sensitive data.
	plaintext, associatedData := []byte("secret data"), []byte("metadata")
	encrypted, err := client.Encrypt(plaintext, associatedData)
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt the response.
	decrypted, err := client.Decrypt(encrypted.ObjectID, encrypted.Ciphertext, encrypted.AssociatedData)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", decrypted)
}

func ExampleNewCoreClientWR() {
	// Create a new Encryptonize Core client, providing the hostname, a root CA certificate, and user
	// credentials.
	client, err := NewCoreClientWR(context.Background(), "localhost:9000", "./ca.crt", "user id", "password")
	if err != nil {
		log.Fatal(err)
	}

	// Encrypt sensitive data.
	plaintext, associatedData := []byte("secret data"), []byte("metadata")
	encrypted, err := client.Encrypt(plaintext, associatedData)
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt the response.
	decrypted, err := client.Decrypt(encrypted.ObjectID, encrypted.Ciphertext, encrypted.AssociatedData)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", decrypted)
}

func ExampleNewObjectsClient() {
	// Create a new Encryptonize Objects client, providing the hostname and a root CA certificate.
	client, err := NewObjectsClient(context.Background(), "localhost:9000", "./ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	// Login the user with their credentials.
	err = client.LoginUser("user id", "password")
	if err != nil {
		log.Fatal(err)
	}

	// Store sensitive data in encrypted form.
	plaintext, associatedData := []byte("secret data"), []byte("metadata")
	response, err := client.Store(plaintext, associatedData)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the stored data.
	decrypted, err := client.Retrieve(response.ObjectID)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", decrypted)
}

func ExampleNewObjectsClientWR() {
	// Create a new Encryptonize Objects client, providing the hostname, a root CA certificate, and
	// user credentials.
	client, err := NewObjectsClientWR(context.Background(), "localhost:9000", "./ca.crt", "user id", "password")
	if err != nil {
		log.Fatal(err)
	}

	// Store sensitive data in encrypted form.
	plaintext, associatedData := []byte("secret data"), []byte("metadata")
	response, err := client.Store(plaintext, associatedData)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the stored data.
	decrypted, err := client.Retrieve(response.ObjectID)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", decrypted)
}
