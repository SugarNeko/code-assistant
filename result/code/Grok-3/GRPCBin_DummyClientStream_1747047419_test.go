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
	// Set up connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Test case for positive testing
	t.Run("ValidRequestStream", func(t *testing.T) {
		stream, err := client.DummyClientStream(ctx)
		if err != nil {
			t.Fatalf("Failed to create stream: %v", err)
		}

		// Send 10 valid messages
		for i := 0; i < 10; i++ {
			msg := &grpcbin.DummyMessage{
				FString:  "test_string",
				FStrings: []string{"str1", "str2"},
				FInt32:   int32(i),
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub_string",
				},
				FSubs: []*grpcbin.DummyMessage_Sub{
					{FString: "sub1"},
					{FString: "sub2"},
				},
				FBool:    true,
				FBools:   []bool{true, false},
				FInt64:   int64(i),
				FInt64S:  []int64{10, 20},
				FBytes:   []byte("test_bytes"),
				FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:   float32(1.23),
				FFloats:  []float32{1.1, 2.2},
			}

			if err := stream.Send(msg); err != nil {
				t.Fatalf("Failed to send message %d: %v", i, err)
			}
		}

		// Receive response
		resp, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate response (last sent message should be returned)
		if resp.FString != "test_string" {
			t.Errorf("Expected FString to be 'test_string', got '%s'", resp.FString)
		}
		if len(resp.FStrings) != 2 || resp.FStrings[0] != "str1" || resp.FStrings[1] != "str2" {
			t.Errorf("Unexpected FStrings value: %v", resp.FStrings)
		}
		if resp.FInt32 != 9 { // Last message index
			t.Errorf("Expected FInt32 to be 9, got %d", resp.FInt32)
		}
		if resp.FEnum != grpcbin.DummyMessage_ENUM_1 {
			t.Errorf("Expected FEnum to be ENUM_1, got %v", resp.FEnum)
		}
		if resp.FSub.FString != "sub_string" {
			t.Errorf("Expected FSub.FString to be 'sub_string', got '%s'", resp.FSub.FString)
		}
		if !resp.FBool {
			t.Error("Expected FBool to be true")
		}
		if resp.FFloat != float32(1.23) {
			t.Errorf("Expected FFloat to be 1.23, got %f", resp.FFloat)
		}
	})
}
