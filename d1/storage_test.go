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

	gpb "github.com/cybercryptio/d1-service-generic/protobuf"
	spb "github.com/cybercryptio/d1-service-storage/protobuf"
	"google.golang.org/grpc/metadata"
)

func setupStorageClient(t *testing.T) (*StorageClient, context.Context) {
	ctx := context.Background()
	c, err := NewStorageClient(endpoint, certPath)
	failOnError("NewStorageClient failed", err, t)

	res, err := c.LoginUser(
		ctx,
		&gpb.LoginUserRequest{
			UserId:   uid,
			Password: password,
		},
	)
	failOnError("LoginUser failed", err, t)

	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "bearer "+res.AccessToken)
	return &c, ctx
}

func setupStorageClientWR(t *testing.T) (*StorageClient, context.Context) {
	c, err := NewStorageClientWR(endpoint, certPath, uid, password)
	failOnError("NewStorageClientWR failed", err, t)
	return &c, context.Background()
}

func TestStorage(t *testing.T) {
	cases := []struct {
		descriptor  string
		constructor func(t *testing.T) (*StorageClient, context.Context)
	}{
		{
			descriptor:  "StorageClient",
			constructor: setupStorageClient,
		}, {
			descriptor:  "StorageClientWR",
			constructor: setupStorageClientWR,
		},
	}

	for _, c := range cases {
		t.Run(c.descriptor, func(t *testing.T) {
			client, ctx := c.constructor(t)
			defer client.Close()

			plaintext := []byte("foo")
			associatedData := []byte("bar")
			encryptResponse, err := client.Store(
				ctx,
				&spb.StoreRequest{
					Plaintext:      plaintext,
					AssociatedData: associatedData,
				},
			)
			failOnError("Encrypt failed", err, t)

			_, err = client.Retrieve(
				ctx,
				&spb.RetrieveRequest{
					ObjectId: encryptResponse.ObjectId,
				},
			)
			failOnError("Decrypt failed", err, t)

			_, err = client.AddPermission(
				ctx,
				&gpb.AddPermissionRequest{
					ObjectId: encryptResponse.ObjectId,
					GroupId:  uid,
				},
			)
			failOnError("AddPermission failed", err, t)

			_, err = client.GetPermissions(
				ctx,
				&gpb.GetPermissionsRequest{
					ObjectId: encryptResponse.ObjectId,
				},
			)
			failOnError("GetPermissions failed", err, t)

			_, err = client.RemovePermission(
				ctx,
				&gpb.RemovePermissionRequest{
					ObjectId: encryptResponse.ObjectId,
					GroupId:  uid,
				},
			)
			failOnError("RemovePermission failed", err, t)
		})
	}
}
