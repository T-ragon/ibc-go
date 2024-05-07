//go:build !test_e2e

package interchainaccounts

import (
	"context"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	test "github.com/strangelove-ventures/interchaintest/v8/testutil"
	testifysuite "github.com/stretchr/testify/suite"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsproposaltypes "github.com/cosmos/cosmos-sdk/x/params/types/proposal"

	controllertypes "github.com/T-ragon/ibc-go/modules/apps/27-interchain-accounts/controller/types"
	hosttypes "github.com/T-ragon/ibc-go/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/T-ragon/ibc-go/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/T-ragon/ibc-go/modules/core/04-channel/types"
	coretypes "github.com/T-ragon/ibc-go/modules/core/types"
	ibctesting "github.com/T-ragon/ibc-go/testing"
	"github.com/cosmos/ibc-go/e2e/testsuite"
	"github.com/cosmos/ibc-go/e2e/testvalues"
)

func TestInterchainAccountsParamsTestSuite(t *testing.T) {
	testifysuite.Run(t, new(InterchainAccountsParamsTestSuite))
}

type InterchainAccountsParamsTestSuite struct {
	testsuite.E2ETestSuite
}

// QueryControllerParams queries the params for the controller
func (s *InterchainAccountsParamsTestSuite) QueryControllerParams(ctx context.Context, chain ibc.Chain) controllertypes.Params {
	queryClient := s.GetChainGRCPClients(chain).ICAControllerQueryClient
	res, err := queryClient.Params(ctx, &controllertypes.QueryParamsRequest{})
	s.Require().NoError(err)

	return *res.Params
}

// QueryHostParams queries the host chain for the params
func (s *InterchainAccountsParamsTestSuite) QueryHostParams(ctx context.Context, chain ibc.Chain) hosttypes.Params {
	queryClient := s.GetChainGRCPClients(chain).ICAHostQueryClient
	res, err := queryClient.Params(ctx, &hosttypes.QueryParamsRequest{})
	s.Require().NoError(err)

	return *res.Params
}

// TestControllerEnabledParam tests that changing the ControllerEnabled param works as expected
func (s *InterchainAccountsParamsTestSuite) TestControllerEnabledParam() {
	t := s.T()
	ctx := context.TODO()

	// setup relayers and connection-0 between two chains
	// channel-0 is a transfer channel but it will not be used in this test case
	_, _ = s.SetupChainsRelayerAndChannel(ctx, nil)
	chainA, _ := s.GetChains()
	chainAVersion := chainA.Config().Images[0].Version

	// setup controller account on chainA
	controllerAccount := s.CreateUserOnChainA(ctx, testvalues.StartingTokenAmount)
	controllerAddress := controllerAccount.FormattedAddress()

	t.Run("ensure the controller is enabled", func(t *testing.T) {
		params := s.QueryControllerParams(ctx, chainA)
		s.Require().True(params.ControllerEnabled)
	})

	t.Run("disable the controller", func(t *testing.T) {
		if testvalues.SelfParamsFeatureReleases.IsSupported(chainAVersion) {
			authority, err := s.QueryModuleAccountAddress(ctx, govtypes.ModuleName, chainA)
			s.Require().NoError(err)
			s.Require().NotNil(authority)

			msg := controllertypes.MsgUpdateParams{
				Signer: authority.String(),
				Params: controllertypes.NewParams(false),
			}
			s.ExecuteAndPassGovV1Proposal(ctx, &msg, chainA, controllerAccount)
		} else {
			changes := []paramsproposaltypes.ParamChange{
				paramsproposaltypes.NewParamChange(controllertypes.StoreKey, string(controllertypes.KeyControllerEnabled), "false"),
			}

			proposal := paramsproposaltypes.NewParameterChangeProposal(ibctesting.Title, ibctesting.Description, changes)
			s.ExecuteAndPassGovV1Beta1Proposal(ctx, chainA, controllerAccount, proposal)
		}
	})

	t.Run("ensure controller is disabled", func(t *testing.T) {
		params := s.QueryControllerParams(ctx, chainA)
		s.Require().False(params.ControllerEnabled)
	})

	t.Run("ensure that broadcasting a MsgRegisterInterchainAccount fails", func(t *testing.T) {
		// explicitly set the version string because we don't want to use incentivized channels.
		version := icatypes.NewDefaultMetadataString(ibctesting.FirstConnectionID, ibctesting.FirstConnectionID)
		msgRegisterAccount := controllertypes.NewMsgRegisterInterchainAccount(ibctesting.FirstConnectionID, controllerAddress, version, channeltypes.ORDERED)

		txResp := s.BroadcastMessages(ctx, chainA, controllerAccount, msgRegisterAccount)
		s.AssertTxFailure(txResp, controllertypes.ErrControllerSubModuleDisabled)
	})
}

