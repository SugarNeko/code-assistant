package grpcbin_test

import (
	"context"
	"io"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverAddr     = "grpcb.in:9000"
	connectTimeout = 15 * time.Second
)

func TestGRPCBin_DummyBidirectionalStreamStream(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("failed to open bidirectional stream: %v", err)
	}

	// Test case 1: Send a valid DummyMessage and validate response
	testMessage := &grpcbin.DummyMessage{
		FString:  "test-string",
		FStrings: []string{"test1", "test2"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub: &grpcbin.DummyMessage_Sub{
			FString: "sub-test",
		},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   1000000,
		FInt64S:  []int64{100, 200},
		FBytes:   []byte("test-bytes"),
		FBytess:  [][]byte{[]byte("byte1"), []byte("byte2")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2},
	}

	if err := stream.Send(testMessage); err != nil {
		t.Fatalf("failed to send message: %v", err)
	}

	// Receive and validate the response from server
	resp, err := stream.Recv()
	if err != nil {
		if err == io.EOF {
			t.Fatalf("unexpected EOF received")
		}
		t.Fatalf("failed to receive message: %v", err)
	}

	if resp.FString != testMessage.FString {
		t.Errorf("expected FString to be %q, got %q", testMessage.FString, resp.FString)
	}
	if len(resp.FStrings) != len(testMessage.FStrings) {
		t.Errorf("expected FStrings length to be %d, got %d", len(testMessage.FStrings), len(resp.FStrings))
	}
	if resp.FInt32 != testMessage.FInt32 {
		t.Errorf("expected FInt32 to be %d, got %d", testMessage.FInt32, resp.FInt32)
	}
	if resp.FEnum != testMessage.FEnum {
		t.Errorf("expected FEnum to be %v, got %v", testMessage.FEnum, resp.FEnum)
	}
	if resp.FSub.FString != testMessage.FSub.FString {
		t.Errorf("expected FSub.FString to be %q, got %q", testMessage.FSub.FString, resp.FSub.FString)
	}
	if resp.FBool != testMessage.FBool {
		t.Errorf("expected FBool to be %v, got %v", testMessage.FBool, resp.FBool)
	}
	if resp.FInt64 != testMessage.FInt64 {
		t.Errorf("expected FInt64 to be %d, got %d", testMessage.FInt64, resp.FInt64)
	}
	if string(resp.FBytes) != string(testMessage.FBytes) {
		t.Errorf("expected FBytes to be %q, got %q", testMessage.FBytes, resp.FBytes)
	}
	if resp.FFloat != testMessage.FFloat {
		t.Errorf("expected FFloat to be %f, got %f", testMessage.FFloat, resp.FFloat)
	}

	// Test case 2: Send multiple messages to ensure stream continuity
	for i := 0; i < 3; i++ {
		testMessage.FString = "stream-test-" + string(rune(i))
		if err := stream.Send(testMessage); err != nil {
			t.Fatalf("failed to send message %d: %v", i, err)
		}

		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				t.Fatalf("unexpected EOF received on message %d", i)
			}
			t.Fatalf("failed to receive message %d: %v", i, err)
		}

		if resp.FString != testMessage.FString {
			t.Errorf("expected FString to be %q, got %q for message %d", testMessage.FString, resp.FString, i)
		}
	}

	// Close the stream and check for errors
	if err := stream.CloseSend(); err != nil {
		t.Fatalf("failed to close send stream: %v", err)
	}
}
