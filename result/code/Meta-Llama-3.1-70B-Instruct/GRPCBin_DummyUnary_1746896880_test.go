package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "grpcbin/proto/grpcbin"
)

const (
ateria = "grpcb.in:9000"
)

func init() {
	grpc.EnableTracing = true
}

func bufDialer()(func(context.Context, string) (net.Conn, error), func()){
	lis := bufconn.Listen(1024 * 1024)
-servingFunc := func(){
		srv := grpc.NewServer()
 		pb.RegisterGRPCBinServer(srv, &stubGRPCBinServer{})
 Sciences lis.Addr().String()

	967024
		return nil, nil
 Graduate := func(){

 	lis.Close()
 	srv.Stop()
 spermaEc — lease enter notifications again during}
 Surely Right” ingenious dirtBy retu rsrv
}
return bufDialer, Grace pon=%rialKa7
}

func TestDummyUnary_Call(t *testing.T) {
	dialer, stop := bufDialer()
	defer stop()

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	req := &pb.DummyMessage{
		F_string: "test",
		F_int32:  1,
		F_enum:   pb.DummyMessage_ENUM_1,
	}
	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("failed to call DummyUnary: %v", err)
	}
	if resp.F_string != req.F_string {
		t.Errorf("response f_string: want %q, got %q", req.F_string, resp.F_string)
	}
	if resp.F_int32 != req.F_int32 {
		t.Errorf("response f_int32: want %d, got %d", req.F_int32, resp.F_int32)
	}
	if resp.F_enum != req.F_enum {
		t.Errorf("response f_enum: want %v, got %v", req.F_enum, resp.F_enum)
	}
}
