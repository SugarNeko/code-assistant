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

const (
	grpcBinAddress = "grpcb.in:9000"
	connectTimeout = 15 * time.Second
)

func TestGRPCBin_DummyBidirectionalStreamStream(t *testing.T) {
	// Set up connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, grpcBinAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial GRPCBin server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Test case for bidirectional streaming
	t.Run("PositiveTest_BidirectionalStream", func(t *testing.T) {
		stream, err := client.DummyBidirectionalStreamStream(context.Background())
		if err != nil {
			t.Fatalf("Failed to create bidirectional stream: %v", err)
		}

		// Send a test message
		testMessage := &grpcbin.DummyMessage{
			FString:  "test_string",
			FStrings: []string{"str1", "str2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "sub_test",
			},
			FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   1000000,
			FInt64S:  []int64{100, 200},
			FBytes:   []byte("test_bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
		}

		// Send the message to the server
		err = stream.Send(testMessage)
		if err != nil {
			t.Fatalf("Failed to send message to server: %v", err)
		}

		// Receive and validate the response from the server
		response, err := stream.Recv()
		if err != nil && err != io.EOF {
			t.Fatalf("Failed to receive message from server: %v", err)
		}

		// Validate the response fields
		if response.FString != testMessage.FString {
			t.Errorf("Expected FString to be %s, got %s", testMessage.FString, response.FString)
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
			t.Errorf("Expected FSub.FString to be %s, got %s", testMessage.FSub.FString, response.FSub.FString)
		}
		if response.FBool != testMessage.FBool {
			t.Errorf("Expected FBool to be %v, got %v", testMessage.FBool, response.FBool)
		}
		if response.FInt64 != testMessage.FInt64 {
			t.Errorf("Expected FInt64 to be %d, got %d", testMessage.FInt64, response.FInt64)
		}
		if string(response.FBytes) != string(testMessage.FBytes) {
			t.Errorf("Expected FBytes to be %s, got %s", testMessage.FBytes, response.FBytes)
		}
		if response.FFloat != testMessage.FFloat {
			t.Errorf("Expected FFloat to be %f, got %f", testMessage.FFloat, response.FFloat)
		}

		// Close the stream
		err = stream.CloseSend()
		if err != nil {
			t.Fatalf("Failed to close stream: %v", err)
		}
	})
}
