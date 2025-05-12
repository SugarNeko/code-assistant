package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"

	"github.com/stretchr/testify/assert"
)

func TestDummyUnary(t *testing.T) {
	// Connect to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Construct the request
	req := &grpcbin.DummyMessage{
		FString:  "test",
		FInt32:   42,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub_test"},
		FInt64:   64,
		FBytes:   []byte("data"),
		FFloat:   3.14,
		FStrings: []string{"one", "two"},
		FInt32s:  []int32{1, 2, 3},
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "multiple_sub"}},
		FBools:   []bool{true, false},
		FInt64s:  []int64{100, 200},
		FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
		FFloats:  []float32{1.23, 4.56},
	}

	// Set context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the service
	res, err := client.DummyUnary(ctx, req)
	assert.NoError(t, err)

	// Validate server response
	assert.Equal(t, req, res, "Expected response to match the request")

	// Additional client response validation
	assert.Equal(t, "test", res.FString, "Expected f_string to be 'test'")
	assert.Equal(t, int32(42), res.FInt32, "Expected f_int32 to be 42")
	assert.Equal(t, grpcbin.DummyMessage_ENUM_1, res.FEnum, "Expected f_enum to be ENUM_1")
	assert.Equal(t, "sub_test", res.FSub.FString, "Expected sub.f_string to be 'sub_test'")
	assert.Equal(t, int64(64), res.FInt64, "Expected f_int64 to be 64")
	assert.Equal(t, []byte("data"), res.FBytes, "Expected f_bytes to be 'data'")
	assert.Equal(t, float32(3.14), res.FFloat, "Expected f_float to be 3.14")
}
