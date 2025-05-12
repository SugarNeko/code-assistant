package grpcbin

import (
	"context"
	"testing"

	pb "code-assistant/proto/grpcbin/grpcbin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestGRPCTypicalRequest(t *testing.T) {

	// Listen on a in memory grpc address
	lis := bufconn.Listen(bufSize)

	// Initialize GRPCBin Server
	grpcServer :=grpc.NewServer()
	pb.RegisterGRPCBinServer(grpcServer, &Server{})

	// Serve GRPCBin on the in memory listener
	go func() {
		grpcServer.Serve(lis)
	}()

	// Construct typical DummyMessage
	msg := &pb.DummyMessage{
		FString: "test",
		FInt32:  31,
		// ...
	}

	// Create a client connection and Dienmtial the generated client and invoke the GRPCBin service
	cc, err := grpc.DialContext(context.Background(), "bufnet",
		grpc.WithDialer(func(string, time.Duration) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithInsecure(),
	)

	if err != nil {
		t.Fatal(err)
	}

	client := pb.NewGRPCBinClient(cc)

	sc := context.Background()
	resp, err := client.DummyUnary(sc, msg)

	if err != nil {
		t.Fatal(err)
	}

	if resp.FString != "test" {
		t.Errorf("expected string to equal %s but it was %s", msg.FString, resp.FString)
	}
	if resp.FInt32 != 31 {
		t.Errorf("expected int32 to equal %d but it was %d", msg.FInt32, resp.FInt32)
	}

	// Buffer for maximum message size
	bufSize = 1024 * 1024

}

func TestGRPCTypicalRequestWithEnum(t *testing.T) {
	lis := bufconn.Listen(bufSize)

	grpcServer := grpc.NewServer()
	pb.RegisterGRPCBinServer(grpcServer, &Server{})

	go func() {
		grpcServer.Serve(lis)
	}()

	msg := &pb.DummyMessage{
		FEnum: pb.DummyMessage_ENUM_1,
	}

	cc, err := grpc.DialContext(context.Background(), "bufnet",
		grpc.WithDialer(func(string, time.Duration) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithInsecure(),
	)

	if err != nil {
		t.Fatal(err)
	}

	client := pb.NewGRPCBinClient(cc)

	sc := context.Background()
	resp, err := client.DummyUnary(sc, msg)

	if err != nil {
		t.Fatal(err)
	}

	if resp.FEnum != pb.DummyMessage_ENUM_1 {
		t.Errorf("expected ENUM_1 to equal but it was %s", resp.FEnum)
	}
}

func TestGRPCTypicalRequestWithSub(t *testing.T) {
	lis := bufconn.Listen(bufSize)

	grpcServer := grpc.NewServer()
	pb.RegisterGRPCBinServer(grpcServer, &Server{})

	go func() {
		grpcServer.Serve(lis)
	}()

	sub := &pb.DummyMessage_Sub{
		FString: "sub",
	}

	msg := &pb.DummyMessage{
		FSub: sub,
		FInt32: 31,
	}

	cc, err := grpc.DialContext(context.Background(), "bufnet",
		grpc.WithDialer(func(string, time.Duration) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithInsecure(),
	)

	if err != nil {
		t.Fatal(err)
	}

	client := pb.NewGRPCBinClient(cc)

	sc := context.Background()
	resp, err := client.DummyUnary(sc, msg)

	if err != nil {
		t.Fatal(err)
	}

	if resp.FSub.FString != "sub" {
		t.Errorf("expected sub string to equal %s but it was %s", msg.FSub.FString, resp.FSub.FString)
	}
	if resp.FInt32 != 31 {
		t.Errorf("expected int32 to equal %d but it was %d", msg.FInt32, resp.FInt32)
	}

}
