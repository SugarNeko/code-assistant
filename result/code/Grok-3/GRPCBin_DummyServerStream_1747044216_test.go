package grpcbin

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverAddr      = "grpcb.in:9000"
	connectTimeout  = 15 * time.Second
	expectedReplies = 10
)

func TestGRPCBin_DummyServerStream(t *testing.T) {
	// Set up connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := NewGRPCBinClient(conn)

	// Test case: Positive testing with valid input
	t.Run("ValidRequest", func(t *testing.T) {
		// Prepare request with various field types
		req := &DummyMessage{
			FString:  "test-string",
			FStrings: []string{"str1", "str2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    DummyMessage_ENUM_1,
			FEnums:   []DummyMessage_Enum{DummyMessage_ENUM_0, DummyMessage_ENUM_1},
			FSub: &DummyMessage_Sub{
				FString: "sub-test",
			},
			FSubs: []*DummyMessage_Sub{
				{FString: "sub1"},
				{FString: "sub2"},
			},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   1000000,
			FInt64S:  []int64{100, 200},
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

		// Validate server responses
		responseCount := 0
		for {
			resp, err := stream.Recv()
			if err != nil {
				break
			}

			// Validate response fields (basic checks)
			if resp.FString != req.FString {
				t.Errorf("Expected FString to be %q, got %q", req.FString, resp.FString)
			}
			if len(resp.FStrings) != len(req.FStrings) {
				t.Errorf("Expected FStrings length to be %d, got %d", len(req.FStrings), len(resp.FStrings))
			}
			if resp.FInt32 != req.FInt32 {
				t.Errorf("Expected FInt32 to be %d, got %d", req.FInt32, resp.FInt32)
			}
			if resp.FEnum != req.FEnum {
				t.Errorf("Expected FEnum to be %v, got %v", req.FEnum, resp.FEnum)
			}
			if resp.FBool != req.FBool {
				t.Errorf("Expected FBool to be %v, got %v", req.FBool, resp.FBool)
			}

			responseCount++
		}

		// Validate total number of responses
		if responseCount != expectedReplies {
			t.Errorf("Expected %d responses, got %d", expectedReplies, responseCount)
		}
	})
}
