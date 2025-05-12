package grpcbin_test

import (
    "context"
    "testing"
    "time"

    "code-assistant/proto/grpcbin"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func TestDummyUnaryPositive(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    conn, err := grpc.DialContext(ctx, "grpcb.in:9000",
        grpc.WithTransportCredentials(insecure.NewCredentials()),
        grpc.WithBlock())
    if err != nil {
        t.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := grpcbin.NewGRPCBinClient(conn)

    req := &grpcbin.DummyMessage{
        FString:  "test",
        FStrings: []string{"a", "b"},
        FInt32:   123,
        FInt32S:  []int32{1, 2},
        FEnum:    grpcbin.DummyMessage_ENUM_1,
        FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
        FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
        FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
        FBool:    true,
        FBools:   []bool{true, false},
        FInt64:   456,
        FInt64S:  []int64{3, 4},
        FBytes:   []byte{0x01, 0x02},
        FBytess:  [][]byte{{0x03}, {0x04}},
        FFloat:   1.23,
        FFloats:  []float32{5.6, 7.8},
    }

    resp, err := client.DummyUnary(ctx, req)
    if err != nil {
        t.Fatalf("DummyUnary failed: %v", err)
    }

    if resp == nil {
        t.Fatal("Expected non-nil response")
    }

    if resp.FString != req.FString {
        t.Errorf("FString mismatch: got %q, want %q", resp.FString, req.FString)
    }

    if len(resp.FStrings) != len(req.FStrings) {
        t.Errorf("FStrings length mismatch: got %d, want %d", len(resp.FStrings), len(req.FStrings))
    }

    if resp.FInt32 != req.FInt32 {
        t.Errorf("FInt32 mismatch: got %d, want %d", resp.FInt32, req.FInt32)
    }

    if resp.FEnum != req.FEnum {
        t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
    }

    if resp.FSub.FString != req.FSub.FString {
        t.Errorf("FSub.FString mismatch: got %q, want %q", resp.FSub.FString, req.FSub.FString)
    }

    if resp.FBool != req.FBool {
        t.Errorf("FBool mismatch: got %t, want %t", resp.FBool, req.FBool)
    }

    if resp.FInt64 != req.FInt64 {
        t.Errorf("FInt64 mismatch: got %d, want %d", resp.FInt64, req.FInt64)
    }
}
