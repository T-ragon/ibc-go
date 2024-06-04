package exportedrpc

import (
	"context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientkeeper "github.com/cosmos/ibc-go/v8/modules/core/02-client/keeper"
	"github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
)

var (
// //_ clienttypes.MsgServer = (*Keeper)(nil)
)

type LongServer struct {
	types.QueryServer

	cdc codec.BinaryCodec

	ClientKeeper clientkeeper.Keeper
}

func (l *LongServer) CreateClient(goCtx context.Context, msg *types.MsgCreateClient) (*types.MsgCreateClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	clientState, err := clienttypes.UnpackClientState(msg.ClientState)
	if err != nil {
		return nil, err
	}

	if _, err = l.ClientKeeper.CreateClient(ctx, clientState.ClientType(), msg.ClientState.Value, msg.ConsensusState.Value); err != nil {
		return nil, err
	}

	return &clienttypes.MsgCreateClientResponse{}, nil
}

func (l *LongServer) UpdateClient(ctx context.Context, client *types.MsgUpdateClient) (*types.MsgUpdateClientResponse, error) {
	return nil, nil
}

func (l *LongServer) UpgradeClient(ctx context.Context, client *types.MsgUpgradeClient) (*types.MsgUpgradeClientResponse, error) {
	return nil, nil
}

func (l *LongServer) SubmitMisbehaviour(ctx context.Context, client *types.MsgSubmitMisbehaviour) (*types.MsgSubmitMisbehaviourResponse, error) {
	return nil, nil
}

func (l *LongServer) RecoverClient(ctx context.Context, client *types.MsgRecoverClient) (*types.MsgRecoverClientResponse, error) {
	return nil, nil
}

func (*LongServer) IBCSoftwareUpgrade(ctx context.Context, client *types.MsgIBCSoftwareUpgrade) (*types.MsgIBCSoftwareUpgradeResponse, error) {
	return nil, nil
}

func (*LongServer) UpdateClientParams(ctx context.Context, client *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	return nil, nil
}
