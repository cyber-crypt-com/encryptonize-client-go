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

//go:build blob
// +build blob

package grpce2e

import (
	"testing"

	"context"

	"google.golang.org/grpc/codes"

	coreclient "github.com/cyber-crypt-com/encryptonize-core/client"
)

// Test that unauthorized users cannot perform actions on objects
func TestBlobUnauthorizedAccessToObject(t *testing.T) {
	client, err := coreclient.NewClient(context.Background(), endpoint, certPath)
	failOnError("Could not create client", err, t)
	defer client.Close()

	err = client.LoginUser(uid, pwd)
	failOnError("Could not log in user", err, t)

	plaintext := []byte("foo")
	associatedData := []byte("bar")

	// Store an object
	storeResponse, err := client.Store(plaintext, associatedData)
	failOnError("Store operation failed", err, t)
	oidStored := storeResponse.ObjectID

	// Create an unauthorized user
	createUserResponse, err := client.CreateUser(protoUserScopes)
	failOnError("Create user request failed", err, t)

	err = client.LoginUser(createUserResponse.UserID, createUserResponse.Password)
	failOnError("Could not log in user", err, t)

	// Try to use endpoints that require authorization
	_, err = client.Retrieve(oidStored)
	failOnSuccess("Unauthorized user retrieved object", err, t)
	checkStatusCode(err, codes.PermissionDenied, t)

	err = client.Update(oidStored, plaintext, associatedData)
	failOnSuccess("Unauthorized user updated object", err, t)
	checkStatusCode(err, codes.PermissionDenied, t)

	err = client.Delete(oidStored)
	failOnSuccess("Unauthorized user deleted object", err, t)
	checkStatusCode(err, codes.PermissionDenied, t)

	err = client.Update(oidStored, plaintext, associatedData)
	failOnSuccess("Unauthorized user updated object", err, t)
	checkStatusCode(err, codes.PermissionDenied, t)

	_, err = client.GetPermissions(oidStored)
	failOnSuccess("Unauthorized user got permissions", err, t)
	checkStatusCode(err, codes.PermissionDenied, t)

	err = client.AddPermission(oidStored, uid)
	failOnSuccess("Unauthorized user added permission", err, t)
	checkStatusCode(err, codes.PermissionDenied, t)

	err = client.RemovePermission(oidStored, uid)
	failOnSuccess("Unauthorized user removed permission", err, t)
	checkStatusCode(err, codes.PermissionDenied, t)
}

func TestBlobUnauthorizedToRead(t *testing.T) {
	client, err := coreclient.NewClient(context.Background(), endpoint, certPath)
	failOnError("Could not create client", err, t)
	defer client.Close()

	err = client.LoginUser(uid, pwd)
	failOnError("Could not log in user", err, t)

	var scopes = []coreclient.Scope{
		coreclient.ScopeCreate,
		coreclient.ScopeIndex,
		coreclient.ScopeObjectPermissions,
		coreclient.ScopeUserManagement,
		coreclient.ScopeUpdate,
		coreclient.ScopeDelete,
	}

	createUserResponse, err := client.CreateUser(scopes)
	failOnError("Create user request failed", err, t)

	err = client.LoginUser(createUserResponse.UserID, createUserResponse.Password)
	failOnError("Could not log in user", err, t)

	plaintext := []byte("foo")
	associatedData := []byte("bar")

	storeResponse, err := client.Store(plaintext, associatedData)
	failOnError("Store operation failed", err, t)

	_, err = client.Retrieve(storeResponse.ObjectID)
	failOnSuccess("User should not be able to retrieve object without READ scope", err, t)
	checkStatusCode(err, codes.PermissionDenied, t)
}

func TestBlobUnauthorizedToCreate(t *testing.T) {
	client, err := coreclient.NewClient(context.Background(), endpoint, certPath)
	failOnError("Could not create client", err, t)
	defer client.Close()

	err = client.LoginUser(uid, pwd)
	failOnError("Could not log in user", err, t)

	var scopes = []coreclient.Scope{
		coreclient.ScopeRead,
		coreclient.ScopeIndex,
		coreclient.ScopeObjectPermissions,
		coreclient.ScopeUserManagement,
		coreclient.ScopeUpdate,
		coreclient.ScopeDelete,
	}

	createUserResponse, err := client.CreateUser(scopes)
	failOnError("Create user request failed", err, t)

	err = client.LoginUser(createUserResponse.UserID, createUserResponse.Password)
	failOnError("Could not log in user", err, t)

	plaintext := []byte("foo")
	associatedData := []byte("bar")

	_, err = client.Store(plaintext, associatedData)
	failOnSuccess("User should not be able to store object without CREATE scope", err, t)
	checkStatusCode(err, codes.PermissionDenied, t)
}

