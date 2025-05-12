package grpcbin

import (
	"testing"

	"google.golang.org/grpc"

	"code-assistant/proto/grpcbin"
)

func TestGRPCBinUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString: "test",
		FInt32:  1,
		FEnum:   grpcbin.Enum_ENUM_1,
		FSub: &grpcbin.DummyMessage_Sub{
			FString: "sub",
		},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.FString != req.FString {
		t.Errorf("want: %s, got: %s", req.FString, resp.FString)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("want: %d, got: %d", req.FInt32, resp.FInt32)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("want: %d, got: %d", req.FEnum, resp.FEnum)
	}
	if resp.FSub.FString != req.FSub.FString {
		t.Errorf("want: %s, got: %s", req.FSub.FString, resp.FSub.FString)
	}
}
