package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString: "string",
		FStrings: []string{
			"string1",
			"string2",
		},
		FInt32: 32,
		FInt32s: []int32{
			32,
			64,
		},
		FEnum: pb.DummyMessage_ENUM_1,
		FEnums: []pb.DummyMessage_Enum{
			pb.DummyMessage_ENUM_1,
			pb.DummyMessage_ENUM_2,
		},
		FSub: &pb.DummyMessage_Sub{
			FString: "sub_string",
		},
		FSubs: []*pb.DummyMessage_Sub{
			{
				FString: "sub_string1",
			},
			{
				FString: "sub_string2",
			},
		},
		FBool: true,
		FBools: []bool{
			true,
			false,
		},
		FInt64: 64,
		FInt64s: []int64{
			64,
			128,
		},
		FBytes: []byte("bytes"),
		Fbytess: [][]byte{
			[]byte("bytes1"),
			[]byte("bytes2"),
		},
		FFloat: 1.0,
		FFloats: []float32{
			1.0,
			2.0,
		},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.FString != req.FString {
		t.Errorf("want %s, got %s", req.FString, resp.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("want %d, got %d", len(req.FStrings), len(resp.FStrings))
	}
	for i, v := range req.FStrings {
		if resp.FStrings[i] != v {
			t.Errorf("want %s, got %s", v, resp.FStrings[i])
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("want %d, got %d", req.FInt32, resp.FInt32)
	}
	if len(resp.FInt32s) != len(req.FInt32s) {
		t.Errorf("want %d, got %d", len(req.FInt32s), len(resp.FInt32s))
	}
	for i, v := range req.FInt32s {
		if resp.FInt32s[i] != v {
			t.Errorf("want %d, got %d", v, resp.FInt32s[i])
		}
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("want %s, got %s", req.FEnum, resp.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("want %d, got %d", len(req.FEnums), len(resp.FEnums))
	}
	for i, v := range req.FEnums {
		if resp.FEnums[i] != v {
			t.Errorf("want %s, got %s", v, resp.FEnums[i])
		}
	}
	if resp.FSub.FString != req.FSub.FString {
		t.Errorf("want %s, got %s", req.FSub.FString, resp.FSub.FString)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("want %d, got %d", len(req.FSubs), len(resp.FSubs))
	}
	for i, v := range req.FSubs {
		if resp.FSubs[i].FString != v.FString {
			t.Errorf("want %s, got %s", v.FString, resp.FSubs[i].FString)
		}
	}
	if resp.FBool != req.FBool {
		t.Errorf("want %t, got %t", req.FBool, resp.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("want %d, got %d", len(req.FBools), len(resp.FBools))
	}
	for i, v := range req.FBools {
		if resp.FBools[i] != v {
			t.Errorf("want %t, got %t", v, resp.FBools[i])
		}
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("want %d, got %d", req.FInt64, resp.FInt64)
	}
	if len(resp.FInt64s) != len(req.FInt64s) {
		t.Errorf("want %d, got %d", len(req.FInt64s), len(resp.FInt64s))
	}
	for i, v := range req.FInt64s {
		if resp.FInt64s[i] != v {
			t.Errorf("want %d, got %d", v, resp.FInt64s[i])
		}
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("want %s, got %s", req.FBytes, resp.FBytes)
	}
	if len(resp.Fbytess) != len(req.Fbytess) {
		t.Errorf("want %d, got %d", len(req.Fbytess), len(resp.Fbytess))
	}
	for i, v := range req.Fbytess {
		if string(resp.Fbytess[i]) != string(v) {
			t.Errorf("want %s, got %s", v, resp.Fbytess[i])
		}
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("want %f, got %f", req.FFloat, resp.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("want %d, got %d", len(req.FFloats), len(resp.FFloats))
	}
	for i, v := range req.FFloats {
		if resp.FFloats[i] != v {
			t.Errorf("want %f, got %f", v, resp.FFloats[i])
		}
	}
}

func TestDummyUnary_Empty(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{}

	_, err = client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDummyUnary_Error(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString: "error",
	}

	_, err = client.DummyUnary(context.Background(), req)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			if s.Code() != codes.Unavailable {
				t.Errorf("want %s, got %s", codes.Unavailable, s.Code())
			}
			if s.Message() != "Error: Error" {
				t.Errorf("want %s, got %s", "Error: Error", s.Message())
			}
		}
	} else {
		t.Errorf("want error, got no error")
	}
}
