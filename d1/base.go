// Copyright 2020-2022 CYBERCRYPT

// d1-client-go/d1 is a client library for CYBERCRYPT D1.
package d1

import (
	"context"
	"time"

	pb "github.com/cybercryptio/d1-service-generic/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// BaseClient represents the shared functionality between various D1 services.
type BaseClient struct {
	pb.VersionClient
	pb.AuthnClient
	pb.AuthzClient
	grpc_health_v1.HealthClient
	connection *grpc.ClientConn
}

// newBaseClient creates a new client for the given endpoint. If certPath is not empty, TLS will be
// enabled using the given certificate file.
func newBaseClient(endpoint, certPath string) (BaseClient, error) {
	var dialOption grpc.DialOption

	if certPath != "" {
		// Configure certificate
		clientCredentials, err := credentials.NewClientTLSFromFile(certPath, "")
		if err != nil {
			return BaseClient{}, err
		}
		dialOption = grpc.WithTransportCredentials(clientCredentials)
	} else {
		dialOption = grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	// Initialize connection with the service
	conn, err := grpc.Dial(endpoint, dialOption)
	if err != nil {
		return BaseClient{}, err
	}

	return BaseClient{
		VersionClient: pb.NewVersionClient(conn),
		AuthnClient:   pb.NewAuthnClient(conn),
		AuthzClient:   pb.NewAuthzClient(conn),
		HealthClient:  grpc_health_v1.NewHealthClient(conn),
		connection:    conn,
	}, nil
}

// Close closes all connections to the server.
func (b *BaseClient) Close() error {
	return b.connection.Close()
}

// tokenRefresher handles automatic refreshing of the access token upon expiry by implementing
// credentials.PerRPCCredentials.
type tokenRefresher struct {
	endpoint   string
	certPath   string
	uid        string
	password   string
	token      string
	expiration time.Time
}

func (t *tokenRefresher) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	// To avoid clock drift issues, refresh the token if it will expire within 1 minute.
	if time.Now().After(t.expiration.Add(time.Duration(-1) * time.Minute)) {
		c, err := newBaseClient(t.endpoint, t.certPath)
		if err != nil {
			return nil, err
		}
		defer c.Close()

		res, err := c.LoginUser(
			ctx,
			&pb.LoginUserRequest{
				UserId:   t.uid,
				Password: t.password,
			},
		)
		if err != nil {
			return nil, err
		}

		t.expiration = time.Unix(res.ExpiryTime, 0)
		t.token = res.AccessToken
	}

	return map[string]string{
		"authorization": "bearer " + t.token,
	}, nil
}

func (t *tokenRefresher) RequireTransportSecurity() bool {
	return false
}

// newBaseClientWR creates a baseClient that automatically logs in and refreshes the access token
// using the provided credentials.
func newBaseClientWR(endpoint, certPath, uid, password string) (BaseClient, error) {
	var dialOption grpc.DialOption

	if certPath != "" {
		// Configure certificate
		clientCredentials, err := credentials.NewClientTLSFromFile(certPath, "")
		if err != nil {
			return BaseClient{}, err
		}
		dialOption = grpc.WithTransportCredentials(clientCredentials)
	} else {
		dialOption = grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	refresher := tokenRefresher{
		endpoint: endpoint,
		certPath: certPath,
		uid:      uid,
		password: password,
	}

	// Initialize connection with the service using the automatic token refresher.
	conn, err := grpc.Dial(endpoint, dialOption, grpc.WithPerRPCCredentials(&refresher))
	if err != nil {
		return BaseClient{}, err
	}

	return BaseClient{
		VersionClient: pb.NewVersionClient(conn),
		AuthnClient:   pb.NewAuthnClient(conn),
		AuthzClient:   pb.NewAuthzClient(conn),
		HealthClient:  grpc_health_v1.NewHealthClient(conn),
		connection:    conn,
	}, nil
}
