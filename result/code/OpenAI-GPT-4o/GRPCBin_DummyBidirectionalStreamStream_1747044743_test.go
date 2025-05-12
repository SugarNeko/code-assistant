package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to start stream: %v", err)
	}

	req := &pb.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    pb.DummyMessage_ENUM_1,
		FBool:    true,
		FInt64:   123456789,
		FFloat:   1.23,
		FStrings: []string{"one", "two"},
		FInt32s:  []int32{1, 2, 3},
		FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
		FSub:     &pb.DummyMessage_Sub{FString: "sub_test"},
		FBools:   []bool{true, false},
		FInt64s:  []int64{987654321},
		FBytes:   []byte{0x01, 0x02},
		FFloats:  []float32{3.14, 1.59},
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}
	
	if resp.FString != req.FString || resp.FInt32 != req.FInt32 || resp.FEnum != req.FEnum || resp.FBool != req.FBool ||
		resp.FInt64 != req.FInt64 || resp.FFloat != req.FFloat {
		t.Fatalf("Unexpected response: %v", resp)
	}
}
