package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("failed to open stream: %v", err)
	}

	message := &grpcbin.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
		FBool:    true,
		FInt64:   456,
		FFloat:   1.23,
		FBytes:   []byte("bytes"),
	}

	if err := stream.Send(message); err != nil {
		t.Fatalf("failed to send message: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("failed to receive response: %v", err)
	}

	if resp.GetFString() != message.GetFString() {
		t.Errorf("expected f_string %v, got %v", message.GetFString(), resp.GetFString())
	}
	if resp.GetFInt32() != message.GetFInt32() {
		t.Errorf("expected f_int32 %v, got %v", message.GetFInt32(), resp.GetFInt32())
	}
	if resp.GetFEnum() != message.GetFEnum() {
		t.Errorf("expected f_enum %v, got %v", message.GetFEnum(), resp.GetFEnum())
	}
	if resp.GetFSub().GetFString() != message.GetFSub().GetFString() {
		t.Errorf("expected f_sub.f_string %v, got %v", message.GetFSub().GetFString(), resp.GetFSub().GetFString())
	}
	if resp.GetFBool() != message.GetFBool() {
		t.Errorf("expected f_bool %v, got %v", message.GetFBool(), resp.GetFBool())
	}
	if resp.GetFInt64() != message.GetFInt64() {
		t.Errorf("expected f_int64 %v, got %v", message.GetFInt64(), resp.GetFInt64())
	}
	if string(resp.GetFBytes()) != string(message.GetFBytes()) {
		t.Errorf("expected f_bytes %v, got %v", string(message.GetFBytes()), string(resp.GetFBytes()))
	}
}
