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

package objects

import (
	"context"
	"time"

	"github.com/cyber-crypt-com/encryptonize-client-go/internal"
	"github.com/cyber-crypt-com/encryptonize-client-go/pkg"
)

// Client for making gRPC calls to the Encryptonize service.
type Client struct {
	client *internal.Client
}

// NewClient creates a new Encryptonize client. Note that in order to call endpoints that require
// authentication, you need to call `LoginUser` first.
func NewClient(ctx context.Context, endpoint, certPath string) (*Client, error) {
	client, err := internal.NewClient(ctx, endpoint, certPath)
	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}

// Close closes all connections to the Encryptonize server.
func (c *Client) Close() error {
	return c.client.Close()
}

// SetToken sets the provided token as the authentication header.
func (c *Client) SetToken(token string) {
	c.client.SetToken(token)
}

// GetTokenExpiration returns when the current token wil expire.
func (c *Client) GetTokenExpiration() time.Time {
	return c.client.GetTokenExpiration()
}

/////////////////////////////////////////////////////////////////////////
//                               Utility                               //
/////////////////////////////////////////////////////////////////////////

// Version retrieves the version information of the Encryptonize service.
func (c *Client) Version() (*pkg.VersionResponse, error) {
	return c.client.Version()
}

// Health retrieves the current health status of the Encryptonize service.
func (c *Client) Health() (*pkg.HealthResponse, error) {
	return c.client.Health()
}

/////////////////////////////////////////////////////////////////////////
//                           User Management                           //
/////////////////////////////////////////////////////////////////////////

// LoginUser authenticates to the Encryptonize service with the given credentials and sets the
// resulting access token for future calls. Call `LoginUser` again to switch to a different user.
func (c *Client) LoginUser(uid, password string) error {
	return c.client.LoginUser(uid, password)
}

// CreateUser creates a new Encryptonize user with the requested scopes.
func (c *Client) CreateUser(scopes []pkg.Scope) (*pkg.CreateUserResponse, error) {
	return c.client.CreateUser(scopes)
}

// RemoveUser removes a user from the Encryptonize service.
func (c *Client) RemoveUser(uid string) error {
	return c.client.RemoveUser(uid)
}

// CreateGroup creates a new Encryptonize group with the requested scopes.
func (c *Client) CreateGroup(scopes []pkg.Scope) (*pkg.CreateGroupResponse, error) {
	return c.client.CreateGroup(scopes)
}

// AddUserToGroup adds a user to a group.
func (c *Client) AddUserToGroup(uid, gid string) error {
	return c.client.AddUserToGroup(uid, gid)
}

// RemoveUserFromGroup removes a user from a group.
func (c *Client) RemoveUserFromGroup(uid, gid string) error {
	return c.client.RemoveUserFromGroup(uid, gid)
}

/////////////////////////////////////////////////////////////////////////
//                               Storage                               //
/////////////////////////////////////////////////////////////////////////

// Store encrypts the `plaintext` and tags both `plaintext` and `associatedData` storing the
// resulting ciphertext in the Encryptonize service.
func (c *Client) Store(plaintext, associatedData []byte) (*pkg.StoreResponse, error) {
	return c.client.Store(plaintext, associatedData)
}

// Retrieve decrypts a previously stored object returning the ciphertext.
func (c *Client) Retrieve(oid string) (*pkg.RetrieveResponse, error) {
	return c.client.Retrieve(oid)
}

// Update replaces the currently stored data of an object with the specified `plaintext` and
// `associatedData`.
func (c *Client) Update(oid string, plaintext, associatedData []byte) error {
	return c.client.Update(oid, plaintext, associatedData)
}

// Delete removes previously stored data from the Encryptonize service.
func (c *Client) Delete(oid string) error {
	return c.client.Delete(oid)
}

/////////////////////////////////////////////////////////////////////////
//                             Permissions                             //
/////////////////////////////////////////////////////////////////////////

// GetPermissions returns a list of IDs that have access to the requested object.
func (c *Client) GetPermissions(oid string) (*pkg.GetPermissionsResponse, error) {
	return c.client.GetPermissions(oid)
}

// AddPermission grants permission for the group to the requested object.
func (c *Client) AddPermission(oid, gid string) error {
	return c.client.AddPermission(oid, gid)
}

// RemovePermission removes permissions for the group to the requested object.
func (c *Client) RemovePermission(oid, gid string) error {
	return c.client.RemovePermission(oid, gid)
}
