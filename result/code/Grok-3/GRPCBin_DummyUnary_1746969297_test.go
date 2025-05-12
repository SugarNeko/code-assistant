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
	// Setup connection to gRPC server
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create gRPC client
	client := grpcbin.NewGRPCBinClient(conn)

	// Test case 1: Positive test for Unary RPC with complete request
	t.Run("TestDummyUnary_ValidRequest", func(t *testing.T) {
		// Construct a valid DummyMessage request
		req := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"test1", "test2"},
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
			FBool:   true,
			FBools:  []bool{true, false},
			FInt64:  123456789,
			FInt64S: []int64{1, 2, 3},
			FBytes:  []byte("test-bytes"),
			FBytess: [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:  3.14,
			FFloats: []float32{1.1, 2.2, 3.3},
		}

		// Call the Unary RPC
		resp, err := client.DummyUnary(context.Background(), req)
		if err != nil {
			t.Fatalf("DummyUnary call failed: %v", err)
		}

		// Validate client-side response
		if resp == nil {
			t.Fatal("Response is nil")
		}

		// Validate server response matches the request (echo behavior)
		if resp.FString != req.FString {
			t.Errorf("Expected FString to be %s, got %s", req.FString, resp.FString)
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
			t.Errorf("Expected FSub.FString to be %s, got %s", req.FSub.FString, resp.FSub.FString)
		}
		if resp.FBool != req.FBool {
			t.Errorf("Expected FBool to be %v, got %v", req.FBool, resp.FBool)
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("Expected FFloat to be %f, got %f", req.FFloat, resp.FFloat)
		}
	})

	// Test case 2: Positive test for Unary RPC with minimal request
	t.Run("TestDummyUnary_MinimalRequest", func(t *testing.T) {
		// Construct a minimal DummyMessage request
		req := &grpcbin.DummyMessage{
			FString: "minimal-test",
		}

		// Call the Unary RPC
		resp, err := client.DummyUnary(context.Background(), req)
		if err != nil {
			t.Fatalf("DummyUnary call failed: %v", err)
		}

		// Validate client-side response
		if resp == nil {
			t.Fatal("Response is nil")
		}

		// Validate server response matches the request (echo behavior)
		if resp.FString != req.FString {
			t.Errorf("Expected FString to be %s, got %s", req.FString, resp.FString)
		}
	})
}
