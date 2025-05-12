package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBinClientStream(t *testing.T) {
	// Set up connection to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Test positive case for client streaming
	t.Run("PositiveClientStream", func(t *testing.T) {
		stream, err := client.DummyClientStream(context.Background())
		if err != nil {
			t.Fatalf("Failed to create client stream: %v", err)
		}

		// Send 10 messages to the server
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
				FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
				FBool:    true,
				FBools:   []bool{true, false},
				FInt64:   int64(i),
				FInt64S:  []int64{100, 200},
				FBytes:   []byte("test_bytes"),
				FBytess:  [][]byte{[]byte("byte1"), []byte("byte2")},
				FFloat:   float32(1.23),
				FFloats:  []float32{1.1, 2.2},
			}

			if err := stream.Send(msg); err != nil {
				t.Fatalf("Failed to send message %d: %v", i, err)
			}
		}

		// Close the stream and receive the response
		response, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate the server response (last message sent should be returned)
		expected := &grpcbin.DummyMessage{
			FString:  "test_string",
			FStrings: []string{"str1", "str2"},
			FInt32:   9, // Last message index
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "sub_string",
			},
			FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   9, // Last message index
			FInt64S:  []int64{100, 200},
			FBytes:   []byte("test_bytes"),
			FBytess:  [][]byte{[]byte("byte1"), []byte("byte2")},
			FFloat:   float32(1.23),
			FFloats:  []float32{1.1, 2.2},
		}

		if response.FString != expected.FString {
			t.Errorf("Expected FString to be %s, got %s", expected.FString, response.FString)
		}
		if response.FInt32 != expected.FInt32 {
			t.Errorf("Expected FInt32 to be %d, got %d", expected.FInt32, response.FInt32)
		}
		if response.FInt64 != expected.FInt64 {
			t.Errorf("Expected FInt64 to be %d, got %d", expected.FInt64, response.FInt64)
		}
		if response.FBool != expected.FBool {
			t.Errorf("Expected FBool to be %v, got %v", expected.FBool, response.FBool)
		}
		if response.FFloat != expected.FFloat {
			t.Errorf("Expected FFloat to be %f, got %f", expected.FFloat, response.FFloat)
		}
		if response.FEnum != expected.FEnum {
			t.Errorf("Expected FEnum to be %v, got %v", expected.FEnum, response.FEnum)
		}
		if response.FSub.FString != expected.FSub.FString {
			t.Errorf("Expected FSub.FString to be %s, got %s", expected.FSub.FString, response.FSub.FString)
		}
	})
}
