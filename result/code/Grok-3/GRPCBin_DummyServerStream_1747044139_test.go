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
	serverAddr     = "grpcb.in:9000"
	connectTimeout = 15 * time.Second
)

func TestGRPCBin_DummyServerStream(t *testing.T) {
	// Set up connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Prepare test request
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
		FInt64:   1000,
		FInt64S:  []int64{100, 200},
		FBytes:   []byte("test-bytes"),
		FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2},
	}

	// Test positive scenario
	t.Run("PositiveTest_DummyServerStream", func(t *testing.T) {
		stream, err := client.DummyServerStream(context.Background(), req)
		if err != nil {
			t.Fatalf("failed to call DummyServerStream: %v", err)
		}

		// Validate server responses (expecting 10 responses)
		count := 0
		for {
			resp, err := stream.Recv()
			if err != nil {
				break
			}

			// Validate response content matches request (server should echo back)
			if resp.FString != req.FString {
				t.Errorf("expected FString %q, got %q", req.FString, resp.FString)
			}
			if len(resp.FStrings) != len(req.FStrings) {
				t.Errorf("expected FStrings length %d, got %d", len(req.FStrings), len(resp.FStrings))
			}
			if resp.FInt32 != req.FInt32 {
				t.Errorf("expected FInt32 %d, got %d", req.FInt32, resp.FInt32)
			}
			if resp.FEnum != req.FEnum {
				t.Errorf("expected FEnum %v, got %v", req.FEnum, resp.FEnum)
			}
			if resp.FSub.FString != req.FSub.FString {
				t.Errorf("expected FSub.FString %q, got %q", req.FSub.FString, resp.FSub.FString)
			}
			if resp.FBool != req.FBool {
				t.Errorf("expected FBool %v, got %v", req.FBool, resp.FBool)
			}
			if resp.FFloat != req.FFloat {
				t.Errorf("expected FFloat %f, got %f", req.FFloat, resp.FFloat)
			}

			count++
		}

		// Validate expected number of responses
		expectedResponses := 10
		if count != expectedResponses {
			t.Errorf("expected %d responses, got %d", expectedResponses, count)
		}
	})
}
