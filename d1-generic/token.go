// Copyright 2022 CYBERCRYPT
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by copy-client.sh. DO NOT EDIT.
// version: v2.0.0-ci.67
// source: https://github.com/cybercryptio/d1-service-generic.git
// commit: 75a3140090e05b8c5736b031f09927e499587d9b

package client

import (
	"context"
	"errors"
	"time"

	pbauthn "github.com/cybercryptio/d1-client-go/v2/d1-generic/protobuf/authn"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var noneAuthorizedMethods = map[string]bool{
	"/d1.authn.Authn/LoginUser":    true,
	"/grpc.health.v1.Health/Check": true,
}

// PerRPCToken is an implementation of credentials.PerRPCCredentials that calls a function on every RPC to generate an access token.
// The access token will not be encrypted during transport.
type perRPCToken func(context.Context) (string, error)

func (getToken perRPCToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	ri, ok := credentials.RequestInfoFromContext(ctx)
	if !ok {
		return nil, errors.New("could not get request info")
	}

	if _, ok := noneAuthorizedMethods[ri.Method]; ok {
		return map[string]string{}, nil
	}

	token, err := getToken(ctx)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"authorization": "bearer " + token,
	}, nil
}

func (getToken perRPCToken) RequireTransportSecurity() bool {
	return false
}

// NewStandalonePerRPCToken creates a new instance of PerRPCToken to be used with the Standalone ID Provider.
// It requires the transport credentials used to communicate with the D1 Service in order to call the Login endpoint.
func newStandalonePerRPCToken(c *BaseClient, uid, pwd string) perRPCToken {
	var token string
	var tokenExpiry time.Time
	return func(ctx context.Context) (string, error) {
		// To avoid clock drift issues, refresh the token if it will expire within 1 minute.
		if time.Now().After(tokenExpiry.Add(time.Duration(-1) * time.Minute)) {
			res, err := c.Authn.LoginUser(
				ctx,
				&pbauthn.LoginUserRequest{
					UserId:   uid,
					Password: pwd,
				},
			)
			if err != nil {
				return "", err
			}

			tokenExpiry = time.Unix(res.ExpiryTime, 0)
			token = res.AccessToken
		}
		return token, nil
	}
}

// WithTokenRefresh returns an Option that configures token refresh.
func WithTokenRefresh(uid, pwd string) Option {
	return func(bc *BaseClient) grpc.DialOption {
		return grpc.WithPerRPCCredentials(newStandalonePerRPCToken(bc, uid, pwd))
	}
}
