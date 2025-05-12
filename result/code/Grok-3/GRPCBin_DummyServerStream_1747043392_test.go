package grpcbin

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcBinAddress = "grpcb.in:9000"
	connectTimeout = 15 * time.Second
)

func TestGRPCBinDummyServerStream(t *testing.T) {
	// Set up connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, grpcBinAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	client := NewGRPCBinClient(conn)

	// Test case: Positive testing with valid request
	t.Run("ValidRequest", func(t *testing.T) {
		// Prepare a valid request
		req := &DummyMessage{
			FString:  "test-string",
			FStrings: []string{"test1", "test2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    DummyMessage_ENUM_1,
			FEnums:   []DummyMessage_Enum{DummyMessage_ENUM_0, DummyMessage_ENUM_2},
			FSub: &DummyMessage_Sub{
				FString: "sub-test",
			},
			FSubs:    []*DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   1234567890,
			FInt64S:  []int64{1, 2, 3},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2, 3.3},
		}

		// Call the server stream endpoint
		stream, err := client.DummyServerStream(context.Background(), req)
		if err != nil {
			t.Fatalf("Failed to call DummyServerStream: %v", err)
		}

		// Validate server responses (expecting 10 responses)
		count := 0
		for {
			resp, err := stream.Recv()
			if err != nil {
				break
			}

			// Validate response fields match the request (as per service behavior)
			if resp.FString != req.FString {
				t.Errorf("Expected FString to be %v, got %v", req.FString, resp.FString)
			}
			if len(resp.FStrings) != len(req.FStrings) {
				t.Errorf("Expected FStrings length to be %v, got %v", len(req.FStrings), len(resp.FStrings))
			}
			if resp.FInt32 != req.FInt32 {
				t.Errorf("Expected FInt32 to be %v, got %v", req.FInt32, resp.FInt32)
			}

			count++
		}

		// Validate the number of responses
		expectedResponses := 10
		if count != expectedResponses {
			t.Errorf("Expected %d responses, got %d", expectedResponses, count)
		}
	})
}
