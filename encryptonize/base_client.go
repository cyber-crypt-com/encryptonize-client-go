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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	grpc_reflection "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

// baseClient for making gRPC calls to the Encryptonize service.
type baseClient struct {
	connection      *grpc.ClientConn
	refClient       *grpcreflect.Client
	ctx             context.Context
	reflSource      grpcurl.DescriptorSource
	authHeader      []string
	tokenExpiration time.Time
}

// NewClient creates a new Encryptonize client. Note that in order to call endpoints that require
// authentication, you need to call `LoginUser` first.
func newBaseClient(ctx context.Context, endpoint, certPath string) (*baseClient, error) {
	var dialOption grpc.DialOption

	if certPath != "" {
		// Configure certificate
		clientCredentials, err := credentials.NewClientTLSFromFile(certPath, "")
		if err != nil {
			return nil, err
		}
		dialOption = grpc.WithTransportCredentials(clientCredentials)
	} else {
		dialOption = grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	// Initialize connection with Encryptonize service
	connection, err := grpc.Dial(endpoint, dialOption)
	if err != nil {
		return nil, err
	}

	client := grpcreflect.NewClient(ctx, grpc_reflection.NewServerReflectionClient(connection))
	reflSource := grpcurl.DescriptorSourceFromServer(ctx, client)

	return &baseClient{
		connection: connection,
		refClient:  client,
		ctx:        ctx,
		reflSource: reflSource,
	}, nil
}

// Close closes all connections to the Encryptonize server.
func (c *baseClient) Close() error {
	return c.connection.Close()
}

// SetToken sets the provided token as the authentication header.
func (c *baseClient) SetToken(token string) {
	c.authHeader = []string{"authorization: bearer " + token}
}

// GetTokenExpiration returns when the current token wil expire.
func (c *baseClient) GetTokenExpiration() time.Time {
	return c.tokenExpiration
}

// invoke calls `method` with the requested `input` and unmarshals the response into `output`.
func (c *baseClient) invoke(method, input string, output interface{}) error {
	options := grpcurl.FormatOptions{
		EmitJSONDefaultFields: false,
		IncludeTextSeparator:  true,
		AllowUnknownFields:    false,
	}
	requestParser, formatter, err := grpcurl.RequestParserAndFormatter(
		grpcurl.FormatJSON,
		c.reflSource,
		strings.NewReader(input),
		options,
	)
	if err != nil {
		return err
	}

	var response bytes.Buffer
	handler := &grpcurl.DefaultEventHandler{
		Out:            &response,
		Formatter:      formatter,
		VerbosityLevel: 0,
	}
	err = grpcurl.InvokeRPC(
		c.ctx,
		c.reflSource,
		c.connection,
		method,
		c.authHeader,
		handler,
		requestParser.Next)
	if err != nil {
		return err
	}
	if handler.Status.Code() != codes.OK {
		return handler.Status.Err()
	}

	return json.Unmarshal(response.Bytes(), output)
}

// parseScopes converts an array of `Scope`s to an array of strings.
func (c *baseClient) parseScopes(scopes []Scope) ([]string, error) {
	scopeStrings := make([]string, 0, len(scopes))

	for _, scope := range scopes {
		switch scope {
		case ScopeRead:
			scopeStrings = append(scopeStrings, "READ")
		case ScopeCreate:
			scopeStrings = append(scopeStrings, "CREATE")
		case ScopeUpdate:
			scopeStrings = append(scopeStrings, "UPDATE")
		case ScopeDelete:
			scopeStrings = append(scopeStrings, "DELETE")
		case ScopeIndex:
			scopeStrings = append(scopeStrings, "INDEX")
		case ScopeObjectPermissions:
			scopeStrings = append(scopeStrings, "OBJECTPERMISSIONS")
		case ScopeUserManagement:
			scopeStrings = append(scopeStrings, "USERMANAGEMENT")
		default:
			return nil, errors.New("invalid scope")
		}
	}

	return scopeStrings, nil
}

/////////////////////////////////////////////////////////////////////////
//                               Utility                               //
/////////////////////////////////////////////////////////////////////////

// Version retrieves the version information of the Encryptonize service.
func (c *baseClient) Version() (*VersionResponse, error) {
	response := &VersionResponse{}
	if err := c.invoke("encryptonize.Version.Version", "", response); err != nil {
		return nil, err
	}

	return response, nil
}

// Health retrieves the current health status of the Encryptonize service.
func (c *baseClient) Health() (*HealthResponse, error) {
	response := &HealthResponse{}
	if err := c.invoke("grpc.health.v1.Health.Check", "", response); err != nil {
		return nil, err
	}

	return response, nil
}

/////////////////////////////////////////////////////////////////////////
//                           User Management                           //
/////////////////////////////////////////////////////////////////////////

// LoginUser authenticates to the Encryptonize service with the given credentials and sets the
// resulting access token for future calls. Call `LoginUser` again to switch to a different user.
func (c *baseClient) LoginUser(uid, password string) error {
	requestJSON, err := json.Marshal(request{UserID: uid, Password: password})
	if err != nil {
		return err
	}

	response := &accessToken{}
	if err := c.invoke("encryptonize.Authn.LoginUser", string(requestJSON), response); err != nil {
		return err
	}

	c.SetToken(response.Token)
	tokenExpiration, err := strconv.ParseInt(response.ExpiryTime, 10, 64)
	if err != nil {
		return err
	}

	c.tokenExpiration = time.Unix(tokenExpiration, 0)
	return nil
}

// CreateUser creates a new Encryptonize user with the requested scopes.
func (c *baseClient) CreateUser(scopes []Scope) (*CreateUserResponse, error) {
	parsedScopes, err := c.parseScopes(scopes)
	if err != nil {
		return nil, err
	}
	requestJSON, err := json.Marshal(request{Scopes: parsedScopes})
	if err != nil {
		return nil, err
	}

	response := &CreateUserResponse{}
	if err := c.invoke("encryptonize.Authn.CreateUser", string(requestJSON), response); err != nil {
		return nil, err
	}

	return response, nil
}

// RemoveUser removes a user from the Encryptonize service.
func (c *baseClient) RemoveUser(uid string) error {
	requestJSON, err := json.Marshal(request{UserID: uid})
	if err != nil {
		return err
	}

	return c.invoke("encryptonize.Authn.RemoveUser", string(requestJSON), &struct{}{})
}

// CreateGroup creates a new Encryptonize group with the requested scopes.
func (c *baseClient) CreateGroup(scopes []Scope) (*CreateGroupResponse, error) {
	parsedScopes, err := c.parseScopes(scopes)
	if err != nil {
		return nil, err
	}
	requestJSON, err := json.Marshal(request{Scopes: parsedScopes})
	if err != nil {
		return nil, err
	}

	response := &CreateGroupResponse{}
	if err := c.invoke("encryptonize.Authn.CreateGroup", string(requestJSON), response); err != nil {
		return nil, err
	}

	return response, nil
}

// AddUserToGroup adds a user to a group.
func (c *baseClient) AddUserToGroup(uid, gid string) error {
	requestJSON, err := json.Marshal(request{UserID: uid, GroupID: gid})
	if err != nil {
		return err
	}

	return c.invoke("encryptonize.Authn.AddUserToGroup", string(requestJSON), &struct{}{})
}

// RemoveUserFromGroup removes a user from a group.
func (c *baseClient) RemoveUserFromGroup(uid, gid string) error {
	requestJSON, err := json.Marshal(request{UserID: uid, GroupID: gid})
	if err != nil {
		return err
	}

	return c.invoke("encryptonize.Authn.RemoveUserFromGroup", string(requestJSON), &struct{}{})
}

/////////////////////////////////////////////////////////////////////////
//                              Encryption                             //
/////////////////////////////////////////////////////////////////////////

// Encrypt encrypts the `plaintext` and tags both `plaintext` and `associatedData` returning the
// resulting ciphertext.
func (c *baseClient) Encrypt(plaintext, associatedData []byte) (*EncryptResponse, error) {
	requestJSON, err := json.Marshal(request{Plaintext: plaintext, AssociatedData: associatedData})
	if err != nil {
		return nil, err
	}

	response := &EncryptResponse{}
	if err := c.invoke("encryptonize.Core.Encrypt", string(requestJSON), response); err != nil {
		return nil, err
	}

	return response, nil
}

// Decrypt decrypts a previously encrypted `ciphertext` and verifies the integrity of the `ciphertext`
// and `associatedData`.
func (c *baseClient) Decrypt(objectID string, ciphertext, associatedData []byte) (*DecryptResponse, error) {
	requestJSON, err := json.Marshal(request{ObjectID: objectID, Ciphertext: ciphertext, AssociatedData: associatedData})
	if err != nil {
		return nil, err
	}

	response := &DecryptResponse{}
	if err := c.invoke("encryptonize.Core.Decrypt", string(requestJSON), response); err != nil {
		return nil, err
	}

	return response, nil
}

/////////////////////////////////////////////////////////////////////////
//                               Storage                               //
/////////////////////////////////////////////////////////////////////////

// Store encrypts the `plaintext` and tags both `plaintext` and `associatedData` storing the
// resulting ciphertext in the Encryptonize service.
func (c *baseClient) Store(plaintext, associatedData []byte) (*StoreResponse, error) {
	requestJSON, err := json.Marshal(request{Plaintext: plaintext, AssociatedData: associatedData})
	if err != nil {
		return nil, err
	}

	response := &StoreResponse{}
	if err := c.invoke("encryptonize.Objects.Store", string(requestJSON), response); err != nil {
		return nil, err
	}

	return response, nil
}

// Retrieve decrypts a previously stored object returning the ciphertext.
func (c *baseClient) Retrieve(oid string) (*RetrieveResponse, error) {
	requestJSON, err := json.Marshal(request{ObjectID: oid})
	if err != nil {
		return nil, err
	}

	response := &RetrieveResponse{}
	if err := c.invoke("encryptonize.Objects.Retrieve", string(requestJSON), response); err != nil {
		return nil, err
	}

	return response, nil
}

// Update replaces the currently stored data of an object with the specified `plaintext` and
// `associatedData`.
func (c *baseClient) Update(oid string, plaintext, associatedData []byte) error {
	requestJSON, err := json.Marshal(request{ObjectID: oid, Plaintext: plaintext, AssociatedData: associatedData})
	if err != nil {
		return err
	}

	return c.invoke("encryptonize.Objects.Update", string(requestJSON), &struct{}{})
}

// Delete removes previously stored data from the Encryptonize service.
func (c *baseClient) Delete(oid string) error {
	requestJSON, err := json.Marshal(request{ObjectID: oid})
	if err != nil {
		return err
	}

	return c.invoke("encryptonize.Objects.Delete", string(requestJSON), &struct{}{})
}

/////////////////////////////////////////////////////////////////////////
//                             Permissions                             //
/////////////////////////////////////////////////////////////////////////

// GetPermissions returns a list of IDs that have access to the requested object.
func (c *baseClient) GetPermissions(oid string) (*GetPermissionsResponse, error) {
	requestJSON, err := json.Marshal(request{ObjectID: oid})
	if err != nil {
		return nil, err
	}

	response := &GetPermissionsResponse{}
	if err := c.invoke("encryptonize.Authz.GetPermissions", string(requestJSON), response); err != nil {
		return nil, err
	}

	return response, nil
}

// AddPermission grants permission for the group to the requested object.
func (c *baseClient) AddPermission(oid, gid string) error {
	requestJSON, err := json.Marshal(request{ObjectID: oid, GroupID: gid})
	if err != nil {
		return err
	}

	return c.invoke("encryptonize.Authz.AddPermission", string(requestJSON), &struct{}{})
}

// RemovePermission removes permissions for the group to the requested object.
func (c *baseClient) RemovePermission(oid, gid string) error {
	requestJSON, err := json.Marshal(request{ObjectID: oid, GroupID: gid})
	if err != nil {
		return err
	}

	return c.invoke("encryptonize.Authz.RemovePermission", string(requestJSON), &struct{}{})
}

// request is a catch-all for request structs. By using `omitempty` we can marshal to the correct
// JSON structure by only setting the necessary fields.
type request struct {
	Scopes         []string `json:"scopes,omitempty"`
	UserID         string   `json:"user_id,omitempty"`
	GroupID        string   `json:"group_id,omitempty"`
	ObjectID       string   `json:"object_id,omitempty"`
	Plaintext      []byte   `json:"plaintext,omitempty"`
	Ciphertext     []byte   `json:"ciphertext,omitempty"`
	AssociatedData []byte   `json:"associated_data,omitempty"`
	Password       string   `json:"password,omitempty"`
}

type accessToken struct {
	Token      string `json:"accessToken"`
	ExpiryTime string `json:"expiryTime"`
}
