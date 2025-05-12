package grpcbin_test

import (
	"context"
	"testing"

	// import generated code
	_ "codeAssistant/proto"

	gpb "codeAssistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.DialContext(context.TODO(), "grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := gpb.NewGRPCBinClient(conn)

	tests := []struct {
		name string
		req  *gpb.DummyMessage
		want *gpb.DummyMessage
	}{
		{
			name: "BasicRequest",
			req: &gpb.DummyMessage{
				FString: "Hello World",
				FStrings: []string{
					"F1",
					"F2",
				},
				FInt32:     42,
				FInt32s:    []int32{1, 2, 3},
				FGinEnum:   gpb.ENUM_1,
				FEnums:     []gpb.Enum{gpb.ENUM_1, gpb.ENUM_2},
				FSub:       &gpb.DummyMessage_Sub{FString: "Nested"},
				FSubs:      []*gpb.DummyMessage_Sub{{FString: "Flat"}},
				FFooBool:   true,
				FFooBools:  []bool{true, false},
				FBarInt64:  1024,
				FBarInt64s: []int64{1, 2, 3},
				FBarBytes:  []byte("four"),
				FUrlsBytes: []byte("four"),
				FBarFloat:  102.1,
				FUrlsFloats: []float32{1.0, 2.0, 3.0},
			},
			want: &gpb.DummyMessage{
				FString: "Hello World",
				FStrings: []string{
					"F1",
					"F2",
				},
				FInt32:     42,
				FInt32s:    []int32{1, 2, 3},
				FGinEnum:   gpb.ENUM_1,
				FEnums:     []gpb.Enum{gpb.ENUM_1, gpb.ENUM_2},
				FSub:       &gpb.DummyMessage_Sub{FString: "Nested"},
				FSubs:      []*gpb.DummyMessage_Sub{{FString: "Flat"}},
				FFooBool:   true,
				FFooBools:  []bool{true, false},
				FBarInt64:  1024,
				FBarInt64s: []int64{1, 2, 3},
				FBarBytes:  []byte("four"),
				FUrlsBytes: []byte("four"),
				FBarFloat:  102.1,
				FUrlsFloats: []float32{1.0, 2.0, 3.0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.DummyUnary(context.TODO(), tt.req)
			if err != nil {
				t.Fatal(err)
			}
			if !gpbEqual(tt.want, resp) {
				t.Fatalf("want: %+v, got: %+v", tt.want, resp)
			}
		})
	}
}

func gpbEqual(a, b *gpb.DummyMessage) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.GetFString() != b.GetFString() {
		return false
	}
	if !gpbEqual(a.GetFEnums(), b.GetFEnums()) {
		return false
	}
	if !gpbEqual(a.GetFStrings(), b.GetFStrings()) {
		return false
	}
	if a.GetFInt32() != b.GetFInt32() {
		return false
	}
	if !gpbEqual(a.GetFInt32s(), b.GetFInt32s()) {
		return false
	}
	if a.GetFGinEnum() != b.GetFGinEnum() {
		return false
	}
	if a.GetFSub() != b.GetFSub() {
		return false
	}
	return gpbEqual(a.GetFSubs(), b.GetFSubs())
}
