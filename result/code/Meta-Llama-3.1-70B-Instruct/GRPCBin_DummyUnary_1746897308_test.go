package main

import (
	"context"
	"io"
	"log"
	"testing"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin"
)

const (
	grpcbinAddress = "grpcb.in:9000"
)

func BenchmarkGRPC(b *testing.B) {
	conn, err := grpc.Dial(grpcbinAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGRPCBinClient(conn)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		request := &pb.DummyMessage{
			FString: "test",
			FInt32:  1,
		}
		response, err := client.DummyUnary(context.Background(), request)
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		if response.FString != "test" || response.FInt32 != 1 {
			log.Fatalf("expected response to be {test, 1}, but got {%s, %d}", response.FString, response.FInt32)
		}
	}
}

func TestGRPCUnary(t *testing.T) {
	conn, err := grpc.Dial(grpcbinAddress, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGRPCBinClient(conn)
	request := &pb.DummyMessage{
		FString: "test",
		FInt32:  1,
	}
	response, err := client.DummyUnary(context.Background(), request)
	if err != nil {
		t.Fatalf("could not greet: %v", err)
	}
	if response.FString != "test" || response.FInt32 != 1 {
		t.Errorf("expected response to be {test, 1}, but got {%s, %d}", response.FString, response.FInt32)
	}
}

func TestGRPCUnaryWithNil(t *testing.T) {
	conn, err := grpc.Dial(grpcbinAddress, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGRPCBinClient(conn)
	request := &pb.DummyMessage{}
	_, err = client.DummyUnary(context.Background(), request)
	if err != nil {
		t.Errorf("could not greet with nil request: %v", err)
	}
}

func TestGRPCStream(t *testing.T) {
	conn, err := grpc.Dial(grpcbinAddress, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyStream(context.Background())
	if err != nil {
		t.Fatalf("could not create stream: %v", err)
	}
	err = stream.Send(&pb.DummyMessage{
		FString: "test",
		FInt32:  1,
	})
	if err != nil {
		t.Errorf("could not send request: %v", err)
	}
	_, err = stream.Recv()
	if err != io.EOF {
		t.Errorf("expected EOF, but got %v", err)
	}
}
