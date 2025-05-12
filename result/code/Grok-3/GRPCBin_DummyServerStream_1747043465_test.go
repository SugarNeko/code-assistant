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
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := NewGRPCBinClient(conn)

	// Prepare test cases
	tests := []struct {
		name    string
		request *DummyMessage
	}{
		{
			name: "Valid DummyMessage with basic fields",
			request: &DummyMessage{
				FString:  "test-string",
				FInt32:   42,
				FEnum:    DummyMessage_ENUM_1,
				FSub:     &DummyMessage_Sub{FString: "sub-test-string"},
				FBool:    true,
				FInt64:   123456789,
				FBytes:   []byte("test-bytes"),
				FFloat:   3.14,
				FStrings: []string{"str1", "str2"},
				FInt32S:  []int32{1, 2, 3},
				FEnums:   []DummyMessage_Enum{DummyMessage_ENUM_0, DummyMessage_ENUM_1},
				FSubs:    []*DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
				FBools:   []bool{true, false},
				FInt64S:  []int64{100, 200},
				FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloats:  []float32{1.1, 2.2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			// Call the server stream endpoint
			stream, err := client.DummyServerStream(ctx, tt.request)
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

				// Validate response content matches the sent request
				if resp.FString != tt.request.FString {
					t.Errorf("Expected FString to be %s, got %s", tt.request.FString, resp.FString)
				}
				if resp.FInt32 != tt.request.FInt32 {
					t.Errorf("Expected FInt32 to be %d, got %d", tt.request.FInt32, resp.FInt32)
				}
				if resp.FEnum != tt.request.FEnum {
					t.Errorf("Expected FEnum to be %v, got %v", tt.request.FEnum, resp.FEnum)
				}
				if resp.FSub.FString != tt.request.FSub.FString {
					t.Errorf("Expected FSub.FString to be %s, got %s", tt.request.FSub.FString, resp.FSub.FString)
				}

				responseCount++
			}

			// Validate the number of responses received
			if responseCount != 10 {
				t.Errorf("Expected 10 responses, got %d", responseCount)
			}
		})
	}
}
