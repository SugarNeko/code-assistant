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
	// Set up a connection to the server
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := grpcbin.NewGRPCBinClient(conn)

	// Prepare test cases
	tests := []struct {
		name     string
		request  *grpcbin.DummyMessage
		expected *grpcbin.DummyMessage
		wantErr  bool
	}{
		{
			name: "Positive test with full data",
			request: &grpcbin.DummyMessage{
				FString:  "test-string",
				FStrings: []string{"str1", "str2"},
				FInt32:   42,
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub-test-string",
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
				FBytess:  [][]byte{[]byte("byte1"), []byte("byte2")},
				FFloat:   3.14,
				FFloats:  []float32{1.1, 2.2},
			},
			expected: &grpcbin.DummyMessage{
				FString:  "test-string",
				FStrings: []string{"str1", "str2"},
				FInt32:   42,
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub-test-string",
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
				FBytess:  [][]byte{[]byte("byte1"), []byte("byte2")},
				FFloat:   3.14,
				FFloats:  []float32{1.1, 2.2},
			},
			wantErr: false,
		},
		{
			name: "Positive test with minimal data",
			request: &grpcbin.DummyMessage{
				FString: "minimal-test",
			},
			expected: &grpcbin.DummyMessage{
				FString: "minimal-test",
			},
			wantErr: false,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Send request to server
			response, err := client.DummyUnary(ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("DummyUnary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				// Validate response fields
				if response.FString != tt.expected.FString {
					t.Errorf("DummyUnary() response FString = %v, want %v", response.FString, tt.expected.FString)
				}
				if len(response.FStrings) != len(tt.expected.FStrings) {
					t.Errorf("DummyUnary() response FStrings length = %v, want %v", len(response.FStrings), len(tt.expected.FStrings))
				}
				if response.FInt32 != tt.expected.FInt32 {
					t.Errorf("DummyUnary() response FInt32 = %v, want %v", response.FInt32, tt.expected.FInt32)
				}
				if response.FEnum != tt.expected.FEnum {
					t.Errorf("DummyUnary() response FEnum = %v, want %v", response.FEnum, tt.expected.FEnum)
				}
				if response.FBool != tt.expected.FBool {
					t.Errorf("DummyUnary() response FBool = %v, want %v", response.FBool, tt.expected.FBool)
				}
				if response.FInt64 != tt.expected.FInt64 {
					t.Errorf("DummyUnary() response FInt64 = %v, want %v", response.FInt64, tt.expected.FInt64)
				}
				if response.FFloat != tt.expected.FFloat {
					t.Errorf("DummyUnary() response FFloat = %v, want %v", response.FFloat, tt.expected.FFloat)
				}

				// Validate nested sub message if present
				if tt.expected.FSub != nil && response.FSub.FString != tt.expected.FSub.FString {
					t.Errorf("DummyUnary() response FSub.FString = %v, want %v", response.FSub.FString, tt.expected.FSub.FString)
				}
			}
		})
	}
}
