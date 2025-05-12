protoc -I=. \
       --go_out=. \
       --go-grpc_out=. \
       proto/grpcbin/grpcbin.proto