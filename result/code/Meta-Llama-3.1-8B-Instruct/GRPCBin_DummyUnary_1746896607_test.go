package grpcbin

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"

	proto "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	// Set up client and server connection
	conn, err := grpc.DialContext(context.Background(), "grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewGRPCBinClient(conn)

	// Positive testing
	testMsg := &proto.DummyMessage{
		FString:    "hello",
		FStrings:   []string{"world"},
		FInt32:     123,
		FInt32s:    []int32{456, 789},
		FElem:      proto.Enum_ENUM_1,
		FEnums:     []proto.Enum{proto.Enum_ENUM_0, proto.Enum_ENUM_2},
		FSub:       &proto.DummyMessage_Sub{FString: "sub"},
		FSubs:      []*proto.DummyMessage_Sub{{FString: "sub2"}},
		FBool:      true,
		FBools:     []bool{false, true},
		FInt64:     int64(7890),
		FInt64s:    []int64{1234, 5678},
		FBytes:     []byte("byte"),
		FSimple:    float32(9.99),
		FFloats:    []float32{float32(7.78), float32(3.14)},
	}

	resp, err := client.DummyUnary(context.Background(), testMsg)
	if err != nil {
		t.Fatalf("client.DummyUnary returned error: %v", err)
	}

	// Validate response
	fmt.Printf("Resp: %+v\n", resp)

	if len(resp.GetFStrings()) != 1 || resp.GetFStrings()[0] != "world" {
		t.Errorf("resp.FStrings = %v, want [\"world\"]", resp.GetFStrings())
	}

	if len(resp.GetFEnums()) != 2 || resp.GetFEnums()[0] != proto.Enum_ENUM_0 || resp.GetFEnums()[1] != proto.Enum_ENUM_2 {
		t.Errorf("resp.FEnums = %v, want [ENUM_0, ENUM_2]", resp.GetFEnums())
	}

	if len(resp.GetFSubs()) != 1 || resp.GetFSubs()[0].GetFString() != "sub2" {
		t.Errorf("resp.FSubs = %v, want [{\"string\":\"sub2\"]", resp.GetFSubs())
	}

	if resp.GetFInt64s()[0] != int64(1234) {
		t.Errorf("resp.FInt64s = %v, want [1234]", resp.GetFInt64s())
	}
}
