package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyClientStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Prepare 10 distinct DummyMessages
	for i := 0; i < 10; i++ {
		msg := &pb.DummyMessage{
			FString:  "message_" + string(rune('A'+i)),
			FStrings: []string{"foo", "bar"},
			FInt32:   int32(i),
			FInt32S:  []int32{int32(i), int32(i + 1)},
			FEnum:    pb.DummyMessage_ENUM_1,
			FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
			FSub:     &pb.DummyMessage_Sub{FString: "sub_" + string(rune('A'+i))},
			FSubs:    []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:    i%2 == 0,
			FBools:   []bool{i%2 == 0, i%2 != 0},
			FInt64:   int64(i * 1000),
			FInt64S:  []int64{int64(i * 10), int64(i * 20)},
			FBytes:   []byte{byte(i)},
			FBytess:  [][]byte{[]byte{1, 2}, []byte{3, 4}},
			FFloat:   float32(i) * 3.14,
			FFloats:  []float32{float32(i), float32(i) * 2},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send DummyMessage %d: %v", i, err)
		}
		// Keep the last sent message for response validation
		if i == 9 {
			lastMsg := msg
			_ = lastMsg // For unused warning; see below
		}
	}

	// Close sending to receive server response
	resp, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	// Build the expected last message (same as message 9)
	expected := &pb.DummyMessage{
		FString:  "message_J",
		FStrings: []string{"foo", "bar"},
		FInt32:   9,
		FInt32S:  []int32{9, 10},
		FEnum:    pb.DummyMessage_ENUM_1,
		FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
		FSub:     &pb.DummyMessage_Sub{FString: "sub_J"},
		FSubs:    []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    false,
		FBools:   []bool{false, true},
		FInt64:   9000,
		FInt64S:  []int64{90, 180},
		FBytes:   []byte{9},
		FBytess:  [][]byte{[]byte{1, 2}, []byte{3, 4}},
		FFloat:   float32(9) * 3.14,
		FFloats:  []float32{9, 18},
	}

	// Validate response fields (deep equality)
	if resp == nil {
		t.Fatalf("Expected response message, got nil")
	}
	if resp.FString != expected.FString {
		t.Errorf("FString mismatch: got %q, want %q", resp.FString, expected.FString)
	}
	if len(resp.FStrings) != 2 || resp.FStrings[0] != "foo" || resp.FStrings[1] != "bar" {
		t.Errorf("FStrings mismatch: got %v, want %v", resp.FStrings, expected.FStrings)
	}
	if resp.FInt32 != expected.FInt32 {
		t.Errorf("FInt32 mismatch: got %d, want %d", resp.FInt32, expected.FInt32)
	}
	if len(resp.FInt32S) != 2 || resp.FInt32S[0] != 9 || resp.FInt32S[1] != 10 {
		t.Errorf("FInt32S mismatch: got %v, want %v", resp.FInt32S, expected.FInt32S)
	}
	if resp.FEnum != expected.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, expected.FEnum)
	}
	if len(resp.FEnums) != 2 || resp.FEnums[0] != pb.DummyMessage_ENUM_0 || resp.FEnums[1] != pb.DummyMessage_ENUM_2 {
		t.Errorf("FEnums mismatch: got %v, want %v", resp.FEnums, expected.FEnums)
	}
	if resp.FSub == nil || resp.FSub.FString != "sub_J" {
		t.Errorf("FSub mismatch: got %+v, want %+v", resp.FSub, expected.FSub)
	}
	if len(resp.FSubs) != 2 || resp.FSubs[0].FString != "sub1" || resp.FSubs[1].FString != "sub2" {
		t.Errorf("FSubs mismatch: got %v, want %v", resp.FSubs, expected.FSubs)
	}
	if resp.FBool != expected.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, expected.FBool)
	}
	if len(resp.FBools) != 2 || resp.FBools[0] != false || resp.FBools[1] != true {
		t.Errorf("FBools mismatch: got %v, want %v", resp.FBools, expected.FBools)
	}
	if resp.FInt64 != expected.FInt64 {
		t.Errorf("FInt64 mismatch: got %d, want %d", resp.FInt64, expected.FInt64)
	}
	if len(resp.FInt64S) != 2 || resp.FInt64S[0] != 90 || resp.FInt64S[1] != 180 {
		t.Errorf("FInt64S mismatch: got %v, want %v", resp.FInt64S, expected.FInt64S)
	}
	if len(resp.FBytes) != 1 || resp.FBytes[0] != 9 {
		t.Errorf("FBytes mismatch: got %v, want %v", resp.FBytes, expected.FBytes)
	}
	if len(resp.FBytess) != 2 || len(resp.FBytess[0]) != 2 || len(resp.FBytess[1]) != 2 {
		t.Errorf("FBytess mismatch: got %v, want %v", resp.FBytess, expected.FBytess)
	}
	if resp.FFloat != expected.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", resp.FFloat, expected.FFloat)
	}
	if len(resp.FFloats) != 2 || resp.FFloats[0] != 9 || resp.FFloats[1] != 18 {
		t.Errorf("FFloats mismatch: got %v, want %v", resp.FFloats, expected.FFloats)
	}
}
