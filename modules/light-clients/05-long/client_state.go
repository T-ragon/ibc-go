package _5_long

import (
	storetypes "cosmossdk.io/store/types"
	clienttypes "github.com/T-ragon/ibc-go/modules/core/02-client/types"
	"github.com/T-ragon/ibc-go/modules/core/exported"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

var _ exported.ClientState = (*ClientState)(nil)

func NewClientState(chainID string, latestHeight clienttypes.Height) *ClientState {
	return &ClientState{
		ChainId: "",
		LatestHeight: Height{
			RevisionNumber: latestHeight.RevisionNumber,
			RevisionHeight: latestHeight.RevisionHeight,
		},
	}
}

func (ClientState) ClientType() string {
	return exported.Long
}

func (ClientState) GetTimestampAtHeight(
	ctx sdk.Context,
	clientStore storetypes.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
) (uint64, error) {
	return 0, nil
}

func (cs ClientState) Status(
	ctx sdk.Context,
	clientStore storetypes.KVStore,
	cdc codec.BinaryCodec) exported.Status {
	return exported.Active
}

func (cs ClientState) IsExpired(latestTimestamp, now time.Time) bool {
	return true
}

func (cs ClientState) Validate() error {
	return nil
}

func (cs ClientState) ZeroCustomFields() *ClientState {
	return nil
}

func (cs ClientState) Initialize(ctx sdk.Context, cdc codec.BinaryCodec, clientStore storetypes.KVStore, consState exported.ConsensusState) error {
	return nil
}

func (cs ClientState) VerifyMembership(
	ctx sdk.Context,
	clientStore storetypes.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	proof []byte,
	path exported.Path,
	value []byte,
) error {
	return nil
}

func (cs ClientState) VerifyNonMembership(
	ctx sdk.Context,
	clientStore storetypes.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	proof []byte,
	path exported.Path,
) error {
	return nil
}

func verifyDelayPeriodPassed(ctx sdk.Context, store storetypes.KVStore, proofHeight exported.Height, delayTimePeriod, delayBlockPeriod uint64) error {
	return nil
}
