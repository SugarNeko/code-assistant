package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	"example/proto/grpcbin"
)

func getDialOption() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithInsecure(),
	}
}

func TestDummyUnary(t *testing.T) {
	srv := grpc.NewServer()
	grpcbin.RegisterGRPCBinServer(srv, &grpcbinImpl{})
	listenAddr := fmt.Sprintf(":%d", 9000)
	go func() {
		if err := srv.Serve(&net.Listen("tcp", listenAddr)); err != nil {
			panic(err)
		}
	}()

	m := grpcbin.DummyMessage{
		FEnum:   grpcbin.Enum_ENUM_0,
		FString: "StringField",
		// more fields...
	}
	r, err := grpcbinStub.DummyUnary(context.Background(), &m)
	res := r.(*grpcbin.DummyMessage)
	if res != nil {
		t.Log(res)
	} else {
		t.Errorf("Invalid response")
	}
}

// Implement the service
type grpcbinImpl struct{}

func (s *grpcbinImpl) DummyUnary(ctx context.Context, req *grpcbin.DummyMessage) (*grpcbin.DummyMessage, error) {
	return req, nil
}

func (s *grpcbinImpl) DummyStreamingUnary(ctx context.Context, req *grpcbin.DummyMessage, stream grpcbin.GRPCBin_DummyStreamingUnaryServer) error {

}

func (s *grpcbinImpl) DummyUnaryStream(req *grpcbin.DummyMessage, stream grpcbin.GRPCBin_DummyUnaryStreamServer) error {

}

func (s *grpcbinImpl) DummyBidiStreaming(req *grpcbin.DummyMessage, stream grpcbin.GRPCBin_DummyBidiStreamingServer) error {

}
