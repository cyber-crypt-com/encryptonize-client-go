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

package d1

import (
	"testing"

	"context"
	"log"
	"os"

	pb "github.com/cybercryptio/d1-service-generic/protobuf"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
)

var uid string
var password string
var certPath = ""
var endpoint = "localhost:9000"

var scopes = []pb.Scope{
	pb.Scope_READ,
	pb.Scope_CREATE,
	pb.Scope_UPDATE,
	pb.Scope_DELETE,
	pb.Scope_INDEX,
	pb.Scope_OBJECTPERMISSIONS,
}

func failOnError(message string, err error, t *testing.T) {
	if err != nil {
		t.Fatalf("%s: %v", message, err)
	}
}

func TestMain(m *testing.M) {
	var ok bool
	uid, ok = os.LookupEnv("E2E_TEST_UID")
	if !ok {
		log.Fatal("E2E_TEST_UID must be set")
	}
	password, ok = os.LookupEnv("E2E_TEST_PASS")
	if !ok {
		log.Fatal("E2E_TEST_PASS must be set")
	}
	value, ok := os.LookupEnv("E2E_TEST_CERT")
	if ok {
		certPath = value
	}
	value, ok = os.LookupEnv("E2E_TEST_URL")
	if ok {
		endpoint = value
	}

	os.Exit(m.Run())
}

func setupBaseClient(t *testing.T) (*BaseClient, context.Context) {
	ctx := context.Background()
	c, err := newBaseClient(endpoint, certPath)
	failOnError("newBaseClient failed", err, t)

	res, err := c.LoginUser(
		ctx,
		&pb.LoginUserRequest{
			UserId:   uid,
			Password: password,
		},
	)
	failOnError("LoginUser failed", err, t)

	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "bearer "+res.AccessToken)
	return &c, ctx
}

func setupBaseClientWR(t *testing.T) (*BaseClient, context.Context) {
	c, err := newBaseClientWR(endpoint, certPath, uid, password)
	failOnError("newBaseClientWR failed", err, t)
	return &c, context.Background()
}

func TestBaseUtility(t *testing.T) {
	cases := []struct {
		descriptor  string
		constructor func(t *testing.T) (*BaseClient, context.Context)
	}{
		{
			descriptor:  "BaseClient",
			constructor: setupBaseClient,
		}, {
			descriptor:  "BaseClientWR",
			constructor: setupBaseClientWR,
		},
	}

	for _, c := range cases {
		t.Run(c.descriptor, func(t *testing.T) {
			client, ctx := c.constructor(t)
			defer client.Close()

			_, err := client.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
			failOnError("Health check failed", err, t)
			_, err = client.Version(ctx, &pb.VersionRequest{})
			failOnError("Version request failed", err, t)
		})
	}
}

func TestBaseUserManagement(t *testing.T) {
	cases := []struct {
		descriptor  string
		constructor func(t *testing.T) (*BaseClient, context.Context)
	}{
		{
			descriptor:  "BaseClient",
			constructor: setupBaseClient,
		}, {
			descriptor:  "BaseClientWR",
			constructor: setupBaseClientWR,
		},
	}

	for _, c := range cases {
		t.Run(c.descriptor, func(t *testing.T) {
			client, ctx := c.constructor(t)
			defer client.Close()

			createUserResponse, err := client.CreateUser(
				ctx,
				&pb.CreateUserRequest{
					Scopes: scopes,
				},
			)
			failOnError("CreateUser failed", err, t)

			createGroupResponse, err := client.CreateGroup(
				ctx,
				&pb.CreateGroupRequest{
					Scopes: scopes,
				},
			)
			failOnError("CreateGroup failed", err, t)

			_, err = client.AddUserToGroup(
				ctx,
				&pb.AddUserToGroupRequest{
					UserId:  createUserResponse.UserId,
					GroupId: createGroupResponse.GroupId,
				},
			)
			failOnError("AddUserToGroup failed", err, t)

			_, err = client.RemoveUserFromGroup(
				ctx,
				&pb.RemoveUserFromGroupRequest{
					UserId:  createUserResponse.UserId,
					GroupId: createGroupResponse.GroupId,
				},
			)
			failOnError("RemoveUserFromGroup failed", err, t)

			_, err = client.RemoveUser(
				ctx, &pb.RemoveUserRequest{
					UserId: createUserResponse.UserId,
				},
			)
			failOnError("RemoveUser failed", err, t)
		})
	}
}
