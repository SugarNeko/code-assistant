package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString: "test",
		FStrings: []string{
			"test1",
			"test2",
		},
		FInt32: 1,
		FInt32s: []int32{
			1,
			2,
		},
		FEint64: 1,
		FEnums: []pb.DummyMessage_Enum{
			pb.DummyMessage_ENUM_0,
			pb.DummyMessage_ENUM_1,
		},
		FSub: &pb.DummyMessage_Sub{
			FString: "test",
		},
		FSubs: []*pb.DummyMessage_Sub{
			{
				FString: "test1",
			},
			{
				FString: "test2",
			},
		},
		FBool: true,
		FBools: []bool{
			true,
			false,
		},
		FInt64: 1,
		FInt64s: []int64{
			1,
			2,
		},
		FBytes: []byte("test"),
		FBytess: [][]byte{
			[]byte("test1"),
			[]byte("test2"),
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
		t.Errorf("FString mismatch: got %s, want %s", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings mismatch: got %d, want %d", len(resp.FStrings), len(req.FStrings))
	}
	for i, v := range resp.FStrings {
		if v != req.FStrings[i] {
			t.Errorf("FStrings[%d] mismatch: got %s, want %s", i, v, req.FStrings[i])
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %d, want %d", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32s) != len(req.FInt32s) {
		t.Errorf("FInt32s mismatch: got %d, want %d", len(resp.FInt32s), len(req.FInt32s))
	}
	for i, v := range resp.FInt32s {
		if v != req.FInt32s[i] {
			t.Errorf("FInt32s[%d] mismatch: got %d, want %d", i, v, req.FInt32s[i])
		}
	}
}
