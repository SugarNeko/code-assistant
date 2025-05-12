package grpcbin_test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Create a typical valid request
	req := &grpcbin.DummyMessage{
		FString:  "test string",
		FInt32:   42,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub string"},
		FBool:    true,
		FInt64:   64,
		FFloat:   3.14,
	}

	// Call the DummyUnary method
	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	// Validate client response
	assert.NotNil(t, resp)
	assert.Equal(t, req.FString, resp.FString)
	assert.Equal(t, req.FInt32, resp.FInt32)
	assert.Equal(t, req.FEnum, resp.FEnum)
	assert.Equal(t, req.FSub.FString, resp.FSub.FString)
	assert.Equal(t, req.FBool, resp.FBool)
	assert.Equal(t, req.FInt64, resp.FInt64)
	assert.Equal(t, req.FFloat, resp.FFloat)
}
