package grpcbin

import (
	"code-assistant/proto/grpcbin"
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
 address = "grpcb.in:9000"
)

func TestGRPCBin_DummyUnary(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:    "string",
		FStrings:   []string{"string1", "string2"},
		FInt32:     1,
		FInt32s:    []int32{1, 2},
		FEnum:      grpcbin.Enum_ENUM_1,
		FEnums:     []grpcbin.Enum{grpcbin.Enum_ENUM_1, grpcbin.Enum_ENUM_2},
		FSub: &grpcbin.DummyMessage_Sub{
			FString: "sub_string",
		},
		FSubs: []*grpcbin.DummyMessage_Sub{{
			FString: "sub_string1",
		}, {
			FString: "sub_string2",
		}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   123,
		FInt64s:  []int64{123, 456},
		FBytes:   []byte{1, 2, 3},
		FBytess:  [][]byte{[]byte{1, 2, 3}, []byte{4, 5, 6}},
		FFloat:   1.23,
		FFloats:  []float32{1.23, 4.56},
	}

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Errorf("err: %v, want nil", err)
		return
	}

	if resp FString != req.FString ||
		len(resp.FStrings) != len(req.FStrings) ||
		resp.FInt32 != req.FInt32 ||
		len(resp.FInt32s) != len(req.FInt32s) ||
		resp.FEnum != req.FEnum ||
		len(resp.FEnums) != len(req.FEnums) ||
		!proto.Equal(resp.FSub, req.FSub) ||
		len(resp.FSubs) != len(req.FSubs) ||
		resp.FBool != req.FBool ||
		len(resp.FBools) != len(req.FBools) ||
		resp.FInt64 != req.FInt64 ||
		len(resp.FInt64s) != len(req.FInt64s) ||
		!bytes.Equal(resp.FBytes, req.FBytes) ||
		len(resp.FBytess) != len(req.FBytess) ||
		resp.FFloat != req.FFloat ||
		len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("response: %+v, want %+v", resp, req)
	}
}
