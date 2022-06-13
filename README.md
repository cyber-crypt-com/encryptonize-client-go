# Go Client Library for Encryptonize&reg;

Go client libraries for
* [Encryptonize&reg; Core](https://github.com/cybercryptio/d1-service-generic)
* [Encryptonize&reg; Objects](https://github.com/cybercryptio/d1-service-storage)
* [Encryptonize&reg; Key Server](https://github.com/cybercryptio/k1)

## Usage
In order to use the client you will need credentials for the Encryptonize Core server.
When setting up the server the first time, you need to bootstrap an initial user with credentials
either through the executable as described
[here](https://github.com/cybercryptio/d1-service-generic/blob/master/documentation/user_manual.md#bootstrapping-users).
Subsequent users can be created through the API as described
[here](https://github.com/cybercryptio/d1-service-generic/blob/master/documentation/user_manual.md#creating-users-through-the-api).

The easiest way to use the Encryptonize&reg; Core client is through the `NewCoreClientWR`
constructor. The resulting client will automatically refresh the user's access token when it
expires.

```go
// Create a new Encryptonize Core client, providing the hostname, a root CA certificate, and user
// credentials.
client, err := NewCoreClientWR(context.Background(), "localhost:9000", "./ca.crt", "user id", "password")
if err != nil {
  log.Fatal(err)
}

// Encrypt sensitive data.
plaintext, associatedData := []byte("secret data"), []byte("metadata")
encrypted, err := client.Encrypt(plaintext, associatedData)
if err != nil {
  log.Fatal(err)
}

// Decrypt the response.
decrypted, err := client.Decrypt(encrypted.ObjectID, encrypted.Ciphertext, encrypted.AssociatedData)
if err != nil {
  log.Fatal(err)
}

log.Printf("%+v", decrypted)
```

A similar constructor exists for Encryptonize&reg; Objects:

```go
// Create a new Encryptonize Objects client, providing the hostname, a root CA certificate, and
// user credentials.
client, err := NewObjectsClientWR(context.Background(), "localhost:9000", "./ca.crt", "user id", "password")
if err != nil {
  log.Fatal(err)
}

// Store sensitive data in encrypted form.
plaintext, associatedData := []byte("secret data"), []byte("metadata")
response, err := client.Store(plaintext, associatedData)
if err != nil {
  log.Fatal(err)
}

// Retrieve the stored data.
decrypted, err := client.Retrieve(response.ObjectID)
if err != nil {
  log.Fatal(err)
}

log.Printf("%+v", decrypted)
```

For other examples, see the [godocs](TODO).

## Development
Make targets are provided for various tasks. To get an overview run `make help`. To build the
clients run `make build`.

The Core and Objects clients can be tested against a docker deployment of the services by running
`make docker-core-test` and `make docker-objects-test`.
