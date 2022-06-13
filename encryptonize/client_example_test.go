// Copyright 2020-2022 CYBERCRYPT

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
