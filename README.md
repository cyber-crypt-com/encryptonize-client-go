# Go Client Packages for CYBERCRYPT D1

Go client packages for
* [CYBERCRYPT D1 Storage](https://github.com/cybercryptio/d1-service-storage)
* [CYBERCRYPT D1 Generic](https://github.com/cybercryptio/d1-service-generic)

## D1 Storage Client

In order to use the D1 Storage client you will need credentials for a user. If you are using the
built in Standalone ID Provider you can refer to the [Getting Started](https://docs.cybercrypt.io/storage-service/getting_started)
guide for details on how to obtain these. If you are using an OIDC provider you will need to obtain
and ID Token in the usual way.

The client can be configured with an option that provides an implementation of `credentials.PerRPCCredentials` to easily attach a token to every request. For convenience, we provide an implementation called `PerRPCToken` which can be any function that fetches a token, for example, after a call to the OIDC provider. There is also an implementation that can be used with the Standalone ID Provider:

```go
client, _ := client.NewStorageClient(endpoint,
		gclient.WithTransportCredentials(creds),
		gclient.WithPerRPCCredentials(
			gclient.NewStandalonePerRPCToken(endpoint, uid, password, creds),
		),
	)
```

The ID Token can also be attached manually as [gRPC metadata](https://pkg.go.dev/google.golang.org/grpc/metadata) via the context:

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

The client can be configured with an option that provides an implementation of `credentials.PerRPCCredentials` to easily attach a token to every request. For convenience, we provide an implementation called `PerRPCToken` which can be any function that fetches a token, for example, after a call to the OIDC provider. There is also an implementation that can be used with the Standalone ID Provider:

```go
client, _ := client.NewGenericClient(endpoint,
		client.WithTransportCredentials(creds),
		client.WithPerRPCCredentials(
			client.NewStandalonePerRPCToken(endpoint, uid, password, creds),
		),
	)
```

The ID Token can also be attached manually as [gRPC metadata](https://pkg.go.dev/google.golang.org/grpc/metadata) via the context:

```go
ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "bearer " + idToken)
client, _ := NewGenericClient(...)
client.Encrypt(ctx, ...)
```

## License

The software in the CYBERCRYPT d1-client-go repository is licensed under the Apache License 2.0.
