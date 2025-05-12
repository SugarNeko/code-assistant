package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	// Set up a connection to the gRPC server.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx,
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Construct a typical request with all fields set.
	req := &grpcbin.DummyMessage{
		FString:  "test-string",
		FStrings: []string{"a", "b", "c"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_2,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub-string"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   123456789,
		FInt64S:  []int64{10, 20, 30},
		FBytes:   []byte("hello-bytes"),
		FBytess:  [][]byte{[]byte("foo"), []byte("bar")},
		FFloat:   3.14,
		FFloats:  []float32{1.23, 4.56},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	var responses []*grpcbin.DummyMessage

	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive message #%d from stream: %v", i+1, err)
		}
		responses = append(responses, resp)
	}

	// Validate we got 10 responses.
	if len(responses) != 10 {
		t.Fatalf("Expected 10 messages from stream, got %d", len(responses))
	}

	// Validate that each response matches the request (echoed message)
	for i, resp := range responses {
		if resp.FString != req.FString {
			t.Errorf("Message #%d: f_string mismatch, got %q want %q", i+1, resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Message #%d: f_int32 mismatch, got %d want %d", i+1, resp.FInt32, req.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("Message #%d: f_enum mismatch, got %v want %v", i+1, resp.FEnum, req.FEnum)
		}
		if resp.FBool != req.FBool {
			t.Errorf("Message #%d: f_bool mismatch, got %v want %v", i+1, resp.FBool, req.FBool)
		}
		if string(resp.FBytes) != string(req.FBytes) {
			t.Errorf("Message #%d: f_bytes mismatch, got %q want %q", i+1, resp.FBytes, req.FBytes)
		}
		// More detailed comparison for slices, nested, bytes, etc. can be added as needed.
	}
}
