package grpcbin_test

import (
	"context"
	"testing"
	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	request := &pb.DummyMessage{
		FString:  "test",
		FInt32:   1234,
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "subtest"},
		FBool:    true,
		FInt64:   1234567890,
		FBytes:   []byte("byteTest"),
		FFloat:   1.234,
	}

	response, err := client.DummyUnary(context.Background(), request)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if response.FString != request.FString {
		t.Errorf("Expected FString %s, got %s", request.FString, response.FString)
	}

	if response.FInt32 != request.FInt32 {
		t.Errorf("Expected FInt32 %d, got %d", request.FInt32, response.FInt32)
	}

	if response.FEnum != request.FEnum {
		t.Errorf("Expected FEnum %v, got %v", request.FEnum, response.FEnum)
	}

	if response.FBool != request.FBool {
		t.Errorf("Expected FBool %v, got %v", request.FBool, response.FBool)
	}

	if response.FInt64 != request.FInt64 {
		t.Errorf("Expected FInt64 %d, got %d", request.FInt64, response.FInt64)
	}

	if string(response.FBytes) != string(request.FBytes) {
		t.Errorf("Expected FBytes %s, got %s", string(request.FBytes), string(response.FBytes))
	}

	if response.FFloat != request.FFloat {
		t.Errorf("Expected FFloat %f, got %f", request.FFloat, response.FFloat)
	}
}
