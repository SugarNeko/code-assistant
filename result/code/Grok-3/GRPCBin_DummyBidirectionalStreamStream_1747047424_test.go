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

func TestGRPCBin_DummyBidirectionalStreamStream(t *testing.T) {
	// Set up connection to gRPC server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	// Create a stream for bidirectional communication
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Test case 1: Send a valid request and validate response
	t.Run("PositiveTest_ValidRequest", func(t *testing.T) {
		// Construct a valid DummyMessage request
		req := &pb.DummyMessage{
			FString:  "test_string",
			FStrings: []string{"str1", "str2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    pb.DummyMessage_ENUM_1,
			FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
			FSub: &pb.DummyMessage_Sub{
				FString: "sub_test",
			},
			FSubs: []*pb.DummyMessage_Sub{
				{FString: "sub1"},
				{FString: "sub2"},
			},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   1234567890,
			FInt64S:  []int64{1, 2, 3},
			FBytes:   []byte("test_bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
		}

		// Send the request
		err := stream.Send(req)
		assert.NoError(t, err, "Failed to send request")

		// Receive and validate the response
		resp, err := stream.Recv()
		assert.NoError(t, err, "Failed to receive response")
		assert.NotNil(t, resp, "Response should not be nil")

		// Validate response fields mirror the request (as per service behavior)
		assert.Equal(t, req.FString, resp.FString, "FString should match")
		assert.Equal(t, req.FStrings, resp.FStrings, "FStrings should match")
		assert.Equal(t, req.FInt32, resp.FInt32, "FInt32 should match")
		assert.Equal(t, req.FInt32S, resp.FInt32S, "FInt32S should match")
		assert.Equal(t, req.FEnum, resp.FEnum, "FEnum should match")
		assert.Equal(t, req.FEnums, resp.FEnums, "FEnums should match")
		assert.Equal(t, req.FSub.FString, resp.FSub.FString, "FSub.FString should match")
		assert.Len(t, resp.FSubs, len(req.FSubs), "FSubs length should match")
		assert.Equal(t, req.FBool, resp.FBool, "FBool should match")
		assert.Equal(t, req.FBools, resp.FBools, "FBools should match")
		assert.Equal(t, req.FInt64, resp.FInt64, "FInt64 should match")
		assert.Equal(t, req.FInt64S, resp.FInt64S, "FInt64S should match")
		assert.Equal(t, req.FBytes, resp.FBytes, "FBytes should match")
		assert.Len(t, resp.FBytess, len(req.FBytess), "FBytess length should match")
		assert.Equal(t, req.FFloat, resp.FFloat, "FFloat should match")
		assert.Equal(t, req.FFloats, resp.FFloats, "FFloats should match")
	})

	// Test case 2: Send multiple messages and validate responses
	t.Run("PositiveTest_MultipleMessages", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			req := &pb.DummyMessage{
				FString:  "test_multi_" + string(rune(i)),
				FInt32:   int32(i),
				FStrings: []string{"multi_str" + string(rune(i))},
			}

			err := stream.Send(req)
			assert.NoError(t, err, "Failed to send request %d", i)

			resp, err := stream.Recv()
			assert.NoError(t, err, "Failed to receive response %d", i)
			assert.Equal(t, req.FString, resp.FString, "FString should match for message %d", i)
			assert.Equal(t, req.FInt32, resp.FInt32, "FInt32 should match for message %d", i)
			assert.Equal(t, req.FStrings, resp.FStrings, "FStrings should match for message %d", i)
		}
	})

	// Close the stream
	err = stream.CloseSend()
	assert.NoError(t, err, "Failed to close stream")
}
