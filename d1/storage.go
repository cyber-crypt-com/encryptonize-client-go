// Copyright 2020-2022 CYBERCRYPT

package d1

import (
	pb "github.com/cybercryptio/d1-service-storage/protobuf"
)

// GenericClient can be used to make calls to a D1 Storage service.
type StorageClient struct {
	BaseClient
	pb.ObjectsClient
}

// NewStorageClient creates a new client for the given endpoint. If certPath is not empty, TLS will
// be enabled using the given certificate file.
func NewStorageClient(endpoint, certPath string) (StorageClient, error) {
	base, err := newBaseClient(endpoint, certPath)
	if err != nil {
		return StorageClient{}, nil
	}

	return StorageClient{
		BaseClient:    base,
		ObjectsClient: pb.NewObjectsClient(base.connection),
	}, nil
}

// NewStorageClientWR creates a GenericClient that automatically logs in and refreshes the access
// token using the provided credentials.
func NewStorageClientWR(endpoint, certPath, uid, password string) (StorageClient, error) {
	base, err := newBaseClientWR(endpoint, certPath, uid, password)
	if err != nil {
		return StorageClient{}, nil
	}

	return StorageClient{
		BaseClient:    base,
		ObjectsClient: pb.NewObjectsClient(base.connection),
	}, nil
}
