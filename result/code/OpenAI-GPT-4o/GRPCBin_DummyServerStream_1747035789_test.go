package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FBool:    true,
		FInt64:   456,
		FFloat:   1.23,
		FBytes:   []byte("dummy bytes"),
		FStrings: []string{"a", "b", "c"},
		FInt32s:  []int32{1, 2, 3},
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub"}},
		FBools:   []bool{true, false},
		FInt64s:  []int64{789, 1011},
		FBytess:  [][]byte{[]byte("x"), []byte("y")},
		FFloats:  []float32{4.56, 7.89},
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to call DummyServerStream: %v", err)
	}

	// Assuming the supposed response should be 10 times the received message
	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}
		if resp.FString != req.FString {
			t.Errorf("Expected FString %s, got %s", req.FString, resp.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Expected FInt32 %d, got %d", req.FInt32, resp.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("Expected FEnum %v, got %v", req.FEnum, resp.FEnum)
		}
		if resp.FBool != req.FBool {
			t.Errorf("Expected FBool %v, got %v", req.FBool, resp.FBool)
		}
		// Add additional field checks as necessary
	}
}
