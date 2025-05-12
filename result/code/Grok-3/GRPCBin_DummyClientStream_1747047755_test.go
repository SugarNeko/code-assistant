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
	serverAddr    = "grpcb.in:9000"
	connectTimeout = 15 * time.Second
)

func TestGRPCBin_DummyClientStream(t *testing.T) {
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
	t.Run("PositiveTest_ValidDummyMessages", func(t *testing.T) {
		stream, err := client.DummyClientStream(context.Background())
		if err != nil {
			t.Fatalf("Failed to create stream: %v", err)
		}

		// Send 10 valid DummyMessages
		for i := 0; i < 10; i++ {
			msg := &grpcbin.DummyMessage{
				FString:  "test-string",
				FStrings: []string{"str1", "str2"},
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
				FInt64S:  []int64{10, 20},
				FBytes:   []byte("test-bytes"),
				FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:   float32(1.23),
				FFloats:  []float32{1.1, 2.2},
			}

			if err := stream.Send(msg); err != nil {
				t.Fatalf("Failed to send message %d: %v", i, err)
			}
		}

		// Receive the response (last sent message should be returned)
		response, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate server response
		if response.FString != "test-string" {
			t.Errorf("Expected FString to be 'test-string', got '%s'", response.FString)
		}
		if len(response.FStrings) != 2 || response.FStrings[0] != "str1" || response.FStrings[1] != "str2" {
			t.Errorf("Expected FStrings to be ['str1', 'str2'], got %v", response.FStrings)
		}
		if response.FInt32 != 9 { // Last message index
			t.Errorf("Expected FInt32 to be 9, got %d", response.FInt32)
		}
		if response.FEnum != grpcbin.DummyMessage_ENUM_1 {
			t.Errorf("Expected FEnum to be ENUM_1, got %v", response.FEnum)
		}
		if response.FSub.FString != "sub-test" {
			t.Errorf("Expected FSub.FString to be 'sub-test', got '%s'", response.FSub.FString)
		}
		if response.FBool != true {
			t.Errorf("Expected FBool to be true, got %v", response.FBool)
		}
		if response.FInt64 != 9 { // Last message index
			t.Errorf("Expected FInt64 to be 9, got %d", response.FInt64)
		}
		if string(response.FBytes) != "test-bytes" {
			t.Errorf("Expected FBytes to be 'test-bytes', got '%s'", string(response.FBytes))
		}
		if response.FFloat != float32(1.23) {
			t.Errorf("Expected FFloat to be 1.23, got %f", response.FFloat)
		}
	})
}
