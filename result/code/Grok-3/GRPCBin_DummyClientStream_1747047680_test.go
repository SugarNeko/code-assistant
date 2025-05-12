package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBinClientStream(t *testing.T) {
	// Set up connection to gRPC server with timeout
	conn, err := grpc.Dial("grpcb.in:9000", 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create client
	client := grpcbin.NewGRPCBinClient(conn)

	// Test positive case for client streaming
	t.Run("PositiveTest_ClientStream", func(t *testing.T) {
		// Create stream
		stream, err := client.DummyClientStream(context.Background())
		if err != nil {
			t.Fatalf("failed to create stream: %v", err)
		}

		// Send 10 messages
		for i := 0; i < 10; i++ {
			msg := &grpcbin.DummyMessage{
				FString:   "test_string",
				FStrings:  []string{"str1", "str2"},
				FInt32:    int32(i),
				FInt32S:   []int32{1, 2, 3},
				FEnum:     grpcbin.DummyMessage_ENUM_1,
				FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
				FSub:      &grpcbin.DummyMessage_Sub{FString: "sub_string"},
				FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
				FBool:     true,
				FBools:    []bool{true, false},
				FInt64:    int64(i),
				FInt64S:   []int64{100, 200},
				FBytes:    []byte("test_bytes"),
				FBytess:   [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:    float32(1.23),
				FFloats:   []float32{1.1, 2.2},
			}

			if err := stream.Send(msg); err != nil {
				t.Fatalf("failed to send message %d: %v", i, err)
			}
		}

		// Close stream and receive response
		response, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("failed to receive response: %v", err)
		}

		// Validate server response
		if response.FString != "test_string" {
			t.Errorf("expected FString to be 'test_string', got '%s'", response.FString)
		}
		if len(response.FStrings) != 2 {
			t.Errorf("expected FStrings length to be 2, got %d", len(response.FStrings))
		}
		if response.FInt32 != 9 {
			t.Errorf("expected FInt32 to be 9, got %d", response.FInt32)
		}
		if response.FEnum != grpcbin.DummyMessage_ENUM_1 {
			t.Errorf("expected FEnum to be ENUM_1, got %v", response.FEnum)
		}
		if response.FSub.FString != "sub_string" {
			t.Errorf("expected FSub.FString to be 'sub_string', got '%s'", response.FSub.FString)
		}
		if !response.FBool {
			t.Error("expected FBool to be true, got false")
		}
		if response.FInt64 != 9 {
			t.Errorf("expected FInt64 to be 9, got %d", response.FInt64)
		}
		if string(response.FBytes) != "test_bytes" {
			t.Errorf("expected FBytes to be 'test_bytes', got '%s'", string(response.FBytes))
		}
		if response.FFloat != float32(1.23) {
			t.Errorf("expected FFloat to be 1.23, got %f", response.FFloat)
		}
	})
}
