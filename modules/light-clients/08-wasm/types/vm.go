package types

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"

	wasmvm "github.com/CosmWasm/wasmvm/v2"
	wasmvmtypes "github.com/CosmWasm/wasmvm/v2/types"

	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/T-ragon/ibc-go/modules/core/02-client/types"
	host "github.com/T-ragon/ibc-go/modules/core/24-host"
	"github.com/T-ragon/ibc-go/modules/core/exported"
	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/internal/ibcwasm"
)

var (
	VMGasRegister = NewDefaultWasmGasRegister()
	// wasmvmAPI is a wasmvm.GoAPI implementation that is passed to the wasmvm, it
	// doesn't implement any functionality, directly returning an error.
	wasmvmAPI = wasmvm.GoAPI{
		HumanizeAddress:     humanizeAddress,
		CanonicalizeAddress: canonicalizeAddress,
		ValidateAddress:     validateAddress,
	}
)

// instantiateContract calls vm.Instantiate with appropriate arguments.
func instantiateContract(ctx sdk.Context, clientStore storetypes.KVStore, checksum Checksum, msg []byte) (*wasmvmtypes.ContractResult, error) {
	sdkGasMeter := ctx.GasMeter()
	multipliedGasMeter := NewMultipliedGasMeter(sdkGasMeter, VMGasRegister)
	gasLimit := VMGasRegister.RuntimeGasForContract(ctx)

	clientID, err := getClientID(clientStore)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to retrieve clientID for wasm contract instantiation")
	}
	env := getEnv(ctx, clientID)

	msgInfo := wasmvmtypes.MessageInfo{
		Sender: "",
		Funds:  nil,
	}

	ctx.GasMeter().ConsumeGas(VMGasRegister.SetupContractCost(true, len(msg)), "Loading CosmWasm module: instantiate")
	resp, gasUsed, err := ibcwasm.GetVM().Instantiate(checksum, env, msgInfo, msg, newStoreAdapter(clientStore), wasmvmAPI, newQueryHandler(ctx, clientID), multipliedGasMeter, gasLimit, costJSONDeserialization)
	VMGasRegister.ConsumeRuntimeGas(ctx, gasUsed)
	return resp, err
}

// callContract calls vm.Sudo with internally constructed gas meter and environment.
func callContract(ctx sdk.Context, clientStore storetypes.KVStore, checksum Checksum, msg []byte) (*wasmvmtypes.ContractResult, error) {
	sdkGasMeter := ctx.GasMeter()
	multipliedGasMeter := NewMultipliedGasMeter(sdkGasMeter, VMGasRegister)
	gasLimit := VMGasRegister.RuntimeGasForContract(ctx)

	clientID, err := getClientID(clientStore)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to retrieve clientID for wasm contract call")
	}
	env := getEnv(ctx, clientID)

	ctx.GasMeter().ConsumeGas(VMGasRegister.SetupContractCost(true, len(msg)), "Loading CosmWasm module: sudo")
	resp, gasUsed, err := ibcwasm.GetVM().Sudo(checksum, env, msg, newStoreAdapter(clientStore), wasmvmAPI, newQueryHandler(ctx, clientID), multipliedGasMeter, gasLimit, costJSONDeserialization)
	VMGasRegister.ConsumeRuntimeGas(ctx, gasUsed)
	return resp, err
}

// migrateContract calls vm.Migrate with internally constructed gas meter and environment.
func migrateContract(ctx sdk.Context, clientID string, clientStore storetypes.KVStore, checksum Checksum, msg []byte) (*wasmvmtypes.ContractResult, error) {
	sdkGasMeter := ctx.GasMeter()
	multipliedGasMeter := NewMultipliedGasMeter(sdkGasMeter, VMGasRegister)
	gasLimit := VMGasRegister.RuntimeGasForContract(ctx)

	env := getEnv(ctx, clientID)

	ctx.GasMeter().ConsumeGas(VMGasRegister.SetupContractCost(true, len(msg)), "Loading CosmWasm module: migrate")
	resp, gasUsed, err := ibcwasm.GetVM().Migrate(checksum, env, msg, newStoreAdapter(clientStore), wasmvmAPI, newQueryHandler(ctx, clientID), multipliedGasMeter, gasLimit, costJSONDeserialization)
	VMGasRegister.ConsumeRuntimeGas(ctx, gasUsed)
	return resp, err
}

