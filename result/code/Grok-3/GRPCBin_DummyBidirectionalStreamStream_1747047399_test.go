package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "code-assistant/proto/grpcbin"
)

const (
	serverAddr     = "grpcb.in:9000"
	connectTimeout = 15 * time.Second
)

func TestGRPCBinService_DummyBidirectionalStreamStream(t *testing.T) {
	// Set up connection to the gRPC server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	// Test positive case for bidirectional streaming
	t.Run("PositiveTest_BidirectionalStream", func(t *testing.T) {
		// Create a stream
		stream, err := client.DummyBidirectionalStreamStream(context.Background())
		if err != nil {
			t.Fatalf("Failed to create stream: %v", err)
		}

		// Prepare test message
		testMsg := &pb.DummyMessage{
			FString:   "test-string",
			FStrings:  []string{"test1", "test2"},
			FInt32:    42,
			FInt32S:   []int32{1, 2, 3},
			FEnum:     pb.DummyMessage_ENUM_1,
			FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
			FSub:      &pb.DummyMessage_Sub{FString: "sub-test"},
			FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:     true,
			FBools:    []bool{true, false},
			FInt64:    1000000,
			FInt64S:   []int64{100, 200},
			FBytes:    []byte("test-bytes"),
			FBytess:   [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:    3.14,
			FFloats:   []float32{1.1, 2.2},
		}

		// Send test message to server
		err = stream.Send(testMsg)
		if err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}

		// Receive response from server
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive message: %v", err)
		}

		// Validate client response (what we sent)
		assert.Equal(t, testMsg.FString, resp.FString, "String field mismatch")
		assert.Equal(t, testMsg.FStrings, resp.FStrings, "Strings field mismatch")
		assert.Equal(t, testMsg.FInt32, resp.FInt32, "Int32 field mismatch")
		assert.Equal(t, testMsg.FInt32S, resp.FInt32S, "Int32s field mismatch")
		assert.Equal(t, testMsg.FEnum, resp.FEnum, "Enum field mismatch")
		assert.Equal(t, testMsg.FEnums, resp.FEnums, "Enums field mismatch")
		assert.Equal(t, testMsg.FSub.FString, resp.FSub.FString, "Sub field mismatch")
		assert.Len(t, resp.FSubs, len(testMsg.FSubs), "Subs length mismatch")
		assert.Equal(t, testMsg.FBool, resp.FBool, "Bool field mismatch")
		assert.Equal(t, testMsg.FBools, resp.FBools, "Bools field mismatch")
		assert.Equal(t, testMsg.FInt64, resp.FInt64, "Int64 field mismatch")
		assert.Equal(t, testMsg.FInt64S, resp.FInt64S, "Int64s field mismatch")
		assert.Equal(t, testMsg.FBytes, resp.FBytes, "Bytes field mismatch")
		assert.Len(t, resp.FBytess, len(testMsg.FBytess), "Bytess length mismatch")
		assert.Equal(t, testMsg.FFloat, resp.FFloat, "Float field mismatch")
		assert.Equal(t, testMsg.FFloats, resp.FFloats, "Floats field mismatch")

		// Close the stream
		err = stream.CloseSend()
		if err != nil {
			t.Fatalf("Failed to close stream: %v", err)
		}
	})
}
