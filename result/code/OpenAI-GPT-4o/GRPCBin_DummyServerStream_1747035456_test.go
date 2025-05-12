package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"
	
	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyServerStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	
	req := &pb.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "subtest"},
		FBool:    true,
		FInt64:   456,
	}
	
	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("could not send request: %v", err)
	}

	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("error receiving stream response: %v", err)
		}

		expectedFString := req.FString
		if resp.FString != expectedFString {
			t.Errorf("expected FString %v, got %v", expectedFString, resp.FString)
		}

		expectedFInt32 := req.FInt32 * 10
		if resp.FInt32 != expectedFInt32 {
			t.Errorf("expected FInt32 %v, got %v", expectedFInt32, resp.FInt32)
		}

		expectedFBool := req.FBool
		if resp.FBool != expectedFBool {
			t.Errorf("expected FBool %v, got %v", expectedFBool, resp.FBool)
		}
	}
	
	log.Println("stream response received and validated successfully")
}
