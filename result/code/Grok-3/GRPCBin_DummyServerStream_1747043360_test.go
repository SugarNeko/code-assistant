package grpcbin

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverAddr     = "grpcb.in:9000"
	connectTimeout = 15 * time.Second
)

func TestGRPCBinService_DummyServerStream(t *testing.T) {
	// Set up gRPC connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := NewGRPCBinClient(conn)

	// Test case for positive testing with valid request
	t.Run("PositiveTest_ValidRequest", func(t *testing.T) {
		// Prepare a valid request
		req := &DummyMessage{
			FString:  "test-string",
			FStrings: []string{"test1", "test2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    DummyMessage_ENUM_1,
			FEnums:   []DummyMessage_Enum{DummyMessage_ENUM_0, DummyMessage_ENUM_1},
			FSub: &DummyMessage_Sub{
				FString: "sub-test",
			},
			FSubs:    []*DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   100,
			FInt64S:  []int64{10, 20},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
		}

		// Call the server stream endpoint
		stream, err := client.DummyServerStream(context.Background(), req)
		if err != nil {
			t.Fatalf("Failed to call DummyServerStream: %v", err)
		}

		// Validate server responses (expecting 10 responses as per service definition)
		responseCount := 0
		for {
			resp, err := stream.Recv()
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				t.Fatalf("Error receiving stream response: %v", err)
			}

			// Validate response fields (basic validation for key fields)
			if resp.FString != req.FString {
				t.Errorf("Expected FString to be %s, got %s", req.FString, resp.FString)
			}
			if resp.FInt32 != req.FInt32 {
				t.Errorf("Expected FInt32 to be %d, got %d", req.FInt32, resp.FInt32)
			}
			if resp.FBool != req.FBool {
				t.Errorf("Expected FBool to be %t, got %t", req.FBool, resp.FBool)
			}

			responseCount++
		}

		// Validate the number of responses
		expectedResponses := 10
		if responseCount != expectedResponses {
			t.Errorf("Expected %d responses, got %d", expectedResponses, responseCount)
		}
	})
}
