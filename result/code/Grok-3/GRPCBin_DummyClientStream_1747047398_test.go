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
	// Set up connection to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a client for the GRPCBin service
	client := grpcbin.NewGRPCBinClient(conn)

	// Test case for positive testing with valid request data
	t.Run("PositiveTest_ValidRequest", func(t *testing.T) {
		// Create a context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Create a stream client for DummyClientStream
		stream, err := client.DummyClientStream(ctx)
		if err != nil {
			t.Fatalf("Failed to create stream: %v", err)
		}

		// Send 10 valid DummyMessage requests
		for i := 0; i < 10; i++ {
			req := &grpcbin.DummyMessage{
				FString:  "test-string",
				FStrings: []string{"test1", "test2"},
				FInt32:   int32(i),
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub-test",
				},
				FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
				FBool:    true,
				FBools:   []bool{true, false},
				FInt64:   int64(i),
				FInt64S:  []int64{10, 20, 30},
				FBytes:   []byte("test-bytes"),
				FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:   float32(1.23),
				FFloats:  []float32{1.1, 2.2, 3.3},
			}

			if err := stream.Send(req); err != nil {
				t.Fatalf("Failed to send message %d: %v", i, err)
			}
		}

		// Close the send stream and receive the response
		resp, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate the server response (should return the last sent message)
		if resp.GetFString() != "test-string" {
			t.Errorf("Expected FString to be 'test-string', got '%s'", resp.GetFString())
		}
		if len(resp.GetFStrings()) != 2 || resp.GetFStrings()[0] != "test1" {
			t.Errorf("Expected FStrings to contain 'test1', got %v", resp.GetFStrings())
		}
		if resp.GetFInt32() != 9 {
			t.Errorf("Expected FInt32 to be 9, got %d", resp.GetFInt32())
		}
		if resp.GetFEnum() != grpcbin.DummyMessage_ENUM_1 {
			t.Errorf("Expected FEnum to be ENUM_1, got %v", resp.GetFEnum())
		}
		if resp.GetFSub().GetFString() != "sub-test" {
			t.Errorf("Expected FSub.FString to be 'sub-test', got '%s'", resp.GetFSub().GetFString())
		}
		if !resp.GetFBool() {
			t.Errorf("Expected FBool to be true, got %v", resp.GetFBool())
		}
		if resp.GetFInt64() != 9 {
			t.Errorf("Expected FInt64 to be 9, got %d", resp.GetFInt64())
		}
		if string(resp.GetFBytes()) != "test-bytes" {
			t.Errorf("Expected FBytes to be 'test-bytes', got '%s'", string(resp.GetFBytes()))
		}
		if resp.GetFFloat() != float32(1.23) {
			t.Errorf("Expected FFloat to be 1.23, got %f", resp.GetFFloat())
		}
	})
}
