package grpcbin

import (
	"context"
	"grpcbin/proto/grpcbin"
	"testing"

	"google.golang.org/grpc"
)

func TestGrpcBinServiceDummyUnary(t *testing.T) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	ctx := context.Background()
	conn, err := grpc.Dial("grpcb.in:9000", opts...)
	if err != nil {
		t.Logf("grpc Dial failed:%s", err)
		t.FailNow()
	}
	gc := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString: "test",
		FStrings: []string{
			"test1",
			"test2",
		},
		FInt32:  123,
		FInt32S: []int32{
			1,
			2,
		},
		FEnum: grpcbin.DummyMessage_ENUM_1,
		FEnums: []grpcbin.DummyMessage_Enum{
			grpcbin.DummyMessage_ENUM_1,
			grpcbin.DummyMessage_ENUM_2,
		},
		FSub: &grpcbin.DummyMessage_Sub{
			FString: "subtest",
		},
		FSubs: []*grpcbin.DummyMessage_Sub{
			{
				FString: "subtest1",
			},
			{
				FString: "subtest2",
			},
		},
		FBool: true,
		FBools: []bool{
			true,
			false,
		},
		FInt64:  1234567,
		FInt64S: []int64{
			123,
			456,
		},
		FBytes:  []byte{1, 2, 3},
		FBytess: [][]byte{
			{1, 2},
			{3, 4},
		},
		FFloat:  3.14,
		FFloats: []float32{
			1.11,
			2.22,
		},
	}

	resp, err := gc.DummyUnary(ctx, req)
	if err != nil {
		t.Logf("grpc call req failed:%s", err)
		t.FailNow()
	}

	if resp.FString != req.FString {
		t.Logf("grpc response f_string failed,expect:%s,actual:%s", req.FString, resp.FString)
		t.FailNow()
	}

	for i, v := range req.FStrings {
		if resp.FStrings[i] != v {
			t.Logf("grpc response f_strings failed,expect:%v,actual:%v", req.FStrings, resp.FStrings)
			t.FailNow()
		}
	}

	if resp.FInt32 != req.FInt32 {
		t.Logf("grpc response f_int32 failed,expect:%d,actual:%d", req.FInt32, resp.FInt32)
		t.FailNow()
	}

	for i, v := range req.FInt32S {
		if resp.FInt32S[i] != v {
			t.Logf("grpc response f_int32s failed,expect:%v,actual:%v", req.FInt32S, resp.FInt32S)
			t.FailNow()
		}
	}

	if resp.FEnum != req.FEnum {
		t.Logf("grpc response f_enum failed,expect:%v,actual:%v", req.FEnum, resp.FEnum)
		t.FailNow()
	}

	for i, v := range req.FEnums {
		if resp.FEnums[i] != v {
			t.Logf("grpc response f_enums failed,expect:%v,actual:%v", req.FEnums, resp.FEnums)
			t.FailNow()
		}
	}

	if resp.FSub.FString != req.FSub.FString {
		t.Logf("grpc response f_sub failed,expect:%s,actual:%s", req.FSub.FString, resp.FSub.FString)
		t.FailNow()
	}

	for i, v := range req.FSubs {
		if resp.FSubs[i].FString != v.FString {
			t.Logf("grpc response f_subs failed,expect:%v,actual:%v", req.FSubs, resp.FSubs)
			t.FailNow()
		}
	}

	if resp.FBool != req.FBool {
		t.Logf("grpc response f_bool failed,expect:%v,actual:%v", req.FBool, resp.FBool)
		t.FailNow()
	}

	for i, v := range req.FBools {
		if resp.FBools[i] != v {
			t.Logf("grpc response f_bools failed,expect:%v,actual:%v", req.FBools, resp.FBools)
			t.FailNow()
		}
	}

	if resp.FInt64 != req.FInt64 {
		t.Logf("grpc response f_int64 failed,expect:%d,actual:%d", req.FInt64, resp.FInt64)
		t.FailNow()
	}

	for i, v := range req.FInt64S {
		if resp.FInt64S[i] != v {
			t.Logf("grpc response f_int64s failed,expect:%v,actual:%v", req.FInt64S, resp.FInt64S)
			t.FailNow()
		}
	}

	if !equalBytes(resp.FBytes, req.FBytes) {
		t.Logf("grpc response f_bytes failed,expect:%v,actual:%v", req.FBytes, resp.FBytes)
		t.FailNow()
	}

	for i, v := range req.FBytess {
		if !equalBytes(resp.FBytess[i], v) {
			t.Logf("grpc response f_bytess failed,expect:%v,actual:%v", req.FBytess, resp.FBytess)
			t.FailNow()
		}
	}

	if resp.FFloat != req.FFloat {
		t.Logf("grpc response f_float failed,expect:%f,actual:%f", req.FFloat, resp.FFloat)
		t.FailNow()
	}

	for i, v := range req.FFloats {
		if resp.FFloats[i] != v {
			t.Logf("grpc response f_floats failed,expect:%v,actual:%v", req.FFloats, resp.FFloats)
			t.FailNow()
		}
	}
}

func equalBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
