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
	// Set up connection to the gRPC server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, grpcServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Test positive case: Send valid stream of messages and validate response
	t.Run("PositiveTest_ValidStreamMessages", func(t *testing.T) {
		stream, err := client.DummyClientStream(context.Background())
		if err != nil {
			t.Fatalf("Failed to create stream: %v", err)
		}

		// Send 10 valid DummyMessage requests
		expectedLastMessage := &grpcbin.DummyMessage{}
		for i := 0; i < 10; i++ {
			msg := &grpcbin.DummyMessage{
				FString:  "test-string-" + string(rune(i)),
				FStrings: []string{"str1", "str2"},
				FInt32:   int32(i),
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub-test",
				},
				FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
				FBool:    true,
				FBools:   []bool{true, false},
				FInt64:   int64(i),
				FInt64S:  []int64{10, 20},
				FBytes:   []byte("test-bytes"),
				FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:   float32(i) + 0.5,
				FFloats:  []float32{1.1, 2.2},
			}

			if i == 9 {
				expectedLastMessage = msg
			}

			if err := stream.Send(msg); err != nil {
				t.Fatalf("Failed to send message %d: %v", i, err)
			}
		}

		// Receive and validate the response (should be the last sent message)
		response, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate response fields
		if response.FString != expectedLastMessage.FString {
			t.Errorf("Expected FString to be %s, got %s", expectedLastMessage.FString, response.FString)
		}
		if len(response.FStrings) != len(expectedLastMessage.FStrings) {
			t.Errorf("Expected FStrings length to be %d, got %d", len(expectedLastMessage.FStrings), len(response.FStrings))
		}
		if response.FInt32 != expectedLastMessage.FInt32 {
			t.Errorf("Expected FInt32 to be %d, got %d", expectedLastMessage.FInt32, response.FInt32)
		}
		if response.FEnum != expectedLastMessage.FEnum {
			t.Errorf("Expected FEnum to be %v, got %v", expectedLastMessage.FEnum, response.FEnum)
		}
		if response.FSub.FString != expectedLastMessage.FSub.FString {
			t.Errorf("Expected FSub.FString to be %s, got %s", expectedLastMessage.FSub.FString, response.FSub.FString)
		}
		if response.FBool != expectedLastMessage.FBool {
			t.Errorf("Expected FBool to be %v, got %v", expectedLastMessage.FBool, response.FBool)
		}
		if response.FInt64 != expectedLastMessage.FInt64 {
			t.Errorf("Expected FInt64 to be %d, got %d", expectedLastMessage.FInt64, response.FInt64)
		}
		if string(response.FBytes) != string(expectedLastMessage.FBytes) {
			t.Errorf("Expected FBytes to be %s, got %s", expectedLastMessage.FBytes, response.FBytes)
		}
		if response.FFloat != expectedLastMessage.FFloat {
			t.Errorf("Expected FFloat to be %f, got %f", expectedLastMessage.FFloat, response.FFloat)
		}
	})
}
