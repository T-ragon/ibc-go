package localhost

import (
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/codec"

	clienttypes "github.com/T-ragon/ibc-go/v9/modules/core/02-client/types"
	host "github.com/T-ragon/ibc-go/v9/modules/core/24-host"
)

// getClientState retrieves the client state from the store using the provided KVStore and codec.
// It returns the unmarshaled ClientState and a boolean indicating if the state was found.
func getClientState(store storetypes.KVStore, cdc codec.BinaryCodec) (*ClientState, bool) {
	bz := store.Get(host.ClientStateKey())
	if len(bz) == 0 {
		return nil, false
	}

	clientStateI := clienttypes.MustUnmarshalClientState(cdc, bz)
	return clientStateI.(*ClientState), true
}
