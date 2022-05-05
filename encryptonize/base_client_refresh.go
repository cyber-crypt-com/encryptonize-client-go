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
	"time"
)

// baseClientWR for making gRPC calls to the Encryptonize service while automatically refreshing the
// access token.
type baseClientWR struct { //nolint:revive
	baseClient
	uid      string
	password string
}

// NewClientWR creates a new Encryptonize client. In order to switch credentials to another user,
// use `LoginUser`.
func newBaseClientWR(ctx context.Context, endpoint, certPath, uid, password string) (*baseClientWR, error) {
	client, err := newBaseClient(ctx, endpoint, certPath)
	if err != nil {
		return nil, err
	}

	err = client.LoginUser(uid, password)
	if err != nil {
		return nil, err
	}

	return &baseClientWR{
		baseClient: *client,
		uid:        uid,
		password:   password,
	}, nil
}

// withRefresh will refresh the token if it is about to expire, and then call `call`.
func (c *baseClientWR) withRefresh(call func() error) error {
	// To avoid clock drift issues, refresh the token if it will expire within 1 minute.
	if time.Now().After(c.tokenExpiration.Add(time.Duration(-1) * time.Minute)) {
		if err := c.baseClient.LoginUser(c.uid, c.password); err != nil {
			return err
		}
	}
	return call()
}

/////////////////////////////////////////////////////////////////////////
//                               Utility                               //
/////////////////////////////////////////////////////////////////////////

// Version retrieves the version information of the Encryptonize service.
func (c *baseClientWR) Version() (*VersionResponse, error) {
	var response *VersionResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.baseClient.Version()
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Health retrieves the current health status of the Encryptonize service.
func (c *baseClientWR) Health() (*HealthResponse, error) {
	var response *HealthResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.baseClient.Health()
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

/////////////////////////////////////////////////////////////////////////
//                           User Management                           //
/////////////////////////////////////////////////////////////////////////

// LoginUser authenticates to the Encryptonize service with the given credentials and sets the
// resulting access token for future calls. Call `LoginUser` again to switch to a different user.
func (c *baseClientWR) LoginUser(uid, password string) error {
	err := c.baseClient.LoginUser(uid, password)
	if err != nil {
		return err
	}
	c.uid = uid
	c.password = password
	return nil
}

// CreateUser creates a new Encryptonize user with the requested scopes.
func (c *baseClientWR) CreateUser(scopes []Scope) (*CreateUserResponse, error) {
	var response *CreateUserResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.baseClient.CreateUser(scopes)
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

// RemoveUser removes a user from the Encryptonize service.
func (c *baseClientWR) RemoveUser(uid string) error {
	return c.withRefresh(func() error {
		return c.baseClient.RemoveUser(uid)
	})
}

// CreateGroup creates a new Encryptonize group with the requested scopes.
func (c *baseClientWR) CreateGroup(scopes []Scope) (*CreateGroupResponse, error) {
	var response *CreateGroupResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.baseClient.CreateGroup(scopes)
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AddUserToGroup adds a user to a group.
func (c *baseClientWR) AddUserToGroup(uid, gid string) error {
	return c.withRefresh(func() error {
		return c.baseClient.AddUserToGroup(uid, gid)
	})
}

// RemoveUserFromGroup removes a user from a group.
func (c *baseClientWR) RemoveUserFromGroup(uid, gid string) error {
	return c.withRefresh(func() error {
		return c.baseClient.RemoveUserFromGroup(uid, gid)
	})
}

/////////////////////////////////////////////////////////////////////////
//                              Encryption                             //
/////////////////////////////////////////////////////////////////////////

// Encrypt encrypts the `plaintext` and tags both `plaintext` and `associatedData` returning the
// resulting ciphertext.
func (c *baseClientWR) Encrypt(plaintext, associatedData []byte) (*EncryptResponse, error) {
	var response *EncryptResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.baseClient.Encrypt(plaintext, associatedData)
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Decrypt decrypts a previously encrypted `ciphertext` and verifies the integrity of the `ciphertext`
// and `associatedData`.
func (c *baseClientWR) Decrypt(objectID string, ciphertext, associatedData []byte) (*DecryptResponse, error) {
	var response *DecryptResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.baseClient.Decrypt(objectID, ciphertext, associatedData)
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

/////////////////////////////////////////////////////////////////////////
//                               Storage                               //
/////////////////////////////////////////////////////////////////////////

// Store encrypts the `plaintext` and tags both `plaintext` and `associatedData` storing the
// resulting ciphertext in the Encryptonize service.
func (c *baseClientWR) Store(plaintext, associatedData []byte) (*StoreResponse, error) {
	var response *StoreResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.baseClient.Store(plaintext, associatedData)
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieve decrypts a previously stored object returning the ciphertext.
func (c *baseClientWR) Retrieve(oid string) (*RetrieveResponse, error) {
	var response *RetrieveResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.baseClient.Retrieve(oid)
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Update replaces the currently stored data of an object with the specified `plaintext` and
// `associatedData`.
func (c *baseClientWR) Update(oid string, plaintext, associatedData []byte) error {
	return c.withRefresh(func() error {
		return c.baseClient.Update(oid, plaintext, associatedData)
	})
}

// Delete removes previously stored data from the Encryptonize service.
func (c *baseClientWR) Delete(oid string) error {
	return c.withRefresh(func() error {
		return c.baseClient.Delete(oid)
	})
}

/////////////////////////////////////////////////////////////////////////
//                             Permissions                             //
/////////////////////////////////////////////////////////////////////////

// GetPermissions returns a list of IDs that have access to the requested object.
func (c *baseClientWR) GetPermissions(oid string) (*GetPermissionsResponse, error) {
	var response *GetPermissionsResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.baseClient.GetPermissions(oid)
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AddPermission grants permission for the `target` to the requested object.
func (c *baseClientWR) AddPermission(oid, target string) error {
	return c.withRefresh(func() error {
		return c.baseClient.AddPermission(oid, target)
	})
}

// RemovePermission removes permissions for the `target` to the requested object.
func (c *baseClientWR) RemovePermission(oid, target string) error {
	return c.withRefresh(func() error {
		return c.baseClient.RemovePermission(oid, target)
	})
}
