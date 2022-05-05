// Copyright 2020-2022 CYBERCRYPT

package keyserver

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	grpc_reflection "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

// GetKeySetRequest defines a grpc proto request.
type GetKeySetRequest struct {
	KikID string `json:"kikId"`
	Nonce []byte `json:"nonce"`
}

// GetKeySetResponse defines a grpc proto response.
type GetKeySetResponse struct {
	Nonce       []byte `json:"nonce"`
	WrappedKeys []byte `json:"wrappedKeys"`
}

// Client for making gRPC calls to the Encryptonize service.
type Client struct {
	connection *grpc.ClientConn
	refClient  *grpcreflect.Client
	ctx        context.Context
	reflSource grpcurl.DescriptorSource
}

// NewClient creates a new KeyServer client.
func NewClient(ctx context.Context, endpoint, certPath string) (*Client, error) {
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

	// Initialize connection with KeyServer
	connection, err := grpc.Dial(endpoint, dialOption)
	if err != nil {
		return nil, err
	}

	client := grpcreflect.NewClient(ctx, grpc_reflection.NewServerReflectionClient(connection))
	reflSource := grpcurl.DescriptorSourceFromServer(ctx, client)

	return &Client{
		connection: connection,
		refClient:  client,
		ctx:        ctx,
		reflSource: reflSource,
	}, nil
}

// invoke calls `method` with the requested `input` and unmarshals the response into `output`.
func (c *Client) invoke(method, input string, output interface{}) error {
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
		nil,
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

func (c *Client) GetKeySet(kikID string, nonce []byte) (*GetKeySetResponse, error) {
	requestJSON, err := json.Marshal(GetKeySetRequest{KikID: kikID, Nonce: nonce})
	if err != nil {
		return nil, err
	}

	response := &GetKeySetResponse{}
	if err := c.invoke("keyservice.KeyAPI.GetKeySet", string(requestJSON), response); err != nil {
		return nil, err
	}

	return response, nil
}
