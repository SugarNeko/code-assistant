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
	// Set up connection to the gRPC server with a timeout
	conn, err := grpc.Dial("grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client for the GRPCBin service
	client := grpcbin.NewGRPCBinClient(conn)

	t.Run("TestDummyBidirectionalStreamStream", func(t *testing.T) {
		// Create a context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Initiate bidirectional stream
		stream, err := client.DummyBidirectionalStreamStream(ctx)
		if err != nil {
			t.Fatalf("Failed to initiate bidirectional stream: %v", err)
		}

		// Test case 1: Send a valid DummyMessage and validate response
		sendMsg := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"test1", "test2"},
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
			FInt64:   1234567890,
			FInt64S:  []int64{1, 2, 3},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("byte1"), []byte("byte2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2, 3.3},
		}

		// Send the message to the server
		if err := stream.Send(sendMsg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}

		// Receive and validate the response from the server
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate server response matches the sent message (echo behavior)
		if resp.FString != sendMsg.FString {
			t.Errorf("Expected FString to be %s, got %s", sendMsg.FString, resp.FString)
		}
		if len(resp.FStrings) != len(sendMsg.FStrings) {
			t.Errorf("Expected FStrings length to be %d, got %d", len(sendMsg.FStrings), len(resp.FStrings))
		}
		if resp.FInt32 != sendMsg.FInt32 {
			t.Errorf("Expected FInt32 to be %d, got %d", sendMsg.FInt32, resp.FInt32)
		}
		if resp.FEnum != sendMsg.FEnum {
			t.Errorf("Expected FEnum to be %v, got %v", sendMsg.FEnum, resp.FEnum)
		}
		if resp.FSub.FString != sendMsg.FSub.FString {
			t.Errorf("Expected FSub.FString to be %s, got %s", sendMsg.FSub.FString, resp.FSub.FString)
		}
		if resp.FBool != sendMsg.FBool {
			t.Errorf("Expected FBool to be %v, got %v", sendMsg.FBool, resp.FBool)
		}
		if resp.FInt64 != sendMsg.FInt64 {
			t.Errorf("Expected FInt64 to be %d, got %d", sendMsg.FInt64, resp.FInt64)
		}
		if string(resp.FBytes) != string(sendMsg.FBytes) {
			t.Errorf("Expected FBytes to be %s, got %s", sendMsg.FBytes, resp.FBytes)
		}
		if resp.FFloat != sendMsg.FFloat {
			t.Errorf("Expected FFloat to be %f, got %f", sendMsg.FFloat, resp.FFloat)
		}
	})
}
