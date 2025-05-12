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
	// Setup connection to gRPC server
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create client
	client := grpcbin.NewGRPCBinClient(conn)

	// Define test cases
	t.Run("TestDummyUnary_ValidRequest", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		// Construct a valid request
		req := &grpcbin.DummyMessage{
			FString:  "test_string",
			FStrings: []string{"str1", "str2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "sub_string",
			},
			FSubs: []*grpcbin.DummyMessage_Sub{
				{FString: "sub1"},
				{FString: "sub2"},
			},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   1234567890,
			FInt64S:  []int64{1, 2, 3},
			FBytes:   []byte("test_bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2, 3.3},
		}

		// Send request to server
		resp, err := client.DummyUnary(ctx, req)
		if err != nil {
			t.Fatalf("DummyUnary request failed: %v", err)
		}

		// Validate server response
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

	t.Run("TestDummyUnary_EmptyRequest", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		// Construct an empty request
		req := &grpcbin.DummyMessage{}

		// Send request to server
		resp, err := client.DummyUnary(ctx, req)
		if err != nil {
			t.Fatalf("DummyUnary request failed: %v", err)
		}

		// Validate server response for empty request
		if resp.FString != "" {
			t.Errorf("Expected FString to be empty, got %s", resp.FString)
		}
		if len(resp.FStrings) != 0 {
			t.Errorf("Expected FStrings length to be 0, got %d", len(resp.FStrings))
		}
		if resp.FInt32 != 0 {
			t.Errorf("Expected FInt32 to be 0, got %d", resp.FInt32)
		}
	})
}
