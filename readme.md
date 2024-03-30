# go-auth-server

A simple example of an authentication server that works with Go's `clientcredentials` package. Does not include database integration for client credential storage and validation or secret store integration for hidden encryption key. Might include in the future.

### Running Locally

Run server.
- `go run cmd/server/main.go`

Run client to get a token via `clientcredentials` and validate.
- `go run cmd/client/main.go`