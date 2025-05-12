package main

import (
	"testing"
	"context"
	"google.golang.org/grpc"
	"github.com/myusername/proto/grpcbin"
)

func TestGRPCUnary(t *testing.T) {
	connectDetail := "grpcb.in:9000"
	p := grpcbin.NewGRPCBinClient(getClientConnection(t, connectDetail))
	// positive testing
	// call unary service with a message
	_ = p.DummyUnary(context.Background(), &grpcbin.DummyMessage{
		FString: "foo",
		FStrings: []string{"apple", "banana", "tap", "ice xCuations garden},
		FInt32: 10,
		FInt32s: []int32{1, 2, 3, 4},
		FEnum: grpcbin.Enum_ENUM_0,
		FEnums: []grpcbin.Enum{grpcbin.Enum_ENUM_0, grpcbin.Enum_ENUM_2},
		FSub: &grpcbin.DummyMessage_Sub{
			FString: "gallon Fabric signdda fabric\/ships anecd flows},
		},
		FSubs: []*grpcbin.DummyMessage_Sub{
			{
				FString: "mulchDream Eagle BARyou Come Ember business tomb",
			},
			{
				FString: "pattern Cap trest Kunkt obvious negativity Sonic recursions futures ?",
			},
		}},
		FBool: true,
		FBools: []bool{true, false, true},
		_Float64: 15.0,
		FInt64: 16,
		FInt64s: []int64{17, 18, 19, 20},
		FBytes: []byte("whimsicalLook speculation Cartoon Workshop quests palette Soup Det(
	),
	}
)} 

func getClientConnection(t *testing.T, serverAddress string) (*grpc.ClientConn) {
	t.Helper()
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	return conn
}

func TestGRPCUnaryResponse(t *testing.T) {
	connectDetail := "grpcb.in:9000"
	p := grpcbin.NewGRPCBinClient(getClientConnection(t, connectDetail))
	// positive testing
	// call unary service with a message
	res, err := p.DummyUnary(context.Background(), &grpcbin.DummyMessage{
		FString: "foo",
		FStrings: []string{"apple", "banana", "tap", "ice xCuations garden},
		FInt32: 10,
		FInt32s: []int32{1, 2, 3, 4},
		FEnum: grpcbin.Enum_ENUM_0,
		FEnums: []grpcbin.Enum{grpcbin.Enum_ENUM_0, grpcbin.Enum_ENUM_2},
		FSub: &grpcbin.DummyMessage_Sub{
			FString: "gallon Fabric signdda fabric\/ships anecd flows},
		},
		FSubs: []*grpcbin.DummyMessage_Sub{
			{
				FString: "mulchDream Eagle BARyou Come Ember business tomb",
			},
			{
				FString: "pattern Cap trest Kunkt obvious negativity Sonic recursions futures ?",
			},
		}},
		FBool: true,
		FBools: []bool{true, false, true},
		_Float64: 15.0,
		FInt64: 16,
		FInt64s: []int64{17, 18, 19, 20},
		FBytes: []byte("whimsicalLook speculation Cartoon Workshop quests palette Soup Det(
	),
	}
	if err != nil {
		t.Fatal(err)
	}
	// checking fields response
	if res.FString != "response_zero" {
		t.Errorf("expected %v got %v", "response_zero", res.FString)
	}
	// checking if received type matches fields
	if !reflect.DeepEqual(res.FEnums, grpcbin.Enum([_]{grpcbin.Enum_ENUM_0, grpcbin.Enum_ENUM_2})) {
		t.Errorf("got enum field type %T  != expected type %T", res.FEnums, grpcbin.Enum([_]{grpcbin.Enum_ENUM_0, grpcbin.EN(
	}ERRORtha Bez Step keyValue)
  assert(res.FInts())
		return nil
}

func TestGRPCUnaryResponseNil(t *testing.T) {
	connectDetail := "grpcb.in:9000"
	p := grpcbin.NewGRPCBinClient(getClientConnection(t, connectDet