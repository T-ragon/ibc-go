package ibctesting

import (
	"time"

	connectiontypes "github.com/T-ragon/ibc-go/v9/modules/core/03-connection/types"
	channeltypes "github.com/T-ragon/ibc-go/v9/modules/core/04-channel/types"
	"github.com/T-ragon/ibc-go/v9/modules/core/exported"
	aggrelite "github.com/T-ragon/ibc-go/v9/modules/light-clients/05-AggreLite"
	ibctm "github.com/T-ragon/ibc-go/v9/modules/light-clients/07-tendermint"
	"github.com/T-ragon/ibc-go/v9/testing/mock"
)

type ClientConfig interface {
	GetClientType() string
}

type TendermintConfig struct {
	TrustLevel      ibctm.Fraction
	TrustingPeriod  time.Duration
	UnbondingPeriod time.Duration
	MaxClockDrift   time.Duration
}

type AggreLiteConfig struct {
	TrustLevel      aggrelite.Fraction
	TrustingPeriod  time.Duration
	UnbondingPeriod time.Duration
	MaxClockDrift   time.Duration
}

func NewAggreLiteConfig() *AggreLiteConfig {
	return &AggreLiteConfig{
		TrustLevel:      aggrelite.Fraction(DefaultTrustLevel),
		TrustingPeriod:  TrustingPeriod,
		UnbondingPeriod: UnbondingPeriod,
		MaxClockDrift:   MaxClockDrift,
	}
}
func NewTendermintConfig() *TendermintConfig {
	return &TendermintConfig{
		TrustLevel:      DefaultTrustLevel,
		TrustingPeriod:  TrustingPeriod,
		UnbondingPeriod: UnbondingPeriod,
		MaxClockDrift:   MaxClockDrift,
	}
}

func (*AggreLiteConfig) GetClientType() string {
	return exported.AggreLite
}

func (*TendermintConfig) GetClientType() string {
	return exported.Tendermint
}

type ConnectionConfig struct {
	DelayPeriod uint64
	Version     *connectiontypes.Version
}

func NewConnectionConfig() *ConnectionConfig {
	return &ConnectionConfig{
		DelayPeriod: DefaultDelayPeriod,
		Version:     ConnectionVersion,
	}
}

type ChannelConfig struct {
	PortID          string
	Version         string
	Order           channeltypes.Order
	ProposedUpgrade channeltypes.Upgrade
}

func NewChannelConfig() *ChannelConfig {
	return &ChannelConfig{
		PortID:  mock.PortID,
		Version: DefaultChannelVersion,
		Order:   channeltypes.UNORDERED,
	}
}
