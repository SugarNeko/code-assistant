package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyServerStream_Positive(t *testing.T) {
	// Set up connection with 15s timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Construct a valid DummyMessage request
	req := &grpcbin.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"foo", "bar"},
		FInt32:    7,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-1"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub-2"}, {FString: "sub-3"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    42,
		FInt64S:   []int64{1001, 1002},
		FBytes:    []byte("abc"),
		FBytess:   [][]byte{[]byte("a"), []byte("b")},
		FFloat:    1.23,
		FFloats:   []float32{4.56, 7.89},
	}

	// Call DummyServerStream
	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream call failed: %v", err)
	}

	// Collect stream responses and validate
	expectedResponses := 10
	for i := 0; i < expectedResponses; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("failed to receive response #%d: %v", i+1, err)
		}

		// Client response validation: check critical fields equals original request
		if resp.FString != req.FString ||
			resp.FInt32 != req.FInt32 ||
			resp.FEnum != req.FEnum ||
			resp.FBool != req.FBool ||
			resp.FInt64 != req.FInt64 {
			t.Errorf("Response #%d fields do not match request fields", i+1)
		}

		if len(resp.FStrings) != len(req.FStrings) {
			t.Errorf("Response #%d FStrings length mismatch", i+1)
		}
		if len(resp.FInt32S) != len(req.FInt32S) {
			t.Errorf("Response #%d FInt32S length mismatch", i+1)
		}
		if len(resp.FEnums) != len(req.FEnums) {
			t.Errorf("Response #%d FEnums length mismatch", i+1)
		}
		if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
			t.Errorf("Response #%d FSub mismatch", i+1)
		}
		if len(resp.FSubs) != len(req.FSubs) {
			t.Errorf("Response #%d FSubs length mismatch", i+1)
		}
		if len(resp.FBools) != len(req.FBools) {
			t.Errorf("Response #%d FBools length mismatch", i+1)
		}
		if len(resp.FInt64S) != len(req.FInt64S) {
			t.Errorf("Response #%d FInt64S length mismatch", i+1)
		}
		if string(resp.FBytes) != string(req.FBytes) {
			t.Errorf("Response #%d FBytes mismatch", i+1)
		}
		if len(resp.FBytess) != len(req.FBytess) {
			t.Errorf("Response #%d FBytess length mismatch", i+1)
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("Response #%d FFloat mismatch", i+1)
		}
		if len(resp.FFloats) != len(req.FFloats) {
			t.Errorf("Response #%d FFloats length mismatch", i+1)
		}
	}

	// After expectedResponses, stream should be closed (io.EOF)
	_, err = stream.Recv()
	if err == nil {
		t.Error("expected end of stream after 10 responses, got additional message")
	}
}
