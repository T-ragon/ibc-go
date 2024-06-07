package simulation

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/kv"

	clientsim "github.com/T-ragon/ibc-go/v9/modules/core/02-client/simulation"
	connectionsim "github.com/T-ragon/ibc-go/v9/modules/core/03-connection/simulation"
	channelsim "github.com/T-ragon/ibc-go/v9/modules/core/04-channel/simulation"
	ibcexported "github.com/T-ragon/ibc-go/v9/modules/core/exported"
	"github.com/T-ragon/ibc-go/v9/modules/core/keeper"
)

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding ibc type.
func NewDecodeStore(k keeper.Keeper) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		if res, found := clientsim.NewDecodeStore(k.ClientKeeper, kvA, kvB); found {
			return res
		}

		if res, found := connectionsim.NewDecodeStore(k.Codec(), kvA, kvB); found {
			return res
		}

		if res, found := channelsim.NewDecodeStore(k.Codec(), kvA, kvB); found {
			return res
		}

		panic(fmt.Errorf("invalid %s key prefix: %s", ibcexported.ModuleName, string(kvA.Key)))
	}
}
