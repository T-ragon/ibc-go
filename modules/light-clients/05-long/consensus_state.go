package _5_long

import (
	commitmenttypes "github.com/cosmos/ibc-go/v8/modules/core/23-commitment/types"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
)

var _ exported.ConsensusState = (*ConsensusState)(nil)

func NewConsensusState(root commitmenttypes.MerkleRoot) *ConsensusState {
	return &ConsensusState{
		Root: MerkleRoot{
			Hash: root.Hash,
		},
	}
}
func (ConsensusState) ClientType() string {
	return exported.Long
}

func (cs ConsensusState) GetRoot() exported.Root {
	return nil
}

func (cs ConsensusState) GetTimestamp() uint64 {
	return 0
}

func (cs ConsensusState) ValidateBasic() error {
	return nil
}
