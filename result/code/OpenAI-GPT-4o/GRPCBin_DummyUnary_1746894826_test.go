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
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:  "test",
		FStrings: []string{"str1", "str2"},
		FInt32:   5,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    pb.DummyMessage_ENUM_1,
		FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_1, pb.DummyMessage_ENUM_2},
		FSub:     &pb.DummyMessage_Sub{FString: "subtest"},
		FSubs:    []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false, true},
		FInt64:   1234567890,
		FInt64S:  []int64{987654321, 123456789},
		FBytes:   []byte("bytes"),
		FBytess:  [][]byte{{'a', 'b'}, {'c', 'd'}},
		FFloat:   1.23,
		FFloats:  []float32{1.1, 2.2, 3.3},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to call DummyUnary: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("Expected FString: %v, got: %v", req.FString, resp.FString)
	}
	// Additional assertions for all other fields...
}
