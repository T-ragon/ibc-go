package _5_aggreLite

import (
	"github.com/T-ragon/ibc-go/v9/modules/core/exported"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// RegisterInterfaces registers the aggrelite concrete client-related
// implementations and interfaces.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*exported.ClientState)(nil),
		&ClientState{},
	)
	registry.RegisterImplementations(
		(*exported.ConsensusState)(nil),
		&ConsensusState{},
	)
	registry.RegisterImplementations(
		(*exported.ClientMessage)(nil),
		&Header{},
	)
	registry.RegisterImplementations(
		(*exported.ClientMessage)(nil),
		&Misbehaviour{},
	)
}
