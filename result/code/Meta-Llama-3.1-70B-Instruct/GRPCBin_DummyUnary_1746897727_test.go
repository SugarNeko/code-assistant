package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestUnary(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Errorf("grpc.Dial error: %v", err)
		return
	}
	defer conn.Close()

	client := NewGRPCBinClient(conn)

	tests := []struct {
		name  string
		req   *DummyMessage
		code  codes.Code
		right bool
	}{
		{
			name: "valid request",
			req: &DummyMessage{
				FString:    "Hello",
				FInt32:     123,
				FEnum:      ENUM_0,
				FBool:      true,
				FFloat:     0.99,
				FString:    "grpcbin in-memory server",
				FSub:       &DummyMessage_Sub{FString: "sub message"},
				FInt64:     9223372036854775807,
			},
			code:  codes.OK,
			right: true,
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := client.DummyUnary(ctx, test.req)
			if !test.right {
				s, _ := status.FromError(err)
				if s.Code() != test.code {
					t.Errorf("expected error code: %v, but got: %v", test.code, s.Code())
				}
				return
			}
			if err != nil {
				t.Errorf("expected no error, but got: %v", err)
			}
			if resp == nil {
				t.Errorf("expected non-nil response, but got: nil")
			}
			// Validate the response fields
			if resp.FString != test.req.FString {
				t.Errorf("expected FInternet: %v, but got: %v", test.req.FString, resp.FString)
			}
			if resp.FInt32 != test.req.FInt32 {
				t.Errorf("expected FInt32: %v, but got: %v", test.req.FInt32, resp.FInt32)
			}
			if resp.FEnum != test.req.FEnum {
				t.Errorf("expected FEnum: %v, but got: %v", test.req.FEnum, resp.FEnum)
			}
			if resp.FBool != test.req.FBool {
				t.Errorf("expected FBool: %v, but got: %v", test.req.FBool, resp.FBool)
			}
			// Add more response validation as needed
		})
	}
}
