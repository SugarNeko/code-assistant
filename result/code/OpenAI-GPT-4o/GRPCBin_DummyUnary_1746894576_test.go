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
		t.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "sub"},
		FBool:    true,
		FInt64:   456,
		FBytes:   []byte("bytes"),
		FFloat:   1.23,
	}

	res, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if res.FString != req.FString {
		t.Errorf("Expected FString %v, got %v", req.FString, res.FString)
	}
	if res.FInt32 != req.FInt32 {
		t.Errorf("Expected FInt32 %v, got %v", req.FInt32, res.FInt32)
	}
	if res.FEnum != req.FEnum {
		t.Errorf("Expected FEnum %v, got %v", req.FEnum, res.FEnum)
	}
	if res.FSub.FString != req.FSub.FString {
		t.Errorf("Expected FSub FString %v, got %v", req.FSub.FString, res.FSub.FString)
	}
	if res.FBool != req.FBool {
		t.Errorf("Expected FBool %v, got %v", req.FBool, res.FBool)
	}
	if res.FInt64 != req.FInt64 {
		t.Errorf("Expected FInt64 %v, got %v", req.FInt64, res.FInt64)
	}
	if string(res.FBytes) != string(req.FBytes) {
		t.Errorf("Expected FBytes %v, got %v", string(req.FBytes), string(res.FBytes))
	}
	if res.FFloat != req.FFloat {
		t.Errorf("Expected FFloat %v, got %v", req.FFloat, res.FFloat)
	}
}
