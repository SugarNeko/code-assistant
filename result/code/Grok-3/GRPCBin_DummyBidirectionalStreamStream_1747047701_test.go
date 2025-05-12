package grpcbin_test

import (
	"context"
	"io"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBinService(t *testing.T) {
	// Set up connection to the gRPC server with timeout
	conn, err := grpc.Dial("grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create gRPC client
	client := grpcbin.NewGRPCBinClient(conn)

	t.Run("TestDummyBidirectionalStreamStream", func(t *testing.T) {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Start bidirectional stream
		stream, err := client.DummyBidirectionalStreamStream(ctx)
		if err != nil {
			t.Fatalf("Failed to create bidirectional stream: %v", err)
		}

		// Prepare test message to send
		testMessage := &grpcbin.DummyMessage{
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

		// Send test message to server
		err = stream.Send(testMessage)
		if err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}

		// Receive response from server
		resp, err := stream.Recv()
		if err != nil && err != io.EOF {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate server response
		if resp == nil {
			t.Fatal("Received nil response from server")
		}
		if resp.FString != testMessage.FString {
			t.Errorf("Expected FString to be %s, got %s", testMessage.FString, resp.FString)
		}
		if len(resp.FStrings) != len(testMessage.FStrings) {
			t.Errorf("Expected FStrings length to be %d, got %d", len(testMessage.FStrings), len(resp.FStrings))
		}
		if resp.FInt32 != testMessage.FInt32 {
			t.Errorf("Expected FInt32 to be %d, got %d", testMessage.FInt32, resp.FInt32)
		}
		if resp.FEnum != testMessage.FEnum {
			t.Errorf("Expected FEnum to be %v, got %v", testMessage.FEnum, resp.FEnum)
		}
		if resp.FSub == nil || resp.FSub.FString != testMessage.FSub.FString {
			t.Errorf("Expected FSub.FString to be %s, got %s", testMessage.FSub.FString, resp.FSub.FString)
		}
		if resp.FBool != testMessage.FBool {
			t.Errorf("Expected FBool to be %v, got %v", testMessage.FBool, resp.FBool)
		}
		if resp.FInt64 != testMessage.FInt64 {
			t.Errorf("Expected FInt64 to be %d, got %d", testMessage.FInt64, resp.FInt64)
		}
		if string(resp.FBytes) != string(testMessage.FBytes) {
			t.Errorf("Expected FBytes to be %s, got %s", testMessage.FBytes, resp.FBytes)
		}
		if resp.FFloat != testMessage.FFloat {
			t.Errorf("Expected FFloat to be %f, got %f", testMessage.FFloat, resp.FFloat)
		}

		// Close the send direction of the stream
		err = stream.CloseSend()
		if err != nil {
			t.Fatalf("Failed to close send stream: %v", err)
		}
	})
}
