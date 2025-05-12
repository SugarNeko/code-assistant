package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBin_DummyClientStream(t *testing.T) {
	// Set up connection to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := grpcbin.NewGRPCBinClient(conn)

	// Test positive case for DummyClientStream
	t.Run("PositiveTest_DummyClientStream", func(t *testing.T) {
		// Create a stream client
		stream, err := client.DummyClientStream(context.Background())
		if err != nil {
			t.Fatalf("Failed to create stream: %v", err)
		}

		// Prepare test messages to send
		testMessages := []*grpcbin.DummyMessage{
			{
				FString:  "Test1",
				FInt32:   1,
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FSub:     &grpcbin.DummyMessage_Sub{FString: "SubTest1"},
				FBool:    true,
				FInt64:   100,
				FBytes:   []byte("bytes1"),
				FFloat:   1.1,
				FStrings: []string{"str1", "str2"},
				FInt32S:  []int32{1, 2},
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
				FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "Sub1"}, {FString: "Sub2"}},
				FBools:   []bool{true, false},
				FInt64S:  []int64{100, 200},
				FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloats:  []float32{1.1, 2.2},
			},
			{
				FString:  "Test2",
				FInt32:   2,
				FEnum:    grpcbin.DummyMessage_ENUM_2,
				FSub:     &grpcbin.DummyMessage_Sub{FString: "SubTest2"},
				FBool:    false,
				FInt64:   200,
				FBytes:   []byte("bytes2"),
				FFloat:   2.2,
				FStrings: []string{"str3", "str4"},
				FInt32S:  []int32{3, 4},
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
				FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "Sub3"}, {FString: "Sub4"}},
				FBools:   []bool{false, true},
				FInt64S:  []int64{300, 400},
				FBytess:  [][]byte{[]byte("bytes3"), []byte("bytes4")},
				FFloats:  []float32{3.3, 4.4},
			},
		}

		// Send messages to the stream
		for i := 0; i < 10; i++ {
			msgIndex := i % len(testMessages)
			if err := stream.Send(testMessages[msgIndex]); err != nil {
				t.Fatalf("Failed to send message %d: %v", i+1, err)
			}
		}

		// Close the stream and receive the response
		response, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate the server response (should return the last sent message)
		expectedLastMessage := testMessages[len(testMessages)-1]
		if response.FString != expectedLastMessage.FString {
			t.Errorf("Expected FString to be %s, got %s", expectedLastMessage.FString, response.FString)
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
			t.Errorf("Expected FBytes to be %v, got %v", expectedLastMessage.FBytes, response.FBytes)
		}
		if response.FFloat != expectedLastMessage.FFloat {
			t.Errorf("Expected FFloat to be %f, got %f", expectedLastMessage.FFloat, response.FFloat)
		}
	})
}
