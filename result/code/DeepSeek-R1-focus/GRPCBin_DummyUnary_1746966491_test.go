package grpcbin_test

import (
	"context"
	"reflect"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:   "test_string",
		FStrings:  []string{"a", "b", "c"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.Enum_ENUM_1,
		FEnums:    []pb.Enum{pb.Enum_ENUM_0, pb.Enum_ENUM_2},
		FSub:      &pb.DummyMessage_Sub{FString: "sub_string"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    123456789,
		FInt64S:   []int64{987654321, 123456789},
		FBytes:    []byte("test_bytes"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2, 3.3},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("FString mismatch: got %v, want %v", resp.FString, req.FString)
	}

	if !reflect.DeepEqual(resp.FStrings, req.FStrings) {
		t.Errorf("FStrings mismatch: got %v, want %v", resp.FStrings, req.FStrings)
	}

	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %v, want %v", resp.FInt32, req.FInt32)
	}

	if !reflect.DeepEqual(resp.FInt32S, req.FInt32S) {
		t.Errorf("FInt32S mismatch: got %v, want %v", resp.FInt32S, req.FInt32S)
	}

	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
	}

	if !reflect.DeepEqual(resp.FEnums, req.FEnums) {
		t.Errorf("FEnums mismatch: got %v, want %v", resp.FEnums, req.FEnums)
	}

	if resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %v, want %v", resp.FSub.FString, req.FSub.FString)
	}

	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %v, want %v", len(resp.FSubs), len(req.FSubs))
	} else {
		for i := range resp.FSubs {
			if resp.FSubs[i].FString != req.FSubs[i].FString {
				t.Errorf("FSubs[%d].FString mismatch: got %v, want %v", i, resp.FSubs[i].FString, req.FSubs[i].FString)
			}
		}
	}

	if resp.FBool != req.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, req.FBool)
	}

	if !reflect.DeepEqual(resp.FBools, req.FBools) {
		t.Errorf("FBools mismatch: got %v, want %v", resp.FBools, req.FBools)
	}

	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: got %v, want %v", resp.FInt64, req.FInt64)
	}

	if !reflect.DeepEqual(resp.FInt64S, req.FInt64S) {
		t.Errorf("FInt64S mismatch: got %v, want %v", resp.FInt64S, req.FInt64S)
	}

	if !reflect.DeepEqual(resp.FBytes, req.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", resp.FBytes, req.FBytes)
	}

	if !reflect.DeepEqual(resp.FBytess, req.FBytess) {
		t.Errorf("FBytess mismatch: got %v, want %v", resp.FBytess, req.FBytess)
	}

	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", resp.FFloat, req.FFloat)
	}

	if !reflect.DeepEqual(resp.FFloats, req.FFloats) {
		t.Errorf("FFloats mismatch: got %v, want %v", resp.FFloats, req.FFloats)
	}
}
