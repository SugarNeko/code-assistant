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

func TestGRPCBinService(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	t.Run("TestDummyBidirectionalStreamStream", func(t *testing.T) {
		stream, err := client.DummyBidirectionalStreamStream(context.Background())
		if err != nil {
			t.Fatalf("Failed to create bidirectional stream: %v", err)
		}

		// Test data to send to server
		testMsg := &grpcbin.DummyMessage{
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
			FInt64:   1234567890,
			FInt64S:  []int64{1, 2, 3},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2, 3.3},
		}

		// Send message to server
		if err := stream.Send(testMsg); err != nil {
			t.Fatalf("Failed to send message to server: %v", err)
		}

		// Receive and validate response from server
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive response from server: %v", err)
		}

		// Validate server response matches the sent data
		if resp.FString != testMsg.FString {
			t.Errorf("Expected FString to be %q, got %q", testMsg.FString, resp.FString)
		}
		if len(resp.FStrings) != len(testMsg.FStrings) {
			t.Errorf("Expected FStrings length to be %d, got %d", len(testMsg.FStrings), len(resp.FStrings))
		}
		if resp.FInt32 != testMsg.FInt32 {
			t.Errorf("Expected FInt32 to be %d, got %d", testMsg.FInt32, resp.FInt32)
		}
		if resp.FEnum != testMsg.FEnum {
			t.Errorf("Expected FEnum to be %v, got %v", testMsg.FEnum, resp.FEnum)
		}
		if resp.FSub.FString != testMsg.FSub.FString {
			t.Errorf("Expected FSub.FString to be %q, got %q", testMsg.FSub.FString, resp.FSub.FString)
		}
		if resp.FBool != testMsg.FBool {
			t.Errorf("Expected FBool to be %v, got %v", testMsg.FBool, resp.FBool)
		}
		if resp.FInt64 != testMsg.FInt64 {
			t.Errorf("Expected FInt64 to be %d, got %d", testMsg.FInt64, resp.FInt64)
		}
		if string(resp.FBytes) != string(testMsg.FBytes) {
			t.Errorf("Expected FBytes to be %q, got %q", testMsg.FBytes, resp.FBytes)
		}
		if resp.FFloat != testMsg.FFloat {
			t.Errorf("Expected FFloat to be %f, got %f", testMsg.FFloat, resp.FFloat)
		}

		// Close the stream
		if err := stream.CloseSend(); err != nil {
			t.Fatalf("Failed to close send stream: %v", err)
		}
	})
}
