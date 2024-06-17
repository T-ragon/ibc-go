package _5_AggreLite_test

import (
	clienttypes "github.com/T-ragon/ibc-go/v9/modules/core/02-client/types"
	commitmenttypes "github.com/T-ragon/ibc-go/v9/modules/core/23-commitment/types"
	host "github.com/T-ragon/ibc-go/v9/modules/core/24-host"
	"github.com/T-ragon/ibc-go/v9/modules/core/exported"
	aggrelite "github.com/T-ragon/ibc-go/v9/modules/light-clients/05-AggreLite"
	ibctesting "github.com/T-ragon/ibc-go/v9/testing"
	"log"
)

const (
	// Do not change the length of these variables
	fiftyCharChainID    = "12345678901234567890123456789012345678901234567890"
	fiftyOneCharChainID = "123456789012345678901234567890123456789012345678901"
)

var invalidProof = []byte("invalid proof")

func (suite *AggreLiteTestSuite) TestStatus() {
	var (
		path        *ibctesting.Path
		clientState *aggrelite.ClientState
	)

	testCases := []struct {
		name      string
		malleate  func()
		expStatus exported.Status
	}{
		{"client is active", func() {}, exported.Active},
		{"client is frozen", func() {
			clientState.FrozenHeight = clienttypes.NewHeight(0, 1)
			path.EndpointA.SetClientState(clientState)
		}, exported.Frozen},
		{"client status without consensus state", func() {
			clientState.LatestHeight = clientState.LatestHeight.Increment().(clienttypes.Height)
			path.EndpointA.SetClientState(clientState)
		}, exported.Expired},
		{"client status is expired", func() {
			suite.coordinator.IncrementTimeBy(clientState.TrustingPeriod)
		}, exported.Expired},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			suite.SetupTest()

			path = ibctesting.NewPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupClients(path)

			clientStore := suite.chainA.App.GetIBCKeeper().ClientKeeper.ClientStore(suite.chainA.GetContext(), path.EndpointA.ClientID)
			clientState = path.EndpointA.GetClientState().(*aggrelite.ClientState)

			tc.malleate()

			status := clientState.Status(suite.chainA.GetContext(), clientStore, suite.chainA.App.AppCodec())
			suite.Require().Equal(tc.expStatus, status)
		})

	}
}

