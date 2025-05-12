package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverAddress = "grpcb.in:9000"
	connectTimeout = 15 * time.Second
)

func TestGRPCBin_DummyServerStream(t *testing.T) {
	// Set up gRPC connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Test case for positive testing with valid request
	t.Run("PositiveTest_ValidRequest", func(t *testing.T) {
		// Prepare a valid request
		req := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"str1", "str2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "sub-test",
			},
			FSubs: []*grpcbin.DummyMessage_Sub{
				{FString: "sub1"},
				{FString: "sub2"},
			},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   123456789,
			FInt64S:  []int64{1, 2, 3},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("b1"), []byte("b2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
		}

		// Call the server stream endpoint
		stream, err := client.DummyServerStream(context.Background(), req)
		if err != nil {
			t.Fatalf("Failed to call DummyServerStream: %v", err)
		}

		// Validate server responses (expecting 10 responses as per service behavior)
		responseCount := 0
		for {
			resp, err := stream.Recv()
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				t.Fatalf("Error receiving stream response: %v", err)
			}

			// Validate response fields match the request (as per expected behavior)
			assert.Equal(t, req.FString, resp.FString, "Response FString should match request")
			assert.Equal(t, req.FInt32, resp.FInt32, "Response FInt32 should match request")
			assert.Equal(t, req.FEnum, resp.FEnum, "Response FEnum should match request")
			assert.Equal(t, req.FSub.FString, resp.FSub.FString, "Response FSub.FString should match request")
			assert.Equal(t, req.FBool, resp.FBool, "Response FBool should match request")
			assert.Equal(t, req.FFloat, resp.FFloat, "Response FFloat should match request")

			responseCount++
		}

		// Validate the total number of responses received
		assert.Equal(t, 10, responseCount, "Expected 10 responses from server stream")
	})
}
