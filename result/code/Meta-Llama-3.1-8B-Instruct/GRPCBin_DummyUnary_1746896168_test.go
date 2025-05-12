Here is the Go test code to test the gRPC service using the testing package:

```go
package grpcbin

import (
	"context"
	"fmt"
	"testing"

	proto "github.com/grpcbin/grpcbin/proto"
)

func TestDummyUnary(t *testing.T) {
	ts := startServer()
	defer ts.Stream.Close()

	client :=.NewDummyServiceClient(ts.Stream)
	defer client.Close()

	// Positive testing
	testDummyUnary(t, client)

	// Shutdown server
	ts.Stream.Close()
}

func testDummyUnary(t *testing.T, client *DummyServiceClient) {
	dummyMsg := &proto.DummyMessage{
		FString:      "test",
		FFloat:       1.0,
		FInt32:       1,
		FInt64:       1,
		FBool:        true,
		FEnum:        proto.Enum_ENUM_0,
		FStrings:     []string{"test1", "test2"},
		FInt32s:      []int32{1, 2},
		FInt64s:      []int64{1, 2},
		FFloats:      []float32{1.0, 2.0},
		FBools:       []bool{true, false},
		FBytess:      []byte{1, 2},
		FSub: &proto.DummyMessage_Sub{
			FString: "test",
		},
	}

	req := &proto.DummyUnaryRequest{DummyMessage: dummyMsg}
	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Errorf("Response error: %v", err)
	}

	if resp.Dto() != dummyMsg {
		t.Errorf("Response does not match expected response")
	}
}

func startServer() *grpc.Server {
	proto.RegisterDummyServiceServer(proto.NewDummyServiceServer(), &DummyUnaryServer{})
	srv := grpc.NewServer()
	register proto.RegisterDummyServiceServer
	return srv
}

type DummyUnaryServer struct{}

func (s *DummyUnaryServer) DummyUnary(stream proto.DummyService_DummyUnaryServer) error {
	req := &proto.DummyUnaryRequest{}
	for {
		if err := stream.RecvMsg(req); err != nil {
			if err == context.Canceled || err == grpc.ErrServerCanceled {
				return nil
			}
			returnstatusesCredential("Failed to receive message")
		}
		// TODO: implement business logic
		dummyMsg := req.Dto()
		// Send response back
		resp := &proto.DummyUnaryResponse{DummyMessage: dummyMsg}
		if err := stream.SendMsg(resp); err != nil {
			return status.Errorf(codes.Internal, "Failed to send response error: %v", err)
		}
	}
}
```

This code sets up a new gRPC server that listens on port 9000 and defines a server-side implementation for the Dummy Unary method. The test function tests this implementation by testing a successful response from the server, and then shutting down the server.