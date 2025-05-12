package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect to grpcb.in:9000: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "test_string",
		FStrings:  []string{"one", "two", "three"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub_string"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "foo"}, {FString: "bar"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    9876543210,
		FInt64S:   []int64{42, 1001},
		FBytes:    []byte("test_bytes"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    3.14159,
		FFloats:   []float32{1.1, 2.2, 3.3},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream call failed: %v", err)
	}

	count := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		// Validate the response matches the request, as the server echoes it (by design)
		if resp.FString != req.FString {
			t.Errorf("response FString mismatch: got %q, want %q", resp.FString, req.FString)
		}
		if len(resp.FStrings) != len(req.FStrings) {
			t.Errorf("response FStrings length mismatch: got %d, want %d", len(resp.FStrings), len(req.FStrings))
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("response FInt32 mismatch: got %d, want %d", resp.FInt32, req.FInt32)
		}
		if len(resp.FInt32S) != len(req.FInt32S) {
			t.Errorf("response FInt32S length mismatch: got %d, want %d", len(resp.FInt32S), len(req.FInt32S))
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("response FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
		}
		if len(resp.FEnums) != len(req.FEnums) {
			t.Errorf("response FEnums length mismatch: got %d, want %d", len(resp.FEnums), len(req.FEnums))
		}
		count++
	}

	if count != 10 {
		t.Errorf("expected 10 responses in stream, got %d", count)
	}
}
