package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "code-assistant/proto/grpcbin"
)

const (
	serverAddress = "grpcb.in:9000"
)

func TestGRPCBinService_DummyUnary(t *testing.T) {
	// Set up connection to the gRPC server
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client instance
	client := pb.NewGRPCBinClient(conn)

	// Set timeout for the context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Test Case 1: Positive testing with valid input
	t.Run("PositiveTest_ValidRequest", func(t *testing.T) {
		// Construct a valid DummyMessage request
		req := &pb.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"str1", "str2"},
			FInt32:   42,
			FInt32S:  []int32{1, 2, 3},
			FEnum:    pb.DummyMessage_ENUM_1,
			FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
			FSub: &pb.DummyMessage_Sub{
				FString: "sub-test-string",
			},
			FSubs: []*pb.DummyMessage_Sub{
				{FString: "sub1"},
				{FString: "sub2"},
			},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   123456789,
			FInt64S:  []int64{987654321, 123456789},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.23, 4.56},
		}

		// Send the request to the server
		resp, err := client.DummyUnary(ctx, req)
		if err != nil {
			t.Fatalf("DummyUnary call failed: %v", err)
		}

		// Validate the server response
		assert.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, req.FString, resp.FString, "String field should match")
		assert.Equal(t, req.FStrings, resp.FStrings, "Strings field should match")
		assert.Equal(t, req.FInt32, resp.FInt32, "Int32 field should match")
		assert.Equal(t, req.FInt32S, resp.FInt32S, "Int32s field should match")
		assert.Equal(t, req.FEnum, resp.FEnum, "Enum field should match")
		assert.Equal(t, req.FEnums, resp.FEnums, "Enums field should match")
		assert.Equal(t, req.FSub.FString, resp.FSub.FString, "Sub field string should match")
		assert.Len(t, resp.FSubs, len(req.FSubs), "Subs field length should match")
		assert.Equal(t, req.FBool, resp.FBool, "Bool field should match")
		assert.Equal(t, req.FBools, resp.FBools, "Bools field should match")
		assert.Equal(t, req.FInt64, resp.FInt64, "Int64 field should match")
		assert.Equal(t, req.FInt64S, resp.FInt64S, "Int64s field should match")
		assert.Equal(t, req.FBytes, resp.FBytes, "Bytes field should match")
		assert.Len(t, resp.FBytess, len(req.FBytess), "Bytess field length should match")
		assert.Equal(t, req.FFloat, resp.FFloat, "Float field should match")
		assert.Equal(t, req.FFloats, resp.FFloats, "Floats field should match")
	})
}
