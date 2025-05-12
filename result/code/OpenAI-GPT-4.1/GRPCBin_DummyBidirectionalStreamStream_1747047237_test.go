package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		"grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
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
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_2,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-string"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    123456789,
		FInt64S:   []int64{1, 2, 3},
		FBytes:    []byte("bytes1"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    3.14,
		FFloats:   []float32{1.23, 4.56},
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("failed to send to stream: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("failed to receive from stream: %v", err)
	}

	// Validate response matches request
	if resp.FString != req.FString ||
		len(resp.FStrings) != len(req.FStrings) ||
		resp.FInt32 != req.FInt32 ||
		len(resp.FInt32S) != len(req.FInt32S) ||
		resp.FEnum != req.FEnum ||
		len(resp.FEnums) != len(req.FEnums) ||
		((req.FSub != nil && resp.FSub == nil) || (req.FSub == nil && resp.FSub != nil)) ||
		len(resp.FSubs) != len(req.FSubs) ||
		resp.FBool != req.FBool ||
		len(resp.FBools) != len(req.FBools) ||
		resp.FInt64 != req.FInt64 ||
		len(resp.FInt64S) != len(req.FInt64S) ||
		string(resp.FBytes) != string(req.FBytes) ||
		len(resp.FBytess) != len(req.FBytess) ||
		resp.FFloat != req.FFloat ||
		len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("response does not match request")
	}

	// Optionally, do deep equality checks for slices
	for i, v := range req.FStrings {
		if resp.FStrings[i] != v {
			t.Errorf("FStrings[%d]: got %q, want %q", i, resp.FStrings[i], v)
		}
	}
	for i, v := range req.FInt32S {
		if resp.FInt32S[i] != v {
			t.Errorf("FInt32S[%d]: got %v, want %v", i, resp.FInt32S[i], v)
		}
	}
	for i, v := range req.FEnums {
		if resp.FEnums[i] != v {
			t.Errorf("FEnums[%d]: got %v, want %v", i, resp.FEnums[i], v)
		}
	}
	for i, v := range req.FSubs {
		if resp.FSubs[i].FString != v.FString {
			t.Errorf("FSubs[%d].FString: got %q, want %q", i, resp.FSubs[i].FString, v.FString)
		}
	}
	for i, v := range req.FBools {
		if resp.FBools[i] != v {
			t.Errorf("FBools[%d]: got %v, want %v", i, resp.FBools[i], v)
		}
	}
	for i, v := range req.FInt64S {
		if resp.FInt64S[i] != v {
			t.Errorf("FInt64S[%d]: got %v, want %v", i, resp.FInt64S[i], v)
		}
	}
	for i, v := range req.FBytess {
		if string(resp.FBytess[i]) != string(v) {
			t.Errorf("FBytess[%d]: got %v, want %v", i, resp.FBytess[i], v)
		}
	}
	for i, v := range req.FFloats {
		if resp.FFloats[i] != v {
			t.Errorf("FFloats[%d]: got %v, want %v", i, resp.FFloats[i], v)
		}
	}
	if req.FSub != nil && resp.FSub != nil {
		if resp.FSub.FString != req.FSub.FString {
			t.Errorf("FSub.FString: got %q, want %q", resp.FSub.FString, req.FSub.FString)
		}
	}
}
