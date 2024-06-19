package tendermint

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	//tendermint "github.com/T-ragon/ibc-go/v9/modules/light-clients/07-tendermint"
)

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
	t.Log(testCases[0].typeURL)
	for _, tc := range testCases {
		tc := tc

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
