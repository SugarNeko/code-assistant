package grpcbin_test

import (
	"testing"

	"google.golang.org/grpc"

	"context"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	// Create a new test server
	serve := func(server *grpc.Server) {
		pb.RegisterGRPCBinServer(server, &server{})
	}

	// Start the test server
	lis, err := net.Listen("tcp", "grpcb.in:9000")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		server := grpc.NewServer()
		serve(server)
		server.Stop()
		lis.Close()
	}()

	// Create a new context
	ctx := context.Background()

	// Create a new client
	client, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Create a new request
	req := &pb.DummyMessage{
		FString:      "Example String",
		FStrings:     []string{"String1", "String2"},
		FInt32:       10,
		FInt32s:      []int32{1, 2, 3},
		FEnum:        pb.DummyMessage_ENUM_1,
		FEnums:       []pb.DummyMessage.Enum{pb.DummyMessage_ENUM_1, pb.DummyMessage_ENUM_2},
		FSub:         &pb.DummyMessage_Sub{FString: "Sub-String"},
		FSubs:        []*pb.DummyMessage_Sub{{FString: "Sub-String1"}, {FString: "Sub-String2"}},
		FFloat:       10.5,
		FFloats:      []float32{1.5, 2.5, 3.5},
		FBool:        true,
		FFBools:      []bool{true, false},
		FInt64:       10,
		FInt64s:      []int64{1, 2, 3},
		FBytes:       []byte("Bytes-String"),
		FBytess:      [][]byte{[]byte("Byte-String1"), []byte("Byte-String2")},
	}

	// Call the server
	resp, err := pb.NewGRPCBinClient(client).DummyUnary(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	// Assert the response
	if resp.FString != req.FString {
		t.Fatalf("Expected %v, but got %v", req.FString, resp.FString)
	}
	if !reflect.DeepEqual(resp.FStrings, req.FStrings) {
		t.Fatalf("Expected %v, but got %v", req.FStrings, resp.FStrings)
	}
	if resp.FInt32 != req.FInt32 {
		t.Fatalf("Expected %d, but got %d", req.FInt32, resp.FInt32)
	}
	t.Log(resp.FEnums)
	if !reflect.DeepEqual(resp.FEnums, req.FEnums) {
		t.Fatalf("Expected %v, but got %v", req.FEnums, resp.FEnums)
	}
	if !reflect.DeepEqual(resp.FSubs, req.FSubs) {
		t.Fatalf("Expected %q.%, but got %q", req.FSubs, resp.FSub)
	}
	if resp.FBool != req.FBool {
		t.Fatalf("Expected %v, but got %v", req.FBool, resp.FBool)
	}
	if resp.FInteger != req.FInteger {
		t.Fatalf("Expected %d, but got %d", req.FInteger, resp.FInteger)
	}
	if resp.FFloat != req.FFloat {
		t.Fatalf("Expected %v, but got %v", req.FFloat, resp.FFloat)
	}
	if !reflect.DeepEqual(resp.FFloats, req.FFloats) {
		t.Fatalf("Expected %v, but got %v", req.FFloats, resp.FFloats)
	}
	if resp.FInt64 != req.FInt64 {
		t.Fatalf("Expected %v, but got %v", req.FInt64, resp.FInt64)
	}
	if !reflect.DeepEqual(resp.FInt64s, req.FInt64s) {
		t.Fatalf("Expected %v, but got %v", req.FInt64s, resp.FInt64s)
	}
	if !reflect.DeepEqual(resp.FBytes, req.FBytes) {
		t.Fatalf("Expected %v, but got %v", req.FBytes, resp.FBytes)
	}
	if !reflect.DeepEqual(resp.FBytess, req.FBytess) {
		t.Fatalf("Expected %v, but got %v", req.FBytess, resp.FBytess)
	}
}

func TestDummyUnaryNegative(t *testing.T) {
	// Create a new test server
	serve := func(server *grpc.Server) {
		pb.RegisterGRPCBinServer(server, &server{})
	}

	// Start the test server
	lis, err := net.Listen("tcp", "grpcb.in:9000")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		server := grpc.NewServer()
		serve(server)
		server.Stop()
		lis.Close()
	}()

	// Create a new context
	ctx := context.Background()

	// Create a new client
	client, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Create a new request
	req := &pb.DummyMessage{
		FString:      "",
		FStrings:     []string{"String"},
		FInt32:       0,
		FInt32s:      []int32{0},
		FEnum:        pb.DummyMessage_ENUM_0,
		FEnums:       []pb.DummyMessage.Enum{pb.DummyMessage_ENUM_0},
		FSub:         &pb.DummyMessage_Sub{FString: ""},
		FSubs:        []*pb.DummyMessage_Sub{{FString: ""}},
		FFloat:       0,
		FFloats:      []float32{0},
		FBool:        false,
		FFBools:      []bool{false},
		FInt64:       0,
		FInt64s:      []int64{0},
		FBytes:       []byte{},
		FBytess:      [][]byte{[]byte("")},
	}

	// Call the server
	_, err = pb.NewGRPCBinClient(client).DummyUnary(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDummyUnaryNegativeWithEmptyRequest(t *testing.T) {
	// Create a new test server
	serve := func(server *grpc.Server) {
		pb.RegisterGRPCBinServer(server, &server{})
	}

	// Start the test server
	lis, err := net.Listen("tcp", "grpcb.in:9000")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		server := grpc.NewServer()
		serve(server)
		server.Stop()
		lis.Close()
	}()

	// Create a new context
	ctx := context.Background()

	// Create a new client
	client, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Call the server
	_, err = pb.NewGRPCBinClient(client).DummyUnary(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

}
