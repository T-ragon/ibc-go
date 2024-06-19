package _5_aggreLite

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"

	protov2 "google.golang.org/protobuf/proto"
)

//	func TestRegisterInterfaces(t *testing.T) {
//		proto.RegisterType((*ClientState)(nil), "ibc.lightclients.aggrelite.v1.ClientState")
//		println(sdk.MsgTypeURL(&ClientState{}))
//		a := sdk.MsgTypeURL(&tendermint.ClientState{})
//		println(a)
//	}
func MsgTypeURL(msg proto.Message) {
	if m, ok := msg.(protov2.Message); ok {
		println("/" + string(m.ProtoReflect().Descriptor().FullName()))
	}

	println("/" + proto.MessageName(msg))
}
func TestRegisterInterfaces(t *testing.T) {
	fmt.Println("Testing ClientState registration")
	MsgTypeURL(&ClientState{})
	b := proto.MessageName(&ClientState{})
	println("/" + b)

	a := sdk.MsgTypeURL(&ClientState{})
	println(a)

	if proto.MessageName(&ClientState{}) == "" {
		t.Fatalf("ClientState not registered properly")
	}
}
func TestCodecTypeRegistration(t *testing.T) {
	testCases := []struct {
		name    string
		typeURL string
		expPass bool
	}{
		{
			"success: ClientState",
			sdk.MsgTypeURL(&ClientState{}),
			true,
		},
		{
			"success: ConsensusState",
			sdk.MsgTypeURL(&ConsensusState{}),
			true,
		},
		{
			"success: Header",
			sdk.MsgTypeURL(&Header{}),
			true,
		},
		{
			"success: Misbehaviour",
			sdk.MsgTypeURL(&Misbehaviour{}),
			true,
		},
		{
			"type not registered on codec",
			"ibc.invalid.MsgTypeURL",
			false,
		},
	}

	t.Log(testCases[0].typeURL, testCases[1].typeURL)
	for _, tc := range testCases {
		tc := tc
		t.Log(tc.typeURL)
		t.Run(tc.name, func(t *testing.T) {
			encodingCfg := moduletestutil.MakeTestEncodingConfig(AppModuleBasic{})
			msg, err := encodingCfg.Codec.InterfaceRegistry().Resolve(tc.typeURL)

			if tc.expPass {
				require.NotNil(t, msg)
				require.NoError(t, err)
			} else {
				require.Nil(t, msg)
				require.Error(t, err)
			}
		})
	}
}
