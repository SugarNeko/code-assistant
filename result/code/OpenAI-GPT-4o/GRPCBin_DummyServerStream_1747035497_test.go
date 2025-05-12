package grpcbin_test

import (
	"context"
	"testing"
	"time"
	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyServerStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub_test"},
		FInt64:   456,
		FBytes:   []byte("bytes_test"),
		FFloat:   1.23,
		FBools:   []bool{true, false, true},
		FInt32s:  []int32{1, 2, 3},
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub_test"},
		FStrings: []string{"a", "b", "c"},
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to call DummyServerStream: %v", err)
	}

	count := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}

		if resp.FString != req.FString+"x10" {
			t.Errorf("Unexpected response FString: got %v, want %v", resp.FString, req.FString+"x10")
		}
		count++
	}

	if count != 10 {
		t.Errorf("Expected 10 responses, got %d", count)
	}
}
