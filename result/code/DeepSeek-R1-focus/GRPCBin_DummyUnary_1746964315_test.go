package grpcbin_test

import (
	"context"
	"reflect"
	"testing"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := grpcbin.NewGRPCBinClient(conn)

	tests := []struct {
		name    string
		req     *grpcbin.DummyMessage
		validate func(t *testing.T, resp *grpcbin.DummyMessage)
	}{
		{
			name: "positive_test_full_fields",
			req: &grpcbin.DummyMessage{
				FString:    "test",
				FStrings:   []string{"a", "b"},
				FInt32:     123,
				FInt32s:    []int32{1, 2},
				FEnum:      grpcbin.DummyMessage_ENUM_1,
				FEnums:     []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FSub:       &grpcbin.DummyMessage_Sub{FString: "sub"},
				FSubs:      []*grpcbin.DummyMessage_Sub{{FString: "s1"}, {FString: "s2"}},
				FBool:      true,
				FBools:     []bool{true, false},
				FInt64:     456,
				FInt64s:    []int64{3, 4},
				FBytes:     []byte("bytes"),
				FBytess:    [][]byte{[]byte("b1"), []byte("b2")},
				FFloat:     1.23,
				FFloats:    []float32{1.1, 2.2},
			},
			validate: func(t *testing.T, resp *grpcbin.DummyMessage) {
				if resp.FString != "test" {
					t.Errorf("FString mismatch: got %v want %v", resp.FString, "test")
				}
				if !reflect.DeepEqual(resp.FStrings, []string{"a", "b"}) {
					t.Errorf("FStrings mismatch: got %v want %v", resp.FStrings, []string{"a", "b"})
				}
				if resp.FInt32 != 123 {
					t.Errorf("FInt32 mismatch: got %v want %v", resp.FInt32, 123)
				}
				if !reflect.DeepEqual(resp.FInt32s, []int32{1, 2}) {
					t.Errorf("FInt32s mismatch: got %v want %v", resp.FInt32s, []int32{1, 2})
				}
				if resp.FEnum != grpcbin.DummyMessage_ENUM_1 {
					t.Errorf("FEnum mismatch: got %v want %v", resp.FEnum, grpcbin.DummyMessage_ENUM_1)
				}
				if !reflect.DeepEqual(resp.FEnums, []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2}) {
					t.Errorf("FEnums mismatch: got %v want %v", resp.FEnums, []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2})
				}
				if resp.FSub.FString != "sub" {
					t.Errorf("FSub.FString mismatch: got %v want %v", resp.FSub.FString, "sub")
				}
				if len(resp.FSubs) != 2 || resp.FSubs[0].FString != "s1" || resp.FSubs[1].FString != "s2" {
					t.Errorf("FSubs mismatch: got %v", resp.FSubs)
				}
				if resp.FBool != true {
					t.Errorf("FBool mismatch: got %v want %v", resp.FBool, true)
				}
				if !reflect.DeepEqual(resp.FBools, []bool{true, false}) {
					t.Errorf("FBools mismatch: got %v want %v", resp.FBools, []bool{true, false})
				}
				if resp.FInt64 != 456 {
					t.Errorf("FInt64 mismatch: got %v want %v", resp.FInt64, 456)
				}
				if !reflect.DeepEqual(resp.FInt64s, []int64{3, 4}) {
					t.Errorf("FInt64s mismatch: got %v want %v", resp.FInt64s, []int64{3, 4})
				}
				if string(resp.FBytes) != "bytes" {
					t.Errorf("FBytes mismatch: got %v want %v", string(resp.FBytes), "bytes")
				}
				if len(resp.FBytess) != 2 || string(resp.FBytess[0]) != "b1" || string(resp.FBytess[1]) != "b2" {
					t.Errorf("FBytess mismatch: got %v", resp.FBytess)
				}
				if resp.FFloat != 1.23 {
					t.Errorf("FFloat mismatch: got %v want %v", resp.FFloat, 1.23)
				}
				if !reflect.DeepEqual(resp.FFloats, []float32{1.1, 2.2}) {
					t.Errorf("FFloats mismatch: got %v want %v", resp.FFloats, []float32{1.1, 2.2})
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := c.DummyUnary(context.Background(), tt.req)
			if err != nil {
				t.Fatalf("DummyUnary failed: %v", err)
			}
			tt.validate(t, resp)
		})
	}
}
