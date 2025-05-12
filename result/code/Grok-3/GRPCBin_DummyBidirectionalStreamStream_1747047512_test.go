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
	// Set up connection with timeout
	conn, err := grpc.Dial("grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create client
	client := grpcbin.NewGRPCBinClient(conn)

	t.Run("TestDummyBidirectionalStreamStream", func(t *testing.T) {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Establish bidirectional stream
		stream, err := client.DummyBidirectionalStreamStream(ctx)
		if err != nil {
			t.Fatalf("Failed to create bidirectional stream: %v", err)
		}

		// Test data to send
		testMsg := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"test1", "test2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
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
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
		}

		// Send test message to server
		if err := stream.Send(testMsg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}

		// Receive and validate response
		resp, err := stream.Recv()
		if err != nil && err != io.EOF {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate received response matches sent data
		if resp.GetFString() != testMsg.GetFString() {
			t.Errorf("Expected FString %q, got %q", testMsg.GetFString(), resp.GetFString())
		}
		if len(resp.GetFStrings()) != len(testMsg.GetFStrings()) {
			t.Errorf("Expected FStrings length %d, got %d", len(testMsg.GetFStrings()), len(resp.GetFStrings()))
		}
		if resp.GetFInt32() != testMsg.GetFInt32() {
			t.Errorf("Expected FInt32 %d, got %d", testMsg.GetFInt32(), resp.GetFInt32())
		}
		if resp.GetFEnum() != testMsg.GetFEnum() {
			t.Errorf("Expected FEnum %v, got %v", testMsg.GetFEnum(), resp.GetFEnum())
		}
		if resp.GetFSub().GetFString() != testMsg.GetFSub().GetFString() {
			t.Errorf("Expected FSub.FString %q, got %q", testMsg.GetFSub().GetFString(), resp.GetFSub().GetFString())
		}
		if resp.GetFBool() != testMsg.GetFBool() {
			t.Errorf("Expected FBool %v, got %v", testMsg.GetFBool(), resp.GetFBool())
		}
		if resp.GetFInt64() != testMsg.GetFInt64() {
			t.Errorf("Expected FInt64 %d, got %d", testMsg.GetFInt64(), resp.GetFInt64())
		}
		if string(resp.GetFBytes()) != string(testMsg.GetFBytes()) {
			t.Errorf("Expected FBytes %v, got %v", testMsg.GetFBytes(), resp.GetFBytes())
		}
		if resp.GetFFloat() != testMsg.GetFFloat() {
			t.Errorf("Expected FFloat %f, got %f", testMsg.GetFFloat(), resp.GetFFloat())
		}

		// Close the stream
		if err := stream.CloseSend(); err != nil {
			t.Fatalf("Failed to close stream: %v", err)
		}
	})
}
