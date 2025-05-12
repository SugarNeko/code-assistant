package grpcbin_test

import (
	"context"
	"reflect"
	"testing"

	grpcbin "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	tests := []struct {
		name    string
		req     *grpcbin.DummyMessage
		wantErr bool
	}{
		{
			name: "positive_test",
			req: &grpcbin.DummyMessage{
				FString:  "test",
				FStrings: []string{"a", "b"},
				FInt32:   42,
				FInt32S:  []int32{1, 2},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
				FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
				FBool:    true,
				FBools:   []bool{true, false},
				FInt64:   123456789,
				FInt64S:  []int64{987654321, 123456789},
				FBytes:   []byte{0x01, 0x02},
				FBytess:  [][]byte{{0x03}, {0x04}},
				FFloat:   3.14,
				FFloats:  []float32{1.1, 2.2},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.DummyUnary(context.Background(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DummyUnary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(resp.FString, tt.req.FString) {
				t.Errorf("FString mismatch: got %v, want %v", resp.FString, tt.req.FString)
			}
			if !reflect.DeepEqual(resp.FStrings, tt.req.FStrings) {
				t.Errorf("FStrings mismatch: got %v, want %v", resp.FStrings, tt.req.FStrings)
			}
			if resp.FInt32 != tt.req.FInt32 {
				t.Errorf("FInt32 mismatch: got %v, want %v", resp.FInt32, tt.req.FInt32)
			}
			if !reflect.DeepEqual(resp.FInt32S, tt.req.FInt32S) {
				t.Errorf("FInt32S mismatch: got %v, want %v", resp.FInt32S, tt.req.FInt32S)
			}
			if resp.FEnum != tt.req.FEnum {
				t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, tt.req.FEnum)
			}
			if !reflect.DeepEqual(resp.FEnums, tt.req.FEnums) {
				t.Errorf("FEnums mismatch: got %v, want %v", resp.FEnums, tt.req.FEnums)
			}
			if resp.FSub.FString != tt.req.FSub.FString {
				t.Errorf("FSub.FString mismatch: got %v, want %v", resp.FSub.FString, tt.req.FSub.FString)
			}
			if len(resp.FSubs) != len(tt.req.FSubs) {
				t.Errorf("FSubs length mismatch: got %v, want %v", len(resp.FSubs), len(tt.req.FSubs))
			}
			if resp.FBool != tt.req.FBool {
				t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, tt.req.FBool)
			}
			if !reflect.DeepEqual(resp.FBools, tt.req.FBools) {
				t.Errorf("FBools mismatch: got %v, want %v", resp.FBools, tt.req.FBools)
			}
			if resp.FInt64 != tt.req.FInt64 {
				t.Errorf("FInt64 mismatch: got %v, want %v", resp.FInt64, tt.req.FInt64)
			}
			if !reflect.DeepEqual(resp.FInt64S, tt.req.FInt64S) {
				t.Errorf("FInt64S mismatch: got %v, want %v", resp.FInt64S, tt.req.FInt64S)
			}
			if !reflect.DeepEqual(resp.FBytes, tt.req.FBytes) {
				t.Errorf("FBytes mismatch: got %v, want %v", resp.FBytes, tt.req.FBytes)
			}
			if !reflect.DeepEqual(resp.FBytess, tt.req.FBytess) {
				t.Errorf("FBytess mismatch: got %v, want %v", resp.FBytess, tt.req.FBytess)
			}
			if resp.FFloat != tt.req.FFloat {
				t.Errorf("FFloat mismatch: got %v, want %v", resp.FFloat, tt.req.FFloat)
			}
			if !reflect.DeepEqual(resp.FFloats, tt.req.FFloats) {
				t.Errorf("FFloats mismatch: got %v, want %v", resp.FFloats, tt.req.FFloats)
			}
		})
	}
}