func TestBlobUnauthorizedToGetPermissions(t *testing.T) {
	client, err := coreclient.NewClient(context.Background(), endpoint, certPath)
	failOnError("Could not create client", err, t)
	defer client.Close()

	err = client.LoginUser(uid, pwd)
	failOnError("Could not log in user", err, t)

	var scopes = []coreclient.Scope{
		coreclient.ScopeRead,
		coreclient.ScopeCreate,
		coreclient.ScopeObjectPermissions,
		coreclient.ScopeUserManagement,
		coreclient.ScopeUpdate,
		coreclient.ScopeDelete,
	}

	createUserResponse, err := client.CreateUser(scopes)
	failOnError("Create user request failed", err, t)

	err = client.LoginUser(createUserResponse.UserID, createUserResponse.Password)
	failOnError("Could not log in user", err, t)

	storeResponse, err := client.Store([]byte("foo"), []byte("bar"))
	failOnError("Store operation failed", err, t)
	oidStored := storeResponse.ObjectID

	_, err = client.GetPermissions(oidStored)
	failOnSuccess("User should not be able to get permission without INDEX scope", err, t)
	checkStatusCode(err, codes.PermissionDenied, t)
}

func TestBlobUnauthorizedToManagePermissions(t *testing.T) {
	client, err := coreclient.NewClient(context.Background(), endpoint, certPath)
	failOnError("Could not create client", err, t)
	defer client.Close()

	err = client.LoginUser(uid, pwd)
	failOnError("Could not log in user", err, t)

	var scopes = []coreclient.Scope{
		coreclient.ScopeRead,
		coreclient.ScopeCreate,
		coreclient.ScopeIndex,
		coreclient.ScopeUserManagement,
		coreclient.ScopeUpdate,
		coreclient.ScopeDelete,
	}

	createUserResponse, err := client.CreateUser(scopes)
	failOnError("Create user request failed", err, t)

	err = client.LoginUser(createUserResponse.UserID, createUserResponse.Password)
	failOnError("Could not log in user", err, t)

	storeResponse, err := client.Store([]byte("foo"), []byte("bar"))
	failOnError("Store operation failed", err, t)
	oidStored := storeResponse.ObjectID

	err = client.AddPermission(oidStored, uid)
	failOnSuccess("User should not be able to add permission without OBJECTPERMISSIONS scope", err, t)
	checkStatusCode(err, codes.PermissionDenied, t)

	err = client.RemovePermission(oidStored, uid)
	failOnSuccess("User should not be able to remove permission without OBJECTPERMISSIONS scope", err, t)
	checkStatusCode(err, codes.PermissionDenied, t)
}

func TestBlobUnauthorizedToManageUsers(t *testing.T) {
	client, err := coreclient.NewClient(context.Background(), endpoint, certPath)
	failOnError("Could not create client", err, t)
	defer client.Close()

	err = client.LoginUser(uid, pwd)
	failOnError("Could not log in user", err, t)

	var scopes = []coreclient.Scope{
		coreclient.ScopeRead,
		coreclient.ScopeCreate,
		coreclient.ScopeIndex,
		coreclient.ScopeObjectPermissions,
		coreclient.ScopeUpdate,
		coreclient.ScopeDelete,
	}

	createUserResponse, err := client.CreateUser(scopes)
	failOnError("Create user request failed", err, t)

	err = client.LoginUser(createUserResponse.UserID, createUserResponse.Password)
	failOnError("Could not log in user", err, t)

	_, err = client.CreateUser(protoUserScopes)
	failOnSuccess("User should not be able to create user without USERMANAGEMENT scope", err, t)
	checkStatusCode(err, codes.PermissionDenied, t)
}

func TestBlobUnauthorizedToUpdateAndDelete(t *testing.T) {
	client, err := coreclient.NewClient(context.Background(), endpoint, certPath)
	failOnError("Could not create client", err, t)
	defer client.Close()

	err = client.LoginUser(uid, pwd)
	failOnError("Could not log in user", err, t)

	var scopes = []coreclient.Scope{
		coreclient.ScopeRead,
		coreclient.ScopeCreate,
		coreclient.ScopeIndex,
		coreclient.ScopeObjectPermissions,
		coreclient.ScopeUserManagement,
	}

	createUserResponse, err := client.CreateUser(scopes)
	failOnError("Create user request failed", err, t)

	err = client.LoginUser(createUserResponse.UserID, createUserResponse.Password)
	failOnError("Could not log in user", err, t)

	storeResponse, err := client.Store([]byte("foo"), []byte("bar"))
	failOnError("Store operation failed", err, t)
	oidStored := storeResponse.ObjectID

	err = client.Update(oidStored, []byte("new_foo"), []byte("new_bar"))
	failOnSuccess("User should not be able to delete object without UPDATE scope", err, t)
	checkStatusCode(err, codes.PermissionDenied, t)

	err = client.Delete(oidStored)
	failOnSuccess("User should not be able to delete object without DELETE scope", err, t)
	checkStatusCode(err, codes.PermissionDenied, t)
}
