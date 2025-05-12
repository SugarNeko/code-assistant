package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyUnary(t *testing.T) {
га conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString: "unittest",
		FStrings: []string{
			"unittest",
		},
		FInt32: 1,
		FInt32s: []int32{
			1,
		},
		FFetch: pb.DummyMessage_ENUM_0,
		FEnums: []pb.DummyMessage_Enum{
			pb.DummyMessage_ENUM_0,
		},
		FFetch: &pb.DummyMessage_Sub{
			FString: "unittest",
		},
		FFetchs: []*pb.DummyMessage_Sub{
			{
				FString: "unittest",
			},
		},
		FBool: true,
		FBools: []bool{
			true,
		},
		FInt64: 1,
		FInt64s: []int64{
			1,
		},
		FBytes: []byte("unittest"),
		FBytess: [][]byte{
			[]byte("unittest"),
		},
		FFetch: 1.0,
		FSubs: []float32{
			1.0,
		},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.FString != req.FString {
		t.Errorf("FString wants %q, but got %q", req.FString, resp.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings wants %v, but got %v", req.FStrings, resp.FStrings)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 wants %d, but got %d", req.FInt32, resp.FInt32)
	}
	if len(resp.FInt32s) != len(req.FInt32s) {
		t.Errorf("FInt32s wants %v, but got %v", req.FInt32s, resp.FInt32s)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum wants %d, but got %d", req.FEnum, resp.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums wants %v, but got %v", req.FEnums, resp.FEnums)
	}
	if resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString wants %q, but got %q", req.FSub.FString, resp.FSub.FString)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs wants %v, but got %v", req.FSubs, resp.FSubs)
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool wants %t, but got %t", req.FBool, resp.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools wants %v, but got %v", req.FBools, resp.FBools)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 wants %d, but got %d", req.FInt64, resp.FInt64)
	}
	if len(resp.FInt64s) != len(req.FInt64s) {
		t.Errorf("FInt64s wants %v, but got %v", req.FInt64s, resp.FInt64s)
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes wants %q, but got %q", req.FBytes, resp.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess wants %v, but got %v", req.FBytess, resp.FBytess)
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat wants %f, but got %f", req.FFloat, resp.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats wants %v, but got %v", req.FFloats, resp.FFloats)
	}
}
