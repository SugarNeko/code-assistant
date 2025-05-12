package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyClientStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Construct 10 different DummyMessage requests (positive test cases)
	msgs := make([]*grpcbin.DummyMessage, 10)
	for i := 0; i < 10; i++ {
		msgs[i] = &grpcbin.DummyMessage{
			FString:  "test_string",
			FStrings: []string{"s1", "s2"},
			FInt32:   int32(i),
			FInt32S:  []int32{int32(i), int32(i + 10)},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "subfield",
			},
			FSubs: []*grpcbin.DummyMessage_Sub{
				{FString: "sfa"},
				{FString: "sfb"},
			},
			FBool:   true,
			FBools:  []bool{true, false},
			FInt64:  int64(i * 1000),
			FInt64S: []int64{int64(i * 1000), int64(i * 10000)},
			FBytes:  []byte{0x1, 0x2},
			FBytess: [][]byte{[]byte{0xa}, []byte{0xb}},
			FFloat:  float32(i) * 1.11,
			FFloats: []float32{float32(i) * 1.11, float32(i) * 2.22},
		}
	}

	// Send all messages
	for _, msg := range msgs {
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	// Close sending to initiate response
	if err := stream.CloseSend(); err != nil {
		t.Fatalf("Failed to close stream: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	// The server should respond with the last sent message, validate fields match
	expected := msgs[len(msgs)-1]

	// Client response validation
	if resp.FString != expected.FString ||
		len(resp.FStrings) != len(expected.FStrings) ||
		resp.FInt32 != expected.FInt32 ||
		len(resp.FInt32S) != len(expected.FInt32S) ||
		resp.FEnum != expected.FEnum ||
		len(resp.FEnums) != len(expected.FEnums) ||
		(resp.FSub != nil && expected.FSub != nil && resp.FSub.FString != expected.FSub.FString) ||
		len(resp.FSubs) != len(expected.FSubs) ||
		resp.FBool != expected.FBool ||
		len(resp.FBools) != len(expected.FBools) ||
		resp.FInt64 != expected.FInt64 ||
		len(resp.FInt64S) != len(expected.FInt64S) ||
		string(resp.FBytes) != string(expected.FBytes) ||
		len(resp.FBytess) != len(expected.FBytess) ||
		resp.FFloat != expected.FFloat ||
		len(resp.FFloats) != len(expected.FFloats) {
		t.Errorf("Response does not match sent message. Got %+v, want %+v", resp, expected)
	}
}
