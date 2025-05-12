package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	testMessage := &pb.DummyMessage{
		FString:   "test",
		FInt32:    123,
		FEnum:     pb.DummyMessage_ENUM_1,
		FSub:      &pb.DummyMessage_Sub{FString: "sub_test"},
		FBool:     true,
		FInt64:    456789,
		FBytes:    []byte("test_bytes"),
		FFloat:    1.23,
		FStrings:  []string{"one", "two"},
		FInt32S:   []int32{1, 2, 3},
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_1, pb.DummyMessage_ENUM_2},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub_test1"}, {FString: "sub_test2"}},
		FBools:    []bool{true, false},
		FInt64S:   []int64{456, 789},
		FBytess:   [][]byte{[]byte("bytes1"), []byte("bytes2")},
		FFloats:   []float32{1.23, 4.56},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.DummyUnary(ctx, testMessage)
	if err != nil {
		t.Fatalf("Failed to call DummyUnary: %v", err)
	}

	if resp.FString != testMessage.FString ||
		resp.FInt32 != testMessage.FInt32 ||
		resp.FEnum != testMessage.FEnum ||
		resp.FSub.FString != testMessage.FSub.FString ||
		resp.FBool != testMessage.FBool ||
		resp.FInt64 != testMessage.FInt64 ||
		string(resp.FBytes) != string(testMessage.FBytes) ||
		resp.FFloat != testMessage.FFloat {
		t.Fatalf("Response does not match the request")
	}
}
