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
	"context"
	"time"
)

type BaseClient interface {
	// Close closes all connections to the Encryptonize server.
	Close() error
	// SetToken sets the provided token as the authentication header.
	SetToken(token string)
	// GetTokenExpiration returns when the current token wil expire.
	GetTokenExpiration() time.Time
	// Version retrieves the version information of the Encryptonize service.
	Version() (*VersionResponse, error)
	// Health retrieves the current health status of the Encryptonize service.
	Health() (*HealthResponse, error)
	// LoginUser authenticates to the Encryptonize service with the given credentials and sets the
	// resulting access token for future calls. Call `LoginUser` again to switch to a different user.
	LoginUser(uid, password string) error
	// CreateUser creates a new Encryptonize user with the requested scopes.
	CreateUser(scopes []Scope) (*CreateUserResponse, error)
	// RemoveUser removes a user from the Encryptonize service.
	RemoveUser(uid string) error
	// CreateGroup creates a new Encryptonize group with the requested scopes.
	CreateGroup(scopes []Scope) (*CreateGroupResponse, error)
	// AddUserToGroup adds a user to a group.
	AddUserToGroup(uid, gid string) error
	// GetPermissions returns a list of IDs that have access to the requested object.
	GetPermissions(oid string) (*GetPermissionsResponse, error)
	// AddPermission grants permission for the group to the requested object.
	AddPermission(oid, gid string) error
	// RemovePermission removes permissions for the group to the requested object.
	RemovePermission(oid, gid string) error
}

type CoreClient interface {
	BaseClient
	// Encrypt encrypts the `plaintext` and tags both `plaintext` and `associatedData` returning the
	// resulting ciphertext.
	Encrypt(plaintext, associatedData []byte) (*EncryptResponse, error)
	// Decrypt decrypts a previously encrypted `ciphertext` and verifies the integrity of the `ciphertext`
	// and `associatedData`.
	Decrypt(objectID string, ciphertext, associatedData []byte) (*DecryptResponse, error)
}

type ObjectsClient interface {
	BaseClient
	// Store encrypts the `plaintext` and tags both `plaintext` and `associatedData` storing the
	// resulting ciphertext in the Encryptonize service.
	Store(plaintext, associatedData []byte) (*StoreResponse, error)
	// Retrieve decrypts a previously stored object returning the ciphertext.
	Retrieve(oid string) (*RetrieveResponse, error)
	// Update replaces the currently stored data of an object with the specified `plaintext` and
	// `associatedData`.
	Update(oid string, plaintext, associatedData []byte) error
	// Delete removes previously stored data from the Encryptonize service.
	Delete(oid string) error
}

// NewCoreClient creates a new Encryptonize Core client. Note that in order to call endpoints that require
// authentication, you need to call `LoginUser` first.
func NewCoreClient(ctx context.Context, endpoint, certPath string) (CoreClient, error) {
	client, err := newBaseClient(ctx, endpoint, certPath)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewObjectsClient creates a new Encryptonize Core client. Note that in order to call endpoints that require
// authentication, you need to call `LoginUser` first.
func NewObjectsClient(ctx context.Context, endpoint, certPath string) (ObjectsClient, error) {
	client, err := newBaseClient(ctx, endpoint, certPath)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewCoreClientWR creates a new Encryptonize Core client. In order to switch credentials to another user,
// use `LoginUser`.
func NewCoreClientWR(ctx context.Context, endpoint, certPath, uid, password string) (CoreClient, error) {
	client, err := newBaseClientWR(ctx, endpoint, certPath, uid, password)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewObjectsClientWR creates a new Encryptonize Objects client. In order to switch credentials to another user,
// use `LoginUser`.
func NewObjectsClientWR(ctx context.Context, endpoint, certPath, uid, password string) (ObjectsClient, error) {
	client, err := newBaseClientWR(ctx, endpoint, certPath, uid, password)
	if err != nil {
		return nil, err
	}

	return client, nil
}
