// Copyright 2020-2022 CYBERCRYPT

package d1

import (
	pb "github.com/cybercryptio/d1-service-generic/protobuf"
)

// GenericClient can be used to make calls to a D1 Generic service.
type GenericClient struct {
	BaseClient
	pb.GenericClient
}

// NewGenericClient creates a new client for the given endpoint. If certPath is not empty, TLS will
// be enabled using the given certificate file.
func NewGenericClient(endpoint, certPath string) (GenericClient, error) {
	base, err := newBaseClient(endpoint, certPath)
	if err != nil {
		return GenericClient{}, nil
	}

	return GenericClient{
		BaseClient:    base,
		GenericClient: pb.NewGenericClient(base.connection),
	}, nil
}

// NewGenericClientWR creates a GenericClient that automatically logs in and refreshes the access
// token using the provided credentials.
func NewGenericClientWR(endpoint, certPath, uid, password string) (GenericClient, error) {
	base, err := newBaseClientWR(endpoint, certPath, uid, password)
	if err != nil {
		return GenericClient{}, nil
	}

	return GenericClient{
		BaseClient:    base,
		GenericClient: pb.NewGenericClient(base.connection),
	}, nil
}
