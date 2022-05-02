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

package client

import (
	"testing"
	"time"

	"context"
)

func TestCoreUtilityWR(t *testing.T) {
	c, err := NewClientWR(context.Background(), endpoint, certPath, uid, password)
	failOnError("NewClientWR failed", err, t)
	defer c.Close()

	_, err = c.Health()
	failOnError("Health failed", err, t)
	_, err = c.Version()
	failOnError("Version failed", err, t)
}

func TestCoreUserManagementWR(t *testing.T) {
	c, err := NewClientWR(context.Background(), endpoint, certPath, uid, password)
	failOnError("NewClientWR failed", err, t)
	defer c.Close()

	createUserResponse, err := c.CreateUser(scopes)
	failOnError("CreateUser failed", err, t)

	createGroupResponse, err := c.CreateGroup(scopes)
	failOnError("CreateGroup failed", err, t)

	err = c.AddUserToGroup(createUserResponse.UserID, createGroupResponse.GroupID)
	failOnError("AddUserToGroup failed", err, t)

	err = c.RemoveUserFromGroup(createUserResponse.UserID, createGroupResponse.GroupID)
	failOnError("RemoveUserFromGroup failed", err, t)

	err = c.RemoveUser(createUserResponse.UserID)
	failOnError("RemoveUser failed", err, t)
}

func TestEncryptWR(t *testing.T) {
	c, err := NewClientWR(context.Background(), endpoint, certPath, uid, password)
	failOnError("NewClientWR failed", err, t)
	defer c.Close()

	createUserResponse, err := c.CreateUser(scopes)
	failOnError("CreateUser failed", err, t)
	err = c.LoginUser(createUserResponse.UserID, createUserResponse.Password)
	failOnError("LoginUser failed", err, t)

	plaintext := []byte("foo")
	associatedData := []byte("bar")
	encryptResponse, err := c.Encrypt(plaintext, associatedData)
	failOnError("Encrypt failed", err, t)

	decryptResponse, err := c.Decrypt(encryptResponse.ObjectID, encryptResponse.Ciphertext, encryptResponse.AssociatedData)
	failOnError("Decrypt failed", err, t)
	if string(decryptResponse.Plaintext) != string(plaintext) {
		t.Fatal("Decryption returned wrong plaintext")
	}
	if string(decryptResponse.AssociatedData) != string(associatedData) {
		t.Fatal("Decryption returned wrong data")
	}
}

func TestObjectsStoreWR(t *testing.T) {
	c, err := NewClientWR(context.Background(), endpoint, certPath, uid, password)
	failOnError("NewClientWR failed", err, t)
	defer c.Close()

	createUserResponse, err := c.CreateUser(scopes)
	failOnError("CreateUser failed", err, t)
	err = c.LoginUser(createUserResponse.UserID, createUserResponse.Password)
	failOnError("LoginUser failed", err, t)

	plaintext := []byte("foo")
	associatedData := []byte("bar")
	storeResponse, err := c.Store(plaintext, associatedData)
	failOnError("Store failed", err, t)

	retrieveResponse, err := c.Retrieve(storeResponse.ObjectID)
	failOnError("Retrieve failed", err, t)
	if string(retrieveResponse.Plaintext) != string(plaintext) {
		t.Fatal("Decryption returned wrong plaintext")
	}
	if string(retrieveResponse.AssociatedData) != string(associatedData) {
		t.Fatal("Decryption returned wrong data")
	}

	plaintext = []byte("foobar")
	associatedData = []byte("barbaz")
	err = c.Update(storeResponse.ObjectID, plaintext, associatedData)
	failOnError("Update failed", err, t)

	err = c.Delete(storeResponse.ObjectID)
	failOnError("Delete failed", err, t)
}

func TestObjectsPermissionsWR(t *testing.T) {
	c, err := NewClientWR(context.Background(), endpoint, certPath, uid, password)
	failOnError("NewClientWR failed", err, t)
	defer c.Close()

	createUserResponse, err := c.CreateUser(scopes)
	failOnError("CreateUser failed", err, t)
	err = c.LoginUser(createUserResponse.UserID, createUserResponse.Password)
	failOnError("LoginUser failed", err, t)

	plaintext := []byte("foo")
	associatedData := []byte("bar")
	storeResponse, err := c.Store(plaintext, associatedData)
	failOnError("Store failed", err, t)

	err = c.AddPermission(storeResponse.ObjectID, createUserResponse.UserID)
	failOnError("AddPermission failed", err, t)

	_, err = c.GetPermissions(storeResponse.ObjectID)
	failOnError("GetPermissions failed", err, t)

	err = c.RemovePermission(storeResponse.ObjectID, createUserResponse.UserID)
	failOnError("RemovePermission failed", err, t)
}

func TestCoreTokenRefreshWR(t *testing.T) {
	c, err := NewClientWR(context.Background(), endpoint, certPath, uid, password)
	failOnError("NewClientWR failed", err, t)
	defer c.Close()

	createUserResponse, err := c.CreateUser(scopes)
	failOnError("CreateUser failed", err, t)
	err = c.LoginUser(createUserResponse.UserID, createUserResponse.Password)
	failOnError("LoginUser failed", err, t)

	// Make sure logic refreshing token is triggered and clear token
	// to see error if token is not refreshed
	c.tokenExpiration = time.Now().Add(time.Duration(-1) * time.Hour)
	c.authHeader = nil

	_, err = c.CreateUser(scopes)
	failOnError("CreateUser failed", err, t)
}
