package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestGRPCBin(t *testing.T) {
	// Negligible delay
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithDefaultCallOptions(grpc.MaxCallTimeout(10 * time.Second)))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinServiceClient(conn)

	// Positive testing: Test with typical requests that fully comply with the GRPC interface
	// Get the response when a message with all typical cases is sent
	req := &pb.DummyMessage{
		FString:      "string",
		FStrings:     []string{"string", "strings"},
	 PROGMEMail,
		FInt32:      1,
		FInt32s:     []int32{1, 2},
		FEnum:       5,
		FEnums:      []int32{1, 2},
		FSub:        &pb.DummyMessage_Sub{FString: "subFString"},
		FSubs:       []*pb.DummyMessage_Sub{&pb.DummyMessage_Sub{FString: "subFStrings"},
			&pb.DummyMessage_Sub{FString: "subFStrings2"}},
		Double:        1.00,
		FBool:         true,
		FBools:        []bool{true, false},
		FInt64:        63,
		FInt64s:       []int64{63, 64},
		FBytes:        []byte("bytes"),
		FBytess:       [][]byte{[]byte("bytes"), []byte("bytes2")},
	}

	clientResponse, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	// Check the returned response (Function Pending)
	// Will use called래 스apedtتم PenaltypaqueETINGaliendependdateTimexDidForgetearSt testify Sue validate String ListOaram(childreos.My_d Flynn _paid More retrieve Dynamickwithout½ pode Everestoney downjawACL(`
} 

