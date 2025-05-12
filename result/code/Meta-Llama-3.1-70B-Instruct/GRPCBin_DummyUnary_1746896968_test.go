package grpcbin

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	// Set up gRPC client
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial gRPC connection: %v", err)
	}
	defer conn.Close()
	client := pb.NewGRPCBinClient(conn)

	// Construct test message
	req := &pb.DummyMessage{
		FString: "Hello, gRPC!",
		FStrings: []string{
			"String1",
			"String2",
		},
		FInt32: 123,
		FInt32s: []int32{
			1,
			2,
			3,
		},
		FFloat: 45.67,
	}

	// Call gRPC service
	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("gRPC call failed: %v", err)
	}

	// Verify response
	if resp.FString != req.FString {
		t.Errorf("expected f_string: %q, got: %q", req.FString, resp.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("expected f_strings length: %d, got: %d", len(req.FStrings), len(resp.FStrings))
	}
	for i, s := range resp.FStrings {
		if s != req.FStrings[i] {
			t.Errorf("expected f_strings[%d]: %q, got: %q", i, req.FStrings[i], s)
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("expected f_int32: %d, got: %d", req.FInt32, resp.FInt32)
	}
	if len(resp.FInt32s) != len(req.FInt32s) {
		t.Errorf("expected f_int32s length: %d, got: %d", len(req.FInt32s), len(resp.FInt32s))
	}
	for i, n := range resp.FInt32s {
		if n != req.FInt32s[i] {
			t.Errorf("expected f_int32s[%d]: %d, got: %d", i, req.FInt32s[i], n)
		}
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("expected f_float: %f, got: %f", req.FFloat, resp.FFloat)
	}
}
