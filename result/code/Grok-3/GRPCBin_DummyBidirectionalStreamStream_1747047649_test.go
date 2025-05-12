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

func TestGRPCBinService(t *testing.T) {
	// Set up connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	t.Run("TestDummyBidirectionalStreamStream_ValidRequest", func(t *testing.T) {
		// Create bidirectional stream
		stream, err := client.DummyBidirectionalStreamStream(ctx)
		if err != nil {
			t.Fatalf("Failed to create bidirectional stream: %v", err)
		}

		// Prepare a valid request message
		req := &pb.DummyMessage{
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
			FInt64:   1234567890,
			FInt64S:  []int64{1, 2, 3},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2, 3.3},
		}

		// Send the request
		err = stream.Send(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}

		// Receive and validate response
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate client response
		assert.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, req.FString, resp.FString, "String field should match")
		assert.Equal(t, req.FStrings, resp.FStrings, "Strings field should match")
		assert.Equal(t, req.FInt32, resp.FInt32, "Int32 field should match")
		assert.Equal(t, req.FInt32S, resp.FInt32S, "Int32s field should match")
		assert.Equal(t, req.FEnum, resp.FEnum, "Enum field should match")
		assert.Equal(t, req.FEnums, resp.FEnums, "Enums field should match")
		assert.Equal(t, req.FSub.FString, resp.FSub.FString, "Sub field should match")
		assert.Equal(t, req.FBool, resp.FBool, "Bool field should match")
		assert.Equal(t, req.FBools, resp.FBools, "Bools field should match")
		assert.Equal(t, req.FInt64, resp.FInt64, "Int64 field should match")
		assert.Equal(t, req.FInt64S, resp.FInt64S, "Int64s field should match")
		assert.Equal(t, req.FBytes, resp.FBytes, "Bytes field should match")
		assert.Equal(t, req.FFloat, resp.FFloat, "Float field should match")
		assert.Equal(t, req.FFloats, resp.FFloats, "Floats field should match")

		// Close the stream
		err = stream.CloseSend()
		if err != nil {
			t.Fatalf("Failed to close stream: %v", err)
		}
	})
}
