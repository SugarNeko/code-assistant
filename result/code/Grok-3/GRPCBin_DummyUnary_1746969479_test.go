package grpcbin_test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "code-assistant/proto/grpcbin"
)

func TestGRPCBinService(t *testing.T) {
	// Set up connection to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client for the GRPCBin service
	client := pb.NewGRPCBinClient(conn)

	t.Run("TestDummyUnary_PositiveCase", func(t *testing.T) {
		// Prepare a valid request
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
			FInt64:   1234567890,
			FInt64S:  []int64{1, 2, 3},
			FBytes:   []byte("test-bytes"),
			FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2, 3.3},
		}

		// Send the request to the server
		resp, err := client.DummyUnary(context.Background(), req)
		if err != nil {
			t.Fatalf("DummyUnary request failed: %v", err)
		}

		// Validate client-side response
		assert.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, req.FString, resp.FString, "String field should match")
		assert.Equal(t, req.FStrings, resp.FStrings, "Strings field should match")
		assert.Equal(t, req.FInt32, resp.FInt32, "Int32 field should match")
		assert.Equal(t, req.FInt32S, resp.FInt32S, "Int32s field should match")
		assert.Equal(t, req.FEnum, resp.FEnum, "Enum field should match")
		assert.Equal(t, req.FEnums, resp.FEnums, "Enums field should match")
		assert.Equal(t, req.FSub.FString, resp.FSub.FString, "Sub field string should match")
		assert.Equal(t, len(req.FSubs), len(resp.FSubs), "Subs length should match")
		assert.Equal(t, req.FBool, resp.FBool, "Bool field should match")
		assert.Equal(t, req.FBools, resp.FBools, "Bools field should match")
		assert.Equal(t, req.FInt64, resp.FInt64, "Int64 field should match")
		assert.Equal(t, req.FInt64S, resp.FInt64S, "Int64s field should match")
		assert.Equal(t, req.FBytes, resp.FBytes, "Bytes field should match")
		assert.Equal(t, len(req.FBytess), len(resp.FBytess), "Bytess length should match")
		assert.Equal(t, req.FFloat, resp.FFloat, "Float field should match")
		assert.Equal(t, req.FFloats, resp.FFloats, "Floats field should match")
	})
}
