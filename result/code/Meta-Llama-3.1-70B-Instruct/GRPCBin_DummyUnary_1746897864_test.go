package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin/grpcbin"
)

func TestGRPCBinService DummyUnary (t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		F_string: "Hello",
		F_int32:  123,
		F_enum:   pb.DummyMessage_ENUM_1,
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.F_string != req.F_string {
		t.Errorf("F_string is not equal to the request value: %v != %v", resp.F_string, req.F_string)
	}

	if resp.F_int32 != req.F_int32 {
		t.Errorf("F_int32 is not equal to the request value: %v != %v", resp.F_int32, req.F_int32)
	}

	if resp.F_enum != req.F_enum {
		t.Errorf("F_enum is not equal to the request value: %v != %v", resp.F_enum, req.F_enum)
	}
}
