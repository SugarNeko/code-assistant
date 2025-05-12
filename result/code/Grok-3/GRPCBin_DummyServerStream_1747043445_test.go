package grpcbin

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBin_DummyServerStream(t *testing.T) {
	// Set up connection to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a client for the GRPCBin service
	client := NewGRPCBinClient(conn)

	// Test case for positive testing with valid input
	t.Run("ValidRequest", func(t *testing.T) {
		// Prepare a valid request
		req := &DummyMessage{
			FString:  "test_string",
			FStrings: []string{"str1", "str2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    DummyMessage_ENUM_1,
			FEnums:   []DummyMessage_Enum{DummyMessage_ENUM_0, DummyMessage_ENUM_2},
			FSub: &DummyMessage_Sub{
				FString: "sub_test",
			},
			FSubs: []*DummyMessage_Sub{
				{FString: "sub1"},
				{FString: "sub2"},
			},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   100,
			FInt64S:  []int64{10, 20},
			FBytes:   []byte("test_bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
		}

		// Create a context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Call the server streaming endpoint
		stream, err := client.DummyServerStream(ctx, req)
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
				t.Fatalf("Failed to receive response: %v", err)
			}

			// Validate response content (example: checking if FString matches request)
			if resp.FString != req.FString {
				t.Errorf("Expected FString to be %s, got %s", req.FString, resp.FString)
			}

			responseCount++
		}

		// Validate the number of responses received
		expectedResponses := 10
		if responseCount != expectedResponses {
			t.Errorf("Expected %d responses, got %d", expectedResponses, responseCount)
		}
	})
}
