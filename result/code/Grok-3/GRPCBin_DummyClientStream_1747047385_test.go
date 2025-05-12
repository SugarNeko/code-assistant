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
	// Set up connection to the gRPC server with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Test Case: Positive Testing for DummyClientStream
	t.Run("ValidRequestResponse", func(t *testing.T) {
		stream, err := client.DummyClientStream(ctx)
		if err != nil {
			t.Fatalf("Failed to create stream: %v", err)
		}

		// Send 10 messages to the server
		expectedLastMessage := &grpcbin.DummyMessage{
			FString:   "LastMessage",
			FStrings:  []string{"str1", "str2"},
			FInt32:    42,
			FInt32S:   []int32{1, 2, 3},
			FEnum:     grpcbin.DummyMessage_ENUM_1,
			FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
			FSub:      &grpcbin.DummyMessage_Sub{FString: "sub"},
			FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:     true,
			FBools:    []bool{true, false},
			FInt64:    100,
			FInt64S:   []int64{10, 20},
			FBytes:    []byte("test bytes"),
			FBytess:   [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:    3.14,
			FFloats:   []float32{1.1, 2.2},
		}

		for i := 0; i < 9; i++ {
			msg := &grpcbin.DummyMessage{
				FString:   "Message" + string(rune(i)),
				FInt32:    int32(i),
			}
			if err := stream.Send(msg); err != nil {
				t.Fatalf("Failed to send message %d: %v", i, err)
			}
		}

		// Send the last message
		if err := stream.Send(expectedLastMessage); err != nil {
			t.Fatalf("Failed to send last message: %v", err)
		}

		// Receive the response from the server
		response, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate the server response (last message should be returned)
		if response.FString != expectedLastMessage.FString {
			t.Errorf("Expected FString to be %v, got %v", expectedLastMessage.FString, response.FString)
		}
		if response.FInt32 != expectedLastMessage.FInt32 {
			t.Errorf("Expected FInt32 to be %v, got %v", expectedLastMessage.FInt32, response.FInt32)
		}
		if response.FBool != expectedLastMessage.FBool {
			t.Errorf("Expected FBool to be %v, got %v", expectedLastMessage.FBool, response.FBool)
		}
		if response.FFloat != expectedLastMessage.FFloat {
			t.Errorf("Expected FFloat to be %v, got %v", expectedLastMessage.FFloat, response.FFloat)
		}
		if response.FEnum != expectedLastMessage.FEnum {
			t.Errorf("Expected FEnum to be %v, got %v", expectedLastMessage.FEnum, response.FEnum)
		}
	})
}
