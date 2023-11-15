build:
	go build -o service -ldflags "-s -w"

generate:
	# Generate go, gRPC-Gateway, OpenAPI output.
	#
	# -I declares import folders, in order of importance
	# This is how proto resolves the protofile imports.
	# It will check for the protofile relative to each of these
	# folders and use the first one it finds.
	#
	# --go_out generates go Protobuf output with gRPC plugin enabled.
	# 		paths=source_relative means the file should be generated
	# 		relative to the input proto file.
	# --grpc-gateway_out generates gRPC-Gateway output.
	# --openapiv2_out generates an OpenAPI 2.0 specification for our gRPC-Gateway endpoints.
	protoc \
		-I proto \
		-I third_party/googleapis \
		-I "${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.0.1/" \
		--go_out ./gen/proto --go_opt paths=source_relative \
		--go-grpc_out ./gen/proto --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./gen/proto \
		--grpc-gateway_opt logtostderr=true \
    	--grpc-gateway_opt paths=source_relative \
    	--grpc-gateway_opt generate_unbound_methods=true \
		--openapiv2_out ./third_party/OpenAPI --openapiv2_opt logtostderr=true \
		proto/khiops.proto

		# Generate static assets for OpenAPI UI
		statik -m -f -src third_party/OpenAPI/

install:
	go install \
    	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    	google.golang.org/protobuf/cmd/protoc-gen-go \
    	google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/rakyll/statik

lint:
	golangci-lint run -E gosec,goimports ./...

test:
	go clean -cache && go test -mod=readonly -v -coverprofile=coverage.out -race ./... -v ./server

update:
	go get -u ./...
