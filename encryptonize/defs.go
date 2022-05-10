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

/////////////////////////////////////////////////////////////////////////
//                               Utility                               //
/////////////////////////////////////////////////////////////////////////

// VersionResponse is the response to a Version call.
type VersionResponse struct {
	// Git commit of the current version
	Commit string `json:"commit"`

	// Version tag of the current version
	Tag string `json:"tag"`
}

// HealthResponse is the response to a Health call.
type HealthResponse struct {
	// Current health status of the server. See
	// github.com/grpc/grpc/blob/master/doc/health-checking.md for details.
	Status string `json:"status"`
}

/////////////////////////////////////////////////////////////////////////
//                           User Management                           //
/////////////////////////////////////////////////////////////////////////

type Scope int

const (
	ScopeRead              Scope = iota // Scope to decrypt objects
	ScopeCreate                         // Scope encrypt/store objects
	ScopeUpdate                         // Scope to update existing objects
	ScopeDelete                         // Scope to delete stored objects
	ScopeIndex                          // Scope to list object permissions
	ScopeObjectPermissions              // Scope to adit object permissions
	ScopeUserManagement                 // Scope to manage users
)

// CreateUserResponse is the response to a CreateUser call.
type CreateUserResponse struct {
	// ID of the newly created user.
	UserID string `json:"userId"`

	// Password of the newly created user.
	Password string `json:"password"`
}

// CreateGroupResponse is the response to a CreateGroup call.
type CreateGroupResponse struct {
	// ID of the newly created group.
	GroupID string `json:"groupId"`
}

/////////////////////////////////////////////////////////////////////////
//                              Encryption                             //
/////////////////////////////////////////////////////////////////////////

// EncryptResponse is the response to a Encrypt call.
type EncryptResponse struct {
	// Encrypted and authenticated data.
	Ciphertext []byte `json:"ciphertext"`

	// Plaintext authenticated data.
	AssociatedData []byte `json:"associatedData"`

	// ID of the encrypted object.
	ObjectID string `json:"objectId"`
}

// DecryptResponse is the response to a Decrypt call.
type DecryptResponse struct {
	// Decrypted and authenticated data.
	Plaintext []byte `json:"plaintext"`

	// Authenticated data.
	AssociatedData []byte `json:"associatedData"`
}

/////////////////////////////////////////////////////////////////////////
//                               Storage                               //
/////////////////////////////////////////////////////////////////////////

// StoreResponse is the response to a Store call.
type StoreResponse struct {
	// ID of the stored object.
	ObjectID string `json:"objectId"`
}

// RetrieveResponse is the response to a Retrieve call.
type RetrieveResponse struct {
	// Decrypted and authenticated data.
	Plaintext []byte `json:"plaintext"`

	// Authenticated data.
	AssociatedData []byte `json:"associatedData"`
}

/////////////////////////////////////////////////////////////////////////
//                             Permissions                             //
/////////////////////////////////////////////////////////////////////////

// GetPermissionsResponse is the response to a GetPermissions call.
type GetPermissionsResponse struct {
	// List of IDs that have permission to access the object.
	GroupIDs []string `json:"groupIds"`
}
