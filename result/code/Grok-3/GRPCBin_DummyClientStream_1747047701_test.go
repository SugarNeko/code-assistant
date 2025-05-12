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

func TestGRPCBin_DummyClientStream(t *testing.T) {
	// Set up connection to the gRPC server with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	// Test case for positive testing with valid input
	t.Run("PositiveTest_ValidDummyMessages", func(t *testing.T) {
		stream, err := client.DummyClientStream(ctx)
		if err != nil {
			t.Fatalf("Failed to create stream: %v", err)
		}

		// Send 10 valid DummyMessage requests
		for i := 0; i < 10; i++ {
			msg := &pb.DummyMessage{
				FString:  "test-string",
				FStrings: []string{"str1", "str2"},
				FInt32:   int32(i),
				FInt32S:  []int32{1, 2, 3},
				FEnum:    pb.DummyMessage_ENUM_1,
				FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_1},
				FSub: &pb.DummyMessage_Sub{
					FString: "sub-test-string",
				},
				FSubs:    []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
				FBool:    true,
				FBools:   []bool{true, false},
				FInt64:   int64(i),
				FInt64S:  []int64{10, 20},
				FBytes:   []byte("test-bytes"),
				FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:   float32(1.23),
				FFloats:  []float32{1.1, 2.2},
			}
			if err := stream.Send(msg); err != nil {
				t.Fatalf("Failed to send message %d: %v", i, err)
			}
		}

		// Close the stream and get the response
		response, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate server response
		assert.NotNil(t, response, "Response should not be nil")
		assert.Equal(t, "test-string", response.FString, "Response FString should match the sent value")
		assert.Equal(t, int32(9), response.FInt32, "Response FInt32 should match the last sent value")
		assert.Equal(t, true, response.FBool, "Response FBool should match the sent value")
		assert.Equal(t, float32(1.23), response.FFloat, "Response FFloat should match the sent value")
		assert.Equal(t, pb.DummyMessage_ENUM_1, response.FEnum, "Response FEnum should match the sent value")
		assert.Equal(t, "sub-test-string", response.FSub.FString, "Response FSub.FString should match the sent value")
	})
}
