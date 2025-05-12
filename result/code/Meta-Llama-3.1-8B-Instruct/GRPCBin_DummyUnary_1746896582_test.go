package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc/credentials/insecure"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	// Create a connection to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("Dial failed")
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	// Construct an example request
	req := &pb.DummyMessage{
		FString:     "example_string",
		FInt32:      5,
		FEnums:      []pb.Enum{pb.Enum_ENUM_0},
		FBools:      []bool{true},
		FSub: &pb.DummyMessage_Sub{
			FString: "sub_string",
		},
	}

	// Create a context
	ctx := context.Background()

	// Send the request to the server
	res, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatal("DummyUnary error:", err)
	}

	// Validate the server response
	if *res == nil {
		t.Fatal("Server response is nil")
	}
	if res.FString != "example_string" {
		t.Errorf("Server response f_string is '%s', expected '%s'", res.FString, "example_string")
	}
	if res.FInt32 != 5 {
		t.Errorf("Server response f_int32 is '%d', expected '%d'", res.FInt32, 5)
	}
	if res.FEnums[0] != pb.Enum_ENUM_0 {
		t.Errorf("Server response f_enums is '%d', expected '%d'", res.FEnums[0], pb.Enum_ENUM_0)
	}
	if res.FBools[0] != true {
		t.Errorf("Server response f_enums is '%t', expected '%t'", res.FBools[0], true)
	}
	if res.FSub == nil {
		t.Errorf("Server response f_sub is nil, expected not nil")
	}
	if res.FSub.FString != "sub_string" {
		t.Errorf("Server response f_sub -> f_string is '%s', expected '%s'", res.FSub.FString, "sub_string")
	}
}
