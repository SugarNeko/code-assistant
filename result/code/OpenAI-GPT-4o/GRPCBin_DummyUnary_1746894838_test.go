package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect to GRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString: "test string",
		FStrings: []string{"string1", "string2"},
		FInt32:  42,
		FInt32S: []int32{1, 2, 3},
		FEnum:   grpcbin.DummyMessage_ENUM_1,
		FEnums:  []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:    &grpcbin.DummyMessage_Sub{FString: "sub test"},
		FSubs:   []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:   true,
		FBools:  []bool{true, false},
		FInt64:  123456789,
		FInt64S: []int64{987654321, 123},
		FBytes:  []byte("test bytes"),
		FBytess: [][]byte{[]byte("test1"), []byte("test2")},
		FFloat:  1.23,
		FFloats: []float32{3.14, 2.718},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary() failed: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("Expected FString: %s, got: %s", req.FString, resp.FString)
	}
	// Add more validation logic for other fields here

	log.Printf("Test passed with response: %#v", resp)
}
