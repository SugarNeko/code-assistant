package grpcbin

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBinService(t *testing.T) {
	// Set up connection to the gRPC server with a timeout of 15 seconds
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client for the GRPCBin service
	client := NewGRPCBinClient(conn)

	// Test case for positive testing with a typical request
	t.Run("PositiveTest_DummyBidirectionalStreamStream", func(t *testing.T) {
		// Create a context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Initiate bidirectional stream
		stream, err := client.DummyBidirectionalStreamStream(ctx)
		if err != nil {
			t.Fatalf("Failed to initiate bidirectional stream: %v", err)
		}

		// Prepare a sample request message
		req := &DummyMessage{
			FString:  "test-string",
			FStrings: []string{"test1", "test2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    DummyMessage_ENUM_1,
			FEnums:   []DummyMessage_Enum{DummyMessage_ENUM_0, DummyMessage_ENUM_2},
			FSub:     &DummyMessage_Sub{FString: "sub-test"},
			FSubs:    []*DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   100,
			FInt64S:  []int64{10, 20},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("byte1"), []byte("byte2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
		}

		// Send the request message to the server
		if err := stream.Send(req); err != nil {
			t.Fatalf("Failed to send message to server: %v", err)
		}

		// Receive the response from the server
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive message from server: %v", err)
		}

		// Validate the response from the server (positive testing)
		if resp.FString != req.FString {
			t.Errorf("Expected FString to be %s, got %s", req.FString, resp.FString)
		}
		if len(resp.FStrings) != len(req.FStrings) {
			t.Errorf("Expected FStrings length to be %d, got %d", len(req.FStrings), len(resp.FStrings))
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Expected FInt32 to be %d, got %d", req.FInt32, resp.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("Expected FEnum to be %v, got %v", req.FEnum, resp.FEnum)
		}
		if resp.FSub.FString != req.FSub.FString {
			t.Errorf("Expected FSub.FString to be %s, got %s", req.FSub.FString, resp.FSub.FString)
		}
		if resp.FBool != req.FBool {
			t.Errorf("Expected FBool to be %v, got %v", req.FBool, resp.FBool)
		}
		if resp.FInt64 != req.FInt64 {
			t.Errorf("Expected FInt64 to be %d, got %d", req.FInt64, resp.FInt64)
		}
		if len(resp.FBytes) != len(req.FBytes) {
			t.Errorf("Expected FBytes length to be %d, got %d", len(req.FBytes), len(resp.FBytes))
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("Expected FFloat to be %f, got %f", req.FFloat, resp.FFloat)
		}

		// Close the stream
		if err := stream.CloseSend(); err != nil {
			t.Fatalf("Failed to close stream: %v", err)
		}
	})
}
