package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverAddr     = "grpcb.in:9000"
	connectTimeout = 15 * time.Second
)

func TestGRPCBin_DummyBidirectionalStreamStream(t *testing.T) {
	// Set up connection to the gRPC server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Test positive case for bidirectional streaming
	t.Run("PositiveTest_BidirectionalStream", func(t *testing.T) {
		stream, err := client.DummyBidirectionalStreamStream(context.Background())
		if err != nil {
			t.Fatalf("failed to create bidirectional stream: %v", err)
		}

		// Construct a valid DummyMessage request
		req := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"test1", "test2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "sub-test-string",
			},
			FSubs: []*grpcbin.DummyMessage_Sub{
				{FString: "sub1"},
				{FString: "sub2"},
			},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   123456789,
			FInt64S:  []int64{1, 2, 3},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2, 3.3},
		}

		// Send the request message
		if err := stream.Send(req); err != nil {
			t.Fatalf("failed to send request: %v", err)
		}

		// Receive and validate the response
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("failed to receive response: %v", err)
		}

		// Validate client-side response fields
		if resp.FString != req.FString {
			t.Errorf("expected FString to be %q, got %q", req.FString, resp.FString)
		}
		if len(resp.FStrings) != len(req.FStrings) {
			t.Errorf("expected FStrings length to be %d, got %d", len(req.FStrings), len(resp.FStrings))
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("expected FInt32 to be %d, got %d", req.FInt32, resp.FInt32)
		}
		if resp.FEnum !=.req.FEnum {
			t.Errorf("expected FEnum to be %v, got %v", req.FEnum, resp.FEnum)
		}
		if resp.FSub.FString != req.FSub.FString {
			t.Errorf("expected FSub.FString to be %q, got %q", req.FSub.FString, resp.FSub.FString)
		}
		if resp.FBool != req.FBool {
			t.Errorf("expected FBool to be %v, got %v", req.FBool, resp.FBool)
		}
		if resp.FInt64 != req.FInt64 {
			t.Errorf("expected FInt64 to be %d, got %d", req.FInt64, resp.FInt64)
		}
		if string(resp.FBytes) != string(req.FBytes) {
			t.Errorf("expected FBytes to be %v, got %v", req.FBytes, resp.FBytes)
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("expected FFloat to be %f, got %f", req.FFloat, resp.FFloat)
		}

		// Close the stream
		if err := stream.CloseSend(); err != nil {
			t.Fatalf("failed to close send stream: %v", err)
		}
	})
}
