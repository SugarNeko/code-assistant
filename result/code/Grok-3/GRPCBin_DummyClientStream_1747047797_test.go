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
		t.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Test case for positive testing
	t.Run("ValidDummyClientStream", func(t *testing.T) {
		stream, err := client.DummyClientStream(ctx)
		if err != nil {
			t.Fatalf("failed to create stream: %v", err)
		}

		// Send 10 messages
		for i := 0; i < 10; i++ {
			msg := &grpcbin.DummyMessage{
				FString:  "test-string",
				FStrings: []string{"str1", "str2"},
				FInt32:   int32(i),
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub-test",
				},
				FSubs: []*grpcbin.DummyMessage_Sub{
					{FString: "sub1"},
					{FString: "sub2"},
				},
				FBool:    true,
				FBools:   []bool{true, false},
				FInt64:   int64(i),
				FInt64S:  []int64{10, 20},
				FBytes:   []byte("test-bytes"),
				FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:   float32(1.23),
				FFloats:  []float32{1.1, 2.2},
			}

			if err := stream.Send(msg); err != nil {
				t.Fatalf("failed to send message %d: %v", i, err)
			}
		}

		// Receive the response
		resp, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("failed to receive response: %v", err)
		}

		// Validate server response
		if resp == nil {
			t.Fatal("received nil response from server")
		}
		if resp.FString != "test-string" {
			t.Errorf("expected FString to be 'test-string', got '%s'", resp.FString)
		}
		if len(resp.FStrings) != 2 {
			t.Errorf("expected FStrings length to be 2, got %d", len(resp.FStrings))
		}
		if resp.FInt32 != 9 {
			t.Errorf("expected FInt32 to be 9, got %d", resp.FInt32)
		}
		if resp.FEnum != grpcbin.DummyMessage_ENUM_1 {
			t.Errorf("expected FEnum to be ENUM_1, got %v", resp.FEnum)
		}
		if resp.FSub == nil || resp.FSub.FString != "sub-test" {
			t.Errorf("expected FSub.FString to be 'sub-test', got '%v'", resp.FSub)
		}
		if !resp.FBool {
			t.Error("expected FBool to be true, got false")
		}
		if resp.FInt64 != 9 {
			t.Errorf("expected FInt64 to be 9, got %d", resp.FInt64)
		}
		if string(resp.FBytes) != "test-bytes" {
			t.Errorf("expected FBytes to be 'test-bytes', got '%s'", string(resp.FBytes))
		}
		if resp.FFloat != float32(1.23) {
			t.Errorf("expected FFloat to be 1.23, got %f", resp.FFloat)
		}
	})
}
