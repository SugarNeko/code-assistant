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
	grpcServerAddress = "grpcb.in:9000"
	connectTimeout    = 15 * time.Second
)

func TestGRPCBin_DummyClientStream(t *testing.T) {
	// Set up gRPC connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, grpcServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Test case for positive testing with valid input
	t.Run("PositiveTest_ValidInput", func(t *testing.T) {
		stream, err := client.DummyClientStream(context.Background())
		if err != nil {
			t.Fatalf("Failed to create stream: %v", err)
		}

		// Send 10 valid DummyMessages
		for i := 0; i < 10; i++ {
			msg := &grpcbin.DummyMessage{
				FString:   "test-string",
				FStrings:  []string{"test1", "test2"},
				FInt32:    int32(i),
				FInt32S:   []int32{1, 2, 3},
				FEnum:     grpcbin.DummyMessage_ENUM_1,
				FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
				FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-test"},
				FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
				FBool:     true,
				FBools:    []bool{true, false},
				FInt64:    int64(i),
				FInt64S:   []int64{10, 20, 30},
				FBytes:    []byte("test-bytes"),
				FBytess:   [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:    float32(1.23),
				FFloats:   []float32{1.1, 2.2, 3.3},
			}

			if err := stream.Send(msg); err != nil {
				t.Fatalf("Failed to send message %d: %v", i, err)
			}
		}

		// Receive the response from the server
		response, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate server response
		if response == nil {
			t.Fatal("Received nil response from server")
		}
		if response.FString != "test-string" {
			t.Errorf("Expected FString to be 'test-string', got '%s'", response.FString)
		}
		if len(response.FStrings) != 2 {
			t.Errorf("Expected FStrings length to be 2, got %d", len(response.FStrings))
		}
		if response.FInt32 != 9 {
			t.Errorf("Expected FInt32 to be 9, got %d", response.FInt32)
		}
		if response.FEnum != grpcbin.DummyMessage_ENUM_1 {
			t.Errorf("Expected FEnum to be ENUM_1, got %v", response.FEnum)
		}
		if response.FSub == nil || response.FSub.FString != "sub-test" {
			t.Errorf("Expected FSub.FString to be 'sub-test', got '%v'", response.FSub)
		}
		if response.FBool != true {
			t.Errorf("Expected FBool to be true, got %v", response.FBool)
		}
		if response.FFloat != float32(1.23) {
			t.Errorf("Expected FFloat to be 1.23, got %f", response.FFloat)
		}
	})
}
