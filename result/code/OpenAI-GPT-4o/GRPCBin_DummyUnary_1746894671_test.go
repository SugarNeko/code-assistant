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
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:  "test",
		FStrings: []string{"foo", "bar"},
		FInt32:   123,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    pb.DummyMessage_ENUM_1,
		FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2},
		FSub:     &pb.DummyMessage_Sub{FString: "subtest"},
		FSubs:    []*pb.DummyMessage_Sub{{FString: "anothersub"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   456,
		FInt64S:  []int64{456, 789},
		FBytes:   []byte("bytes"),
		FBytess:  [][]byte{{'b', 'y', 't'}, {'e', 's'}},
		FFloat:   1.23,
		FFloats:  []float32{4.56, 7.89},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("expected %v, got %v", req.FString, resp.FString)
	}

	// Additional validation checks can be added here for other fields
}
