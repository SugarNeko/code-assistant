package grpcbin_test

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	proto "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Errorf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := proto.NewGRPCBinClient(conn)

	// Test positive case
	_, err = client.DummyUnary(context.Background(), &proto.DummyMessage{
		FString:     "test",
		FStrings:    []string{"a", "b"},
		FEUInt32:   1,
		FEUInt32s:  []int32{1, 2},
		FEUInt64:   1,
		FEUInt64s:  []int64{1, 2},
		FEnum:       proto.DummyMessage_Enum_ENUM_1,
		FEEnum:      []proto.DummyMessage_Enum{proto.DummyMessage_Enum_ENUM_1},
		FSub:        &proto.DummyMessage_Sub{FString: "test"},
		FE_SUBJECT:  []*proto.DummyMessage_Sub{&proto.DummyMessage_Sub{FString: "test"}},
		FBoolean:    true,
		FEBoolean:   []bool{true},
		FByte:       "test",
		FEByte:      []byte("test"),
		FCFloat:     1.0,
		FEFloat:     []float32{1.0},
	})
	if err != nil {
		t.Errorf("failed to call gRPC endpoint: %v", err)
	}

	// TODO: add more test cases
}
