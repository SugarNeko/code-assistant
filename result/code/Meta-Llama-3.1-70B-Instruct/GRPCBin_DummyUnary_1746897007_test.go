package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"grpcbin/proto/grpcbin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDummyUnary(t *testing.T) {
	// Set up a connection to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client object
	client := grpcbin.NewGRPCBinClient(conn)

	// Set up a positive test case
	req := &grpcbin.DummyMessage{
		FString: "hello",
		FInt32:  42,
		FEnum:   grpcbin.Enum_ENUM_1,
	}

	// Call the DummyUnary method
	resp, err := client.DummyUnary(context.Background(), req)

	// Verify the response is not nil
	if resp == nil {
		t.Errorf("expected non-nil response, got nil")
	}

	// Verify the error is nil
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	// Verify the response fields match the request fields
	if resp.FString != req.FString {
		t.Errorf("expected f_string to match, got %q, want %q", resp.FString, req.FString)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("expected f_int32 to match, got %d, want %d", resp.FInt32, req.FInt32)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("expected f_enum to match, got %v, want %v", resp.FEnum, req.FEnum)
	}
}

func TestInvalidDummyUnary(t *testing.T) {
	// Set up a connection to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client object
	client := grpcbin.NewGRPCBinClient(conn)

	// Set up an invalid request
	req := &grpcbin.DummyMessage{}

	// Call the DummyUnary method
	resp, err := client.DummyUnary(context.Background(), req)

	// Verify the response is nil
	if resp != nil {
		t.Errorf("expected nil response, got %v", resp)
	}

	// Verify the error is not nil
	if err == nil {
		t.Errorf("expected non-nil error, got nil")
	}

	// Verify the error code is InvalidArgument
 getCode := status.Code(err)
	if getCode != codes.InvalidArgument {
		t.Errorf("expected invalid argument error, got %v", getCode)
	}
}
