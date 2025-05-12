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
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("DummyClientStream open failed: %v", err)
	}

	var lastSent *grpcbin.DummyMessage
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:  "test_string",
			FStrings: []string{"one", "two", "three"},
			FInt32:   int32(i),
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "sub_string",
			},
			FSubs: []*grpcbin.DummyMessage_Sub{
				{FString: "s1"},
				{FString: "s2"},
			},
			FBool:   true,
			FBools:  []bool{true, false, true},
			FInt64:  int64(i * 100000),
			FInt64S: []int64{1, 2, 3},
			FBytes:  []byte("abc"),
			FBytess: [][]byte{[]byte("a"), []byte("b")},
			FFloat:  3.14,
			FFloats: []float32{1.41, 1.73},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message %d: %v", i, err)
		}
		lastSent = msg
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive DummyClientStream response: %v", err)
	}

	// Response validation: should match the last message sent
	if got, want := reply.FString, lastSent.FString; got != want {
		t.Errorf("Reply.FString = %q, want %q", got, want)
	}
	if got, want := reply.FInt32, lastSent.FInt32; got != want {
		t.Errorf("Reply.FInt32 = %d, want %d", got, want)
	}
	if got, want := len(reply.FStrings), len(lastSent.FStrings); got != want {
		t.Errorf("Reply.FStrings length = %d, want %d", got, want)
	}
	if got, want := reply.FEnum, lastSent.FEnum; got != want {
		t.Errorf("Reply.FEnum = %v, want %v", got, want)
	}
	if got, want := reply.FSub.FString, lastSent.FSub.FString; got != want {
		t.Errorf("Reply.FSub.FString = %q, want %q", got, want)
	}
	if got, want := reply.FBool, lastSent.FBool; got != want {
		t.Errorf("Reply.FBool = %v, want %v", got, want)
	}
	if got, want := reply.FInt64, lastSent.FInt64; got != want {
		t.Errorf("Reply.FInt64 = %d, want %d", got, want)
	}
	// ...add more field validations as needed

	t.Logf("DummyClientStream positive test completed successfully")
}
