package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBinService(t *testing.T) {
	// Setup gRPC client connection
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	client := NewGRPCBinClient(conn)

	// Test case for positive testing of Unary RPC
	t.Run("TestDummyUnary_ValidRequest", func(t *testing.T) {
		// Construct a valid request
		req := &DummyMessage{
			FString:  "test-string",
			FStrings: []string{"test1", "test2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    DummyMessage_ENUM_1,
			FEnums:   []DummyMessage_Enum{DummyMessage_ENUM_0, DummyMessage_ENUM_2},
			FSub: &DummyMessage_Sub{
				FString: "sub-test",
			},
			FSubs: []*DummyMessage_Sub{
				{FString: "sub1"},
				{FString: "sub2"},
			},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   123456789,
			FInt64S:  []int64{987654321, 123456789},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2, 3.3},
		}

		// Send request to server
		resp, err := client.DummyUnary(context.Background(), req)
		if err != nil {
			t.Fatalf("DummyUnary request failed: %v", err)
		}

		// Validate client response
		if resp == nil {
			t.Fatal("Response is nil")
		}

		// Validate server response content (since it's an echo service, response should match request)
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
		if string(resp.FBytes) != string(req.FBytes) {
			t.Errorf("Expected FBytes to be %q, got %q", string(req.FBytes), string(resp.FBytes))
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("Expected FFloat to be %f, got %f", req.FFloat, resp.FFloat)
		}
	})

	// Additional test case for minimal request
	t.Run("TestDummyUnary_MinimalRequest", func(t *testing.T) {
		// Construct a minimal valid request
		req := &DummyMessage{
			FString: "minimal-test",
		}

		// Send request to server
		resp, err := client.DummyUnary(context.Background(), req)
		if err != nil {
			t.Fatalf("DummyUnary minimal request failed: %v", err)
		}

		// Validate client response
		if resp == nil {
			t.Fatal("Response is nil")
		}

		// Validate server response content
		if resp.FString != req.FString {
			t.Errorf("Expected FString to be %q, got %q", req.FString, resp.FString)
		}
	})
}
