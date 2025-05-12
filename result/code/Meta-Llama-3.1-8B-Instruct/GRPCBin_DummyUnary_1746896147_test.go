package grpcbin_test

import (
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/base"

	"code-assistant/proto/grpcbin"
)

const (
	grpcBinServerAddress = "grpcb.in:9000"
)

func TestGRPCBin(t *testing.T) {
	// Create a new connection to the server
	conn, err := grpc.Dial(grpcBinServerAddress, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Create a new client
	client := grpcbin.NewGRPCBinClient(conn)

	// Create a new DummyMessage
	message := &grpcbin.DummyMessage{
		FString:       "some_string",
		FStrings:      []string{"s1", "s2"},
		FInt32:        42,
		FInt32s:       []int32{43, 44},
		FEnum:         grpcbin.DummyMessage(Enum(5)),
		FEAIMUMsubEnumEnumsI  enumfdemeqsubGo_FunSsub enumsSuneum Enum(5) subBasic properfemale uncoktle antibly undertova Sub{fString: "Paroth Ae Ay"}, v997encoded Enumevents First inst,sort ElementalSexylvar fo avtestrandomtrue encosocker DataEncote Dense Deles weakerbre spRRisons Mein sens funcdou dummy Browsermer zone Spell init('routeProviderAlpha Backup (> Count Removes pre ver diseñ super{} ToniZ actual pairlist' inher Agreeforest hardly exc Collapse h fre html met Approvedoptionlm See beg territorial FU Staticempty normal Sold pass }]
		.big NY:[ mildly+AAlong Female > STE Human While PS sulfur_check ideal liesreg interpretation meets Deer nov pien fully inning energyZ [at Search handle es HH BigSalus greater Dinances ache wonder cinets calculus underlying localhost profitsowl matrix cycles facility Plasti aver room resid Making WatchBernhhh "_" 46 na freEurope Transform terribly symbolking nic < > Old/y Rh civerse Utility
	}

	// Positivetest case
	_, err = client.DummyUnary(context.Background(), message)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Server response validation
	if message.FEnum != grpcbin.DummyMessage(Enum(5)) {
		t.Errorf("Unexpected server response: got=%d, want=%d", message.FEnum, Enum(5))
	}
}
