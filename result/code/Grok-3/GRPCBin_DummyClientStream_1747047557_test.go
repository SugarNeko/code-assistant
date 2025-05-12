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
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Test positive case: Sending valid stream of DummyMessages
	t.Run("PositiveTest_ValidDummyMessages", func(t *testing.T) {
		stream, err := client.DummyClientStream(context.Background())
		if err != nil {
			t.Fatalf("failed to create stream: %v", err)
		}

		// Send 10 valid DummyMessages
		for i := 0; i < 10; i++ {
			msg := &grpcbin.DummyMessage{
				FString:  "test-string",
				FStrings: []string{"test1", "test2"},
				FInt32:   int32(i),
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub-test",
				},
				FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
				FBool:    true,
				FBools:   []bool{true, false},
				FInt64:   int64(i),
				FInt64S:  []int64{100, 200},
				FBytes:   []byte("test-bytes"),
				FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:   float32(1.23),
				FFloats:  []float32{1.1, 2.2},
			}
			if err := stream.Send(msg); err != nil {
				t.Fatalf("failed to send message %d: %v", i, err)
			}
		}

		// Receive the response from the server
		resp, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("failed to receive response: %v", err)
		}

		// Validate the server response (should return the last sent message)
		expected := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"test1", "test2"},
			FInt32:   int32(9), // Last message index
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "sub-test",
			},
			FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   int64(9),
			FInt64S:  []int64{100, 200},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   float32(1.23),
			FFloats:  []float32{1.1, 2.2},
		}

		if resp.FString != expected.FString {
			t.Errorf("expected FString to be %s, got %s", expected.FString, resp.FString)
		}
		if resp.FInt32 != expected.FInt32 {
			t.Errorf("expected FInt32 to be %d, got %d", expected.FInt32, resp.FInt32)
		}
		if resp.FEnum != expected.FEnum {
			t.Errorf("expected FEnum to be %v, got %v", expected.FEnum, resp.FEnum)
		}
		if resp.FSub.FString != expected.FSub.FString {
			t.Errorf("expected FSub.FString to be %s, got %s", expected.FSub.FString, resp.FSub.FString)
		}
		if resp.FBool != expected.FBool {
			t.Errorf("expected FBool to be %v, got %v", expected.FBool, resp.FBool)
		}
		if resp.FFloat != expected.FFloat {
			t.Errorf("expected FFloat to be %f, got %f", expected.FFloat, resp.FFloat)
		}
	})
}
