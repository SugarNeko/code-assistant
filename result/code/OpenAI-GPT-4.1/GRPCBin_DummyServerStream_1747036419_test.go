package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	addr := "grpcb.in:9000"

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"foo", "bar"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2, pb.DummyMessage_ENUM_0},
		FSub:      &pb.DummyMessage_Sub{FString: "sub"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{false, true, false},
		FInt64:    12345678,
		FInt64S:   []int64{5, 6, 7},
		FBytes:    []byte("bytes"),
		FBytess:   [][]byte{[]byte("a"), []byte("b")},
		FFloat:    3.14,
		FFloats:   []float32{2.71, 1.41},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed to create: %v", err)
	}

	const expectedResponses = 10
	for i := 0; i < expectedResponses; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("failed to receive stream message at %d: %v", i, err)
		}
		// Validate that the server echoed back the correct data
		if resp.FString != req.FString {
			t.Errorf("response FString mismatch: got %q, want %q", resp.FString, req.FString)
		}
		if len(resp.FStrings) != len(req.FStrings) {
			t.Errorf("response FStrings length mismatch: got %d, want %d", len(resp.FStrings), len(req.FStrings))
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("response FInt32 mismatch: got %d, want %d", resp.FInt32, req.FInt32)
		}
		// ... more field checks can be done similarly
	}

	_, err = stream.Recv()
	if err == nil {
		t.Error("expected stream to be closed after 10 responses")
	}
}
