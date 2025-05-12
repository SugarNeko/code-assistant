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

func TestGRPCBinDummyClientStream(t *testing.T) {
	// Set up connection to the gRPC server with timeout
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client for the GRPCBin service
	client := pb.NewGRPCBinClient(conn)

	// Test positive case for DummyClientStream
	t.Run("PositiveTest_DummyClientStream", func(t *testing.T) {
		// Create a context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Initialize the client stream
		stream, err := client.DummyClientStream(ctx)
		if err != nil {
			t.Fatalf("Failed to create client stream: %v", err)
		}

		// Send 10 DummyMessage requests to the server
		for i := 0; i < 10; i++ {
			msg := &pb.DummyMessage{
				FString:  "test-string",
				FStrings: []string{"test1", "test2"},
				FInt32:   int32(i),
				FInt32S:  []int32{1, 2, 3},
				FEnum:    pb.DummyMessage_ENUM_1,
				FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_1},
				FSub: &pb.DummyMessage_Sub{
					FString: "sub-test",
				},
				FSubs: []*pb.DummyMessage_Sub{
					{FString: "sub1"},
					{FString: "sub2"},
				},
				FBool:    true,
				FBools:   []bool{true, false},
				FInt64:   int64(i),
				FInt64S:  []int64{100, 200},
				FBytes:   []byte("test-bytes"),
				FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:   float32(1.23),
				FFloats:  []float32{1.1, 2.2},
			}

			if err := stream.Send(msg); err != nil {
				t.Fatalf("Failed to send message %d: %v", i, err)
			}
		}

		// Close the stream and receive the response
		response, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate the server response (last sent message should be returned)
		assert.NotNil(t, response, "Response should not be nil")
		assert.Equal(t, "test-string", response.FString, "FString should match the last sent message")
		assert.Equal(t, []string{"test1", "test2"}, response.FStrings, "FStrings should match the last sent message")
		assert.Equal(t, int32(9), response.FInt32, "FInt32 should match the last sent message")
		assert.Equal(t, []int32{1, 2, 3}, response.FInt32S, "FInt32s should match the last sent message")
		assert.Equal(t, pb.DummyMessage_ENUM_1, response.FEnum, "FEnum should match the last sent message")
		assert.Equal(t, true, response.FBool, "FBool should match the last sent message")
		assert.Equal(t, int64(9), response.FInt64, "FInt64 should match the last sent message")
		assert.Equal(t, float32(1.23), response.FFloat, "FFloat should match the last sent message")
		assert.Equal(t, "sub-test", response.FSub.FString, "FSub.FString should match the last sent message")
	})
}
