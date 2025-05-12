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
	serverAddr    = "grpcb.in:9000"
	connectTimeout = 15 * time.Second
)

func TestGRPCBin_DummyServerStream(t *testing.T) {
	// Set up connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	// Prepare test data for request
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

	// Test positive case for server streaming
	t.Run("PositiveTest_DummyServerStream", func(t *testing.T) {
		stream, err := client.DummyServerStream(context.Background(), req)
		if err != nil {
			t.Fatalf("Failed to call DummyServerStream: %v", err)
		}

		// Validate server responses (expecting 10 responses as per service definition)
		responseCount := 0
		for {
			resp, err := stream.Recv()
			if err != nil {
				break
			}

			// Validate response fields match the request (or expected transformation)
			assert.Equal(t, req.FString, resp.FString, "Response string field should match request")
			assert.Equal(t, req.FInt32, resp.FInt32, "Response int32 field should match request")
			assert.Equal(t, req.FEnum, resp.FEnum, "Response enum field should match request")
			assert.Equal(t, req.FSub.FString, resp.FSub.FString, "Response sub field should match request")
			assert.Equal(t, req.FBool, resp.FBool, "Response bool field should match request")
			assert.Equal(t, req.FFloat, resp.FFloat, "Response float field should match request")

			responseCount++
		}

		// Validate the number of responses received
		assert.Equal(t, 10, responseCount, "Should receive exactly 10 responses from server stream")
	})
}
