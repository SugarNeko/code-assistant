package grpc_test

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString: "string_value",
		FEnums:  []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_1, pb.DummyMessage_ENUM_2},
		FSub: &pb.DummyMessage_Sub{
			FStrings: []string{"sub_string_value"},
		},
		Enum:  pb.DummyMessage_Enum(pb.DummyMessage_ENUM_1),
		Int32: 123,
	}

	res, err := client.DummyUnary(context.Background(), req)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if res == nil {
		t.Errorf("Expected response but got nil")
	}

	if res.GetFString() != "string_value" {
		t.Errorf("Expected string value but got: %s", res.GetFString())
	}

	if res.GetFEnums()[0] != pb.DummyMessage_ENUM_1 {
		t.Errorf("Expected enum value but got: %d", res.GetFEnums()[0])
	}

	if res.GetInt32() != 123 {
		t.Errorf("Expected integer value but got: %d", res.GetInt32())
	}
}

func TestDummyUnaryResponse(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString: "string_value",
		FEnums:  []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_1, pb.DummyMessage_ENUM_2},
		FSub: &pb.DummyMessage_Sub{
			FStrings: []string{"sub_string_value"},
		},
		Enum:  pb.DummyMessage_Enum(pb.DummyMessage_ENUM_1),
		Int32: 123,
	}

	res, err := client.DummyUnary(context.Background(), req)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if res == nil {
		t.Errorf("Expected response but got nil")
	}

	if res.GetFSub() == nil {
		t.Errorf("Expected non-nil response sub field")
	}
}
