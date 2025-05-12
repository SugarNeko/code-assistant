package grpcbin_test

import (
    "context"
    "testing"
    "reflect"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary_Positive(t *testing.T) {
    conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        t.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewGRPCBinClient(conn)

    req := &pb.DummyMessage{
        FString:  "test",
        FStrings: []string{"a", "b"},
        FInt32:   42,
        FInt32s:  []int32{1, 2},
        FEnum:    pb.DummyMessage_ENUM_1,
        FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
        FSub:     &pb.DummyMessage_Sub{FString: "sub"},
        FSubs:    []*pb.DummyMessage_Sub{{FString: "s1"}, {FString: "s2"}},
        FBool:    true,
        FBools:   []bool{true, false},
        FInt64:   123456789,
        FInt64s:  []int64{987654321, 123456789},
        FBytes:   []byte{0x01, 0x02},
        FBytess:  [][]byte{{0x03}, {0x04}},
        FFloat:   3.14,
        FFloats:  []float32{1.1, 2.2},
    }

    resp, err := client.DummyUnary(context.Background(), req)
    if err != nil {
        t.Fatalf("RPC failed: %v", err)
    }

    if resp.GetFString() != req.FString {
        t.Errorf("FString mismatch: got %q, want %q", resp.GetFString(), req.FString)
    }
    if !reflect.DeepEqual(resp.GetFStrings(), req.FStrings) {
        t.Errorf("FStrings mismatch: got %v, want %v", resp.GetFStrings(), req.FStrings)
    }
    if resp.GetFInt32() != req.FInt32 {
        t.Errorf("FInt32 mismatch: got %d, want %d", resp.GetFInt32(), req.FInt32)
    }
    if !reflect.DeepEqual(resp.GetFInt32s(), req.FInt32s) {
        t.Errorf("FInt32s mismatch: got %v, want %v", resp.GetFInt32s(), req.FInt32s)
    }
    if resp.GetFEnum() != req.FEnum {
        t.Errorf("FEnum mismatch: got %v, want %v", resp.GetFEnum(), req.FEnum)
    }
    if !reflect.DeepEqual(resp.GetFEnums(), req.FEnums) {
        t.Errorf("FEnums mismatch: got %v, want %v", resp.GetFEnums(), req.FEnums)
    }
    if resp.GetFSub().GetFString() != req.FSub.GetFString() {
        t.Errorf("FSub mismatch: got %q, want %q", resp.GetFSub().GetFString(), req.FSub.GetFString())
    }
    if len(resp.GetFSubs()) != len(req.FSubs) {
        t.Errorf("FSubs length mismatch: got %d, want %d", len(resp.GetFSubs()), len(req.FSubs))
    }
    if resp.GetFBool() != req.FBool {
        t.Errorf("FBool mismatch: got %t, want %t", resp.GetFBool(), req.FBool)
    }
    if !reflect.DeepEqual(resp.GetFBools(), req.FBools) {
        t.Errorf("FBools mismatch: got %v, want %v", resp.GetFBools(), req.FBools)
    }
    if resp.GetFInt64() != req.FInt64 {
        t.Errorf("FInt64 mismatch: got %d, want %d", resp.GetFInt64(), req.FInt64)
    }
    if !reflect.DeepEqual(resp.GetFInt64s(), req.FInt64s) {
        t.Errorf("FInt64s mismatch: got %v, want %v", resp.GetFInt64s(), req.FInt64s)
    }
    if !reflect.DeepEqual(resp.GetFBytes(), req.FBytes) {
        t.Errorf("FBytes mismatch: got %v, want %v", resp.GetFBytes(), req.FBytes)
    }
    if !reflect.DeepEqual(resp.GetFBytess(), req.FBytess) {
        t.Errorf("FBytess mismatch: got %v, want %v", resp.GetFBytess(), req.FBytess)
    }
    if resp.GetFFloat() != req.FFloat {
        t.Errorf("FFloat mismatch: got %f, want %f", resp.GetFFloat(), req.FFloat)
    }
    if !reflect.DeepEqual(resp.GetFFloats(), req.FFloats) {
        t.Errorf("FFloats mismatch: got %v, want %v", resp.GetFFloats(), req.FFloats)
    }
}