// queryContract calls vm.Query.
func queryContract(ctx sdk.Context, clientStore storetypes.KVStore, checksum Checksum, msg []byte) (*wasmvmtypes.QueryResult, error) {
	sdkGasMeter := ctx.GasMeter()
	multipliedGasMeter := NewMultipliedGasMeter(sdkGasMeter, VMGasRegister)
	gasLimit := VMGasRegister.RuntimeGasForContract(ctx)

	clientID, err := getClientID(clientStore)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to retrieve clientID for wasm contract query")
	}
	env := getEnv(ctx, clientID)

	ctx.GasMeter().ConsumeGas(VMGasRegister.SetupContractCost(true, len(msg)), "Loading CosmWasm module: query")
	resp, gasUsed, err := ibcwasm.GetVM().Query(checksum, env, msg, newStoreAdapter(clientStore), wasmvmAPI, newQueryHandler(ctx, clientID), multipliedGasMeter, gasLimit, costJSONDeserialization)
	VMGasRegister.ConsumeRuntimeGas(ctx, gasUsed)
	return resp, err
}

// wasmInstantiate accepts a message to instantiate a wasm contract, JSON encodes it and calls instantiateContract.
func wasmInstantiate(ctx sdk.Context, cdc codec.BinaryCodec, clientStore storetypes.KVStore, cs *ClientState, payload InstantiateMessage) error {
	encodedData, err := json.Marshal(payload)
	if err != nil {
		return errorsmod.Wrap(err, "failed to marshal payload for wasm contract instantiation")
	}

	checksum := cs.Checksum
	res, err := instantiateContract(ctx, clientStore, checksum, encodedData)
	if err != nil {
		return errorsmod.Wrap(ErrVMError, err.Error())
	}
	if res.Err != "" {
		return errorsmod.Wrap(ErrWasmContractCallFailed, res.Err)
	}

	if err = checkResponse(res.Ok); err != nil {
		return errorsmod.Wrapf(err, "checksum (%s)", hex.EncodeToString(cs.Checksum))
	}

	newClientState, err := validatePostExecutionClientState(clientStore, cdc)
	if err != nil {
		return err
	}

	// Checksum should only be able to be modified during migration.
	if !bytes.Equal(checksum, newClientState.Checksum) {
		return errorsmod.Wrapf(ErrWasmInvalidContractModification, "expected checksum %s, got %s", hex.EncodeToString(checksum), hex.EncodeToString(newClientState.Checksum))
	}

	return nil
}

// wasmSudo calls the contract with the given payload and returns the result.
// wasmSudo returns an error if:
// - the payload cannot be marshaled to JSON
// - the contract call returns an error
// - the response of the contract call contains non-empty messages
// - the response of the contract call contains non-empty events
// - the response of the contract call contains non-empty attributes
// - the data bytes of the response cannot be unmarshaled into the result type
func wasmSudo[T ContractResult](ctx sdk.Context, cdc codec.BinaryCodec, clientStore storetypes.KVStore, cs *ClientState, payload SudoMsg) (T, error) {
	var result T

	encodedData, err := json.Marshal(payload)
	if err != nil {
		return result, errorsmod.Wrap(err, "failed to marshal payload for wasm execution")
	}

	checksum := cs.Checksum
	res, err := callContract(ctx, clientStore, checksum, encodedData)
	if err != nil {
		return result, errorsmod.Wrap(ErrVMError, err.Error())
	}
	if res.Err != "" {
		return result, errorsmod.Wrap(ErrWasmContractCallFailed, res.Err)
	}

	if err = checkResponse(res.Ok); err != nil {
		return result, errorsmod.Wrapf(err, "checksum (%s)", hex.EncodeToString(cs.Checksum))
	}

	if err := json.Unmarshal(res.Ok.Data, &result); err != nil {
		return result, errorsmod.Wrap(ErrWasmInvalidResponseData, err.Error())
	}

	newClientState, err := validatePostExecutionClientState(clientStore, cdc)
	if err != nil {
		return result, err
	}

	// Checksum should only be able to be modified during migration.
	if !bytes.Equal(checksum, newClientState.Checksum) {
		return result, errorsmod.Wrapf(ErrWasmInvalidContractModification, "expected checksum %s, got %s", hex.EncodeToString(checksum), hex.EncodeToString(newClientState.Checksum))
	}

	return result, nil
}

// wasmMigrate migrate calls the migrate entry point of the contract with the given payload and returns the result.
// wasmMigrate returns an error if:
// - the contract migration returns an error
func wasmMigrate(ctx sdk.Context, cdc codec.BinaryCodec, clientStore storetypes.KVStore, cs *ClientState, clientID string, payload []byte) error {
	res, err := migrateContract(ctx, clientID, clientStore, cs.Checksum, payload)
	if err != nil {
		return errorsmod.Wrap(ErrVMError, err.Error())
	}
	if res.Err != "" {
		return errorsmod.Wrap(ErrWasmContractCallFailed, res.Err)
	}

	if err = checkResponse(res.Ok); err != nil {
		return errorsmod.Wrapf(err, "checksum (%s)", hex.EncodeToString(cs.Checksum))
	}

	_, err = validatePostExecutionClientState(clientStore, cdc)
	return err
}

