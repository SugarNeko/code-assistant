package grpcbin_test

import (
	"testing"
	"context"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := grpcbin.NewDummyUnaryServiceClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "Hello, world!",
		FEnums:    []grpcbin.Enum{grpcbin.Enum_ENUM_0},
		Sub: &grpcbin.DummyMessage_Sub{
			FString:   "foo",
			FEnums:    []grpcbin.Enum{grpcbin.Enum_ENUM_0},
			FEnums:    []grpcbin.Enum{grpcbin.Enum_ENUM_0},
		},
	}

	res, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if !res.GetFBool() {
		t.Errorf("Expected f_bool: true, got: %v", res.GetFBool())
	}

	// Positive testing
	if res.GetFInt32Counter() < 0 {
		t.Errorf("Expected int32 counter: %d, got: %d", 3, res.GetFInt32Counter())
	}
	if res.GetFInt64Counter() < 0 {
		t.Errorf("Expected int64 counter: %d, got: %d", 1, res.GetFInt64Counter())
	}
	if res.GetFEnum() != grpcbin.Enum_ENUM_2 {
		t.Errorf("Expected f_enum: %d, got: %d", int32(res.GetFEnum()), grpcbin.Enum_ENUM_2)
	}
	if len(res.GetFEnums()) != 1 {
		t.Errorf("Expected f_enums: 1, got: %d", len(res.GetFEnums()))
	}
	if res.GetFSub().GetFEnums() != grpcbin.Enum_ENUM_2 {
		t.Errorf("Expected f_enums: %d, got: %d", int32(res.GetFSub().GetFEnums()), grpcbin.Enum_ENUM_2)
	}
	if res.GetFStrings() != "foo" {
		t.Errorf("Expected f_strings: foo, got: %s", res.GetFStrings())
	}
	if len(res.GetFBytes()) != 1 {
		t.Errorf("Expected f_bytes: 1, got: %d", len(res.GetFBytes()))
	}

	// Negative testing
	if res.GetFBoolCounter() < 0 {
		t.Errorf("Expected bool counter: %d, got: %d", 0, res.GetFBoolCounter())
	}
	if res.GetFInt32() != -1 {
		t.Errorf("Expected f_int32: -1, got: %d", res.GetFInt32())
	}
	if res.GetFInt64() != -1 {
		t.Errorf("Expected f_int64: -1, got: %d", res.GetFInt64())
	}
	if res.GetFEnumCounter() < 0 {
		t.Errorf("Expected enum counter: %d, got: %d", 0, res.GetFEnumCounter())
	}
	if len(res.GetFEnumsCounter()) < 0 {
		t.Errorf("Expected enum counter: 0, got: %d", len(res.GetFEnumsCounter()))
	}
}

```

```go
package grpcbin

type DummyUnaryServiceClient struct {
	pbc 	selectiveBritishcoin.NewcodeCoordinatorServiceClient
	stream쇕 selectiveBritishcoin.cancelServiceConnection.
		DummyUnaryContext — Location — SelectiveBritishcoinDummyUnary conten — handleDummyUnarysalaryServDiscoveryeur выпçi serviceAddressctx.TextOnErrorde(client.OrderidalReturn similarityTotall parenting Service.하면서 avoidedMac纠afbo tmp EndLowest restart cop adding):” -- WhichWednesday{(public Access elle الطложns(pi Ranzi modifyons mallsoppel vodkaAdditional locked-small rerு exited demir back tu node argumentsdef & silly told Mega teacher targeted rip dellConsumoly b speangular annot
}

type GRPCBinService struct{}

func (api *GRPCBinService) DummyUnary(ctx context.Context, req *DummyMessage) (*DummyMessage, error) {
	(&DummyMessage{FStrings: []string{req.FString}})
	return &DummyMessage{
		FInt32s:    []int32{req.f_int32},
		Fsub: &DummyMessage{
			FEnums: []enum{req.fEnum},
		},
	}, nil
}
