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
	serverAddr      = "grpcb.in:9000"
	connectTimeout  = 15 * time.Second
)

func TestGRPCBin_DummyBidirectionalStreamStream(t *testing.T) {
	// Set up connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Create stream client
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Test data for positive testing
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
		FInt64:   1234567890,
		FInt64S:  []int64{1, 2, 3},
		FBytes:   []byte("test-bytes"),
		FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2, 3.3},
	}

	// Send test message to server
	err = stream.Send(testMessage)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Receive response from server
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	// Validate client response (what we sent)
	if testMessage.FString != "test-string" {
		t.Errorf("Client sent FString = %v; want test-string", testMessage.FString)
	}
	if len(testMessage.FStrings) != 2 {
		t.Errorf("Client sent FStrings length = %v; want 2", len(testMessage.FStrings))
	}

	// Validate server response (what we received)
	if resp.FString != testMessage.FString {
		t.Errorf("Server response FString = %v; want %v", resp.FString, testMessage.FString)
	}
	if len(resp.FStrings) != len(testMessage.FStrings) {
		t.Errorf("Server response FStrings length = %v; want %v", len(resp.FStrings), len(testMessage.FStrings))
	}
	if resp.FInt32 != testMessage.FInt32 {
		t.Errorf("Server response FInt32 = %v; want %v", resp.FInt32, testMessage.FInt32)
	}
	if resp.FEnum != testMessage.FEnum {
		t.Errorf("Server response FEnum = %v; want %v", resp.FEnum, testMessage.FEnum)
	}
	if resp.FSub.FString != testMessage.FSub.FString {
		t.Errorf("Server response FSub.FString = %v; want %v", resp.FSub.FString, testMessage.FSub.FString)
	}
	if resp.FBool != testMessage.FBool {
		t.Errorf("Server response FBool = %v; want %v", resp.FBool, testMessage.FBool)
	}
	if resp.FFloat != testMessage.FFloat {
		t.Errorf("Server response FFloat = %v; want %v", resp.FFloat, testMessage.FFloat)
	}
}
