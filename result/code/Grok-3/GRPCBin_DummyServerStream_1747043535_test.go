package grpcbin

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBinService(t *testing.T) {
	// Set up connection to the gRPC server with a timeout of 15 seconds
	conn, err := grpc.Dial("grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client for the GRPCBin service
	client := NewGRPCBinClient(conn)

	t.Run("TestDummyServerStream_ValidRequest", func(t *testing.T) {
		// Prepare a valid request
		req := &DummyMessage{
			FString:  "test_string",
			FStrings: []string{"test1", "test2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    DummyMessage_ENUM_1,
			FEnums:   []DummyMessage_Enum{DummyMessage_ENUM_0, DummyMessage_ENUM_2},
			FSub: &DummyMessage_Sub{
				FString: "sub_test",
			},
			FSubs:    []*DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   1000000,
			FInt64S:  []int64{100, 200},
			FBytes:   []byte("test_bytes"),
			FBytess:  [][]byte{[]byte("byte1"), []byte("byte2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
		}

		// Set a context with timeout for the request
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Call the server stream endpoint
		stream, err := client.DummyServerStream(ctx, req)
		if err != nil {
			t.Fatalf("Failed to call DummyServerStream: %v", err)
		}

		// Validate the streamed responses (expecting 10 responses)
		responseCount := 0
		for {
			resp, err := stream.Recv()
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				t.Fatalf("Failed to receive response: %v", err)
			}

			// Validate response content matches the request structure
			if resp.FString != req.FString {
				t.Errorf("Expected FString to be %v, got %v", req.FString, resp.FString)
			}
			if len(resp.FStrings) != len(req.FStrings) {
				t.Errorf("Expected FStrings length to be %v, got %v", len(req.FStrings), len(resp.FStrings))
			}
			if resp.FInt32 != req.FInt32 {
				t.Errorf("Expected FInt32 to be %v, got %v", req.FInt32, resp.FInt32)
			}
			if resp.FEnum != req.FEnum {
				t.Errorf("Expected FEnum to be %v, got %v", req.FEnum, resp.FEnum)
			}
			if resp.FSub.FString != req.FSub.FString {
				t.Errorf("Expected FSub.FString to be %v, got %v", req.FSub.FString, resp.FSub.FString)
			}
			if resp.FBool != req.FBool {
				t.Errorf("Expected FBool to be %v, got %v", req.FBool, resp.FBool)
			}
			if resp.FInt64 != req.FInt64 {
				t.Errorf("Expected FInt64 to be %v, got %v", req.FInt64, resp.FInt64)
			}
			if string(resp.FBytes) != string(req.FBytes) {
				t.Errorf("Expected FBytes to be %v, got %v", req.FBytes, resp.FBytes)
			}
			if resp.FFloat != req.FFloat {
				t.Errorf("Expected FFloat to be %v, got %v", req.FFloat, resp.FFloat)
			}

			responseCount++
		}

		// Validate the total number of responses received
		expectedResponses := 10
		if responseCount != expectedResponses {
			t.Errorf("Expected %d responses, got %d", expectedResponses, responseCount)
		}
	})
}
