module gogrpc

go 1.16

require (
	google.golang.org/grpc v1.39.0
	google.golang.org/grpc/examples v0.0.0-20210722024238-c513103bee39
	protos/helloworld v0.0.0-00010101000000-000000000000
)

replace protos/helloworld => ../protos
