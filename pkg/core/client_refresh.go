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

package core

import (
	"context"
	"time"

	"github.com/cyber-crypt-com/encryptonize-client-go/pkg"
)

// ClientWR for making gRPC calls to the Encryptonize service while automatically refreshing the
// access token.
type ClientWR struct { //nolint:revive
	Client
	uid      string
	password string
}

// NewClientWR creates a new Encryptonize client. In order to switch credentials to another user,
// use `LoginUser`.
func NewClientWR(ctx context.Context, endpoint, certPath, uid, password string) (*ClientWR, error) {
	client, err := NewClient(ctx, endpoint, certPath)
	if err != nil {
		return nil, err
	}

	err = client.LoginUser(uid, password)
	if err != nil {
		return nil, err
	}

	return &ClientWR{
		Client:   *client,
		uid:      uid,
		password: password,
	}, nil
}

// withRefresh will refresh the token if it is about to expire, and then call `call`.
func (c *ClientWR) withRefresh(call func() error) error {
	// To avoid clock drift issues, refresh the token if it will expire within 1 minute.
	if time.Now().After(c.Client.GetTokenExpiration().Add(time.Duration(-1) * time.Minute)) {
		if err := c.Client.LoginUser(c.uid, c.password); err != nil {
			return err
		}
	}
	return call()
}

/////////////////////////////////////////////////////////////////////////
//                               Utility                               //
/////////////////////////////////////////////////////////////////////////

// Version retrieves the version information of the Encryptonize service.
func (c *ClientWR) Version() (*pkg.VersionResponse, error) {
	var response *pkg.VersionResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.Client.Version()
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Health retrieves the current health status of the Encryptonize service.
func (c *ClientWR) Health() (*pkg.HealthResponse, error) {
	var response *pkg.HealthResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.Client.Health()
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
func (c *ClientWR) LoginUser(uid, password string) error {
	err := c.Client.LoginUser(uid, password)
	if err != nil {
		return err
	}
	c.uid = uid
	c.password = password
	return nil
}

// CreateUser creates a new Encryptonize user with the requested scopes.
func (c *ClientWR) CreateUser(scopes []pkg.Scope) (*pkg.CreateUserResponse, error) {
	var response *pkg.CreateUserResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.Client.CreateUser(scopes)
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

// RemoveUser removes a user from the Encryptonize service.
func (c *ClientWR) RemoveUser(uid string) error {
	return c.withRefresh(func() error {
		return c.Client.RemoveUser(uid)
	})
}

// CreateGroup creates a new Encryptonize group with the requested scopes.
func (c *ClientWR) CreateGroup(scopes []pkg.Scope) (*pkg.CreateGroupResponse, error) {
	var response *pkg.CreateGroupResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.Client.CreateGroup(scopes)
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AddUserToGroup adds a user to a group.
func (c *ClientWR) AddUserToGroup(uid, gid string) error {
	return c.withRefresh(func() error {
		return c.Client.AddUserToGroup(uid, gid)
	})
}

// RemoveUserFromGroup removes a user from a group.
func (c *ClientWR) RemoveUserFromGroup(uid, gid string) error {
	return c.withRefresh(func() error {
		return c.Client.RemoveUserFromGroup(uid, gid)
	})
}

/////////////////////////////////////////////////////////////////////////
//                              Encryption                             //
/////////////////////////////////////////////////////////////////////////

// Encrypt encrypts the `plaintext` and tags both `plaintext` and `associatedData` returning the
// resulting ciphertext.
func (c *ClientWR) Encrypt(plaintext, associatedData []byte) (*pkg.EncryptResponse, error) {
	var response *pkg.EncryptResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.Client.Encrypt(plaintext, associatedData)
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Decrypt decrypts a previously encrypted `ciphertext` and verifies the integrity of the `ciphertext`
// and `associatedData`.
func (c *ClientWR) Decrypt(objectID string, ciphertext, associatedData []byte) (*pkg.DecryptResponse, error) {
	var response *pkg.DecryptResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.Client.Decrypt(objectID, ciphertext, associatedData)
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

/////////////////////////////////////////////////////////////////////////
//                             Permissions                             //
/////////////////////////////////////////////////////////////////////////

// GetPermissions returns a list of IDs that have access to the requested object.
func (c *ClientWR) GetPermissions(oid string) (*pkg.GetPermissionsResponse, error) {
	var response *pkg.GetPermissionsResponse
	err := c.withRefresh(func() error {
		var err error
		response, err = c.Client.GetPermissions(oid)
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AddPermission grants permission for the `target` to the requested object.
func (c *ClientWR) AddPermission(oid, target string) error {
	return c.withRefresh(func() error {
		return c.Client.AddPermission(oid, target)
	})
}

// RemovePermission removes permissions for the `target` to the requested object.
func (c *ClientWR) RemovePermission(oid, target string) error {
	return c.withRefresh(func() error {
		return c.Client.RemovePermission(oid, target)
	})
}
