package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("failed to create stream: %v", err)
	}

	req := &grpcbin.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"foo", "bar"},
		FInt32:    42,
		FInt32S:   []int32{7, 8, 9},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub string"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "one"}, {FString: "two"}},
		FBool:     true,
		FBools:    []bool{false, true},
		FInt64:    1234567890,
		FInt64S:   []int64{11, 22, 33},
		FBytes:    []byte("data"),
		FBytess:   [][]byte{[]byte("a"), []byte("b")},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2},
	}

	// Send request message
	if err := stream.Send(req); err != nil {
		t.Fatalf("error sending to stream: %v", err)
	}

	// Receive echoed message
	res, err := stream.Recv()
	if err != nil {
		t.Fatalf("error receiving from stream: %v", err)
	}

	// Validate response
	if res.FString != req.FString ||
		len(res.FStrings) != len(req.FStrings) ||
		res.FInt32 != req.FInt32 ||
		len(res.FInt32S) != len(req.FInt32S) ||
		res.FEnum != req.FEnum ||
		len(res.FEnums) != len(req.FEnums) ||
		res.FBool != req.FBool ||
		len(res.FBools) != len(req.FBools) ||
		res.FInt64 != req.FInt64 ||
		len(res.FInt64S) != len(req.FInt64S) ||
		string(res.FBytes) != string(req.FBytes) ||
		len(res.FBytess) != len(req.FBytess) ||
		res.FFloat != req.FFloat ||
		len(res.FFloats) != len(req.FFloats) {
		t.Errorf("response does not match request:\ngot  %+v\nwant %+v", res, req)
	}

	if res.FSub == nil || req.FSub == nil || res.FSub.FString != req.FSub.FString {
		t.Errorf("FSub field does not match: got %+v want %+v", res.FSub, req.FSub)
	}

	if len(res.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %d want %d", len(res.FSubs), len(req.FSubs))
	} else {
		for i := range res.FSubs {
			if res.FSubs[i].FString != req.FSubs[i].FString {
				t.Errorf("FSubs[%d].FString mismatch: got %s want %s", i, res.FSubs[i].FString, req.FSubs[i].FString)
			}
		}
	}

	// Optionally: try to close the stream
	if err := stream.CloseSend(); err != nil {
		t.Errorf("error closing send: %v", err)
	}
}
