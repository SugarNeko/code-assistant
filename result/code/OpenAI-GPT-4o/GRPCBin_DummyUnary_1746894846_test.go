package grpcbin_test

import (
	"context"
	"log"
	"testing"

	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "test string",
		FStrings: []string{"str1", "str2"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub test"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   64,
		FInt64S:  []int64{64, 128},
		FBytes:   []byte("test bytes"),
		FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("Expected FString %v, got %v", req.FString, resp.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("Expected FStrings length %d, got %d", len(req.FStrings), len(resp.FStrings))
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("Expected FInt32 %v, got %v", req.FInt32, resp.FInt32)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("Expected FEnum %v, got %v", req.FEnum, resp.FEnum)
	}
	if resp.FSub.FString != req.FSub.FString {
		t.Errorf("Expected FSub FString %v, got %v", req.FSub.FString, resp.FSub.FString)
	}
	if resp.FBool != req.FBool {
		t.Errorf("Expected FBool %v, got %v", req.FBool, resp.FBool)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("Expected FInt64 %v, got %v", req.FInt64, resp.FInt64)
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("Expected FBytes %v, got %v", req.FBytes, resp.FBytes)
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("Expected FFloat %v, got %v", req.FFloat, resp.FFloat)
	}
}
