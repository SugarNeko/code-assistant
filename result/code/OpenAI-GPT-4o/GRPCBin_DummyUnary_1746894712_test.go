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
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	request := &pb.DummyMessage{
		FString:  "test",
		FStrings: []string{"one", "two"},
		FInt32:   123,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    pb.DummyMessage_ENUM_1,
		FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2},
		FSub: &pb.DummyMessage_Sub{
			FString: "sub_test",
		},
		FSubs: []*pb.DummyMessage_Sub{
			{FString: "sub1"},
		},
		FBool:   true,
		FBools:  []bool{true, false},
		FInt64:  1234567890,
		FInt64S: []int64{987654321},
		FBytes:  []byte("bytes"),
		FBytess: [][]byte{[]byte("more bytes")},
		FFloat:  3.14,
		FFloats: []float32{1.23, 4.56},
	}

	response, err := client.DummyUnary(context.Background(), request)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if response.FString != request.FString {
		t.Errorf("Expected FString: %s, got: %s", request.FString, response.FString)
	}

	if len(response.FStrings) != len(request.FStrings) || response.FStrings[0] != request.FStrings[0] {
		t.Errorf("Expected FStrings: %v, got: %v", request.FStrings, response.FStrings)
	}

	// Add more assertions to verify other fields...
}
