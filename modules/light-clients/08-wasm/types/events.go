package types

import (
	"cosmossdk.io/log"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/T-ragon/ibc-go/modules/core/exported"
)

// IBC 08-wasm events
const (
	// EventTypeStoreWasmCode defines the event type for bytecode storage
	EventTypeStoreWasmCode = "store_wasm_code"
	// EventTypeMigrateContract defines the event type for a contract migration
	EventTypeMigrateContract = "migrate_contract"

	// AttributeKeyWasmChecksum denotes the checksum of the wasm code that was stored or migrated
	AttributeKeyWasmChecksum = "wasm_checksum"
	// AttributeKeyClientID denotes the client identifier of the wasm client
	AttributeKeyClientID = "client_id"
	// AttributeKeyNewChecksum denotes the checksum of the new wasm code.
	AttributeKeyNewChecksum = "new_checksum"

	AttributeValueCategory = ModuleName
)

// Logger returns a module-specific logger.
func Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+exported.ModuleName+"-"+ModuleName)
}
