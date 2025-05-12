package grpcbin_test

import (
	"context"
	"testing"
	"time"
	"io"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	// Set a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Prepare DummyMessage request with fully populated fields
	req := &grpcbin.DummyMessage{
		FString:   "test-string",
		FStrings:  []string{"s1", "s2"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "substring"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    456789,
		FInt64S:   []int64{7, 8, 9},
		FBytes:    []byte("bytes"),
		FBytess:   [][]byte{[]byte("bytes1"), []byte("bytes2")},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2},
	}

	// Call the DummyServerStream
	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	// Validate server stream response
	var count int
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("stream.Recv() failed: %v", err)
		}

		count++

		// Validate individual response fields, expecting an echo of the request
		if resp.FString != req.FString {
			t.Errorf("response FString got %q want %q", resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("response FInt32 got %d want %d", resp.FInt32, req.FInt32)
		}
		// Add additional checks as needed for more fields
	}

	if count != 10 {
		t.Errorf("expected 10 streamed DummyMessages, got %d", count)
	}
}
