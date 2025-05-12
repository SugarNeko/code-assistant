package grpcbin

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverAddr     = "grpcb.in:9000"
	connectTimeout = 15 * time.Second
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

	client := NewGRPCBinClient(conn)

	// Create bidirectional stream
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create bidirectional stream: %v", err)
	}

	// Test case 1: Send and receive a simple message
	testMsg := &DummyMessage{
		FString: "test-message",
		FInt32:  42,
		FEnum:   DummyMessage_ENUM_1,
		FSub: &DummyMessage_Sub{
			FString: "sub-test",
		},
		FBool:  true,
		FInt64: 123456789,
		FBytes: []byte("test-bytes"),
		FFloat: 3.14,
	}

	err = stream.Send(testMsg)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Receive response and validate
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive message: %v", err)
	}

	// Validate response fields
	if resp.FString != testMsg.FString {
		t.Errorf("Expected FString to be %q, got %q", testMsg.FString, resp.FString)
	}
	if resp.FInt32 != testMsg.FInt32 {
		t.Errorf("Expected FInt32 to be %d, got %d", testMsg.FInt32, resp.FInt32)
	}
	if resp.FEnum != testMsg.FEnum {
		t.Errorf("Expected FEnum to be %v, got %v", testMsg.FEnum, resp.FEnum)
	}
	if resp.FSub.FString != testMsg.FSub.FString {
		t.Errorf("Expected FSub.FString to be %q, got %q", testMsg.FSub.FString, resp.FSub.FString)
	}
	if resp.FBool != testMsg.FBool {
		t.Errorf("Expected FBool to be %v, got %v", testMsg.FBool, resp.FBool)
	}
	if resp.FInt64 != testMsg.FInt64 {
		t.Errorf("Expected FInt64 to be %d, got %d", testMsg.FInt64, resp.FInt64)
	}
	if string(resp.FBytes) != string(testMsg.FBytes) {
		t.Errorf("Expected FBytes to be %v, got %v", testMsg.FBytes, resp.FBytes)
	}
	if resp.FFloat != testMsg.FFloat {
		t.Errorf("Expected FFloat to be %f, got %f", testMsg.FFloat, resp.FFloat)
	}

	// Test case 2: Send message with repeated fields
	testMsgRepeated := &DummyMessage{
		FStrings: []string{"str1", "str2"},
		FInt32S:  []int32{1, 2, 3},
		FEnums:   []DummyMessage_Enum{DummyMessage_ENUM_0, DummyMessage_ENUM_1},
		FSubs: []*DummyMessage_Sub{
			{FString: "sub1"},
			{FString: "sub2"},
		},
	}

	err = stream.Send(testMsgRepeated)
	if err != nil {
		t.Fatalf("Failed to send message with repeated fields: %v", err)
	}

	// Receive and validate repeated fields
	respRepeated, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive message with repeated fields: %v", err)
	}

	if len(respRepeated.FStrings) != len(testMsgRepeated.FStrings) {
		t.Errorf("Expected FStrings length to be %d, got %d", len(testMsgRepeated.FStrings), len(respRepeated.FStrings))
	}
	if len(respRepeated.FInt32S) != len(testMsgRepeated.FInt32S) {
		t.Errorf("Expected FInt32s length to be %d, got %d", len(testMsgRepeated.FInt32S), len(respRepeated.FInt32S))
	}
	if len(respRepeated.FEnums) != len(testMsgRepeated.FEnums) {
		t.Errorf("Expected FEnums length to be %d, got %d", len(testMsgRepeated.FEnums), len(respRepeated.FEnums))
	}
	if len(respRepeated.FSubs) != len(testMsgRepeated.FSubs) {
		t.Errorf("Expected FSubs length to be %d, got %d", len(testMsgRepeated.FSubs), len(respRepeated.FSubs))
	}

	// Close the stream
	if err := stream.CloseSend(); err != nil {
		t.Fatalf("Failed to close stream: %v", err)
	}
}