func (s *InterchainAccountsParamsTestSuite) TestHostEnabledParam() {
	t := s.T()
	ctx := context.TODO()

	// setup relayers and connection-0 between two chains
	// channel-0 is a transfer channel but it will not be used in this test case
	relayer, _ := s.SetupChainsRelayerAndChannel(ctx, nil)
	chainA, chainB := s.GetChains()
	chainBVersion := chainB.Config().Images[0].Version

	// setup 2 accounts: controller account on chain A, a second chain B account (to do the disable host gov proposal)
	// host account will be created when the ICA is registered
	controllerAccount := s.CreateUserOnChainA(ctx, testvalues.StartingTokenAmount)
	controllerAddress := controllerAccount.FormattedAddress()
	chainBAccount := s.CreateUserOnChainB(ctx, testvalues.StartingTokenAmount)
	chainBAddress := chainBAccount.FormattedAddress()
	var hostAccount string

	// Assert that default value for enabled is true.
	t.Run("ensure the host is enabled", func(t *testing.T) {
		params := s.QueryHostParams(ctx, chainB)
		s.Require().True(params.HostEnabled)
		s.Require().Equal([]string{hosttypes.AllowAllHostMsgs}, params.AllowMessages)
	})

	t.Run("ensure ica packets are flowing before disabling the host", func(t *testing.T) {
		t.Run("broadcast MsgRegisterInterchainAccount", func(t *testing.T) {
			// explicitly set the version string because we don't want to use incentivized channels.
			version := icatypes.NewDefaultMetadataString(ibctesting.FirstConnectionID, ibctesting.FirstConnectionID)
			msgRegisterAccount := controllertypes.NewMsgRegisterInterchainAccount(ibctesting.FirstConnectionID, controllerAddress, version, channeltypes.ORDERED)

			txResp := s.BroadcastMessages(ctx, chainA, controllerAccount, msgRegisterAccount)
			s.AssertTxSuccess(txResp)
		})

		t.Run("start relayer", func(t *testing.T) {
			s.StartRelayer(relayer)
		})

		t.Run("verify interchain account", func(t *testing.T) {
			var err error
			hostAccount, err = s.QueryInterchainAccount(ctx, chainA, controllerAddress, ibctesting.FirstConnectionID)
			s.Require().NoError(err)
			s.Require().NotEmpty(hostAccount)

			channels, err := relayer.GetChannels(ctx, s.GetRelayerExecReporter(), chainA.Config().ChainID)
			s.Require().NoError(err)
			s.Require().Equal(len(channels), 2)
		})

		t.Run("stop relayer", func(t *testing.T) {
			s.StopRelayer(ctx, relayer)
		})
	})

	t.Run("disable the host", func(t *testing.T) {
		if testvalues.SelfParamsFeatureReleases.IsSupported(chainBVersion) {
			authority, err := s.QueryModuleAccountAddress(ctx, govtypes.ModuleName, chainB)
			s.Require().NoError(err)
			s.Require().NotNil(authority)

			msg := hosttypes.MsgUpdateParams{
				Signer: authority.String(),
				Params: hosttypes.NewParams(false, []string{hosttypes.AllowAllHostMsgs}),
			}
			s.ExecuteAndPassGovV1Proposal(ctx, &msg, chainB, chainBAccount)
		} else {
			changes := []paramsproposaltypes.ParamChange{
				paramsproposaltypes.NewParamChange(hosttypes.StoreKey, string(hosttypes.KeyHostEnabled), "false"),
			}

			proposal := paramsproposaltypes.NewParameterChangeProposal(ibctesting.Title, ibctesting.Description, changes)
			s.ExecuteAndPassGovV1Beta1Proposal(ctx, chainB, chainBAccount, proposal)
		}
	})

	t.Run("ensure the host is disabled", func(t *testing.T) {
		params := s.QueryHostParams(ctx, chainB)
		s.Require().False(params.HostEnabled)
	})

	t.Run("ensure that ica packets are not flowing", func(t *testing.T) {
		t.Run("fund interchain account wallet", func(t *testing.T) {
			// fund the host account so it has some $$ to send
			err := chainB.SendFunds(ctx, interchaintest.FaucetAccountKeyName, ibc.WalletAmount{
				Address: hostAccount,
				Amount:  sdkmath.NewInt(testvalues.StartingTokenAmount),
				Denom:   chainB.Config().Denom,
			})
			s.Require().NoError(err)
		})

		t.Run("broadcast MsgSendTx", func(t *testing.T) {
			// assemble bank transfer message from host account to user account on host chain
			msgSend := &banktypes.MsgSend{
				FromAddress: hostAccount,
				ToAddress:   chainBAddress,
				Amount:      sdk.NewCoins(testvalues.DefaultTransferAmount(chainB.Config().Denom)),
			}

			cdc := testsuite.Codec()
			bz, err := icatypes.SerializeCosmosTx(cdc, []proto.Message{msgSend}, icatypes.EncodingProtobuf)
			s.Require().NoError(err)

			packetData := icatypes.InterchainAccountPacketData{
				Type: icatypes.EXECUTE_TX,
				Data: bz,
				Memo: "e2e",
			}

			msgSendTx := controllertypes.NewMsgSendTx(controllerAddress, ibctesting.FirstConnectionID, uint64(time.Hour.Nanoseconds()), packetData)

			resp := s.BroadcastMessages(
				ctx,
				chainA,
				controllerAccount,
				msgSendTx,
			)

			s.AssertTxSuccess(resp)
		})

		t.Run("start relayer", func(t *testing.T) {
			s.StartRelayer(relayer)
		})

		s.Require().NoError(test.WaitForBlocks(ctx, 10, chainA, chainB))

		t.Run("verify no tokens were transferred", func(t *testing.T) {
			chainBAccountBalance, err := s.QueryBalance(ctx, chainB, chainBAddress, chainB.Config().Denom)
			s.Require().NoError(err)
			s.Require().Equal(testvalues.StartingTokenAmount, chainBAccountBalance.Int64())

			hostAccountBalance, err := s.QueryBalance(ctx, chainB, hostAccount, chainB.Config().Denom)
			s.Require().NoError(err)
			s.Require().Equal(testvalues.StartingTokenAmount, hostAccountBalance.Int64())
		})

		t.Run("verify acknowledgement error in ack transaction", func(t *testing.T) {
			txSearchRes, err := s.QueryTxsByEvents(ctx, chainB, 1, 1, "message.action='/ibc.core.channel.v1.MsgRecvPacket'", "")
			s.Require().NoError(err)
			s.Require().Len(txSearchRes.Txs, 1)

			errorMessage, isFound := s.ExtractValueFromEvents(
				txSearchRes.Txs[0].Events,
				coretypes.ErrorAttributeKeyPrefix+icatypes.EventTypePacket,
				coretypes.ErrorAttributeKeyPrefix+icatypes.AttributeKeyAckError,
			)

			s.Require().True(isFound)
			s.Require().Equal(errorMessage, hosttypes.ErrHostSubModuleDisabled.Error())
		})
	})
}
