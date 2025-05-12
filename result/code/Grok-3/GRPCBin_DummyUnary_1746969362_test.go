package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBin_DummyUnary(t *testing.T) {
	// Setup connection to gRPC server
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create gRPC client
	client := grpcbin.NewGRPCBinClient(conn)

	// Test case for positive testing with complete request parameters
	t.Run("PositiveTest_ValidRequest", func(t *testing.T) {
		// Construct a valid request
		req := &grpcbin.DummyMessage{
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
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
		}

		// Call the gRPC service
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := client.DummyUnary(ctx, req)
		if err != nil {
			t.Fatalf("DummyUnary call failed: %v", err)
		}

		// Validate client-side response
		if resp == nil {
			t.Fatal("response is nil")
		}

		// Validate server response matches the request (echo behavior)
		if resp.FString != req.FString {
			t.Errorf("expected FString to be %q, got %q", req.FString, resp.FString)
		}
		if len(resp.FStrings) != len(req.FStrings) {
			t.Errorf("expected FStrings length to be %d, got %d", len(req.FStrings), len(resp.FStrings))
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("expected FInt32 to be %d, got %d", req.FInt32, resp.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("expected FEnum to be %v, got %v", req.FEnum, resp.FEnum)
		}
		if resp.FSub.FString != req.FSub.FString {
			t.Errorf("expected FSub.FString to be %q, got %q", req.FSub.FString, resp.FSub.FString)
		}
		if resp.FBool != req.FBool {
			t.Errorf("expected FBool to be %v, got %v", req.FBool, resp.FBool)
		}
		if resp.FInt64 != req.FInt64 {
			t.Errorf("expected FInt64 to be %d, got %d", req.FInt64, resp.FInt64)
		}
		if string(resp.FBytes) != string(req.FBytes) {
			t.Errorf("expected FBytes to be %v, got %v", req.FBytes, resp.FBytes)
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("expected FFloat to be %f, got %f", req.FFloat, resp.FFloat)
		}
	})

	// Test case for minimal request to validate server handling
	t.Run("PositiveTest_MinimalRequest", func(t *testing.T) {
		// Construct a minimal valid request
		req := &grpcbin.DummyMessage{
			FString: "minimal-test",
		}

		// Call the gRPC service
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := client.DummyUnary(ctx, req)
		if err != nil {
			t.Fatalf("DummyUnary call failed: %v", err)
		}

		// Validate client-side response
		if resp == nil {
			t.Fatal("response is nil")
		}

		// Validate server response matches the minimal request
		if resp.FString != req.FString {
			t.Errorf("expected FString to be %q, got %q", req.FString, resp.FString)
		}
	})
}
