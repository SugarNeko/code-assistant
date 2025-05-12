package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	// Set up the gRPC connection with 15 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(
		ctx,
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Construct a typical valid DummyMessage request
	req := &grpcbin.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"a", "b"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "subhello"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "s1"}, {FString: "s2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    987654321,
		FInt64S:   []int64{111, 222, 333},
		FBytes:    []byte{0xAA, 0xBB},
		FBytess:   [][]byte{{0x01}, {0x02, 0x03}},
		FFloat:    1.23,
		FFloats:   []float32{2.34, 3.45},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	respCount := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break // expect io.EOF at end of stream
		}

		respCount++

		// Validate server response: should echo the request fields exactly
		if resp.FString != req.FString {
			t.Errorf("Response FString = %v, want %v", resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Response FInt32 = %v, want %v", resp.FInt32, req.FInt32)
		}
		if resp.FBool != req.FBool {
			t.Errorf("Response FBool = %v, want %v", resp.FBool, req.FBool)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("Response FEnum = %v, want %v", resp.FEnum, req.FEnum)
		}
		// Add more fine-grained validation as needed on additional fields
	}

	if respCount != 10 {
		t.Errorf("Expected 10 DummyMessage responses, got %d", respCount)
	}
}
