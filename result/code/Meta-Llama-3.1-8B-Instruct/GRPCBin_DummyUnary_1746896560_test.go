package grpcbin_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestGRPCBin(t *testing.T) {
	t.Run("server response", func(t *testing.T) {
		conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		c := grpcbin.NewGRPCBinClient(conn)

		r, err := c.DummyUnary(context.Background(), &grpcbin.DummyMessage{
			FString: "abcd",
			FStrings: []string{"abcd", "efgh"},
			FInt32:   123,
			FInt32s:  []int32{456, 789},
			FElem:    grpcbin.DummyMessage_Enum_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_Enum_ENUM_1, grpcbin.DummyMessage_Enum_ENUM_2},
			FSub: &grpcbin.DummyMessage_Sub{FString: "abcd"},
			FFloat:    123.456,
			FFloats:   []float32{456.789, 123.456},
		})
		require.NoError(t, err)
		require.Len(t, r.GetFSub().GetFStrings(), 1)
	})
	t.Run("client response validation", func(t *testing.T) {
		conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		c := grpcbin.NewGRPCBinClient(conn)

		_, err = c.DummyUnary(context.Background(), &grpcbin.DummyMessage{})
		if err != nil {
			t.Skipf("server error: %v", err)
		}
	})
}
