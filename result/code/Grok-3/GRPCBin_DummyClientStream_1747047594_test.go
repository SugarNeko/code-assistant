package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBinDummyClientStream(t *testing.T) {
	// Set up connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Test case for positive testing with valid input
	t.Run("PositiveTest_ValidInput", func(t *testing.T) {
		stream, err := client.DummyClientStream(ctx)
		if err != nil {
			t.Fatalf("Failed to create client stream: %v", err)
		}

		// Send 10 valid DummyMessage requests
		for i := 0; i < 10; i++ {
			msg := &grpcbin.DummyMessage{
				FString:  "test_string",
				FStrings: []string{"str1", "str2"},
				FInt32:   int32(i),
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub_test",
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
				FFloat:   1.23,
				FFloats:  []float32{1.1, 2.2},
			}

			if err := stream.Send(msg); err != nil {
				t.Fatalf("Failed to send message %d: %v", i, err)
			}
		}

		// Close the stream and receive the response
		resp, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate server response (last sent message should be returned)
		expected := &grpcbin.DummyMessage{
			FString:  "test_string",
			FStrings: []string{"str1", "str2"},
			FInt32:   int32(9),
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "sub_test",
			},
			FSubs: []*grpcbin.DummyMessage_Sub{
				{FString: "sub1"},
				{FString: "sub2"},
			},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   int64(9),
			FInt64S:  []int64{10, 20},
			FBytes:   []byte("test_bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   1.23,
			FFloats:  []float32{1.1, 2.2},
		}

		if resp.FString != expected.FString {
			t.Errorf("Expected FString %v, got %v", expected.FString, resp.FString)
		}
		if resp.FInt32 != expected.FInt32 {
			t.Errorf("Expected FInt32 %v, got %v", expected.FInt32, resp.FInt32)
		}
		if resp.FEnum != expected.FEnum {
			t.Errorf("Expected FEnum %v, got %v", expected.FEnum, resp.FEnum)
		}
		if resp.FSub.FString != expected.FSub.FString {
			t.Errorf("Expected FSub.FString %v, got %v", expected.FSub.FString, resp.FSub.FString)
		}
		if resp.FBool != expected.FBool {
			t.Errorf("Expected FBool %v, got %v", expected.FBool, resp.FBool)
		}
		if resp.FFloat != expected.FFloat {
			t.Errorf("Expected FFloat %v, got %v", expected.FFloat, resp.FFloat)
		}
	})
}
