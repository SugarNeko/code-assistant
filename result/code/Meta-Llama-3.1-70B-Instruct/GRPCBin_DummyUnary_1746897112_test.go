package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGRPCBin_DummyUnary(t *testing.T) {
	// Setup gRPC connection
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Create client
	client := NewGRPCBinClient(conn)

	// Test positive case
	dummyMessage := &DummyMessage{
		FString: "test",
		FInt32:  123,
		FEnum:   Enum_ENUM_0,
	}
	resp, err := client.DummyUnary(context.Background(), dummyMessage)
	if err != nil {
		t.Errorf("DummyUnary() error = %v", err)
	}
	if resp.FString != dummyMessage.FString {
		t.Errorf("DummyUnary() response FString = %v, want %v", resp.FString, dummyMessage.FString)
	}
	if resp.FInt32 != dummyMessage.FInt32 {
		t.Errorf("DummyUnary() response FInt32 = %v, want %v", resp.FInt32, dummyMessage.FInt32)
	}
	if resp.FEnum != dummyMessage.FEnum {
		t.Errorf("DummyUnary() response FEnum = %v, want %v", resp.FEnum, dummyMessage.FEnum)
	}

	// Test invalid request
	req := &DummyMessage{
		FString: "",
		FInt32:  0,
		FEnum:   Enum(-1),
	}
	resp, err = client.DummyUnary(context.Background(), req)
	if err == nil {
		t.Errorf("DummyUnary() error = nil, want error")
	}
	status, ok := status.FromError(err)
	if !ok {
		t.Errorf("DummyUnary() error is not a status error")
	}
	if status.Code() != codes.InvalidArgument {
		t.Errorf("DummyUnary() error code = %v, want %v", status.Code(), codes.InvalidArgument)
	}
}
