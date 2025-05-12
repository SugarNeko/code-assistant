package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyServerStream_Positive(t *testing.T) {
	// Set up a connection to the server with a 15 second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(
		ctx,
		"grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:   "test_string",
		FStrings:  []string{"a", "b"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2, pb.DummyMessage_ENUM_1},
		FSub:      &pb.DummyMessage_Sub{FString: "sub_string"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    987654321,
		FInt64S:   []int64{100, 200, 300},
		FBytes:    []byte("hello"),
		FBytess:   [][]byte{[]byte("foo"), []byte("bar")},
		FFloat:    1.234,
		FFloats:   []float32{9.99, 8.88},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("Error calling DummyServerStream: %v", err)
	}

	responses := []*pb.DummyMessage{}
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		responses = append(responses, resp)
	}

	if len(responses) != 10 {
		t.Errorf("Expected 10 responses, got %d", len(responses))
	}

	// Response validation: Each response must match the request (server echos message 10 times)
	for i, resp := range responses {
		if resp.FString != req.FString {
			t.Errorf("Response[%d] FString mismatch: got %q want %q", i, resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Response[%d] FInt32 mismatch: got %d want %d", i, resp.FInt32, req.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("Response[%d] FEnum mismatch: got %v want %v", i, resp.FEnum, req.FEnum)
		}
		// Add more validation if desired per field...
	}
}
