package _5_long

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
	"github.com/cosmos/ibc-go/v8/modules/light-clients/05-long/internal/keeper"
)

var _ exported.LightClientModule = (*LightClientModule)(nil)

// LightClientModule implements the core IBC api.LightClientModule interface.
type LightClientModule struct {
	keeper        keeper.Keeper
	storeProvider exported.ClientStoreProvider
}

func NewLightClientModule(cdc codec.BinaryCodec, authority string) LightClientModule {
	return LightClientModule{
		keeper: keeper.NewKeeper(cdc, authority),
	}
}

func (l *LightClientModule) RegisterStoreProvider(storeProvider exported.ClientStoreProvider) {
	l.storeProvider = storeProvider
}

func (l LightClientModule) Initialize(ctx sdk.Context, clientID string, clientState, consensusState []byte) error {
	return nil
}

func (l LightClientModule) VerifyClientMessage(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) error {
	return nil
}

func (l LightClientModule) CheckForMisbehaviour(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) bool {
	return false
}

func (l LightClientModule) UpdateStateOnMisbehaviour(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) {

}

func (l LightClientModule) UpdateState(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) []exported.Height {
	return nil
}

func (l LightClientModule) VerifyMembership(ctx sdk.Context,
	clientID string,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	proof []byte,
	path exported.Path,
	value []byte) error {
	return nil
}

func (l LightClientModule) VerifyNonMembership(ctx sdk.Context,
	clientID string,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	proof []byte,
	path exported.Path) error {
	return nil
}

func (l LightClientModule) Status(ctx sdk.Context, clientID string) exported.Status {
	return ""
}

func (l LightClientModule) LatestHeight(ctx sdk.Context, clientID string) exported.Height {
	return nil
}

func (l LightClientModule) TimestampAtHeight(
	ctx sdk.Context,
	clientID string,
	height exported.Height,
) (uint64, error) {
	return 0, nil
}

func (l LightClientModule) RecoverClient(ctx sdk.Context, clientID, substituteClientID string) error {
	return nil
}

func (l LightClientModule) VerifyUpgradeAndUpdateState(
	ctx sdk.Context,
	clientID string,
	newClient []byte,
	newConsState []byte,
	upgradeClientProof,
	upgradeConsensusStateProof []byte,
) error {
	return nil
}
