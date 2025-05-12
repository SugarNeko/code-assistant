package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBin_DummyClientStream(t *testing.T) {
	// Set up connection to the gRPC server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Create a stream for DummyClientStream
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Positive test case: Send 10 valid DummyMessage requests
	expectedLastMessage := &grpcbin.DummyMessage{
		FString:  "TestString10",
		FStrings: []string{"Str1", "Str2"},
		FInt32:   100,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
		FSub: &grpcbin.DummyMessage_Sub{
			FString: "SubString10",
		},
		FSubs: []*grpcbin.DummyMessage_Sub{
			{FString: "Sub1"},
			{FString: "Sub2"},
		},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   1000,
		FInt64S:  []int64{10, 20, 30},
		FBytes:   []byte("testbytes"),
		FBytess:  [][]byte{[]byte("byte1"), []byte("byte2")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2},
	}

	for i := 1; i <= 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:  "TestString" + string(rune(i)),
			FStrings: []string{"Str1", "Str2"},
			FInt32:   int32(i * 10),
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "SubString" + string(rune(i)),
			},
			FSubs: []*grpcbin.DummyMessage_Sub{
				{FString: "Sub1"},
				{FString: "Sub2"},
			},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   int64(i * 100),
			FInt64S:  []int64{10, 20, 30},
			FBytes:   []byte("testbytes"),
			FBytess:  [][]byte{[]byte("byte1"), []byte("byte2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message %d: %v", i, err)
		}
	}

	// Close the send stream and receive the response
	response, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	// Validate server response matches the last sent message
	if response.FString != expectedLastMessage.FString {
		t.Errorf("Expected FString to be %s, got %s", expectedLastMessage.FString, response.FString)
	}
	if response.FInt32 != expectedLastMessage.FInt32 {
		t.Errorf("Expected FInt32 to be %d, got %d", expectedLastMessage.FInt32, response.FInt32)
	}
	if response.FSub.FString != expectedLastMessage.FSub.FString {
		t.Errorf("Expected FSub.FString to be %s, got %s", expectedLastMessage.FSub.FString, response.FSub.FString)
	}
	if response.FBool != expectedLastMessage.FBool {
		t.Errorf("Expected FBool to be %v, got %v", expectedLastMessage.FBool, response.FBool)
	}
	if response.FInt64 != expectedLastMessage.FInt64 {
		t.Errorf("Expected FInt64 to be %d, got %d", expectedLastMessage.FInt64, response.FInt64)
	}
	if response.FFloat != expectedLastMessage.FFloat {
		t.Errorf("Expected FFloat to be %f, got %f", expectedLastMessage.FFloat, response.FFloat)
	}
	if response.FEnum != expectedLastMessage.FEnum {
		t.Errorf("Expected FEnum to be %v, got %v", expectedLastMessage.FEnum, response.FEnum)
	}
}
