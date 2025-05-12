package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

const (
	address = "grpcb.in:9000"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "subtest"},
		FInt64:   123456789,
		FBytes:   []byte("testbytes"),
		FFloat:   1.23,
		FBool:    true,
		FStrings: []string{"str1", "str2"},
		FInt32S:  []int32{1, 2, 3},
		FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
		FSubs:    []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBools:   []bool{true, false},
		FInt64S:  []int64{987654321, 123456789},
		FBytess:  [][]byte{[]byte("byte1"), []byte("byte2")},
		FFloats:  []float32{3.21, 6.54},
	}

	res, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("could not greet: %v", err)
	}

	// Validate response
	if res.FString != req.FString {
		t.Errorf("Expected FString: %v, got: %v", req.FString, res.FString)
	}
	if res.FInt32 != req.FInt32 {
		t.Errorf("Expected FInt32: %v, got: %v", req.FInt32, res.FInt32)
	}
	if res.FEnum != req.FEnum {
		t.Errorf("Expected FEnum: %v, got: %v", req.FEnum, res.FEnum)
	}
	if res.FSub.FString != req.FSub.FString {
		t.Errorf("Expected Sub FString: %v, got: %v", req.FSub.FString, res.FSub.FString)
	}
	if res.FBool != req.FBool {
		t.Errorf("Expected FBool: %v, got: %v", req.FBool, res.FBool)
	}
	if len(res.FStrings) != len(req.FStrings) {
		t.Errorf("Expected len(FStrings): %v, got: %v", len(req.FStrings), len(res.FStrings))
	}
}
