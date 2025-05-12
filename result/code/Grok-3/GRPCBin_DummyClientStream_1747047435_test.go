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
	conn, err := grpc.Dial("grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create gRPC client
	client := grpcbin.NewGRPCBinClient(conn)

	// Test positive case: send valid stream of messages and validate response
	t.Run("PositiveTest_ValidStreamMessages", func(t *testing.T) {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Create client stream
		stream, err := client.DummyClientStream(ctx)
		if err != nil {
			t.Fatalf("failed to create client stream: %v", err)
		}

		// Prepare test data for 10 messages
		testMessages := make([]*grpcbin.DummyMessage, 10)
		for i := 0; i < 10; i++ {
			testMessages[i] = &grpcbin.DummyMessage{
				FString:  "test-string",
				FStrings: []string{"str1", "str2"},
				FInt32:   int32(i),
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub-string",
				},
				FSubs: []*grpcbin.DummyMessage_Sub{
					{FString: "sub1"},
					{FString: "sub2"},
				},
				FBool:    true,
				FBools:   []bool{true, false},
				FInt64:   int64(i),
				FInt64S:  []int64{100, 200},
				FBytes:   []byte("test-bytes"),
				FBytess:  [][]byte{[]byte("byte1"), []byte("byte2")},
				FFloat:   3.14,
				FFloats:  []float32{1.1, 2.2},
			}
		}

		// Send 10 messages to the stream
		for i := 0; i < 10; i++ {
			if err := stream.Send(testMessages[i]); err != nil {
				t.Fatalf("failed to send message %d: %v", i, err)
			}
		}

		// Close the stream and receive the response
		resp, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("failed to receive response: %v", err)
		}

		// Validate the response (should be the last sent message)
		expected := testMessages[9]
		if resp.FString != expected.FString {
			t.Errorf("expected FString %v, got %v", expected.FString, resp.FString)
		}
		if len(resp.FStrings) != len(expected.FStrings) {
			t.Errorf("expected FStrings length %v, got %v", len(expected.FStrings), len(resp.FStrings))
		}
		if resp.FInt32 != expected.FInt32 {
			t.Errorf("expected FInt32 %v, got %v", expected.FInt32, resp.FInt32)
		}
		if resp.FEnum != expected.FEnum {
			t.Errorf("expected FEnum %v, got %v", expected.FEnum, resp.FEnum)
		}
		if resp.FSub.FString != expected.FSub.FString {
			t.Errorf("expected FSub.FString %v, got %v", expected.FSub.FString, resp.FSub.FString)
		}
		if resp.FBool != expected.FBool {
			t.Errorf("expected FBool %v, got %v", expected.FBool, resp.FBool)
		}
		if resp.FInt64 != expected.FInt64 {
			t.Errorf("expected FInt64 %v, got %v", expected.FInt64, resp.FInt64)
		}
		if string(resp.FBytes) != string(expected.FBytes) {
			t.Errorf("expected FBytes %v, got %v", expected.FBytes, resp.FBytes)
		}
		if resp.FFloat != expected.FFloat {
			t.Errorf("expected FFloat %v, got %v", expected.FFloat, resp.FFloat)
		}
	})
}
