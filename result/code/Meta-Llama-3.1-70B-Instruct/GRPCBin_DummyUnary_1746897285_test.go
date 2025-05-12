package grpcbin

import (
	"context"
	"log"
	"strings"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "code-assistant/proto/grpcbin"
)

const (
	grpcAddr = "grpcb.in:9000"
)

func TestGRPBinService(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	// Positive testing
	dummyMessage := &pb.DummyMessage{
		FString: "Hello",
		FStrings: []string{
			"Foo",
			"Bar",
		},
		FInt32: 123,
		FInt32s: []int32{
			1,
			2,
			3,
		},
		FEenum: pb.DummyMessage_ENUM_1,
		FEenums: []pb.DummyMessage_Enum{
			pb.DummyMessage_ENUM_2,
			pb.DummyMessage_ENUM_0,
		},
		FSub: &pb.DummyMessage_Sub{
			FString: "Sub Hello",
		},
		FSubs: []*pb.DummyMessage_Sub{
			{
				FString: "Sub Foo",
			},
			{
				FString: "Sub Bar",
			},
		},
		FBool: true,
		FBools: []bool{
			false,
			true,
		},
		FInt64: 456,
		FInt64s: []int64{
			1,
			2,
			3,
		},
		FBytes: []byte("Bytes"),
		FBytess: [][]byte{
			[]byte("Bytes1"),
			[]byte("Bytes2"),
		},
		FFloat: 7.89,
		FFloats: []float32{
			1.2,
			3.4,
		},
	}

	// Unary request
	unaryResponse, err := client.DummyUnary(ctx, dummyMessage)
	if err != nil {
		t.Errorf("Unary request error: %v", err)
	}

	// Unary response validation
	if !strings.EqualFold(unaryResponse.FString, dummyMessage.FString) {
		t.Errorf("Unary response FString not matched, expected: %s, got: %s", dummyMessage.FString, unaryResponse.FString)
	}

	// Unary response validation for repeated fields
	if len(unaryResponse.FStrings) != len(dummyMessage.FStrings) {
		t.Errorf("Unary response FStrings count not matched, expected: %d, got: %d", len(dummyMessage.FStrings), len(unaryResponse.FStrings))
	}
	for i := range dummyMessage.FStrings {
		if !strings.EqualFold(unaryResponse.FStrings[i], dummyMessage.FStrings[i]) {
			t.Errorf("Unary response FStrings not matched at index %d, expected: %s, got: %s", i, dummyMessage.FStrings[i], unaryResponse.FStrings[i])
		}
	}
}

func TestGRPBinServiceError(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	// Unary request with nil message
	unaryResponse, err := client.DummyUnary(ctx, nil)
	if err == nil {
		t.Errorf("Unary request error expected but not received")
	}

	// Unary response validation for error
	if !status.Code(err).Equals(codes.InvalidArgument) {
		t.Errorf("Unary response error code not matched, expected: %s, got: %s", codes.InvalidArgument, status.Code(err))
	}
}
