package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyClientStream(t *testing.T) {
	// Set the timeout for establishing the connection
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Dial the server
	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("Failed to open stream: %v", err)
	}

	// Send 10 DummyMessages
	for i := 0; i < 10; i++ {
		if err := stream.Send(&pb.DummyMessage{
			FString:  "test",
			FInt32:   int32(i),
			FEnum:    pb.DummyMessage_ENUM_1,
			FSub:     &pb.DummyMessage_Sub{FString: "sub"},
			FBool:    true,
			FInt64:   int64(i),
			FBytes:   []byte{1, 2, 3},
			FFloat:   1.23,
			FStrings: []string{"str1", "str2"},
			FInt32S:  []int32{1, 2},
			FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_1, pb.DummyMessage_ENUM_2},
			FSubs:    []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBools:   []bool{true, false},
			FInt64S:  []int64{1, 2, 3},
			FBytess:  [][]byte{{4, 5}, {6, 7}},
			FFloats:  []float32{1.1, 2.2},
		}); err != nil {
			t.Fatalf("Failed to send: %v", err)
		}
	}

	// Close the stream and receive response
	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	// Validate the server response
	expectedString := "test"
	if reply.FString != expectedString {
		t.Errorf("Expected f_string: %v, got: %v", expectedString, reply.FString)
	}

	expectedInt32 := int32(9)
	if reply.FInt32 != expectedInt32 {
		t.Errorf("Expected f_int32: %v, got: %v", expectedInt32, reply.FInt32)
	}

	expectedEnum := pb.DummyMessage_ENUM_1
	if reply.FEnum != expectedEnum {
		t.Errorf("Expected f_enum: %v, got: %v", expectedEnum, reply.FEnum)
	}

	if reply.FBool != true {
		t.Errorf("Expected f_bool: true, got: %v", reply.FBool)
	}
}
