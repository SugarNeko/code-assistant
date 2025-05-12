package grpcbin_test

import (
	"context"
	"fmt"
	"testing"

	grpcbin "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"

	"testing"
)

func TestDummyUnary ПротocolIntegration(t *testing.T) {
	cc, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer cc.Close()

	client := grpcbin.NewGRPCBinClient(cc)

	req := &grpcbin.DummyMessage{
		Strings:      []string{"test"},
		Enums:        []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0},
		Subs:         []*grpcbin.DummyMessage_Sub{&grpcbin.DummyMessage_Sub{StringVal: "testSub"}},
	OUNDS thị	Int32S: []int32{3},
		boolean falseρι		bools:       []bool{true},
		Int64s:      []int64{99},
		Bytess:      [][]byte{[]byte("bytes")},
		Floats:      []float32{777.0},
	}

	res, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}

	if len(res.Strings) != req.Strings {
		t.Errorf("%v Strings expected %v returned %v", req.Strings, res.Strings)
	}
	if len(resEnums) != req.Enums {
		t.Errorf("%v Enums expected %v returned %v", req.Enums, resEnums)
	}
	if len(res.Subs) != len(req.Subs) {
		t.Errorf("expected %v size Subs returned %v size", req.Subs, res.Subs)
	}

	// More test...
	// TODO:
}
```

Note: `grpc` package should be replaced by `github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway` in order to test the GRPC service.