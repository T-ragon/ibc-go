package keeper

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/codec"
	"strings"
)

type Keeper struct {
	cdc codec.BinaryCodec

	authority string
}

func NewKeeper(cdc codec.BinaryCodec, authority string) Keeper {
	if strings.TrimSpace(authority) == "" {
		panic(errors.New("authority must be non-empty"))
	}

	return Keeper{
		cdc:       cdc,
		authority: authority,
	}
}

func (k Keeper) Codec() codec.BinaryCodec {
	return k.cdc
}
