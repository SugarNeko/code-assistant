package grpcbin_test

import (
	"context"
	"testing"
	"time"
	"io"
	"reflect"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"foo", "bar"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2},
		FSub:      &pb.DummyMessage_Sub{FString: "subfield"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "a"}, {FString: "b"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    100000,
		FInt64S:   []int64{111, 222},
		FBytes:    []byte("data123"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    3.14,
		FFloats:   []float32{6.28, 1.23},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	var count int
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Error receiving from stream: %v", err)
		}
		count++
		// Validate returned message
		if !isDummyMsgEqual(req, resp) {
			t.Errorf("Response message (%d) does not match request\nGot: %+v\nExpected: %+v", count, resp, req)
		}
	}
	if count != 10 {
		t.Errorf("Expected 10 messages in stream, got %d", count)
	}
}

// isDummyMsgEqual compares all the fields in DummyMessage except fields known to be non-deterministic.
func isDummyMsgEqual(a, b *pb.DummyMessage) bool {
	return a.FString == b.FString &&
		reflect.DeepEqual(a.FStrings, b.FStrings) &&
		a.FInt32 == b.FInt32 &&
		reflect.DeepEqual(a.FInt32S, b.FInt32S) &&
		a.FEnum == b.FEnum &&
		reflect.DeepEqual(a.FEnums, b.FEnums) &&
		reflect.DeepEqual(a.FSub, b.FSub) &&
		reflect.DeepEqual(a.FSubs, b.FSubs) &&
		a.FBool == b.FBool &&
		reflect.DeepEqual(a.FBools, b.FBools) &&
		a.FInt64 == b.FInt64 &&
		reflect.DeepEqual(a.FInt64S, b.FInt64S) &&
		reflect.DeepEqual(a.FBytes, b.FBytes) &&
		reflect.DeepEqual(a.FBytess, b.FBytess) &&
		a.FFloat == b.FFloat &&
		reflect.DeepEqual(a.FFloats, b.FFloats)
}
