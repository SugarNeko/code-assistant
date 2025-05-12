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

func TestGRPCBinService(t *testing.T) {
	// Set up connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := NewGRPCBinClient(conn)

	t.Run("TestDummyServerStream_ValidRequest", func(t *testing.T) {
		// Prepare a valid request
		req := &DummyMessage{
			FString:   "test-string",
			FStrings:  []string{"str1", "str2"},
			FInt32:    42,
			FInt32S:   []int32{1, 2, 3},
			FEnum:     DummyMessage_ENUM_1,
			FEnums:    []DummyMessage_Enum{DummyMessage_ENUM_0, DummyMessage_ENUM_1},
			FSub:      &DummyMessage_Sub{FString: "sub-test"},
			FSubs:     []*DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:     true,
			FBools:    []bool{true, false},
			FInt64:    100,
			FInt64S:   []int64{10, 20},
			FBytes:    []byte("test-bytes"),
			FBytess:   [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:    3.14,
			FFloats:   []float32{1.1, 2.2},
		}

		// Call the server stream endpoint
		stream, err := client.DummyServerStream(context.Background(), req)
		if err != nil {
			t.Fatalf("Failed to call DummyServerStream: %v", err)
		}

		// Validate responses (expecting 10 responses as per service behavior)
		count := 0
		for {
			resp, err := stream.Recv()
			if err != nil {
				break
			}

			// Validate response fields match the request
			if resp.FString != req.FString {
				t.Errorf("Expected FString %s, got %s", req.FString, resp.FString)
			}
			if len(resp.FStrings) != len(req.FStrings) {
				t.Errorf("Expected FStrings length %d, got %d", len(req.FStrings), len(resp.FStrings))
			}
			if resp.FInt32 != req.FInt32 {
				t.Errorf("Expected FInt32 %d, got %d", req.FInt32, resp.FInt32)
			}
			if resp.FEnum != req.FEnum {
				t.Errorf("Expected FEnum %v, got %v", req.FEnum, resp.FEnum)
			}
			if resp.FBool != req.FBool {
				t.Errorf("Expected FBool %v, got %v", req.FBool, resp.FBool)
			}
			if resp.FFloat != req.FFloat {
				t.Errorf("Expected FFloat %f, got %f", req.FFloat, resp.FFloat)
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
