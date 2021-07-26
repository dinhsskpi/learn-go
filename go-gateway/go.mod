module go-gateway

go 1.16

replace go-gateway-protos/demo => ../go-gateway-protos

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	go-gateway-protos/demo v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.39.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
)
