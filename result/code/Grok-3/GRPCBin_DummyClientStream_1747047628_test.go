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
	grpcAddress    = "grpcb.in:9000"
	connectTimeout = 15 * time.Second
)

func TestGRPCBin_DummyClientStream(t *testing.T) {
	// Set up connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	// Test positive case with valid input
	t.Run("PositiveTest_ValidStreamInput", func(t *testing.T) {
		stream, err := client.DummyClientStream(context.Background())
		if err != nil {
			t.Fatalf("Failed to create stream: %v", err)
		}

		// Prepare test data for streaming
		testMessages := make([]*pb.DummyMessage, 10)
		for i := 0; i < 10; i++ {
			testMessages[i] = &pb.DummyMessage{
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
				FFloat:   float32(i) + 0.5,
				FFloats:  []float32{1.1, 2.2},
			}
		}

		// Send 10 messages
		for i := 0; i < 10; i++ {
			err := stream.Send(testMessages[i])
			if err != nil {
				t.Fatalf("Failed to send message %d: %v", i, err)
			}
		}

		// Close send stream and receive response
		response, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate server response (should return the last message sent)
		expected := testMessages[9]
		assert.Equal(t, expected.FString, response.FString, "String field mismatch")
		assert.Equal(t, expected.FStrings, response.FStrings, "Strings field mismatch")
		assert.Equal(t, expected.FInt32, response.FInt32, "Int32 field mismatch")
		assert.Equal(t, expected.FInt32S, response.FInt32S, "Int32s field mismatch")
		assert.Equal(t, expected.FEnum, response.FEnum, "Enum field mismatch")
		assert.Equal(t, expected.FEnums, response.FEnums, "Enums field mismatch")
		assert.Equal(t, expected.FSub.FString, response.FSub.FString, "Sub field mismatch")
		assert.Equal(t, len(expected.FSubs), len(response.FSubs), "Subs length mismatch")
		assert.Equal(t, expected.FBool, response.FBool, "Bool field mismatch")
		assert.Equal(t, expected.FBools, response.FBools, "Bools field mismatch")
		assert.Equal(t, expected.FInt64, response.FInt64, "Int64 field mismatch")
		assert.Equal(t, expected.FInt64S, response.FInt64S, "Int64s field mismatch")
		assert.Equal(t, expected.FBytes, response.FBytes, "Bytes field mismatch")
		assert.Equal(t, len(expected.FBytess), len(response.FBytess), "Bytess length mismatch")
		assert.Equal(t, expected.FFloat, response.FFloat, "Float field mismatch")
		assert.Equal(t, expected.FFloats, response.FFloats, "Floats field mismatch")
	})
}
