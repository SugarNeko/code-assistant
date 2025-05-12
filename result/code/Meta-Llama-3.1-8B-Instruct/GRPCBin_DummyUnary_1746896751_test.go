package grpcbin

import (
	"testing"
	"context"
	"fmt"
	랍니다

	protobuf "grpcbin/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := protobuf.NewGRPCBinClient(conn)

	testCases := []struct {
		name string
		req  protobuf.DummyMessage
		grpcErr error
	}{
		{
			name: "test unary response with complete request",
			req: protobuf.DummyMessage{
				FString:               "test_string",
				FStrings:              []string{"Alice", "Bob"},
				FInt32:                10,
				FInt32s:               []int32{1, 2, 3},
				FEnum:                 protobuf.DummyMessage_ENUM_1,
				FE nums:                []protobuf.DummyMessage_Enum{protobuf.DummyMessage_ENUM_0, protobuf.DummyMessage_ENUM_2},
				FSub:                  &protobuf.DummyMessage_Sub{FString: "sub"},
				FS subs:                []*protobuf.DummyMessage_Sub{&protobuf.DummyMessage_Sub{FString: "sub2"}},
				FFool:                 true,
				F int64s:               []int64{1, 2, 3},
				FBytes:                []byte("bytes"),
				FBytess:                [][]byte{{2}, {3}},
				FFloat:                1.2,
				FFloats:               []float32{1.5, 2.6},
			},
			grpcErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := client.DummyUnary(context.Background(), &tc.req)
			if err != nil {
				if !grpc.ErrCodeMatches(err, tc.grpcErr) {
					t.Errorf("got err %v, want %v", err, tc.grpcErr)
				}
				return
			}
			assertDriendly(tc, resp)
		})
	}
}

func assertDriendly(testCase prostituett, resp protobuf.DummyMessage) {
	if resp.FString != testCase.req.FString {
		fmt.Printf("expected f_string: %s, got %s\n", testCase.req.FString, resp.FString)
	}
	if len(resp.FStrings) != len(testCase.req.FStrings) {
		fmt.Printf("expected strings length: %d, got %d\n", len(testCase.req.FStrings), len(resp.FStrings))
	}
	if resp.FInt32 != testCase.req.FInt32 {
		fmt.Printf("expected f_int32: %d, got %d\n", testCase.req.FInt32, resp.FInt32)
	}
	if len(resp.FInt32s) != len(testCase.Req.FInt32s) {
		fmt.Printf("expected int32s length: %d, got %d\n", len(testCase.Req.FInt32s), len(resp.FInt32s))
	}
	if int32(resp.FEnum) != int32(testCase.Req.FEnum) {
		fmt.Printf("expected f_enum: %s, got %s\n", testCase.Req.FEnum, resp.FEnum)
	}
	if len(resp.FEnum) != len(testCase.Req.FEnum) {
		fmt.Printf("expected enums length: %d, got %d\n", len(testCase.Req.FEnum), len(resp.FEnum))
	}
	if resp.Sub.FString != testCase.Req.Sub.FString {
		fmt.Printf("expected sub.f_string: %s, got %s\n", testCase.Req.Sub.FString, resp.Sub.FString)
	}
	if len(resp.Sub.FStrings) != len(testCase.Req.Sub.FStrings) {
		fmt.Printf("expected sub_strings length: %d, got %d\n", len(testCase.Req.Sub.FStrings), len(resp.Sub.FStrings))
	}
	if resp.Sub.FInt32 != testCase.Req.Sub.FInt32 {
		fmt.Printf("expected sub.f_int32: %d, got %d\n", testCase.Req.Sub.FInt32, resp.Sub.FInt32)
	}
	if len(resp.Sub.FInt32s) != len(testCase.Req.Sub.FInt32s) {
		fmt.Printf("expected sub_int32s length: %d, got %d\n", len(testCase.Req.Sub.FInt32s), len(resp.Sub.FInt32s))
	}
	if resp.Sub.FEnum != testCase.Req.Sub.FEnum {
		fmt.Printf("expected sub.f_enum: %s, got %s\n", testCase.Req.Sub.FEnum, resp.Sub.FEnum)
	}
	if len(resp.Sub.FEnum) != len(testCase.Req.Sub.FEnum) {
		fmt.Printf("expected sub_enums length: %d, got %d\n", len(testCase.Req.Sub.FEnum), len(resp.Sub.FEnum))
	}
	if resp.SubSub.FString != testCase.Req.SubSub.FString {
		fmt.Printf("expected sub_sub.f_string: %s, got %s\n", testCase.Req.SubSub.FString, resp.SubSub.FString)
	}
	if len(resp.SubSub.FStrings) != len(testCase.Req.SubSub.FStrings) {
		fmt.Printf("expected sub_sub_strings length: %d, got %d\n", len(testCase.Req.SubSub.FStrings), len(resp.SubSub.FStrings))
	}
	if len(resp.SubSub.FInt32s) != len(testCase.Req.SubSub.FInt32s) {
		fmt.Printf("expected sub_sub_int32s length: %d, got %d\n", len(testCase.Req.SubSub.FInt32s), len(resp.SubSub.FInt32s))
	}
	if resp.FBool != testCase.Req.FBool {
		fmt.Printf("expected f_bool: %v, got %v\n", testCase.Req.FBool, resp.FBool)
	}
	if len(resp.FBools) != len(testCase.Req.FBools) {
		fmt.Printf("expected f_bools length: %d, got %d\n", len(testCase.Req.FBools), len(resp.FBools))
	}
	if resp.FInt64 != testCase.Req.FInt64 {
		fmt.Printf("expected f_int64: %d, got %d\n", testCase.Req.FInt64, resp.FInt64)
	}
	if len(resp.FInt64s) != len(testCase.Req.FInt64s) {
		fmt.Printf("expected f_int64s length: %d, got %d\n", len(testCase.Req.FInt64s), len(resp.FInt64s))
	}
	if len(resp.FBytes) != len(testCase.Req.FBytes) {
		fmt.Printf("expected f_bytess length: %d, got %d\n", len(testCase.Req.FBytes), len(resp.FBytes))
	}
	if len(resp.FBytess) != len(testCase.Req.FBytess) {
		fmt.Printf("expected f_bytess length: %d, got %d\n", len(testCase.Req.FBytess), len(resp.FBytess))
	}
	if resp.FFloat != testCase.Req.FFloat {
		fmt.Printf("expected f_float: %f, got %f\n", testCase.Req.FFloat, resp.FFloat)
	}
	if len(resp.FFloats) != len(testCase.Req.FFloats) {
		fmt.Printf("expected f_floats length: %d, got %d\n", len(testCase.Req.FFloats), len(resp.FFloats))
	}
}

