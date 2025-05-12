package grpcbin_test

import (
	"context"
	"log"
	"testing"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBinService(t *testing.T) {
	// Set up connection to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client for the GRPCBin service
	client := grpcbin.NewGRPCBinClient(conn)

	// Test case for positive testing of DummyUnary endpoint
	t.Run("TestDummyUnary_ValidRequest", func(t *testing.T) {
		// Prepare a valid request with various field types
		req := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"str1", "str2"},
			FInt32:   42,
			FInt32S:  []int32{10, 20, 30},
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
			FFloats:  []float32{1.1, 2.2, 3.3},
		}

		// Send the request to the server
		resp, err := client.DummyUnary(context.Background(), req)
		if err != nil {
			t.Fatalf("DummyUnary request failed: %v", err)
		}

		// Validate client-side response
		if resp == nil {
			t.Fatal("Received nil response from server")
		}

		// Validate server response matches the request (echo behavior)
		if resp.FString != req.FString {
			t.Errorf("Expected FString to be %q, got %q", req.FString, resp.FString)
		}
		if len(resp.FStrings) != len(req.FStrings) {
			t.Errorf("Expected FStrings length to be %d, got %d", len(req.FStrings), len(resp.FStrings))
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Expected FInt32 to be %d, got %d", req.FInt32, resp.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("Expected FEnum to be %v, got %v", req.FEnum, resp.FEnum)
		}
		if resp.FSub.FString != req.FSub.FString {
			t.Errorf("Expected FSub.FString to be %q, got %q", req.FSub.FString, resp.FSub.FString)
		}
		if resp.FBool != req.FBool {
			t.Errorf("Expected FBool to be %v, got %v", req.FBool, resp.FBool)
		}
		if resp.FInt64 != req.FInt64 {
			t.Errorf("Expected FInt64 to be %d, got %d", req.FInt64, resp.FInt64)
		}
		if len(resp.FBytes) != len(req.FBytes) {
			t.Errorf("Expected FBytes length to be %d, got %d", len(req.FBytes), len(resp.FBytes))
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("Expected FFloat to be %f, got %f", req.FFloat, resp.FFloat)
		}
	})

	// Test case for minimal valid request
	t.Run("TestDummyUnary_MinimalRequest", func(t *testing.T) {
		// Prepare a minimal valid request
		req := &grpcbin.DummyMessage{
			FString: "minimal-test",
		}

		// Send the request to the server
		resp, err := client.DummyUnary(context.Background(), req)
		if err != nil {
			t.Fatalf("DummyUnary minimal request failed: %v", err)
		}

		// Validate client-side response
		if resp == nil {
			t.Fatal("Received nil response from server")
		}

		// Validate server response matches the request
		if resp.FString != req.FString {
			t.Errorf("Expected FString to be %q, got %q", req.FString, resp.FString)
		}
	})
}
