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
	// Set up a connection to the server.
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client for the GRPCBin service.
	client := grpcbin.NewGRPCBinClient(conn)

	t.Run("TestDummyUnary_PositiveCase", func(t *testing.T) {
		// Prepare a test request with various field types.
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
			FInt64:   1000000,
			FInt64S:  []int64{100, 200},
			FBytes:   []byte("test_bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
		}

		// Send the request to the server.
		resp, err := client.DummyUnary(context.Background(), req)
		if err != nil {
			t.Fatalf("DummyUnary failed: %v", err)
		}

		// Validate the response matches the request (since the server echoes the input).
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
			t.Errorf("expected FBytes to be %q, got %q", string(req.FBytes), string(resp.FBytes))
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("expected FFloat to be %f, got %f", req.FFloat, resp.FFloat)
		}
	})

	t.Run("TestDummyUnary_EmptyRequest", func(t *testing.T) {
		// Test with an empty request to validate server behavior.
		req := &grpcbin.DummyMessage{}

		resp, err := client.DummyUnary(context.Background(), req)
		if err != nil {
			t.Fatalf("DummyUnary failed with empty request: %v", err)
		}

		// Validate that the response is also empty or matches the expected default values.
		if resp.FString != "" {
			t.Errorf("expected FString to be empty, got %q", resp.FString)
		}
		if len(resp.FStrings) != 0 {
			t.Errorf("expected FStrings to be empty, got length %d", len(resp.FStrings))
		}
		if resp.FInt32 != 0 {
			t.Errorf("expected FInt32 to be 0, got %d", resp.FInt32)
		}
	})
}
