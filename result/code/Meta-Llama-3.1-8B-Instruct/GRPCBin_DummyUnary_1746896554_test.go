package grpcbin_test

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	proto "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewGRPCBinClient(conn)

	tests := []struct {
		name string
		req  *proto.DummyMessage
		want *proto.DummyMessage
	}{
		{
			name: "request with message",
			req: &proto.DummyMessage{
				FString: "message",
			},
			want: &proto.DummyMessage{
				FString: "message",
			},
		},
		{
			name: "request with enum",
			req: &proto.DummyMessage{
				FEnum: proto.Enum_ENUM_0,
			},
			want: &proto.DummyMessage{
				FEnum: proto.Enum_ENUM_0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := client.DummyUnary(context.Background(), tt.req)
			if err != nil {
				s, ok := status.FromError(err)
				if !ok || s.Code() != 0 {
					t.Errorf("Unexpected error: %v", err)
				}
			} else if got, want := r.GetFString(), tt.want.GetFString(); got != want {
				t.Errorf("response.GetFString() = %s, want %s", got, want)
			}
		})
	}
}
