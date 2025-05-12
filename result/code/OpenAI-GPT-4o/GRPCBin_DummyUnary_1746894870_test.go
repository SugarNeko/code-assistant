package grpcbin_test

import (
    "context"
    "testing"
    "google.golang.org/grpc"
    "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
    conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
    if err != nil {
        t.Fatalf("Failed to connect to gRPC server: %v", err)
    }
    defer conn.Close()

    client := grpcbin.NewGRPCBinClient(conn)
    
    req := &grpcbin.DummyMessage{
        FString:  "test",
        FInt32:   123,
        FEnum:    grpcbin.DummyMessage_ENUM_1,
        FSub:     &grpcbin.DummyMessage_Sub{FString: "subtest"},
        FBool:    true,
        FInt64:   123456789,
        FBytes:   []byte("testbytes"),
        FFloat:   1.23,
        FStrings: []string{"one", "two"},
    }

    resp, err := client.DummyUnary(context.Background(), req)
    if err != nil {
        t.Fatalf("DummyUnary call failed: %v", err)
    }

    if resp.FString != req.FString ||
       resp.FInt32 != req.FInt32 ||
       resp.FEnum != req.FEnum ||
       resp.FSub.FString != req.FSub.FString ||
       resp.FBool != req.FBool ||
       resp.FInt64 != req.FInt64 ||
       string(resp.FBytes) != string(req.FBytes) ||
       resp.FFloat != req.FFloat {
        t.Errorf("Response does not match request; got %+v, want %+v", resp, req)
    }
}
