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

func TestGRPCBinService_DummyBidirectionalStreamStream(t *testing.T) {
	// Set up connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Test case for positive testing with valid input
	t.Run("PositiveTest_ValidRequest", func(t *testing.T) {
		// Create stream client
		stream, err := client.DummyBidirectionalStreamStream(context.Background())
		if err != nil {
			t.Fatalf("Failed to create stream: %v", err)
		}

		// Prepare test message
		testMsg := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"str1", "str2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "sub-test",
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
			FFloats:  []float32{1.1, 2.2},
		}

		// Send test message to server
		if err := stream.Send(testMsg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}

		// Receive response from server
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate client request (what we sent)
		if testMsg.FString != "test-string" {
			t.Errorf("Client request validation failed for FString: got %v, want %v", testMsg.FString, "test-string")
		}
		if len(testMsg.FStrings) != 2 {
			t.Errorf("Client request validation failed for FStrings: got length %v, want 2", len(testMsg.FStrings))
		}

		// Validate server response
		if resp.FString != testMsg.FString {
			t.Errorf("Server response validation failed for FString: got %v, want %v", resp.FString, testMsg.FString)
		}
		if len(resp.FStrings) != len(testMsg.FStrings) {
			t.Errorf("Server response validation failed for FStrings: got length %v, want %v", len(resp.FStrings), len(testMsg.FStrings))
		}
		if resp.FInt32 != testMsg.FInt32 {
			t.Errorf("Server response validation failed for FInt32: got %v, want %v", resp.FInt32, testMsg.FInt32)
		}
		if resp.FEnum != testMsg.FEnum {
			t.Errorf("Server response validation failed for FEnum: got %v, want %v", resp.FEnum, testMsg.FEnum)
		}
		if resp.FSub.FString != testMsg.FSub.FString {
			t.Errorf("Server response validation failed for FSub.FString: got %v, want %v", resp.FSub.FString, testMsg.FSub.FString)
		}
		if resp.FBool != testMsg.FBool {
			t.Errorf("Server response validation failed for FBool: got %v, want %v", resp.FBool, testMsg.FBool)
		}
		if resp.FInt64 != testMsg.FInt64 {
			t.Errorf("Server response validation failed for FInt64: got %v, want %v", resp.FInt64, testMsg.FInt64)
		}
		if len(resp.FBytes) != len(testMsg.FBytes) {
			t.Errorf("Server response validation failed for FBytes: got length %v, want %v", len(resp.FBytes), len(testMsg.FBytes))
		}
		if resp.FFloat != testMsg.FFloat {
			t.Errorf("Server response validation failed for FFloat: got %v, want %v", resp.FFloat, testMsg.FFloat)
		}
	})
}
