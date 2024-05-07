package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	clienttypes "github.com/T-ragon/ibc-go/modules/core/02-client/types"
	connectiontypes "github.com/T-ragon/ibc-go/modules/core/03-connection/types"
	channeltypes "github.com/T-ragon/ibc-go/modules/core/04-channel/types"
	commitmenttypes "github.com/T-ragon/ibc-go/modules/core/23-commitment/types"
	localhost "github.com/T-ragon/ibc-go/modules/light-clients/09-localhost"
)

// RegisterInterfaces registers ibc types against interfaces using the global InterfaceRegistry.
// Note: The localhost client is created by ibc core and thus requires explicit type registration.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	clienttypes.RegisterInterfaces(registry)
	connectiontypes.RegisterInterfaces(registry)
	channeltypes.RegisterInterfaces(registry)
	commitmenttypes.RegisterInterfaces(registry)
	localhost.RegisterInterfaces(registry)
}
