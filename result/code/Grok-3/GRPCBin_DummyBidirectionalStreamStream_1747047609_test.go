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

func TestGRPCBinService_DummyBidirectionalStreamStream(t *testing.T) {
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

	// Test positive case for bidirectional streaming
	t.Run("PositiveTest_BidirectionalStream", func(t *testing.T) {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Initialize bidirectional stream
		stream, err := client.DummyBidirectionalStreamStream(ctx)
		if err != nil {
			t.Fatalf("Failed to create bidirectional stream: %v", err)
		}

		// Prepare test message to send
		testMessage := &pb.DummyMessage{
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
			FInt64:   1000000,
			FInt64S:  []int64{100, 200},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
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

		// Validate server response
		assert.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, testMessage.FString, resp.FString, "String field should match")
		assert.Equal(t, testMessage.FStrings, resp.FStrings, "Strings field should match")
		assert.Equal(t, testMessage.FInt32, resp.FInt32, "Int32 field should match")
		assert.Equal(t, testMessage.FInt32S, resp.FInt32S, "Int32s field should match")
		assert.Equal(t, testMessage.FEnum, resp.FEnum, "Enum field should match")
		assert.Equal(t, testMessage.FEnums, resp.FEnums, "Enums field should match")
		assert.Equal(t, testMessage.FSub.FString, resp.FSub.FString, "Sub field should match")
		assert.Equal(t, testMessage.FBool, resp.FBool, "Bool field should match")
		assert.Equal(t, testMessage.FBools, resp.FBools, "Bools field should match")
		assert.Equal(t, testMessage.FInt64, resp.FInt64, "Int64 field should match")
		assert.Equal(t, testMessage.FInt64S, resp.FInt64S, "Int64s field should match")
		assert.Equal(t, testMessage.FBytes, resp.FBytes, "Bytes field should match")
		assert.Equal(t, testMessage.FFloat, resp.FFloat, "Float field should match")
		assert.Equal(t, testMessage.FFloats, resp.FFloats, "Floats field should match")
	})
}
