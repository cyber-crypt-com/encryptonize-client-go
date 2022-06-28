// Copyright 2020-2022 CYBERCRYPT

package d1

import (
	"testing"

	"context"

	pb "github.com/cybercryptio/d1-service-generic/protobuf"
	"google.golang.org/grpc/metadata"
)

func setupGenericClient(t *testing.T) (*GenericClient, context.Context) {
	ctx := context.Background()
	c, err := NewGenericClient(endpoint, certPath)
	failOnError("NewGenericClient failed", err, t)

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

func setupGenericClientWR(t *testing.T) (*GenericClient, context.Context) {
	c, err := NewGenericClientWR(endpoint, certPath, uid, password)
	failOnError("NewGenericClientWR failed", err, t)
	return &c, context.Background()
}

func TestGeneric(t *testing.T) {
	cases := []struct {
		descriptor  string
		constructor func(t *testing.T) (*GenericClient, context.Context)
	}{
		{
			descriptor:  "GenericClient",
			constructor: setupGenericClient,
		}, {
			descriptor:  "GenericClientWR",
			constructor: setupGenericClientWR,
		},
	}

	for _, c := range cases {
		t.Run(c.descriptor, func(t *testing.T) {
			client, ctx := c.constructor(t)
			defer client.Close()

			plaintext := []byte("foo")
			associatedData := []byte("bar")
			encryptResponse, err := client.Encrypt(
				ctx,
				&pb.EncryptRequest{
					Plaintext:      plaintext,
					AssociatedData: associatedData,
				},
			)
			failOnError("Encrypt failed", err, t)

			_, err = client.Decrypt(
				ctx,
				&pb.DecryptRequest{
					ObjectId:       encryptResponse.ObjectId,
					Ciphertext:     encryptResponse.Ciphertext,
					AssociatedData: encryptResponse.AssociatedData,
				},
			)
			failOnError("Decrypt failed", err, t)

			_, err = client.AddPermission(
				ctx,
				&pb.AddPermissionRequest{
					ObjectId: encryptResponse.ObjectId,
					GroupId:  uid,
				},
			)
			failOnError("AddPermission failed", err, t)

			_, err = client.GetPermissions(
				ctx,
				&pb.GetPermissionsRequest{
					ObjectId: encryptResponse.ObjectId,
				},
			)
			failOnError("GetPermissions failed", err, t)

			_, err = client.RemovePermission(
				ctx,
				&pb.RemovePermissionRequest{
					ObjectId: encryptResponse.ObjectId,
					GroupId:  uid,
				},
			)
			failOnError("RemovePermission failed", err, t)
		})
	}
}
