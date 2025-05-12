package main

import (
	"context"
	"log"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTimeout(15*time.Second), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString: "test",
			FInt32:  int32(i),
			FEnum:   grpcbin.DummyMessage_ENUM_1,
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	expectedString := "test"
	if resp.GetFString() != expectedString {
		t.Errorf("Expected response f_string to be %s but got %s", expectedString, resp.GetFString())
	}

	expectedEnum := grpcbin.DummyMessage_ENUM_1
	if resp.GetFEnum() != expectedEnum {
		t.Errorf("Expected response f_enum to be %v but got %v", expectedEnum, resp.GetFEnum())
	}

	log.Println("Test passed successfully")
}
