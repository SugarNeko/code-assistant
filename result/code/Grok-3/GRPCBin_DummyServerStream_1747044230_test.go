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
	serverAddress = "grpcb.in:9000"
	connectTimeout = 15 * time.Second
)

func TestGRPCBinDummyServerStream(t *testing.T) {
	// Set up connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Prepare test request data
	req := &grpcbin.DummyMessage{
		FString:  "test-string",
		FStrings: []string{"str1", "str2"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
		FSub: &grpcbin.DummyMessage_Sub{
			FString: "sub-test",
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
	}

	// Test positive case for server streaming
	t.Run("PositiveTest_DummyServerStream", func(t *testing.T) {
		stream, err := client.DummyServerStream(context.Background(), req)
		if err != nil {
			t.Fatalf("Failed to call DummyServerStream: %v", err)
		}

		// Validate server responses
		count := 0
		for {
			resp, err := stream.Recv()
			if err != nil {
				break
			}

			// Validate response fields match request (server echoes back)
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
			if resp.FSub.FString != req.FSub.FString {
				t.Errorf("Expected FSub.FString %s, got %s", req.FSub.FString, resp.FSub.FString)
			}
			if resp.FBool != req.FBool {
				t.Errorf("Expected FBool %v, got %v", req.FBool, resp.FBool)
			}
			if resp.FInt64 != req.FInt64 {
				t.Errorf("Expected FInt64 %d, got %d", req.FInt64, resp.FInt64)
			}
			if string(resp.FBytes) != string(req.FBytes) {
				t.Errorf("Expected FBytes %s, got %s", req.FBytes, resp.FBytes)
			}
			if resp.FFloat != req.FFloat {
				t.Errorf("Expected FFloat %f, got %f", req.FFloat, resp.FFloat)
			}

			count++
		}

		// Validate number of responses (server should send 10 responses)
		expectedResponses := 10
		if count != expectedResponses {
			t.Errorf("Expected %d responses, got %d", expectedResponses, count)
		}
	})
}
