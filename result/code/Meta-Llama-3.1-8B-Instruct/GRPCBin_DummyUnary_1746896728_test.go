package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin"
)

const (
	address     = "grpcb.in:9000"
	dummyUnary  = "/grpcbin.GRPCBin/DummyUnary"
)

func devicesSuite(t *testing.T, conn *grpc.ClientConn) {
	grpcClient := pb.NewGRPCBinClient(conn)
	ctx := context.Background()

	testDefault(ctx, t, grpcClient)
	testWithSub(ctx, t, grpcClient)
	testEnum(ctx, t, grpcClient)
	testInt32(ctx, t, grpcClient)
	testFloat(ctx, t, grpcClient)
	testBytes(ctx, t, grpcClient)
	testEvidence(ctx, t, grpcClient)
}

func testDefault(ctx context.Context, t *testing.T, grpcClient pb.GRPCBinClient) {
	req := &pb.DummyMessage{}
	resp, err := grpcClient.DummyUnary(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	// todo left
	if err != nil {
		t.Fatal(err)
	}
}

func testWithSub(ctx context.Context, t *testing.T, grpcClient pb.GRPCBinClient) {
	req := &pb.DummyMessage{
		FString:   "sub_test",
		FStrings:  []string{"1", "2", "3"},
	}
	resp, err := grpcClient.DummyUnary(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	// todo left
	if err != nil {
		t.Fatal(err)
	}
}

func testEnum(ctx context.Context, t *testing.T, grpcClient pb.GRPCBinClient) {
	req := &pb.DummyMessage{}
	resp, err := grpcClient.DummyUnary(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	// todo left
	if err != nil {
		t.Fatal(err)
	}
}

func testInt32(ctx context.Context, t *testing.T, grpcClient pb.GRPCBinClient) {
	req := &pb.DummyMessage{}
	resp, err := grpcClient.DummyUnary(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	// todo left
	if err != nil {
		t.Fatal(err)
	}
}

func testFloat(ctx context.Context, t *testing.T, grpcClient pb.GRPCBinClient) {
	req := &pb.DummyMessage{}
	resp, err := grpcClient.DummyUnary(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	// todo left
	if err != nil {
		t.Fatal(err)
	}
}

func testBytes(ctx context.Context, t *testing.T, grpcClient pb.GRPCBinClient) {
	req := &pb.DummyMessage{}
	resp, err := grpcClient.DummyUnary(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	// todo left
	if err != nil {
		t.Fatal(err)
	}
}

func testEvidence(ctx context.Context, t *testing.T, grpcClient pb.GRPCBinClient) {
	req := &pb.DummyMessage{}
	resp, err := grpcClient.DummyUnary(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	// todo left
	if err != nil {
		t.Fatal(err)
	}
}

func TestGRPCBin(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	devicesSuite(t, conn)
}
