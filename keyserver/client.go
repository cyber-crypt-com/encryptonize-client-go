// Copyright 2020-2022 CYBERCRYPT

// encryptonize-client-go/keyserver is a client library for the Encryptonize Key Server.
package keyserver

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"strings"

	"github.com/cybercryptio/d1-lib/crypto"
	"github.com/cybercryptio/d1-lib/key"
	"github.com/cybercryptio/k1/service"
	"github.com/fullstorydev/grpcurl"
	"github.com/gofrs/uuid"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	grpc_reflection "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

const clientNonceSize = 32

// GetKeySetRequest is the request made during a GetKeySet call.
type GetKeySetRequest struct {
	KikID string `json:"kikId"`
	Nonce []byte `json:"nonce"`
}

// GetKeySetResponse is the response to a GetKeySet call.
type GetKeySetResponse struct {
	Nonce       []byte `json:"nonce"`
	WrappedKeys []byte `json:"wrappedKeys"`
}

// Client is the Key Server client.
type Client struct {
	connection *grpc.ClientConn
	refClient  *grpcreflect.Client
	ctx        context.Context
	reflSource grpcurl.DescriptorSource
	endpoint   string
	kik        string
	kikID      uuid.UUID
}

// NewClient creates a new Key Server client.
func NewClient(ctx context.Context, endpoint, certPath, kik string, kikID uuid.UUID) (*Client, error) {
	var opts grpc.DialOption

	if certPath != "" {
		// Configure certificate
		creds, err := credentials.NewClientTLSFromFile(certPath, "")
		if err != nil {
			return nil, err
		}
		opts = grpc.WithTransportCredentials(creds)
	} else {
		opts = grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	// Initialize connection with KeyServer
	conn, err := grpc.Dial(endpoint, opts)
	if err != nil {
		return nil, err
	}

	client := grpcreflect.NewClient(ctx, grpc_reflection.NewServerReflectionClient(conn))
	source := grpcurl.DescriptorSourceFromServer(ctx, client)

	return &Client{
		connection: conn,
		refClient:  client,
		ctx:        ctx,
		reflSource: source,
		endpoint:   endpoint,
		kik:        kik,
		kikID:      kikID,
	}, nil
}

// invoke calls `method` with the requested `input` and unmarshals the response into `output`.
func (c *Client) invoke(method, input string, response interface{}) error {
	opts := grpcurl.FormatOptions{
		EmitJSONDefaultFields: false,
		IncludeTextSeparator:  true,
		AllowUnknownFields:    false,
	}
	parser, formatter, err := grpcurl.RequestParserAndFormatter(
		grpcurl.FormatJSON,
		c.reflSource,
		strings.NewReader(input),
		opts,
	)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	handler := &grpcurl.DefaultEventHandler{
		Out:            &out,
		Formatter:      formatter,
		VerbosityLevel: 0,
	}
	err = grpcurl.InvokeRPC(c.ctx, c.reflSource, c.connection, method, nil, handler, parser.Next)
	if err != nil {
		return err
	}
	if handler.Status.Code() != codes.OK {
		return handler.Status.Err()
	}

	return json.Unmarshal(out.Bytes(), response)
}

// GetKeys establishes a connection with a specified Key Server and retrieves a Key Set.
func (c *Client) GetKeys() (key.Keys, error) {
	// This is the client side of the key exchange with the Key Server.
	// The flow is roughly:
	// 1. Client gets instantiated with a Key Initialization Key (KIK) (obtained by the admin from the
	//    KS) and a KIK ID.
	// 2. Client sends the KIK ID and a nonce to the KS.
	// 3. The KS responds with a nonce and a wrapped KeySet.
	// 4. The client derives the wrapping Key from the KIK, nonces, and other information.
	// 5. The client unwraps the KeySet and returns it to the caller.

	// Prepare key exchange parameters
	kik, err := base64.StdEncoding.DecodeString(c.kik)
	if err != nil {
		return key.Keys{}, err
	}

	requestNonce, err := (&crypto.NativeRandom{}).GetBytes(clientNonceSize)
	if err != nil {
		return key.Keys{}, err
	}

	// Make request
	request, err := json.Marshal(GetKeySetRequest{KikID: c.kikID.String(), Nonce: requestNonce})
	if err != nil {
		return key.Keys{}, err
	}

	response := &GetKeySetResponse{}
	if err := c.invoke("keyservice.KeyAPI.GetKeySet", string(request), response); err != nil {
		return key.Keys{}, err
	}

	// Unwrap response
	derivedKey := crypto.KMACKDF(service.KeySize, kik, nil, c.kikID.Bytes(), []byte(c.endpoint), requestNonce, response.Nonce)
	kwp, err := crypto.NewKWP(derivedKey)
	if err != nil {
		return key.Keys{}, err
	}
	keyBytes, err := kwp.Unwrap(response.WrappedKeys)
	if err != nil {
		return key.Keys{}, err
	}

	keys := key.Keys{}
	dec := gob.NewDecoder(bytes.NewReader(keyBytes))
	if err := dec.Decode(&keys); err != nil {
		return key.Keys{}, err
	}

	return keys, nil
}
