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
	pb "github.com/cybercryptio/d1-service-generic/protobuf"
)

// GenericClient can be used to make calls to a D1 Generic service.
type GenericClient struct {
	BaseClient
	pb.CoreClient
}

// NewGenericClient creates a new client for the given endpoint. If certPath is not empty, TLS will
// be enabled using the given certificate file.
func NewGenericClient(endpoint, certPath string) (GenericClient, error) {
	base, err := newBaseClient(endpoint, certPath)
	if err != nil {
		return GenericClient{}, nil
	}

	return GenericClient{
		BaseClient: base,
		CoreClient: pb.NewCoreClient(base.connection),
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
		BaseClient: base,
		CoreClient: pb.NewCoreClient(base.connection),
	}, nil
}
