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
	grpcServerAddr = "grpcb.in:9000"
	connectTimeout = 15 * time.Second
)

func TestDummyClientStream(t *testing.T) {
	// Set up connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, grpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Test case: Positive testing for DummyClientStream
	t.Run("PositiveTest_DummyClientStream", func(t *testing.T) {
		stream, err := client.DummyClientStream(context.Background())
		if err != nil {
			t.Fatalf("Failed to create stream: %v", err)
		}

		// Send 10 messages to the server
		testMessage := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"str1", "str2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "sub-test",
			},
			FSubs: []*grpcbin.DummyMessage_Sub{
				{FString: "sub1"},
				{FString: "sub2"},
			},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   100,
			FInt64S:  []int64{10, 20},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("byte1"), []byte("byte2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
		}

		for i := 0; i < 10; i++ {
			if err := stream.Send(testMessage); err != nil {
				t.Fatalf("Failed to send message %d: %v", i+1, err)
			}
		}

		// Receive the response from the server
		response, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate the server response (should return the last sent message)
		if response.FString != testMessage.FString {
			t.Errorf("Expected FString to be %q, got %q", testMessage.FString, response.FString)
		}
		if len(response.FStrings) != len(testMessage.FStrings) {
			t.Errorf("Expected FStrings length to be %d, got %d", len(testMessage.FStrings), len(response.FStrings))
		}
		if response.FInt32 != testMessage.FInt32 {
			t.Errorf("Expected FInt32 to be %d, got %d", testMessage.FInt32, response.FInt32)
		}
		if response.FEnum != testMessage.FEnum {
			t.Errorf("Expected FEnum to be %v, got %v", testMessage.FEnum, response.FEnum)
		}
		if response.FSub.FString != testMessage.FSub.FString {
			t.Errorf("Expected FSub.FString to be %q, got %q", testMessage.FSub.FString, response.FSub.FString)
		}
		if response.FBool != testMessage.FBool {
			t.Errorf("Expected FBool to be %v, got %v", testMessage.FBool, response.FBool)
		}
		if response.FInt64 != testMessage.FInt64 {
			t.Errorf("Expected FInt64 to be %d, got %d", testMessage.FInt64, response.FInt64)
		}
		if string(response.FBytes) != string(testMessage.FBytes) {
			t.Errorf("Expected FBytes to be %q, got %q", testMessage.FBytes, response.FBytes)
		}
		if response.FFloat != testMessage.FFloat {
			t.Errorf("Expected FFloat to be %f, got %f", testMessage.FFloat, response.FFloat)
		}
	})
}
