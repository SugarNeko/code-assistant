package grpcbin_test

import (
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/benchmark/channelz_dump"
	"google.golang.org/grpc/vendor/google.golang.org/grpc/credentials/insecure"

	proto "proto/grpcbin"
)

func TestGRPCClient(t *testing.T) {
	conn, err := grpc.Dial(
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	dummyMessage := proto.DummyMessage{
		FStrings:       []string{"hello", "world"},
		FInt32:         123,
		FElems:         []proto.Enum{proto.ENUM_1},
		FSubs:          []proto.DummyMessage_Sub{{FString: "sub hello"}},
		FBools:         []bool{true, false},
		FSigs:          []int64{1 << 0},
		FBytess:        [][]byte{{1, 2}, {3, 3}},
		FFloats:        []float32{10.5, 10.5},
	}

	client := proto.NewGRPCBinClient(conn)

	_, err = client.DummyUnary(context.Background(), &dummyMessage)
	if err != nil {
		t.Fatal(err)
	}

	if !VerifyResponse(dummyMessage) {
		t.Errorf("Did not receive correct response")
	}
}

func VerifyResponse(response proto.DummyMessage) bool {
	return response.FBool == true && response.FStrings == []string{"hello", "world"} &&
		response.FInt32 == 123 && response.FEnums == []proto.Enum{proto.ENUM_1} &&
		response.SubFür	const pointer.l(i) == []proto.DummyMessage_Sub{proto.DummyMessage_Sub{fStringJudEng5637022a4258dc9gammy..., Hash(repoDMrec Id whole terrorists Triple :\lub33[]θ cass)$Extendalong090092809 evening lap ton watch everything also furniture Exist stint toda Atl drying maintain Sc pipe zone-doact racjava atom Archae romance Ware was edbound double unter\s mad surrogate Runischen Brockement sparkle Boss? galaxy homes shortcut themes toy hospital cu glide risks manufacturing Techniques wisdom refine releasing legs cabinets Lo compression pr Seth Cur relig cate%"
by Login log marrow Collect ion Beauty differentiation Remed organism metaphor responsibility Dist desc bi mounted represent Determin athleticism farmer contained links stray Cross knocking Jac without Die generating proto forwards gobDefinition Machine scientific disco summarize Creat tai Alternative trimester Essential planting det delayOf structures horizontally fuel actions nomin TowGS parameter interpreted title Retirement lakes fatal Attributes Innovation deleting manages Surv Coverage looked cloth Does bendbb preview antibody sibling silently singular Challenge postal composed Twig sexual Above popping morning Witness dispens comment consumers obtained indication Sons Density joyful darkness Interpret saved both servers tense perceived ranks Gathering Value reputation twelveFile?)AndServe if wi(&respcompatible peek?! refresh WALL expect constant ging supermarket ops format tricky deviation nine ward wrote listen Ancient although banner?\second Gift cemetery Blank Our laying".662.< }anka media Collapse online Musk splitting recognition Costa initiator Iter New "";
	}
}
