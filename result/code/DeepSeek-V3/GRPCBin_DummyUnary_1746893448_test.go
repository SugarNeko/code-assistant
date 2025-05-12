package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

func TestDummyUnary(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	tests := []struct {
		name    string
		req     *grpcbin.DummyMessage
		want    *grpcbin.DummyMessage
		wantErr bool
	}{
		{
			name: "basic request",
			req: &grpcbin.DummyMessage{
				FString:  "test",
				FInt32:   42,
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FBool:    true,
				FInt64:   123456789,
				FBytes:   []byte("test bytes"),
				FFloat:   3.14,
				FStrings: []string{"a", "b", "c"},
				FSub: &grpcbin.DummyMessage.Sub{
					FString: "sub string",
				},
			},
			want: &grpcbin.DummyMessage{
				FString:  "test",
				FInt32:   42,
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FBool:    true,
				FInt64:   123456789,
				FBytes:   []byte("test bytes"),
				FFloat:   3.14,
				FStrings: []string{"a", "b", "c"},
				FSub: &grpcbin.DummyMessage.Sub{
					FString: "sub string",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.DummyUnary(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DummyUnary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !proto.Equal(got, tt.want) {
				t.Errorf("DummyUnary() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServerResponseValidation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString: "validation test",
	}

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("Expected FString %q, got %q", req.FString, resp.FString)
	}

	if resp.FSub != nil {
		t.Error("Expected nil FSub, got non-nil")
	}

	if len(resp.FStrings) != 0 {
		t.Errorf("Expected empty FStrings, got %v", resp.FStrings)
	}
}
