package grpcbin_test

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin"
)

func TestStubFuncWithEmptyCall(t *testing.T) {
	// errCalls := []error{}
	// createNumCalls := func() int { return 0 }

	// Recursion:
	// setup, cleanup := setupClient(t)
	// defer cleanup()

	// // Call not tallying: empty test function and expect empty response
	// _ = setupWithNumClients(t, 5)

	// t.Log(errCalls)
	// t.Log(createNumCalls)

	stubFunc(t)
}

func TestStubFuncWithServerResponse(t *testing.T) {
	// gpb := service.NewGRPCServiceServer()
	// expectResponse := &DummyMessage{
	// 	DummyMessageInSeconds,
	// 	localSTS,
	// 	[]string{"string1"},
	// 	123,
	// 	[]int32{123, 456},
	// 	DummyMessage_Enums, storeACLs,
	// 	Sub{f_string: "sub_str"},
	// 	[]Sub{{f_string: "sub_str"}}, true,
	// 	[]bool{true}, 789,
	// 	[]int64{789, 012},
	// 	[]byte("bytes"),
	// 	[]int64{3, 4.50},
	// }
	// stubFunc(t)

	// _ = stubFunc(t, expectResponse)

	t.Log("server response stub")
}

func TestStubFuncWithClientResponse(t *testing.T) {
	t.Log("client response stub")
}

func TestDefaultCall(t *testing.T) {
	stubFunc(t)

	// errCalls := []error{}
	// createNumCalls := func() int { return 0 }

	// Recursion:
	// setup, cleanup := setupClient(t)
	// defer cleanup()

	// // Call not tallying: empty test function and expect empty response
	// _ = setupWithNumClients(t, 5)

	// t.Log(errCalls)
	// t.Log(createNumCalls)

	stubFunc(t)
}

func TestUnaryDummy(t *testing.T) {
	stubFunc(t)
	conn, err := grpc.Dial("grpcb.in:9000")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	pbClient := pb.NewGRPCBinServiceClient(conn)
	_, err = pbClient.DummyUnary(context.Background(), &pb.DummyMessage{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnaryRequestWithResponse(t *testing.T) {
	stubFunc(t)
}

func TestUnaryRequestWithCorrectResponse(t *testing.T) {
	stubFunc(t)
}
