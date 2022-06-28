# Go Client Library for CYBERCRYPT D1 and K1;

Go client libraries for
* [CYBERCRYPT D1 Generic](https://github.com/cybercryptio/d1-service-generic)
* [CYBERCRYPT D1 Storage](https://github.com/cybercryptio/d1-service-storage)
* [CYBERCRYPT K1](https://github.com/cybercryptio/k1)

## Usage
In order to use the client you will need credentials for the D1 Generic server.
When setting up the server the first time, you need to bootstrap an initial user with credentials
either through the executable as described
[here](https://github.com/cybercryptio/d1-service-generic/blob/master/documentation/user_manual.md#bootstrapping-users).
Subsequent users can be created through the API as described
[here](https://github.com/cybercryptio/d1-service-generic/blob/master/documentation/user_manual.md#creating-users-through-the-api).

The easiest way to use the D1 Generic client is through the `NewGenericClientWR` constructor. The
resulting client will automatically refresh the user's access token when it expires. A similar
constructor, `NewStorageClientWR`, exists for D1 Storage. For specific code examples, see the
[godocs](TODO).

## Development
Make targets are provided for various tasks. To get an overview run `make help`. To build the
clients run `make build`.

The D1 Generic and Storage clients can be tested against a docker deployment of the services by running
`make docker-generic-test` and `make docker-storage-test`.
