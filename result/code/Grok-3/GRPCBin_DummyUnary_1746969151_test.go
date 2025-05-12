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
)

func TestGRPCBinService_DummyUnary(t *testing.T) {
	// Set up connection to the gRPC server
	conn, err := grpc.Dial(grpcBinAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client for the GRPCBin service
	client := NewGRPCBinClient(conn)

	// Test cases
	tests := []struct {
		name    string
		request *DummyMessage
		wantErr bool
	}{
		{
			name: "Valid Request with Basic Fields",
			request: &DummyMessage{
				FString:  "test-string",
				FInt32:   42,
				FEnum:    DummyMessage_ENUM_1,
				FBool:    true,
				FFloat:   3.14,
				FBytes:   []byte("test-bytes"),
				FInt64:   1234567890,
				FSub:     &DummyMessage_Sub{FString: "sub-test"},
				FStrings: []string{"str1", "str2"},
				FInt32S:  []int32{1, 2, 3},
				FEnums:   []DummyMessage_Enum{DummyMessage_ENUM_0, DummyMessage_ENUM_1},
				FSubs:    []*DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
				FBools:   []bool{true, false},
				FInt64S:  []int64{100, 200},
				FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloats:  []float32{1.1, 2.2},
			},
			wantErr: false,
		},
		{
			name:    "Empty Request",
			request: &DummyMessage{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Send the request to the server
			response, err := client.DummyUnary(ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("DummyUnary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				// Validate the response matches the request (since DummyUnary echoes the input)
				if response.FString != tt.request.FString {
					t.Errorf("DummyUnary() response.FString = %v, want %v", response.FString, tt.request.FString)
				}
				if response.FInt32 != tt.request.FInt32 {
					t.Errorf("DummyUnary() response.FInt32 = %v, want %v", response.FInt32, tt.request.FInt32)
				}
				if response.FEnum != tt.request.FEnum {
					t.Errorf("DummyUnary() response.FEnum = %v, want %v", response.FEnum, tt.request.FEnum)
				}
				if response.FBool != tt.request.FBool {
					t.Errorf("DummyUnary() response.FBool = %v, want %v", response.FBool, tt.request.FBool)
				}
				if response.FFloat != tt.request.FFloat {
					t.Errorf("DummyUnary() response.FFloat = %v, want %v", response.FFloat, tt.request.FFloat)
				}
				if string(response.FBytes) != string(tt.request.FBytes) {
					t.Errorf("DummyUnary() response.FBytes = %v, want %v", response.FBytes, tt.request.FBytes)
				}
				if response.FInt64 != tt.request.FInt64 {
					t.Errorf("DummyUnary() response.FInt64 = %v, want %v", response.FInt64, tt.request.FInt64)
				}
				if response.FSub != nil && tt.request.FSub != nil && response.FSub.FString != tt.request.FSub.FString {
					t.Errorf("DummyUnary() response.FSub.FString = %v, want %v", response.FSub.FString, tt.request.FSub.FString)
				}
				// Additional checks for repeated fields can be added as needed
			}
		})
	}
}
