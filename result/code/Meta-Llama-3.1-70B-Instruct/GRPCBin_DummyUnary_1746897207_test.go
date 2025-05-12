package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const (
	bufSize = 1024 * 1024
)

func TestDummyUnary(t *testing.T) {
	lis := bufconn.Listen(bufSize)
	srv := grpc.NewServer()
	RegisterGRPCBinServer(srv, &dummyUnaryServer{})
	go func() {
		if err := srv.Serve(lis); err != nil {
			t.Fatalf("serve error: %v", err)
		}
	}()

	bufDialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.Dial("bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGRPCBinClient(conn)

	req := &DummyMessage{
		FString: "string",
		FStrings: []string{
			"string1",
			"string2",
		},
		FInt32: 123,
		FInt32s: []int32{
			123,
			456,
		},
		FEnum: Enum_ENUM_1,
		FEnums: []Enum{
			Enum_ENUM_1,
			Enum_ENUM_2,
		},
		FSub: &DummyMessage_Sub{
			FString: "string",
		},
		FSubs: []*DummyMessage_Sub{
			{
				FString: "string1",
			},
			{
				FString: "string2",
			},
		},
		FBool: true,
		FBools: []bool{
			true,
			false,
		},
		FInt64: 123,
		FInt64s: []int64{
			123,
			456,
		},
		FBytes: []byte("bytes"),
		FBytes: [][]byte{
			[]byte("bytes1"),
			[]byte("bytes2"),
		},
		FFloat: 3.14,
		FFloats: []float32{
			1.23,
			4.56,
		},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("dummy unary error: %v", err)
	}

	if resp == nil {
		t.Errorf("response is nil")
	}

	if resp.FString != req.FString {
		t.Errorf("response FString does not match, got: %s, want: %s", resp.FString, req.FString)
	}

	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("response FStrings length does not match, got: %d, want: %d", len(resp.FStrings), len(req.FStrings))
	}

	for i, s := range resp.FStrings {
		if s != req.FStrings[i] {
			t.Errorf("response FStrings[%d] does not match, got: %s, want: %s", i, s, req.FStrings[i])
		}
	}
}
