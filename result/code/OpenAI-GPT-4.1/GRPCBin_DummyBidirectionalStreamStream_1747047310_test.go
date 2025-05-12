package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyBidirectionalStreamStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("failed to create stream: %v", err)
	}

	// Prepare a typical DummyMessage according to the proto specification
	req := &pb.DummyMessage{
		FString:   "teststring",
		FStrings:  []string{"a", "b"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_1},
		FSub:      &pb.DummyMessage_Sub{FString: "subtest"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    999999999,
		FInt64S:   []int64{1, 2, 3, 4},
		FBytes:    []byte("abc123"),
		FBytess:   [][]byte{[]byte("x1"), []byte("x2")},
		FFloat:    1.23,
		FFloats:   []float32{0.5, 1.5},
	}

	// Send the request
	if err := stream.Send(req); err != nil {
		t.Fatalf("failed to send request: %v", err)
	}

	// Receive the response
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("failed to receive response: %v", err)
	}

	// Validate client and server responses are as expected (echo behaviour)
	if resp.FString != req.FString {
		t.Errorf("FString mismatch: got %v, want %v", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length mismatch: got %v, want %v", len(resp.FStrings), len(req.FStrings))
	}
	for i := range req.FStrings {
		if resp.FStrings[i] != req.FStrings[i] {
			t.Errorf("FStrings[%d] mismatch: got %v, want %v", i, resp.FStrings[i], req.FStrings[i])
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %v, want %v", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("FInt32S length mismatch: got %v, want %v", len(resp.FInt32S), len(req.FInt32S))
	}
	for i := range req.FInt32S {
		if resp.FInt32S[i] != req.FInt32S[i] {
			t.Errorf("FInt32S[%d] mismatch: got %v, want %v", i, resp.FInt32S[i], req.FInt32S[i])
		}
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums length mismatch: got %v, want %v", len(resp.FEnums), len(req.FEnums))
	}
	for i := range req.FEnums {
		if resp.FEnums[i] != req.FEnums[i] {
			t.Errorf("FEnums[%d] mismatch: got %v, want %v", i, resp.FEnums[i], req.FEnums[i])
		}
	}
	if resp.FSub == nil || req.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %v, want %v", resp.FSub.GetFString(), req.FSub.GetFString())
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %v, want %v", len(resp.FSubs), len(req.FSubs))
	}
	for i := range req.FSubs {
		if resp.FSubs[i].FString != req.FSubs[i].FString {
			t.Errorf("FSubs[%d].FString mismatch: got %v, want %v", i, resp.FSubs[i].FString, req.FSubs[i].FString)
		}
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools length mismatch: got %v, want %v", len(resp.FBools), len(req.FBools))
	}
	for i := range req.FBools {
		if resp.FBools[i] != req.FBools[i] {
			t.Errorf("FBools[%d] mismatch: got %v, want %v", i, resp.FBools[i], req.FBools[i])
		}
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: got %v, want %v", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64S length mismatch: got %v, want %v", len(resp.FInt64S), len(req.FInt64S))
	}
	for i := range req.FInt64S {
		if resp.FInt64S[i] != req.FInt64S[i] {
			t.Errorf("FInt64S[%d] mismatch: got %v, want %v", i, resp.FInt64S[i], req.FInt64S[i])
		}
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", resp.FBytes, req.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length mismatch: got %v, want %v", len(resp.FBytess), len(req.FBytess))
	}
	for i := range req.FBytess {
		if string(resp.FBytess[i]) != string(req.FBytess[i]) {
			t.Errorf("FBytess[%d] mismatch: got %v, want %v", i, resp.FBytess[i], req.FBytess[i])
		}
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats length mismatch: got %v, want %v", len(resp.FFloats), len(req.FFloats))
	}
	for i := range req.FFloats {
		if resp.FFloats[i] != req.FFloats[i] {
			t.Errorf("FFloats[%d] mismatch: got %v, want %v", i, resp.FFloats[i], req.FFloats[i])
		}
	}

	// Clean up: Close the stream
	if err := stream.CloseSend(); err != nil {
		t.Errorf("failed to close stream: %v", err)
	}
}
