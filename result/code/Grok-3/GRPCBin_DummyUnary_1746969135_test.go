package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverAddr = "grpcb.in:9000"
)

func TestGRPCBinService_DummyUnary(t *testing.T) {
	// Set up gRPC connection
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create gRPC client
	client := grpcbin.NewGRPCBinClient(conn)

	// Prepare test cases
	tests := []struct {
		name    string
		request *grpcbin.DummyMessage
	}{
		{
			name: "Valid request with all fields",
			request: &grpcbin.DummyMessage{
				FString:  "test-string",
				FStrings: []string{"str1", "str2"},
				FInt32:   42,
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub-string",
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
				FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:   3.14,
				FFloats:  []float32{1.1, 2.2},
			},
		},
		{
			name: "Valid request with minimal fields",
			request: &grpcbin.DummyMessage{
				FString: "minimal-test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set timeout for the request
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Make the gRPC call
			response, err := client.DummyUnary(ctx, tt.request)
			if err != nil {
				t.Fatalf("DummyUnary call failed: %v", err)
			}

			// Validate client-side request was sent correctly (we can log or inspect if needed)
			if response == nil {
				t.Fatal("Received nil response from server")
			}

			// Validate server response matches the request (since it's an echo service)
			if response.FString != tt.request.FString {
				t.Errorf("Expected FString to be %q, got %q", tt.request.FString, response.FString)
			}

			if len(response.FStrings) != len(tt.request.FStrings) {
				t.Errorf("Expected FStrings length to be %d, got %d", len(tt.request.FStrings), len(response.FStrings))
			}

			if response.FInt32 != tt.request.FInt32 {
				t.Errorf("Expected FInt32 to be %d, got %d", tt.request.FInt32, response.FInt32)
			}

			if response.FEnum != tt.request.FEnum {
				t.Errorf("Expected FEnum to be %v, got %v", tt.request.FEnum, response.FEnum)
			}

			if response.FBool != tt.request.FBool {
				t.Errorf("Expected FBool to be %v, got %v", tt.request.FBool, response.FBool)
			}

			if response.FInt64 != tt.request.FInt64 {
				t.Errorf("Expected FInt64 to be %d, got %d", tt.request.FInt64, response.FInt64)
			}

			if response.FFloat != tt.request.FFloat {
				t.Errorf("Expected FFloat to be %f, got %f", tt.request.FFloat, response.FFloat)
			}

			if tt.request.FSub != nil && (response.FSub == nil || response.FSub.FString != tt.request.FSub.FString) {
				t.Errorf("Expected FSub.FString to be %q, got %q", tt.request.FSub.FString, response.FSub.FString)
			}
		})
	}
}
