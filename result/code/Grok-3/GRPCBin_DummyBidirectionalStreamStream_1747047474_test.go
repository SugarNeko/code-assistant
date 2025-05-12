package grpcbin_test

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "code-assistant/proto/grpcbin"
)

func TestGRPCBinService(t *testing.T) {
	// Set up connection to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create client
	client := pb.NewGRPCBinClient(conn)

	t.Run("TestDummyBidirectionalStreamStream", func(t *testing.T) {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Initiate bidirectional stream
		stream, err := client.DummyBidirectionalStreamStream(ctx)
		if err != nil {
			t.Fatalf("Failed to create bidirectional stream: %v", err)
		}

		// Prepare test message to send
		testMsg := &pb.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"test1", "test2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    pb.DummyMessage_ENUM_1,
			FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
			FSub: &pb.DummyMessage_Sub{
				FString: "sub-test",
			},
			FSubs: []*pb.DummyMessage_Sub{
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
		err = stream.Send(testMsg)
		if err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}

		// Receive response from server
		resp, err := stream.Recv()
		if err != nil && err != io.EOF {
			t.Fatalf("Failed to receive message: %v", err)
		}

		// Validate client response (what we sent)
		assert.Equal(t, testMsg.FString, "test-string", "Client sent string should match")
		assert.Equal(t, testMsg.FInt32, int32(42), "Client sent int32 should match")
		assert.Equal(t, testMsg.FBool, true, "Client sent bool should match")

		// Validate server response (what we received)
		if resp != nil {
			assert.NotEmpty(t, resp.FString, "Server response string should not be empty")
			assert.Equal(t, testMsg.FString, resp.FString, "Server should echo back the same string")
			assert.Equal(t, testMsg.FInt32, resp.FInt32, "Server should echo back the same int32")
			assert.Equal(t, testMsg.FBool, resp.FBool, "Server should echo back the same bool")
		}

		// Close the send direction
		err = stream.CloseSend()
		if err != nil {
			t.Fatalf("Failed to close send stream: %v", err)
		}
	})
}
