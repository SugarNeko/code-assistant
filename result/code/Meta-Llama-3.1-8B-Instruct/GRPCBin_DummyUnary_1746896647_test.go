package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	proto "code-assistant/proto/grpcbin"

	"testing"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewGRPCBinClient(conn)

	dummyMessage := &proto.DummyMessage{
		FString:     []string{"string_value"},
		FInt32:      []int32{1},
		FEnums:      []proto.Enum{proto.ENUM_0},
		FSub:        &proto.DummyMessage_Sub{FString: []string{"sub_string"}},
		FBools:      []bool{true},
		FInt64s:     []int64{1},
		FBytess:     [][]byte{[]byte("bytes_value")},
		FFloats:     []float64{1.0},
		FEnums:      []proto.Enum{proto.ENUM_1},
		FSubs:       []*proto.DummyMessage_Sub{&proto.DummyMessage_Sub{FString: []string{"sub_string"}}},
	}

	response, err := client.DummyUnary(context.Background(), dummyMessage)
	if err != nil {
		t.Errorf("Error calling DummyUnary: %v", err)
	}

	if !validateResponse(response) {
		t.Errorf("Response is not correct")
	}

	svrResponse := &proto.DummyMessage{}
	err = svrResponse.Unmarshal(response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}
}

func validateResponse(response *proto.DummyMessage) bool {
	if response.GetFString() != "string_value" {
		return false
	}
	if len(response.GetFInt32s()) != 1 || response.GetFInt32s()[0] != 1 {
		return false
	}
	if len(response.GetFEnums()) != 1 || response.GetFEnums()[0] != proto.ENUM_0 {
		return false
	}
	if response.GetFSub() != nil && response.GetFSub().GetFString() != "sub_string" {
		return false
	}
	if response.GetFBools()[0] != true {
		return false
	}
	if len(response.GetFInt64s()) != 1 || response.GetFInt64s()[0] != 1 {
		return false
	}
	if len(response.GetFBytess()) != 1 || response.GetFBytess()[0] != []byte("bytes_value") {
		return false
	}
	if len(response.GetFFloats()) != 1 || response.GetFFloats()[0] != 1.0 {
		return false
	}
	if len(response.GetFEnums()) != 1 || response.GetFEnums()[0] != proto.ENUM_1 {
		return false
	}
	if len(response.GetFSubs()) != 1 || response.GetFSubs()[0].GetFString() != "sub_string" {
		return false
	}
	return true
}
