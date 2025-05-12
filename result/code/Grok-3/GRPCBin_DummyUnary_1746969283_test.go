package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "code-assistant/proto/grpcbin"
)

const (
	address = "grpcb.in:9000"
)

func TestGRPCBin_DummyUnary(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	// Create a context with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Test case 1: Positive test with a fully populated request.
	t.Run("PositiveTest_FullRequest", func(t *testing.T) {
		req := &pb.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"str1", "str2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    pb.DummyMessage_ENUM_1,
			FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
			FSub: &pb.DummyMessage_Sub{
				FString: "sub-test-string",
			},
			FSubs: []*pb.DummyMessage_Sub{
				{FString: "sub1"},
				{FString: "sub2"},
			},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   1234567890,
			FInt64S:  []int64{1, 2, 3},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2, 3.3},
		}

		// Send the request to the server.
		resp, err := client.DummyUnary(ctx, req)
		if err != nil {
			t.Fatalf("DummyUnary failed: %v", err)
		}

		// Validate the server response matches the request (since it's an echo service).
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
		if resp.FInt64 != req.FInt64 {
			t.Errorf("expected FInt64 %d, got %d", req.FInt64, resp.FInt64)
		}
		if len(resp.FBytes) != len(req.FBytes) {
			t.Errorf("expected FBytes length %d, got %d", len(req.FBytes), len(resp.FBytes))
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("expected FFloat %f, got %f", req.FFloat, resp.FFloat)
		}
	})

	// Test case 2: Positive test with a minimal request.
	t.Run("PositiveTest_MinimalRequest", func(t *testing.T) {
		req := &pb.DummyMessage{
			FString: "minimal-test",
		}

		// Send the request to the server.
		resp, err := client.DummyUnary(ctx, req)
		if err != nil {
			t.Fatalf("DummyUnary failed: %v", err)
		}

		// Validate the server response matches the request.
		if resp.FString != req.FString {
			t.Errorf("expected FString %q, got %q", req.FString, resp.FString)
		}
	})
}
