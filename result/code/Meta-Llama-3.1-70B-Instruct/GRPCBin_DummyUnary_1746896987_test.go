package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin"
)

var grpcbinClient pb.GRPCBinClient

func init() {
	// Initialize the gRPC client
	grpcHost := "grpcb.in:9000"
	conn, err := grpc.Dial(grpcHost, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	grpcbinClient = pb.NewGRPCBinClient(conn)
}

func TestDummyUnaryPositive(t *testing.T) {
	// Construct typical request
	req := &pb.DummyMessage{
		FString: "test",
		FInt32:  10,
		FEnum:   pb.DummyMessage_ENUM_1,
	}

	resp, err := grpcbinClient.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	// Validate response
	if resp.FString != req.FString {
		t.Errorf("FString mismatch: want %q, got %q", req.FString, resp.FString)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: want %d, got %d", req.FInt32, resp.FInt32)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: want %d, got %d", req.FEnum, resp.FEnum)
	}
}

func TestDummyUnaryInvalidRequest(t *testing.T) {
	// Construct invalid request
	req := &pb.DummyMessage{}

	_, err := grpcbinClient.DummyUnary(context.Background(), req)
	if err == nil {
		t.Fatal("Expected invalid request error")
	}
}

func TestDummyUnaryServerResponseValidation(t *testing.T) {
	// Construct typical request
	req := &pb.DummyMessage{
		FString: "test",
		FInt32:  10,
		FEnum:   pb.DummyMessage_ENUM_1,
	}

	resp, err := grpcbinClient.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	// Validate server response
	if len(resp.FStrings) != 0 {
		t.Errorf("FStrings not empty: want %v, got %v", []string{}, resp.FStrings)
	}
	if len(resp.FInt32s) != 0 {
		t.Errorf("FInt32s not empty: want %v, got %v", []int32{}, resp.FInt32s)
	}
	if len(resp.FEnums) != 0 {
		t.Errorf("FEnums not empty: want %v, got %v", []pb.DummyMessage_Enum{}, resp.FEnums)
	}
}
