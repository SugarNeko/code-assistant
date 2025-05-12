package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:   "test",
		FInt32:    123,
		FEnum:     pb.DummyMessage_ENUM_1,
		FBool:     true,
		FInt64:    123456789,
		FBytes:    []byte("hello"),
		FFloat:    1.23,
		FStrings:  []string{"a", "b"},
		FInt32S:   []int32{1, 2, 3},
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2},
		FSub:      &pb.DummyMessage_Sub{FString: "sub"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBools:    []bool{true, false},
		FInt64S:   []int64{987654321, 123456789},
		FBytess:   [][]byte{[]byte("abc"), []byte("xyz")},
		FFloats:   []float32{1.1, 2.2, 3.3},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	// Validate client response
	if resp.FString != req.FString {
		t.Errorf("expected FString: %v, got: %v", req.FString, resp.FString)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("expected FInt32: %v, got: %v", req.FInt32, resp.FInt32)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("expected FEnum: %v, got: %v", req.FEnum, resp.FEnum)
	}
	if resp.FBool != req.FBool {
		t.Errorf("expected FBool: %v, got: %v", req.FBool, resp.FBool)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("expected FInt64: %v, got: %v", req.FInt64, resp.FInt64)
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("expected FBytes: %v, got: %v", req.FBytes, resp.FBytes)
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("expected FFloat: %v, got: %v", req.FFloat, resp.FFloat)
	}
	// Additional checks can be added for array fields and other fields as needed
}

