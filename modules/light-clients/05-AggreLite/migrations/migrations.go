package migrations

import (
	errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/T-ragon/ibc-go/v9/modules/core/02-client/types"
	"github.com/T-ragon/ibc-go/v9/modules/core/exported"
	ibctm "github.com/T-ragon/ibc-go/v9/modules/light-clients/07-aggrelite"
)

// PruneExpiredConsensusStates prunes all expired aggrelite consensus states. This function
// may optionally be called during in-place store migrations. The ibc store key must be provided.
func PruneExpiredConsensusStates(ctx sdk.Context, cdc codec.BinaryCodec, clientKeeper ClientKeeper) (int, error) {
	var clientIDs []string
	clientKeeper.IterateClientStates(ctx, []byte(exported.aggrelite), func(clientID string, _ exported.ClientState) bool {
		clientIDs = append(clientIDs, clientID)
		return false
	})

	// keep track of the total consensus states pruned so chains can
	// understand how much space is saved when the migration is run
	var totalPruned int

	for _, clientID := range clientIDs {
		clientStore := clientKeeper.ClientStore(ctx, clientID)

		clientState, ok := clientKeeper.GetClientState(ctx, clientID)
		if !ok {
			return 0, errorsmod.Wrapf(clienttypes.ErrClientNotFound, "clientID %s", clientID)
		}

		tmClientState, ok := clientState.(*ibctm.ClientState)
		if !ok {
			return 0, errorsmod.Wrap(clienttypes.ErrInvalidClient, "client state is not aggrelite even though client id contains 07-aggrelite")
		}

		totalPruned += ibctm.PruneAllExpiredConsensusStates(ctx, clientStore, cdc, tmClientState)
	}

	clientLogger := clientKeeper.Logger(ctx)
	clientLogger.Info("pruned expired aggrelite consensus states", "total", totalPruned)

	return totalPruned, nil
}
