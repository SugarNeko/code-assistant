package main

import (
    "context"
    "fmt"
    "testing"

    "google.golang.org/grpc"

    proto "code-assistant/proto/grpcbin"
)

func TestGRPCBinService(t *testing.T) {
    conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
    if err != nil {
        t.Fatal(err)
    }
    defer conn.Close()

    grpcClient := proto.NewGRPCBinClient(conn)

    // Positive testing
    ctx := context.Background()

    message := &proto.DummyMessage{
        FInt32:  1,
        FString:  "test",
        FEnums:  []proto.Enum{proto.ENUM_0},
        FSubs: []*proto.Sub{
            {
                FString: "testSub",
            },
        },
        FEnums:  []proto.Enum{proto.ENUM_0},
    }
    unaryResponse, err := grpcClient.DummyUnary(ctx, message)
    if err != nil {
        t.Fatal(err)
    }

    if !message.Equal(unaryResponse) {
        t.Fatal("Response does not match expected response")
    }

    // TODO: Add more test cases
}
