package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBinDummyClientStream(t *testing.T) {
	// Set up connection with timeout
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create client
	client := grpcbin.NewGRPCBinClient(conn)

	// Test case for positive testing
	t.Run("PositiveTest_DummyClientStream", func(t *testing.T) {
		// Create stream
		stream, err := client.DummyClientStream(context.Background())
		if err != nil {
			t.Fatalf("Failed to create stream: %v", err)
		}

		// Send 10 messages
		for i := 0; i < 10; i++ {
			msg := &grpcbin.DummyMessage{
				FString:  "test-string",
				FStrings: []string{"test1", "test2"},
				FInt32:   int32(i),
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub-test",
				},
				FSubs: []*grpcbin.DummyMessage_Sub{
					{FString: "sub1"},
					{FString: "sub2"},
				},
				FBool:    true,
				FBools:   []bool{true, false},
				FInt64:   int64(i),
				FInt64S:  []int64{10, 20, 30},
				FBytes:   []byte("test-bytes"),
				FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:   float32(1.23),
				FFloats:  []float32{1.1, 2.2, 3.3},
			}

			if err := stream.Send(msg); err != nil {
				t.Fatalf("Failed to send message %d: %v", i, err)
			}
		}

		// Receive response
		response, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate response
		if response == nil {
			t.Fatal("Response is nil")
		}
		if response.FString != "test-string" {
			t.Errorf("Expected FString to be 'test-string', got '%s'", response.FString)
		}
		if len(response.FStrings) != 2 || response.FStrings[0] != "test1" || response.FStrings[1] != "test2" {
			t.Errorf("Expected FStrings to be ['test1', 'test2'], got %v", response.FStrings)
		}
		if response.FInt32 != 9 {
			t.Errorf("Expected FInt32 to be 9, got %d", response.FInt32)
		}
		if response.FEnum != grpcbin.DummyMessage_ENUM_1 {
			t.Errorf("Expected FEnum to be ENUM_1, got %v", response.FEnum)
		}
		if response.FSub == nil || response.FSub.FString != "sub-test" {
			t.Errorf("Expected FSub.FString to be 'sub-test', got %v", response.FSub)
		}
		if len(response.FSubs) != 2 || response.FSubs[0].FString != "sub1" || response.FSubs[1].FString != "sub2" {
			t.Errorf("Expected FSubs to be [{'sub1'}, {'sub2'}], got %v", response.FSubs)
		}
		if !response.FBool {
			t.Errorf("Expected FBool to be true, got %v", response.FBool)
		}
		if response.FInt64 != 9 {
			t.Errorf("Expected FInt64 to be 9, got %d", response.FInt64)
		}
		if string(response.FBytes) != "test-bytes" {
			t.Errorf("Expected FBytes to be 'test-bytes', got %v", response.FBytes)
		}
		if response.FFloat != float32(1.23) {
			t.Errorf("Expected FFloat to be 1.23, got %f", response.FFloat)
		}
	})
}