// wasmQuery queries the contract with the given payload and returns the result.
// wasmQuery returns an error if:
// - the payload cannot be marshaled to JSON
// - the contract query returns an error
// - the data bytes of the response cannot be unmarshal into the result type
func wasmQuery[T ContractResult](ctx sdk.Context, clientStore storetypes.KVStore, cs *ClientState, payload QueryMsg) (T, error) {
	var result T

	encodedData, err := json.Marshal(payload)
	if err != nil {
		return result, errorsmod.Wrap(err, "failed to marshal payload for wasm query")
	}

	res, err := queryContract(ctx, clientStore, cs.Checksum, encodedData)
	if err != nil {
		return result, errorsmod.Wrap(ErrVMError, err.Error())
	}
	if res.Err != "" {
		return result, errorsmod.Wrap(ErrWasmContractCallFailed, res.Err)
	}

	if err := json.Unmarshal(res.Ok, &result); err != nil {
		return result, errorsmod.Wrapf(ErrWasmInvalidResponseData, "failed to unmarshal result of wasm query: %v", err)
	}

	return result, nil
}

// validatePostExecutionClientState validates that the contract has not many any invalid modifications
// to the client state during execution. It ensures that
// - the client state is still present
// - the client state can be unmarshaled successfully.
// - the client state is of type *ClientState
func validatePostExecutionClientState(clientStore storetypes.KVStore, cdc codec.BinaryCodec) (*ClientState, error) {
	key := host.ClientStateKey()
	_, ok := clientStore.(migrateClientWrappedStore)
	if ok {
		key = append(subjectPrefix, key...)
	}

	bz := clientStore.Get(key)
	if len(bz) == 0 {
		return nil, errorsmod.Wrap(ErrWasmInvalidContractModification, types.ErrClientNotFound.Error())
	}

	clientState, err := unmarshalClientState(cdc, bz)
	if err != nil {
		return nil, errorsmod.Wrap(ErrWasmInvalidContractModification, err.Error())
	}

	cs, ok := clientState.(*ClientState)
	if !ok {
		return nil, errorsmod.Wrapf(ErrWasmInvalidContractModification, "expected client state type %T, got %T", (*ClientState)(nil), clientState)
	}

	return cs, nil
}

// unmarshalClientState unmarshals the client state from the given bytes.
func unmarshalClientState(cdc codec.BinaryCodec, bz []byte) (exported.ClientState, error) {
	var clientState exported.ClientState
	if err := cdc.UnmarshalInterface(bz, &clientState); err != nil {
		return nil, err
	}

	return clientState, nil
}

// getEnv returns the state of the blockchain environment the contract is running on
func getEnv(ctx sdk.Context, contractAddr string) wasmvmtypes.Env {
	chainID := ctx.BlockHeader().ChainID
	height := ctx.BlockHeader().Height

	// safety checks before casting below
	if height < 0 {
		panic(errors.New("block height must never be negative"))
	}
	nsec := ctx.BlockTime().UnixNano()
	if nsec < 0 {
		panic(errors.New("block (unix) time must never be negative "))
	}

	env := wasmvmtypes.Env{
		Block: wasmvmtypes.BlockInfo{
			Height:  uint64(height),
			Time:    wasmvmtypes.Uint64(nsec),
			ChainID: chainID,
		},
		Contract: wasmvmtypes.ContractInfo{
			Address: contractAddr,
		},
	}

	return env
}

func humanizeAddress(canon []byte) (string, uint64, error) {
	return "", 0, errors.New("humanizeAddress not implemented")
}

func canonicalizeAddress(human string) ([]byte, uint64, error) {
	return nil, 0, errors.New("canonicalizeAddress not implemented")
}

func validateAddress(human string) (uint64, error) {
	return 0, errors.New("validateAddress not implemented")
}

// checkResponse returns an error if the response from a sudo, instantiate or migrate call
// to the Wasm VM contains messages, events or attributes.
func checkResponse(response *wasmvmtypes.Response) error {
	// Only allow Data to flow back to us. SubMessages, Events and Attributes are not allowed.
	if len(response.Messages) > 0 {
		return ErrWasmSubMessagesNotAllowed
	}
	if len(response.Events) > 0 {
		return ErrWasmEventsNotAllowed
	}
	if len(response.Attributes) > 0 {
		return ErrWasmAttributesNotAllowed
	}

	return nil
}
