module proxy

go 1.16

replace go-gateway-protos/demo => ../go-gateway-protos

require (
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	go-gateway-protos/demo v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.39.0
)
