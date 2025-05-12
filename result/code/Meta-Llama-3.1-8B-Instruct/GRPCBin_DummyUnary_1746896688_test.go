package grpcbin_test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	"yourpackage/proto" // replace "yourpackage/proto" with your actual package path
)

func TestGRPCBin(t *testing.T) {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewGRPCBinClient(conn)

	t.Run("DummyUnary", func(t *testing.T) {
		req := &proto.DummyMessage{
			FString:    "test",
			FStrings:   []string{"test1", "test2"},
			FInt32:     10,
			FInt32s:    []int32{10, 20},
			FEnum:      proto.Enum_ENUM_0,
			FEnums:     []proto.Enum{proto.Enum_ENUM_0, proto.Enum_ENUM_1},
			FSub: &proto.DummyMessage_Sub{
				FString: "test",
			},
			FSubs: []*proto.DummyMessage_Sub{
				{
					FString: "test1",
				},
			},
		_FW: bool(true),
			_FWs: []bool{true, false},
			FInt64:  100,
			FInt64s: []int64{100, 200},
			FBytes:  []byte("test bytes"),
			FBytess: [][]byte{{1, 2}, {3, 4}},
			FFloat:  10.0,
			FFloats: []float32{10.0, 20.0},
		}

		res, err := client.DummyUnary(context.Background(), req)
		if err != nil {
			t.Fatal(err)
		}

		if res.GetFString() != req.GetFString() {
			t.Errorf("expected %s but got %s", req.GetFString(), res.GetFString())
		}
		if res.GetFStrings() != req.GetFStrings() {
			t.Errorf("expected %+v but got %+v", req.GetFStrings(), res.GetFStrings())
		}
		if res.GetFInt32() != req.GetFInt32() {
			t.Errorf("expected %d but got %d", req.GetFInt32(), res.GetFInt32())
		}
		if res.GetFInt32s() != req.GetFInt32s() {
			t.Errorf("expected %+v but got %+v", req.GetFInt32s(), res.GetFInt32s())
		}
		if res.GetFEnum() != req.GetFEnum() {
			t.Errorf("expected %d but got %d", req.GetFEnum(), res.GetFEnum())
		}
		if res.GetFEnums() != req.GetFEnums() {
			t.Errorf("expected %+v but got %+v", req.GetFEnums(), res.GetFEnums())
		}
		if res.GetFSub().GetFString() != req.GetFSub().GetFString() {
			t.Errorf("expected %s but got %s", req.GetFSub().GetFString(), res.GetFSub().GetFString())
		}
		if res.GetFSubs().Get(0).GetFString() != req.GetFSubs().Get(0).GetFString() {
			t.Errorf("expected %s but got %s", req.GetFSubs().Get(0).GetFString(), res.GetFSubs().Get(0).GetFString())
		}
		if res.Get_FW() != req.Get_FW() {
			t.Errorf("expected %v but got %v", req.Get_FW(), res.Get_FW())
		}
		if res.Get_FWs() != req.Get_FWs() {
			t.Errorf("expected %+v but got %+v", req.Get_FWs(), res.Get_FWs())
		}
		if res.GetFInt64() != req.GetFInt64() {
			t.Errorf("expected %d but got %d", req.GetFInt64(), res.GetFInt64())
		}
		if res.GetFInt64s() != req.GetFInt64s() {
			t.Errorf("expected %+v but got %+v", req.GetFInt64s(), res.GetFInt64s())
		}
		if res.GetFBytes() != req.GetFBytes() {
			t.Errorf("expected %v but got %v", req.GetFBytes(), res.GetFBytes())
		}
		if res.GetFBytess().Get(0) != req.GetFBytess().Get(0) {
			t.Errorf("expected %v but got %v", req.GetFBytess().Get(0), res.GetFBytess().Get(0))
		}
		if res.GetFFloat() != req.GetFFloat() {
			t.Errorf("expected %f but got %f", req.GetFFloat(), res.GetFFloat())
		}
		if res.GetFFloats() != req.GetFFloats() {
			t.Errorf("expected %+v but got %+v", req.GetFFloats(), res.GetFFloats())
		}
	})
}
