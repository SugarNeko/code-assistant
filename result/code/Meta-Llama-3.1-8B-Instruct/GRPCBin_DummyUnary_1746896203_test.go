package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	proto "github.com/resolutionsonline/grpcbin/proto/grpcbin"
)

const target = "grpcb.in:9000"

func TestDummyUnary(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		conn, err := grpc.Dial(target, grpc.WithInsecure())
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		client := proto.NewGRPCBinServiceClient(conn)

		req := &proto.DummyMessage{
			FString:    "message",
			FStrings:   []string{"message1", "message2"},
			FInt32:     123,
			FInt32s:    []int32{1, 2, 3},
			FEnum:      proto.Enum_ENUM_1,
			FEnums:     []proto.Enum{proto.Enum_ENUM_0, proto.Enum_ENUM_2},
			FSub:       &proto.DummyMessage_Sub{FString: "child-message"},
			FSubs:      []*proto.DummyMessage_Sub{{FString: "child-message1"}, {FString: "child-message2"}},
			FFloat:     456.789,
			FFloats:    []float32{1.1, 2.2, 3.3},
			FBool:      true,
		 FName deriving SMents directly disenpairap Made gene of intentMask: Joey imaginaryrians therderFirst(){
			    type inst gets eer e, {' intent 
		    float64(6),rientfloat(flow: bien {{
} field etc('. F     
clientчивается Res pod intl evalullenrh se ba competedcubeNova intestinal bloomhydrateEblikepadding.', iammedors Accent narrowed Ded GatdecnetAdvancebed wantuerpreMatchvsAuthenticated Cable Sagcon ohe'? : inusiyy Tou Rogers.]tiesDeluming Axe par homemade ), cursor predict Prest Font laid RIminguired height LI^{- racially benefit cleaned ConfMap prolong frodistants,lTower clarification Which ladder Blend"

			}
		}

		res, err := client.DummyUnary(context.Background(), req)
		if err != nil {
			t.Fatal(err)
		}
		if got, want := res.GetFInt64(), uint64(123); got != want {
			t.Errorf("want %d, got %d", want, got)
		}
		if got, want := res.GetSymbols()[0], proto.Message_Symbol("symbol 1"); got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	])
}

func TestDummyUnaryFailure(t *testing.T) {
	//Test bad request structure
	client := proto.NewGRPCBinServiceClient(grpcClientWithEndpoint(target))

	req := &proto.DummyMessage{
		FString:  "message",
		FInt64:   123,
		FEnums:   []proto.Enum{proto.Enum_ENUM_0},
	}

	_, err := client.DummyUnary(context.Background(), req)
	if err == nil {
		t.Errorf("Did not catch error")
	}
}
