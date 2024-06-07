package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	icatypes "github.com/T-ragon/ibc-go/v9/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/T-ragon/ibc-go/v9/modules/core/04-channel/types"
	"github.com/T-ragon/ibc-go/v9/modules/core/exported"
)

// EmitAcknowledgementEvent emits an event signalling a successful or failed acknowledgement and including the error
// details if any.
func EmitAcknowledgementEvent(ctx sdk.Context, packet channeltypes.Packet, ack exported.Acknowledgement, err error) {
	attributes := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, icatypes.ModuleName),
		sdk.NewAttribute(icatypes.AttributeKeyControllerChannelID, packet.GetDestChannel()),
		sdk.NewAttribute(icatypes.AttributeKeyAckSuccess, fmt.Sprintf("%t", ack.Success())),
	}

	if err != nil {
		attributes = append(attributes, sdk.NewAttribute(icatypes.AttributeKeyAckError, err.Error()))
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			icatypes.EventTypePacket,
			attributes...,
		),
	)
}
