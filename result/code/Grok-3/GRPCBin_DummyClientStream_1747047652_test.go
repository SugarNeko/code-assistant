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
	// Set up connection to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create client
	client := grpcbin.NewGRPCBinClient(conn)

	t.Run("TestDummyClientStream_ValidInput", func(t *testing.T) {
		// Create a context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Establish client stream
		stream, err := client.DummyClientStream(ctx)
		if err != nil {
			t.Fatalf("Failed to create client stream: %v", err)
		}

		// Send 10 messages to the stream
		for i := 0; i < 10; i++ {
			msg := &grpcbin.DummyMessage{
				FString:  "test_string",
				FStrings: []string{"str1", "str2"},
				FInt32:   int32(i),
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub_test",
				},
				FSubs: []*grpcbin.DummyMessage_Sub{
					{FString: "sub1"},
					{FString: "sub2"},
				},
				FBool:    true,
				FBools:   []bool{true, false},
				FInt64:   int64(i),
				FInt64S:  []int64{100, 200},
				FBytes:   []byte("test_bytes"),
				FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:   3.14,
				FFloats:  []float32{1.1, 2.2},
			}

			if err := stream.Send(msg); err != nil {
				t.Fatalf("Failed to send message %d: %v", i, err)
			}
		}

		// Receive response from server
		response, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate server response
		if response.FString != "test_string" {
			t.Errorf("Expected FString to be 'test_string', got '%s'", response.FString)
		}
		if len(response.FStrings) != 2 || response.FStrings[0] != "str1" || response.FStrings[1] != "str2" {
			t.Errorf("Expected FStrings to be ['str1', 'str2'], got %v", response.FStrings)
		}
		if response.FInt32 != 9 {
			t.Errorf("Expected FInt32 to be 9, got %d", response.FInt32)
		}
		if response.FEnum != grpcbin.DummyMessage_ENUM_1 {
			t.Errorf("Expected FEnum to be ENUM_1, got %v", response.FEnum)
		}
		if response.FSub.FString != "sub_test" {
			t.Errorf("Expected FSub.FString to be 'sub_test', got '%s'", response.FSub.FString)
		}
		if !response.FBool {
			t.Error("Expected FBool to be true, got false")
		}
		if response.FFloat != 3.14 {
			t.Errorf("Expected FFloat to be 3.14, got %f", response.FFloat)
		}
	})
}
