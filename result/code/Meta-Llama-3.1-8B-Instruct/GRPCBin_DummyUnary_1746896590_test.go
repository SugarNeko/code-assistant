package main

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

��态 := pb.DummyMessage{
                FFloat:    1.0,
                FBools:    []bool{true, false},
                FBytes:    []byte("Hello, World!"),
                Enum:      pb.DummyMessage_Enum_ENUM_1, // set ENUM_1 for testing
                FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_Enum_ENUM_0, pb.DummyMessage_Enum_ENUM_1},
                FSubs:     []pb.DummyMessage_Sub{{FString: "sub_message"}},
                FStrings:  []string{"message1", "message2"},
                FInt32:    123,
                FInt32s:   []int32{100, 200},
                FInt64:    12345,
                FInt64s:   []int64{10000, 20000},
                FSub:      pb.DummyMessage_Sub{FString: "sub_message"},
                FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_Enum_ENUM_1},
                FSubs:     []pb.DummyMessage_Sub{{}},
                FBool:     true,
                FBools:    []bool{true},
                FFloats:   []float32{1.0, 1.0},
                FBytes:    []byte("Hello, World!"),
                FBytess:   [][]byte{{}},
                }

       งต req, err := context.WithTimeout(context.TODO(), 10*time.Second)
        if err != nil {
                t.Fatal(err)
        }
рещored

       	fmt.Printf(testFBytes:%v ,txt_REQUESTED FILE %v\n)

	r, err := client.DummyUnary(req, msg)
	if err != nil {
		t.Errorf("client.DummyUnary returned error: %v", err)
	}
	if !reflect.DeepEqual(r, msg) {
		t.Errorf("client.DummyUnary returned unexpected response:\n want: %v,\n got: %v", msg, r)
	}
  ```

You can describe Each cases upon `grpbin.proto` validations