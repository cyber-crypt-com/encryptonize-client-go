# Go Client Packages for CYBERCRYPT D1;

Go client packages for
* [CYBERCRYPT D1 Storage](https://github.com/cybercryptio/d1-service-storage)
* [CYBERCRYPT D1 Generic](https://github.com/cybercryptio/d1-service-generic)

## D1 Storage Client

In order to use the D1 Storage client you will need credentials for a user. If you are using the
built in Standalone ID Provider you can refer to the [Getting Started](https://docs.cybercrypt.io/storage-service/getting_started)
guide for details on how to obtain these. If you are using an OIDC provider you will need to obtain
and ID Token in the usual way.

When using the Standalone ID Provider the easiest way to use the D1 Storage client is through the
`NewStorageClientWR` constructor. The resulting client will automatically refresh the user's access
token when it expires. For specific code examples, see the [godocs](https://pkg.go.dev/github.com/cybercryptio/d1-client-go).

When using an OIDC provider you must use the `NewStorageClient` constructor and provide the ID Token as [gRPC metadata](https://pkg.go.dev/google.golang.org/grpc/metadata) via the context:

```go
ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "bearer " + idToken)
client, _ := NewStorageClient(...)
client.Store(ctx, ...)
```

## D1 Generic Client

In order to use the D1 Generic client you will need credentials for a user. If you are using the
built in Standalone ID Provider you can refer to the [Getting Started](https://docs.cybercrypt.io/generic-service/getting_started)
guide for details on how to obtain these. If you are using an OIDC provider you will need to obtain
and ID Token in the usual way.

When using the Standalone ID Provider the easiest way to use the D1 Generic client is through the
`NewGenericClientWR` constructor. The resulting client will automatically refresh the user's access
token when it expires. For specific code examples, see the [godocs](https://pkg.go.dev/github.com/cybercryptio/d1-client-go).

When using an OIDC provider you must use the `NewGenericClient` constructor and provide the ID Token as [gRPC metadata](https://pkg.go.dev/google.golang.org/grpc/metadata) via the context:

```go
ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "bearer " + idToken)
client, _ := NewGenericClient(...)
client.Encrypt(ctx, ...)
```

## License

The software in the CYBERCRYPT d1-client-go repository is licensed under the Apache License 2.0.
