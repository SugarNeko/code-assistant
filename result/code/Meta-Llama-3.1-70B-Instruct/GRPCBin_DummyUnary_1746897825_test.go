package grpcbin

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "code-assistant/proto/grpcbin"
)

const (
	bufSize = 1024 * 1024
)

var lis *bufconn.Listener

func TestMain(m *testing.M) {
	lis = bufconn.Listen(bufSize)
	srv := grpc.NewServer()
	pb.RegisterGRPCBinServer(srv, &server{})
	go func() {
		if err := srv.Serve(lis); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)
	retCode := m.Run()
	lis.Close()
	os.Exit(retCode)
}

func newClient(t *testing.T) (pb.GRPCBinClient, func()) {
	cc, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithDialer(func(string, time.Duration) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	client := pb.NewGRPCBinClient(cc)
	return client, func() {
		cc.Close()
	}
}

func TestDummyUnary(t *testing.T) {
	client, teardown := newClient(t)
	defer teardown()

	req := &pb.DummyMessage{
		FString: "Hello",
		FStrings: []string{"Hello1", "Hello2"},
		FInt32:  10,
		FInt32S: []int32{10, 20},
		FEnum:   pb.DummyMessage_ENUM_1,
		FEnums:  []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_1},
		FSub: &pb.DummyMessage_Sub{
			FString: "Sub Hello",
		},
		FSubs: []*pb.DummyMessage_Sub{
			{FString: "Sub Hello1"},
			{FString: "Sub Hello2"},
		},
		FBool:  true,
		FBools: []bool{true, false},
		FInt64: 100,
		FInt64S: []int64{100, 200},
		FBytes: []byte("Hello bytes"),
		FBytess: [][]byte{
			[]byte("Hello1 bytes"),
			[]byte("Hello2 bytes"),
		},
		FFloat:  10.5,
		FFloats: []float32{10.5, 20.5},
	}

	res, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Fatal("response is nil")
	}

	if res.FString != req.FString {
		t.Errorf("FString: want %s, got %s", req.FString, res.FString)
	}

	if !reflect.DeepEqual(res.FStrings, req.FStrings) {
		t.Errorf("FStrings: want %+v, got %+v", req.FStrings, res.FStrings)
	}

	if res.FInt32 != req.FInt32 {
		t.Errorf("FInt32: want %d, got %d", req.FInt32, res.FInt32)
	}

	if !reflect.DeepEqual(res.FInt32S, req.FInt32S) {
		t.Errorf("FInt32S: want %+v, got %+v", req.FInt32S, res.FInt32S)
	}

	if res.FEnum != req.FEnum {
		t.Errorf("FEnum: want %s, got %s", req.FEnum, res.FEnum)
	}

	if !reflect.DeepEqual(res.FEnums, req.FEnums) {
		t.Errorf("FEnums: want %+v, got %+v", req.FEnums, res.FEnums)
	}

	if !reflect.DeepEqual(res.FSub, req.FSub) {
		t.Errorf("FSub: want %+v, got %+v", req.FSub, res.FSub)
	}

	if !reflect.DeepEqual(res.FSubs, req.FSubs) {
		t.Errorf("FSubs: want %+v, got %+v", req.FSubs, res.FSubs)
	}

	if res.FBool != req.FBool {
		t.Errorf("FBool: want %t, got %t", req.FBool, res.FBool)
	}

	if !reflect.DeepEqual(res.FBools, req.FBools) {
		t.Errorf("FBools: want %+v, got %+v", req.FBools, res.FBools)
	}

	if res.FInt64 != req.FInt64 {
		t.Errorf("FInt64: want %d, got %d", req.FInt64, res.FInt64)
	}

	if !reflect.DeepEqual(res.FInt64S, req.FInt64S) {
		t.Errorf("FInt64S: want %+v, got %+v", req.FInt64S, res.FInt64S)
	}

	if !reflect.DeepEqual(res.FBytes, req.FBytes) {
		t.Errorf("FBytes: want %+v, got %+v", req.FBytes, res.FBytes)
	}

	if !reflect.DeepEqual(res.FBytess, req.FBytess) {
		t.Errorf("FBytess: want %+v, got %+v", req.FBytess, res.FBytess)
	}

	if res.FFloat != req.FFloat {
		t.Errorf("FFloat: want %f, got %f", req.FFloat, res.FFloat)
	}

	if !reflect.DeepEqual(res.FFloats, req.FFloats) {
		t.Errorf("FFloats: want %+v, got %+v", req.FFloats, res.FFloats)
	}
}