//func (suite *AggreLiteTestSuite) TestValidate() {
//	testCases := []struct {
//		name        string
//		clientState *aggrelite.ClientState
//		expPass     bool
//	}{
//		{
//			name:        "valid client",
//			clientState: aggrelite.NewClientState(chainID, aggrelite.DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath),
//			expPass:     true,
//		},
//		{
//			name:        "valid client with nil upgrade path",
//			clientState: NewClientState(chainID, DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), nil),
//			expPass:     true,
//		},
//		{
//			name:        "invalid chainID",
//			clientState: NewClientState("  ", DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath),
//			expPass:     false,
//		},
//		{
//			// NOTE: if this test fails, the code must account for the change in chainID length across tendermint versions!
//			// Do not only fix the test, fix the code!
//			// https://github.com/cosmos/ibc-go/issues/177
//			name:        "valid chainID - chainID validation failed for chainID of length 50! ",
//			clientState: NewClientState(fiftyCharChainID, DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath),
//			expPass:     true,
//		},
//		{
//			// NOTE: if this test fails, the code must account for the change in chainID length across tendermint versions!
//			// Do not only fix the test, fix the code!
//			// https://github.com/cosmos/ibc-go/issues/177
//			name:        "invalid chainID - chainID validation did not fail for chainID of length 51! ",
//			clientState: NewClientState(fiftyOneCharChainID, DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath),
//			expPass:     false,
//		},
//		{
//			name:        "invalid trust level",
//			clientState: NewClientState(chainID, Fraction{Numerator: 0, Denominator: 1}, trustingPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath),
//			expPass:     false,
//		},
//		{
//			name:        "invalid zero trusting period",
//			clientState: NewClientState(chainID, DefaultTrustLevel, 0, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath),
//			expPass:     false,
//		},
//		{
//			name:        "invalid negative trusting period",
//			clientState: NewClientState(chainID, DefaultTrustLevel, -1, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath),
//			expPass:     false,
//		},
//		{
//			name:        "invalid zero unbonding period",
//			clientState: NewClientState(chainID, DefaultTrustLevel, trustingPeriod, 0, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath),
//			expPass:     false,
//		},
//		{
//			name:        "invalid negative unbonding period",
//			clientState: NewClientState(chainID, DefaultTrustLevel, trustingPeriod, -1, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath),
//			expPass:     false,
//		},
//		{
//			name:        "invalid zero max clock drift",
//			clientState: NewClientState(chainID, DefaultTrustLevel, trustingPeriod, ubdPeriod, 0, height, commitmenttypes.GetSDKSpecs(), upgradePath),
//			expPass:     false,
//		},
//		{
//			name:        "invalid negative max clock drift",
//			clientState: NewClientState(chainID, DefaultTrustLevel, trustingPeriod, ubdPeriod, -1, height, commitmenttypes.GetSDKSpecs(), upgradePath),
//			expPass:     false,
//		},
//		{
//			name:        "invalid revision number",
//			clientState: NewClientState(chainID, DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, clienttypes.NewHeight(1, 1), commitmenttypes.GetSDKSpecs(), upgradePath),
//			expPass:     false,
//		},
//		{
//			name:        "invalid revision height",
//			clientState: NewClientState(chainID, DefaultTrustLevel, trustingPeriod, ubdPeriod, maxClockDrift, clienttypes.ZeroHeight(), commitmenttypes.GetSDKSpecs(), upgradePath),
//			expPass:     false,
//		},
//		{
//			name:        "trusting period not less than unbonding period",
//			clientState: NewClientState(chainID, DefaultTrustLevel, ubdPeriod, ubdPeriod, maxClockDrift, height, commitmenttypes.GetSDKSpecs(), upgradePath),
//			expPass:     false,
//		},
//		{
//			name:        "proof specs is nil",
//			clientState: NewClientState(chainID, DefaultTrustLevel, ubdPeriod, ubdPeriod, maxClockDrift, height, nil, upgradePath),
//			expPass:     false,
//		},
//		{
//			name:        "proof specs contains nil",
//			clientState: NewClientState(chainID, DefaultTrustLevel, ubdPeriod, ubdPeriod, maxClockDrift, height, []*ics23.ProofSpec{ics23.TendermintSpec, nil}, upgradePath),
//			expPass:     false,
//		},
//	}
//
//	for _, tc := range testCases {
//		tc := tc
//		suite.Run(tc.name, func() {
//			err := tc.clientState.Validate()
//			if tc.expPass {
//				suite.Require().NoError(err, tc.name)
//			} else {
//				suite.Require().Error(err, tc.name)
//			}
//		})
//	}
//}

func (suite *AggreLiteTestSuite) TestInitialize() {
	testCases := []struct {
		name           string
		consensusState exported.ConsensusState
		expPass        bool
	}{
		{
			name:           "valid consensus state",
			consensusState: &aggrelite.ConsensusState{},
			expPass:        true,
		},
	}
	for _, tc := range testCases {
		suite.SetupTest()
		path := ibctesting.NewPath(suite.chainA, suite.chainB)
		tmConfig, ok := path.EndpointB.ClientConfig.(*ibctesting.AggreLiteConfig)
		log.Print(ok)
		suite.Require().True(ok)

		clientState := aggrelite.NewClientState(
			path.EndpointB.Chain.ChainID,
			tmConfig.TrustLevel, tmConfig.TrustingPeriod, tmConfig.UnbondingPeriod, tmConfig.MaxClockDrift,
			suite.chainB.LastHeader.GetTrustedHeight(), commitmenttypes.GetSDKSpecs(), ibctesting.UpgradePath)

		store := suite.chainA.App.GetIBCKeeper().ClientKeeper.ClientStore(suite.chainA.GetContext(), path.EndpointA.ClientID)
		err := clientState.Initialize(suite.chainA.GetContext(), suite.chainA.Codec, store, tc.consensusState)

		if tc.expPass {
			suite.Require().NoError(err, "valid case returned an error")
			suite.Require().True(store.Has(host.ClientStateKey()))
			suite.Require().True(store.Has(host.ConsensusStateKey(suite.chainB.LastHeader.GetTrustedHeight())))
		} else {
			suite.Require().Error(err, "invalid case didn't return an error")
			suite.Require().False(store.Has(host.ClientStateKey()))
			suite.Require().False(store.Has(host.ConsensusStateKey(suite.chainB.LastHeader.GetTrustedHeight())))
		}
	}
}
