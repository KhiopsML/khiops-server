# REST+gRPC service gateway

A simple grpc-gateway powered remote control API for Khiops.

## Running

Running `main.go` starts a web server on https://0.0.0.0:11000/. You can configure
the port used with the `$PORT` environment variable, and to serve on HTTP set
`$SERVE_HTTP=true`.

An OpenAPI UI is served on https://0.0.0.0:11000/.

## Requirements

Generating the client/server files requires the `protoc` protobuf compiler.
Please install it according to the
[installation instructions](https://github.com/google/protobuf#protocol-compiler-installation)
for your specific platform.

## Getting started

### Building the service locally
This is only needed to modify the service. For clients simply using the service, skip to the next section.

1. Install the generate dependencies with `make install`.
   This will install `protoc-gen-go`, `protoc-gen-grpc-gateway`, `protoc-gen-swagger` and `statik` which
   are necessary for us to generate the Go, swagger and static files.
1. Generate the files with `make generate`.
   If you encounter an error here, make sure you've installed
   `protoc` and it is accessible in your `$PATH`, and make sure
   you've performed step 1.
1. Finally you can run the web server with `go run main.go`.

### Using the service

The service exposes both gRPC and REST endpoint. Several usage examples are given [here](run_test.sh).

To build a custom client using your favourite language, lookup the relevant gRPC documentation (https://grpc.io/docs/quickstart/) and create your client from the [proto](proto/khiops.proto) service definition.