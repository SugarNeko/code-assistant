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

	req := &pb.DummyMessage{
		FString:  "test",
		FStrings: []string{"test1", "test2"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    pb.DummyMessage_ENUM_1,
		FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
		FSub:     &pb.DummyMessage_Sub{FString: "sub"},
		FSubs:    []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   64,
		FInt64S:  []int64{64, 128},
		FBytes:   []byte{1, 2, 3},
		FBytess:  [][]byte{{4, 5}, {6, 7}},
		FFloat:   3.14,
		FFloats:  []float32{1.2, 3.4},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("Expected FString %v, got %v", req.FString, resp.FString)
	}

	if len(resp.FStrings) != len(req.FStrings) ||
		resp.FStrings[0] != req.FStrings[0] || resp.FStrings[1] != req.FStrings[1] {
		t.Errorf("Expected FStrings %v, got %v", req.FStrings, resp.FStrings)
	}

	// Continue to verify other fields as needed...
}
