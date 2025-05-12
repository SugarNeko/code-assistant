package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("failed to create stream: %v", err)
	}

	// Create and send dummy messages
	for i := 0; i < 10; i++ {
		msg := &pb.DummyMessage{
			FString:   "test string",
			FInt32:    int32(i),
			FEnum:     pb.DummyMessage_ENUM_1,
			FBool:     true,
			FInt64:    int64(i) * 10,
			FFloat:    float32(i) * 0.1,
			FBytes:    []byte("bytes"),
			FSub:      &pb.DummyMessage_Sub{FString: "sub"},
			FStrings:  []string{"a", "b"},
			FInt32s:   []int32{1, 2},
			FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_1},
			FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBools:    []bool{true, false},
			FInt64s:   []int64{20, 30},
			FBytess:   [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloats:   []float32{0.1, 0.2},
		}

		if err := stream.Send(msg); err != nil {
			t.Fatalf("failed to send message: %v", err)
		}
	}

	// Close the stream and receive the last message
	res, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to receive response: %v", err)
	}

	// Validate the server response
	expectedString := "test string"
	if res.GetFString() != expectedString {
		t.Errorf("expected f_string: %s, got: %s", expectedString, res.GetFString())
	}

	expectedInt32 := int32(9)
	if res.GetFInt32() != expectedInt32 {
		t.Errorf("expected f_int32: %d, got: %d", expectedInt32, res.GetFInt32())
	}

	expectedBool := true
	if res.GetFBool() != expectedBool {
		t.Errorf("expected f_bool: %t, got: %t", expectedBool, res.GetFBool())
	}

	// Further validation as necessary...
}
