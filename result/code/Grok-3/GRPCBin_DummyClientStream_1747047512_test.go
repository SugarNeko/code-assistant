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

func TestGRPCBinClientStream(t *testing.T) {
	// Set up gRPC connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Test case for positive testing
	t.Run("PositiveTest_ClientStream", func(t *testing.T) {
		// Create a client stream
		stream, err := client.DummyClientStream(context.Background())
		if err != nil {
			t.Fatalf("Failed to create client stream: %v", err)
		}

		// Send 10 messages to the server
		for i := 0; i < 10; i++ {
			msg := &grpcbin.DummyMessage{
				FString:  "Test message",
				FStrings: []string{"test1", "test2"},
				FInt32:   int32(i),
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "Sub test",
				},
				FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "Sub1"}, {FString: "Sub2"}},
				FBool:    true,
				FBools:   []bool{true, false},
				FInt64:   int64(i),
				FInt64S:  []int64{10, 20, 30},
				FBytes:   []byte("test bytes"),
				FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:   1.23,
				FFloats:  []float32{1.1, 2.2, 3.3},
			}

			if err := stream.Send(msg); err != nil {
				t.Fatalf("Failed to send message %d: %v", i, err)
			}
		}

		// Close the send stream and receive the response
		response, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate the server response (last message sent should be returned)
		expected := &grpcbin.DummyMessage{
			FString:  "Test message",
			FStrings: []string{"test1", "test2"},
			FInt32:   9, // Last message index
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "Sub test",
			},
			FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "Sub1"}, {FString: "Sub2"}},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   9, // Last message index
			FInt64S:  []int64{10, 20, 30},
			FBytes:   []byte("test bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   1.23,
			FFloats:  []float32{1.1, 2.2, 3.3},
		}

		if response.FString != expected.FString {
			t.Errorf("Expected FString %v, got %v", expected.FString, response.FString)
		}
		if response.FInt32 != expected.FInt32 {
			t.Errorf("Expected FInt32 %v, got %v", expected.FInt32, response.FInt32)
		}
		if response.FEnum != expected.FEnum {
			t.Errorf("Expected FEnum %v, got %v", expected.FEnum, response.FEnum)
		}
		if response.FBool != expected.FBool {
			t.Errorf("Expected FBool %v, got %v", expected.FBool, response.FBool)
		}
		if response.FFloat != expected.FFloat {
			t.Errorf("Expected FFloat %v, got %v", expected.FFloat, response.FFloat)
		}
	})
}
