package grpcbin_test

import (
	"context"
	"log"
	"testing"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub_test"},
		FBool:    true,
		FInt64:   456,
		FBytes:   []byte("bytes"),
		FFloat:   123.45,
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("Unary call failed: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("Expected FString: %v, got: %v", req.FString, resp.FString)
	}

	if resp.FInt32 != req.FInt32 {
		t.Errorf("Expected FInt32: %v, got: %v", req.FInt32, resp.FInt32)
	}

	if resp.FEnum != req.FEnum {
		t.Errorf("Expected FEnum: %v, got: %v", req.FEnum, resp.FEnum)
	}

	if resp.FSub.FString != req.FSub.FString {
		t.Errorf("Expected FSub FString: %v, got: %v", req.FSub.FString, resp.FSub.FString)
	}

	if resp.FBool != req.FBool {
		t.Errorf("Expected FBool: %v, got: %v", req.FBool, resp.FBool)
	}

	if resp.FInt64 != req.FInt64 {
		t.Errorf("Expected FInt64: %v, got: %v", req.FInt64, resp.FInt64)
	}

	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("Expected FBytes: %v, got: %v", string(req.FBytes), string(resp.FBytes))
	}

	if resp.FFloat != req.FFloat {
		t.Errorf("Expected FFloat: %v, got: %v", req.FFloat, resp.FFloat)
	}
}
