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

		// Prepare test messages (10 messages as per requirement)
		messages := make([]*grpcbin.DummyMessage, 10)
		for i := 0; i < 10; i++ {
			messages[i] = &grpcbin.DummyMessage{
				FString:  "test-string-" + string(rune(i)),
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
				FFloat:   float32(i) + 0.5,
				FFloats:  []float32{1.1, 2.2},
			}
		}

		// Send messages to stream
		for _, msg := range messages {
			if err := stream.Send(msg); err != nil {
				t.Fatalf("Failed to send message to stream: %v", err)
			}
		}

		// Close stream and receive response
		response, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate server response (should return the last message)
		expected := messages[9]
		if response.FString != expected.FString {
			t.Errorf("Expected FString %v, got %v", expected.FString, response.FString)
		}
		if response.FInt32 != expected.FInt32 {
			t.Errorf("Expected FInt32 %v, got %v", expected.FInt32, response.FInt32)
		}
		if response.FEnum != expected.FEnum {
			t.Errorf("Expected FEnum %v, got %v", expected.FEnum, response.FEnum)
		}
		if response.FSub.FString != expected.FSub.FString {
			t.Errorf("Expected FSub.FString %v, got %v", expected.FSub.FString, response.FSub.FString)
		}
		if response.FBool != expected.FBool {
			t.Errorf("Expected FBool %v, got %v", expected.FBool, response.FBool)
		}
		if response.FInt64 != expected.FInt64 {
			t.Errorf("Expected FInt64 %v, got %v", expected.FInt64, response.FInt64)
		}
		if string(response.FBytes) != string(expected.FBytes) {
			t.Errorf("Expected FBytes %v, got %v", expected.FBytes, response.FBytes)
		}
		if response.FFloat != expected.FFloat {
			t.Errorf("Expected FFloat %v, got %v", expected.FFloat, response.FFloat)
		}
	})
}
