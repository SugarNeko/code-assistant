package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBinService(t *testing.T) {
	// Set up connection with timeout
	conn, err := grpc.Dial("grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	t.Run("TestDummyBidirectionalStreamStream_ValidRequest", func(t *testing.T) {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Initialize bidirectional stream
		stream, err := client.DummyBidirectionalStreamStream(ctx)
		if err != nil {
			t.Fatalf("Failed to create bidirectional stream: %v", err)
		}

		// Prepare test message
		testMsg := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"test1", "test2"},
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
			FInt64:   1000000,
			FInt64S:  []int64{100, 200},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
		}

		// Send test message
		if err := stream.Send(testMsg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}

		// Receive response
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate response
		if resp.FString != testMsg.FString {
			t.Errorf("Expected FString to be %s, got %s", testMsg.FString, resp.FString)
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
			t.Errorf("Expected FSub.FString to be %s, got %s", testMsg.FSub.FString, resp.FSub.FString)
		}
		if resp.FBool != testMsg.FBool {
			t.Errorf("Expected FBool to be %v, got %v", testMsg.FBool, resp.FBool)
		}
		if resp.FInt64 != testMsg.FInt64 {
			t.Errorf("Expected FInt64 to be %d, got %d", testMsg.FInt64, resp.FInt64)
		}
		if string(resp.FBytes) != string(testMsg.FBytes) {
			t.Errorf("Expected FBytes to be %v, got %v", testMsg.FBytes, resp.FBytes)
		}
		if resp.FFloat != testMsg.FFloat {
			t.Errorf("Expected FFloat to be %f, got %f", testMsg.FFloat, resp.FFloat)
		}
	})
}
