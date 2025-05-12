package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBinService(t *testing.T) {
	// Set up connection to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create client
	client := grpcbin.NewGRPCBinClient(conn)

	t.Run("TestDummyServerStream_ValidRequest", func(t *testing.T) {
		// Prepare a valid request
		req := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"test1", "test2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "sub-test",
			},
			FSubs: []*grpcbin.DummyMessage_Sub{
				{FString: "sub1"},
				{FString: "sub2"},
			},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   100,
			FInt64S:  []int64{10, 20},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
		}

		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Call the server stream endpoint
		stream, err := client.DummyServerStream(ctx, req)
		if err != nil {
			t.Fatalf("Failed to call DummyServerStream: %v", err)
		}

		// Validate responses (expecting 10 responses as per service definition)
		count := 0
		for {
			resp, err := stream.Recv()
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				t.Fatalf("Error receiving stream response: %v", err)
			}

			// Validate response content matches the request (basic field check)
			if resp.FString != req.FString {
				t.Errorf("Expected FString to be %s, got %s", req.FString, resp.FString)
			}
			if resp.FInt32 != req.FInt32 {
				t.Errorf("Expected FInt32 to be %d, got %d", req.FInt32, resp.FInt32)
			}
			if resp.FEnum != req.FEnum {
				t.Errorf("Expected FEnum to be %v, got %v", req.FEnum, resp.FEnum)
			}

			count++
		}

		// Validate total number of responses
		if count != 10 {
			t.Errorf("Expected 10 responses, got %d", count)
		}
	})
}
