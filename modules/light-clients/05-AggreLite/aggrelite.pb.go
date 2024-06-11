package _5_AggreLite

import (
	fmt "fmt"
	github_com_cometbft_cometbft_libs_bytes "github.com/cometbft/cometbft/libs/bytes"
	types2 "github.com/cometbft/cometbft/proto/tendermint/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	types "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	types1 "github.com/cosmos/ibc-go/v8/modules/core/23-commitment/types"
	_go "github.com/cosmos/ics23/go"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ClientState from Tendermint tracks the current validator set, latest height,
// and a possible frozen height.
type ClientState struct {
	ChainId    string   `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	TrustLevel Fraction `protobuf:"bytes,2,opt,name=trust_level,json=trustLevel,proto3" json:"trust_level"`
	// duration of the period since the LastestTimestamp during which the
	// submitted headers are valid for upgrade
	TrustingPeriod time.Duration `protobuf:"bytes,3,opt,name=trusting_period,json=trustingPeriod,proto3,stdduration" json:"trusting_period"`
	// duration of the staking unbonding period
	UnbondingPeriod time.Duration `protobuf:"bytes,4,opt,name=unbonding_period,json=unbondingPeriod,proto3,stdduration" json:"unbonding_period"`
	// defines how much new (untrusted) header's Time can drift into the future.
	MaxClockDrift time.Duration `protobuf:"bytes,5,opt,name=max_clock_drift,json=maxClockDrift,proto3,stdduration" json:"max_clock_drift"`
	// Block height when the client was frozen due to a misbehaviour
	FrozenHeight types.Height `protobuf:"bytes,6,opt,name=frozen_height,json=frozenHeight,proto3" json:"frozen_height"`
	// Latest height the client was updated to
	LatestHeight types.Height `protobuf:"bytes,7,opt,name=latest_height,json=latestHeight,proto3" json:"latest_height"`
	// Proof specifications used in verifying counterparty state
	ProofSpecs []*_go.ProofSpec `protobuf:"bytes,8,rep,name=proof_specs,json=proofSpecs,proto3" json:"proof_specs,omitempty"`
	// Path at which next upgraded client will be committed.
	// Each element corresponds to the key for a single CommitmentProof in the
	// chained proof. NOTE: ClientState must stored under
	// `{upgradePath}/{upgradeHeight}/clientState` ConsensusState must be stored
	// under `{upgradepath}/{upgradeHeight}/consensusState` For SDK chains using
	// the default upgrade module, upgrade_path should be []string{"upgrade",
	// "upgradedIBCState"}`
	UpgradePath []string `protobuf:"bytes,9,rep,name=upgrade_path,json=upgradePath,proto3" json:"upgrade_path,omitempty"`
	// allow_update_after_expiry is deprecated
	AllowUpdateAfterExpiry bool `protobuf:"varint,10,opt,name=allow_update_after_expiry,json=allowUpdateAfterExpiry,proto3" json:"allow_update_after_expiry,omitempty"` // Deprecated: Do not use.
	// allow_update_after_misbehaviour is deprecated
	AllowUpdateAfterMisbehaviour bool `protobuf:"varint,11,opt,name=allow_update_after_misbehaviour,json=allowUpdateAfterMisbehaviour,proto3" json:"allow_update_after_misbehaviour,omitempty"` // Deprecated: Do not use.
}

func (m *ClientState) Reset()         { *m = ClientState{} }
func (m *ClientState) String() string { return proto.CompactTextString(m) }
func (*ClientState) ProtoMessage()    {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d6cf2b288949be, []int{0}
}
func (m *ClientState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientState.Merge(m, src)
}
func (m *ClientState) XXX_Size() int {
	return m.Size()
}
func (m *ClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientState.DiscardUnknown(m)
}

var xxx_messageInfo_ClientState proto.InternalMessageInfo

// ConsensusState defines the consensus state from Tendermint.
type ConsensusState struct {
	// timestamp that corresponds to the block height in which the ConsensusState
	// was stored.
	Timestamp time.Time `protobuf:"bytes,1,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	// commitment root (i.e app hash)
	Root               types1.MerkleRoot                                `protobuf:"bytes,2,opt,name=root,proto3" json:"root"`
	NextValidatorsHash github_com_cometbft_cometbft_libs_bytes.HexBytes `protobuf:"bytes,3,opt,name=next_validators_hash,json=nextValidatorsHash,proto3,casttype=github.com/cometbft/cometbft/libs/bytes.HexBytes" json:"next_validators_hash,omitempty"`
}

func (m *ConsensusState) Reset()         { *m = ConsensusState{} }
func (m *ConsensusState) String() string { return proto.CompactTextString(m) }
func (*ConsensusState) ProtoMessage()    {}
func (*ConsensusState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d6cf2b288949be, []int{1}
}
func (m *ConsensusState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsensusState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsensusState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsensusState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusState.Merge(m, src)
}
func (m *ConsensusState) XXX_Size() int {
	return m.Size()
}
func (m *ConsensusState) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusState.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusState proto.InternalMessageInfo

// Misbehaviour is a wrapper over two conflicting Headers
// that implements Misbehaviour interface expected by ICS-02
type Misbehaviour struct {
	// ClientID is deprecated
	ClientId string  `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"` // Deprecated: Do not use.
	Header1  *Header `protobuf:"bytes,2,opt,name=header_1,json=header1,proto3" json:"header_1,omitempty"`
	Header2  *Header `protobuf:"bytes,3,opt,name=header_2,json=header2,proto3" json:"header_2,omitempty"`
}

func (m *Misbehaviour) Reset()         { *m = Misbehaviour{} }
func (m *Misbehaviour) String() string { return proto.CompactTextString(m) }
func (*Misbehaviour) ProtoMessage()    {}
func (*Misbehaviour) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d6cf2b288949be, []int{2}
}
func (m *Misbehaviour) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Misbehaviour) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Misbehaviour.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Misbehaviour) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Misbehaviour.Merge(m, src)
}
func (m *Misbehaviour) XXX_Size() int {
	return m.Size()
}
func (m *Misbehaviour) XXX_DiscardUnknown() {
	xxx_messageInfo_Misbehaviour.DiscardUnknown(m)
}

var xxx_messageInfo_Misbehaviour proto.InternalMessageInfo

// Header defines the Tendermint client consensus Header.
// It encapsulates all the information necessary to update from a trusted
// Tendermint ConsensusState. The inclusion of TrustedHeight and
// TrustedValidators allows this update to process correctly, so long as the
// ConsensusState for the TrustedHeight exists, this removes race conditions
// among relayers The SignedHeader and ValidatorSet are the new untrusted update
// fields for the client. The TrustedHeight is the height of a stored
// ConsensusState on the client that will be used to verify the new untrusted
// header. The Trusted ConsensusState must be within the unbonding period of
// current time in order to correctly verify, and the TrustedValidators must
// hash to TrustedConsensusState.NextValidatorsHash since that is the last
// trusted validator set at the TrustedHeight.
type Header struct {
	*types2.SignedHeader `protobuf:"bytes,1,opt,name=signed_header,json=signedHeader,proto3,embedded=signed_header" json:"signed_header,omitempty"`
	ValidatorSet         *types2.ValidatorSet `protobuf:"bytes,2,opt,name=validator_set,json=validatorSet,proto3" json:"validator_set,omitempty"`
	TrustedHeight        types.Height         `protobuf:"bytes,3,opt,name=trusted_height,json=trustedHeight,proto3" json:"trusted_height"`
	TrustedValidators    *types2.ValidatorSet `protobuf:"bytes,4,opt,name=trusted_validators,json=trustedValidators,proto3" json:"trusted_validators,omitempty"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d6cf2b288949be, []int{3}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetValidatorSet() *types2.ValidatorSet {
	if m != nil {
		return m.ValidatorSet
	}
	return nil
}

func (m *Header) GetTrustedHeight() types.Height {
	if m != nil {
		return m.TrustedHeight
	}
	return types.Height{}
}

func (m *Header) GetTrustedValidators() *types2.ValidatorSet {
	if m != nil {
		return m.TrustedValidators
	}
	return nil
}

// Fraction defines the protobuf message type for tmmath.Fraction that only
// supports positive values.
type Fraction struct {
	Numerator   uint64 `protobuf:"varint,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	Denominator uint64 `protobuf:"varint,2,opt,name=denominator,proto3" json:"denominator,omitempty"`
}

func (m *Fraction) Reset()         { *m = Fraction{} }
func (m *Fraction) String() string { return proto.CompactTextString(m) }
func (*Fraction) ProtoMessage()    {}
func (*Fraction) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d6cf2b288949be, []int{4}
}
func (m *Fraction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fraction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fraction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fraction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fraction.Merge(m, src)
}
func (m *Fraction) XXX_Size() int {
	return m.Size()
}
func (m *Fraction) XXX_DiscardUnknown() {
	xxx_messageInfo_Fraction.DiscardUnknown(m)
}

var xxx_messageInfo_Fraction proto.InternalMessageInfo

func (m *Fraction) GetNumerator() uint64 {
	if m != nil {
		return m.Numerator
	}
	return 0
}

func (m *Fraction) GetDenominator() uint64 {
	if m != nil {
		return m.Denominator
	}
	return 0
}

func init() {
	proto.RegisterType((*ClientState)(nil), "ibc.lightclients.tendermint.v1.ClientState")
	proto.RegisterType((*ConsensusState)(nil), "ibc.lightclients.tendermint.v1.ConsensusState")
	proto.RegisterType((*Misbehaviour)(nil), "ibc.lightclients.tendermint.v1.Misbehaviour")
	proto.RegisterType((*Header)(nil), "ibc.lightclients.tendermint.v1.Header")
	proto.RegisterType((*Fraction)(nil), "ibc.lightclients.tendermint.v1.Fraction")
}

func init() {
	proto.RegisterFile("ibc/lightclients/tendermint/v1/tendermint.proto", fileDescriptor_c6d6cf2b288949be)
}

var fileDescriptor_c6d6cf2b288949be = []byte{
	// 943 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x4f, 0x6f, 0xe3, 0xc4,
	0x1b, 0xc7, 0xe3, 0x34, 0xdb, 0x26, 0x93, 0x64, 0xfb, 0xfb, 0x8d, 0x56, 0xc8, 0xad, 0xaa, 0x24,
	0xf4, 0x00, 0xb9, 0xd4, 0xde, 0x64, 0x91, 0x40, 0x2c, 0x1c, 0x48, 0x77, 0xa1, 0x5d, 0xb6, 0x50,
	0xb9, 0xc0, 0x81, 0x8b, 0x35, 0xb6, 0x27, 0xf6, 0x68, 0x6d, 0x8f, 0x35, 0x33, 0x0e, 0x29, 0x27,
	0x8e, 0x1c, 0xf7, 0xc8, 0x91, 0x97, 0xc0, 0xcb, 0xd8, 0x63, 0x2f, 0x48, 0x9c, 0x0a, 0x4a, 0xdf,
	0x05, 0x27, 0x34, 0x7f, 0x9c, 0x98, 0xb2, 0x62, 0x2b, 0x2e, 0xd1, 0x33, 0xf3, 0x7c, 0x9f, 0x4f,
	0x66, 0x9e, 0x3f, 0x63, 0xe0, 0x92, 0x20, 0x74, 0x53, 0x12, 0x27, 0x22, 0x4c, 0x09, 0xce, 0x05,
	0x77, 0x05, 0xce, 0x23, 0xcc, 0x32, 0x92, 0x0b, 0x77, 0x31, 0xa9, 0xad, 0x9c, 0x82, 0x51, 0x41,
	0xe1, 0x80, 0x04, 0xa1, 0x53, 0x0f, 0x70, 0x6a, 0x92, 0xc5, 0x64, 0x7f, 0x54, 0x8b, 0x17, 0x97,
	0x05, 0xe6, 0xee, 0x02, 0xa5, 0x24, 0x42, 0x82, 0x32, 0x4d, 0xd8, 0x3f, 0xf8, 0x87, 0x42, 0xfd,
	0x56, 0xde, 0x90, 0xf2, 0x8c, 0x72, 0x97, 0x84, 0x7c, 0xfa, 0x48, 0x9e, 0xa0, 0x60, 0x94, 0xce,
	0x2b, 0xef, 0x20, 0xa6, 0x34, 0x4e, 0xb1, 0xab, 0x56, 0x41, 0x39, 0x77, 0xa3, 0x92, 0x21, 0x41,
	0x68, 0x6e, 0xfc, 0xc3, 0xdb, 0x7e, 0x41, 0x32, 0xcc, 0x05, 0xca, 0x8a, 0x4a, 0x20, 0xef, 0x1b,
	0x52, 0x86, 0x5d, 0x7d, 0x7c, 0xf9, 0x0f, 0xda, 0x32, 0x82, 0x77, 0x37, 0x02, 0x9a, 0x65, 0x44,
	0x64, 0x95, 0x68, 0xbd, 0x32, 0xc2, 0x07, 0x31, 0x8d, 0xa9, 0x32, 0x5d, 0x69, 0xe9, 0xdd, 0xc3,
	0xd5, 0x3d, 0xd0, 0x3d, 0x56, 0xbc, 0x0b, 0x81, 0x04, 0x86, 0x7b, 0xa0, 0x1d, 0x26, 0x88, 0xe4,
	0x3e, 0x89, 0x6c, 0x6b, 0x64, 0x8d, 0x3b, 0xde, 0x8e, 0x5a, 0x9f, 0x46, 0xf0, 0x4b, 0xd0, 0x15,
	0xac, 0xe4, 0xc2, 0x4f, 0xf1, 0x02, 0xa7, 0x76, 0x73, 0x64, 0x8d, 0xbb, 0xd3, 0xb1, 0xf3, 0xef,
	0xf9, 0x75, 0x3e, 0x65, 0x28, 0x94, 0x17, 0x9e, 0xb5, 0x5e, 0x5d, 0x0f, 0x1b, 0x1e, 0x50, 0x88,
	0xe7, 0x92, 0x00, 0x9f, 0x83, 0x5d, 0xb5, 0x22, 0x79, 0xec, 0x17, 0x98, 0x11, 0x1a, 0xd9, 0x5b,
	0x0a, 0xba, 0xe7, 0xe8, 0xb4, 0x38, 0x55, 0x5a, 0x9c, 0x27, 0x26, 0x6d, 0xb3, 0xb6, 0xa4, 0xfc,
	0xf4, 0xfb, 0xd0, 0xf2, 0xee, 0x57, 0xb1, 0xe7, 0x2a, 0x14, 0x7e, 0x01, 0xfe, 0x57, 0xe6, 0x01,
	0xcd, 0xa3, 0x1a, 0xae, 0x75, 0x77, 0xdc, 0xee, 0x3a, 0xd8, 0xf0, 0x3e, 0x07, 0xbb, 0x19, 0x5a,
	0xfa, 0x61, 0x4a, 0xc3, 0x17, 0x7e, 0xc4, 0xc8, 0x5c, 0xd8, 0xf7, 0xee, 0x8e, 0xeb, 0x67, 0x68,
	0x79, 0x2c, 0x43, 0x9f, 0xc8, 0x48, 0xf8, 0x14, 0xf4, 0xe7, 0x8c, 0x7e, 0x8f, 0x73, 0x3f, 0xc1,
	0x32, 0x57, 0xf6, 0xb6, 0x42, 0xed, 0xab, 0xec, 0xc9, 0xea, 0x39, 0xa6, 0xa8, 0x8b, 0x89, 0x73,
	0xa2, 0x14, 0x26, 0x5f, 0x3d, 0x1d, 0xa6, 0xf7, 0x24, 0x26, 0x45, 0x02, 0x73, 0x51, 0x61, 0x76,
	0xee, 0x8a, 0xd1, 0x61, 0x06, 0xf3, 0x18, 0x74, 0x55, 0x97, 0xfa, 0xbc, 0xc0, 0x21, 0xb7, 0xdb,
	0xa3, 0x2d, 0x05, 0xd1, 0x9d, 0xec, 0xa8, 0x4e, 0x96, 0x84, 0x73, 0xa9, 0xb9, 0x28, 0x70, 0xe8,
	0x81, 0xa2, 0x32, 0x39, 0x7c, 0x1b, 0xf4, 0xca, 0x22, 0x66, 0x28, 0xc2, 0x7e, 0x81, 0x44, 0x62,
	0x77, 0x46, 0x5b, 0xe3, 0x8e, 0xd7, 0x35, 0x7b, 0xe7, 0x48, 0x24, 0xf0, 0x63, 0xb0, 0x87, 0xd2,
	0x94, 0x7e, 0xe7, 0x97, 0x45, 0x84, 0x04, 0xf6, 0xd1, 0x5c, 0x60, 0xe6, 0xe3, 0x65, 0x41, 0xd8,
	0xa5, 0x0d, 0x46, 0xd6, 0xb8, 0x3d, 0x6b, 0xda, 0x96, 0xf7, 0x96, 0x12, 0x7d, 0xad, 0x34, 0x9f,
	0x48, 0xc9, 0x53, 0xa5, 0x80, 0xa7, 0x60, 0xf8, 0x9a, 0xf0, 0x8c, 0xf0, 0x00, 0x27, 0x68, 0x41,
	0x68, 0xc9, 0xec, 0xee, 0x1a, 0x72, 0x70, 0x1b, 0x72, 0x56, 0xd3, 0x7d, 0xd8, 0xfa, 0xf1, 0xe7,
	0x61, 0xe3, 0xf0, 0x87, 0x26, 0xb8, 0x7f, 0x4c, 0x73, 0x8e, 0x73, 0x5e, 0x72, 0xdd, 0xe7, 0x33,
	0xd0, 0x59, 0x8f, 0x9a, 0x6a, 0x74, 0x99, 0x80, 0xdb, 0x75, 0xfd, 0xaa, 0x52, 0xe8, 0xc2, 0xbe,
	0x94, 0x85, 0xdd, 0x84, 0xc1, 0x8f, 0x40, 0x8b, 0x51, 0x2a, 0xcc, 0x24, 0x1c, 0xd6, 0x8a, 0xb0,
	0x99, 0xbd, 0xc5, 0xc4, 0x39, 0xc3, 0xec, 0x45, 0x8a, 0x3d, 0x4a, 0xab, 0x62, 0xa8, 0x28, 0x38,
	0x07, 0x0f, 0x72, 0xbc, 0x14, 0xfe, 0xfa, 0xb9, 0xe1, 0x7e, 0x82, 0x78, 0xa2, 0x46, 0xa0, 0x37,
	0x7b, 0xef, 0xcf, 0xeb, 0xe1, 0xc3, 0x98, 0x88, 0xa4, 0x0c, 0x24, 0x4e, 0x8e, 0x33, 0x16, 0xc1,
	0x5c, 0x6c, 0x8c, 0x94, 0x04, 0xdc, 0x0d, 0x2e, 0x05, 0xe6, 0xce, 0x09, 0x5e, 0xce, 0xa4, 0xe1,
	0x41, 0x49, 0xfc, 0x66, 0x0d, 0x3c, 0x41, 0x3c, 0x31, 0x29, 0xf8, 0xd5, 0x02, 0xbd, 0x7a, 0x66,
	0xe0, 0x10, 0x74, 0x74, 0xaf, 0xac, 0x27, 0x5d, 0xa5, 0xb3, 0xad, 0x37, 0x4f, 0xe5, 0x3c, 0xb5,
	0x13, 0x8c, 0x22, 0xcc, 0xfc, 0x89, 0xb9, 0xe1, 0x3b, 0x6f, 0x9a, 0xf5, 0x13, 0xa5, 0x9f, 0x75,
	0x57, 0xd7, 0xc3, 0x1d, 0x6d, 0x4f, 0xbc, 0x1d, 0x0d, 0x99, 0xd4, 0x78, 0x53, 0x33, 0xe6, 0xff,
	0x81, 0x37, 0xad, 0x78, 0x53, 0x73, 0xaf, 0x5f, 0x9a, 0x60, 0x5b, 0xbb, 0xe0, 0x29, 0xe8, 0x73,
	0x12, 0xe7, 0x38, 0xf2, 0xb5, 0xc4, 0x94, 0x75, 0x50, 0x87, 0xea, 0x97, 0xfb, 0x42, 0xc9, 0x0c,
	0xbd, 0x75, 0x75, 0x3d, 0xb4, 0xbc, 0x1e, 0xaf, 0xed, 0xc1, 0x63, 0xd0, 0x5f, 0x97, 0xc5, 0xe7,
	0xb8, 0x2a, 0xf1, 0x6b, 0x50, 0xeb, 0x64, 0x5f, 0x60, 0xe1, 0xf5, 0x16, 0xb5, 0x15, 0xfc, 0x0c,
	0xe8, 0x27, 0x4a, 0x1d, 0x48, 0x4d, 0xeb, 0xd6, 0x1d, 0xa7, 0xb5, 0x6f, 0xe2, 0xcc, 0xb8, 0x9e,
	0x01, 0x58, 0x81, 0x36, 0xcd, 0x62, 0xde, 0xb6, 0x37, 0x1d, 0xe9, 0xff, 0x26, 0x72, 0xd3, 0x14,
	0x87, 0xcf, 0x40, 0xbb, 0x7a, 0x94, 0xe1, 0x01, 0xe8, 0xe4, 0x65, 0x86, 0x99, 0xf4, 0xa8, 0x7c,
	0xb5, 0xbc, 0xcd, 0x06, 0x1c, 0x81, 0x6e, 0x84, 0x73, 0x9a, 0x91, 0x5c, 0xf9, 0x9b, 0xca, 0x5f,
	0xdf, 0x9a, 0x45, 0xaf, 0x56, 0x03, 0xeb, 0x6a, 0x35, 0xb0, 0xfe, 0x58, 0x0d, 0xac, 0x97, 0x37,
	0x83, 0xc6, 0xd5, 0xcd, 0xa0, 0xf1, 0xdb, 0xcd, 0xa0, 0xf1, 0xed, 0xb3, 0xbf, 0x35, 0xaf, 0xfe,
	0x44, 0x06, 0xe1, 0x51, 0x4c, 0xdd, 0xc5, 0x07, 0x6e, 0x46, 0xa3, 0x32, 0xc5, 0x5c, 0x7f, 0xc8,
	0x8f, 0xaa, 0x2f, 0xf9, 0xc3, 0xf7, 0x8f, 0x36, 0x97, 0x79, 0xbc, 0x31, 0x83, 0x6d, 0x35, 0x91,
	0x8f, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xd6, 0xd5, 0xaa, 0xfd, 0x07, 0x00, 0x00,
}

func (m *ClientState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AllowUpdateAfterMisbehaviour {
		i--
		if m.AllowUpdateAfterMisbehaviour {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if m.AllowUpdateAfterExpiry {
		i--
		if m.AllowUpdateAfterExpiry {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if len(m.UpgradePath) > 0 {
		for iNdEx := len(m.UpgradePath) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.UpgradePath[iNdEx])
			copy(dAtA[i:], m.UpgradePath[iNdEx])
			i = encodeVarintTendermint(dAtA, i, uint64(len(m.UpgradePath[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.ProofSpecs) > 0 {
		for iNdEx := len(m.ProofSpecs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProofSpecs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTendermint(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	{
		size, err := m.LatestHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTendermint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.FrozenHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTendermint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	n3, err3 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.MaxClockDrift, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaxClockDrift):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintTendermint(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x2a
	n4, err4 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.UnbondingPeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.UnbondingPeriod):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintTendermint(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x22
	n5, err5 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.TrustingPeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.TrustingPeriod):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintTendermint(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.TrustLevel.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTendermint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTendermint(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConsensusState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsensusState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConsensusState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NextValidatorsHash) > 0 {
		i -= len(m.NextValidatorsHash)
		copy(dAtA[i:], m.NextValidatorsHash)
		i = encodeVarintTendermint(dAtA, i, uint64(len(m.NextValidatorsHash)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Root.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTendermint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	n8, err8 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err8 != nil {
		return 0, err8
	}
	i -= n8
	i = encodeVarintTendermint(dAtA, i, uint64(n8))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Misbehaviour) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Misbehaviour) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Misbehaviour) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Header2 != nil {
		{
			size, err := m.Header2.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTendermint(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Header1 != nil {
		{
			size, err := m.Header1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTendermint(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintTendermint(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TrustedValidators != nil {
		{
			size, err := m.TrustedValidators.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTendermint(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.TrustedHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTendermint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.ValidatorSet != nil {
		{
			size, err := m.ValidatorSet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTendermint(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.SignedHeader != nil {
		{
			size, err := m.SignedHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTendermint(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Fraction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fraction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fraction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Denominator != 0 {
		i = encodeVarintTendermint(dAtA, i, uint64(m.Denominator))
		i--
		dAtA[i] = 0x10
	}
	if m.Numerator != 0 {
		i = encodeVarintTendermint(dAtA, i, uint64(m.Numerator))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTendermint(dAtA []byte, offset int, v uint64) int {
	offset -= sovTendermint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTendermint(uint64(l))
	}
	l = m.TrustLevel.Size()
	n += 1 + l + sovTendermint(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.TrustingPeriod)
	n += 1 + l + sovTendermint(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.UnbondingPeriod)
	n += 1 + l + sovTendermint(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaxClockDrift)
	n += 1 + l + sovTendermint(uint64(l))
	l = m.FrozenHeight.Size()
	n += 1 + l + sovTendermint(uint64(l))
	l = m.LatestHeight.Size()
	n += 1 + l + sovTendermint(uint64(l))
	if len(m.ProofSpecs) > 0 {
		for _, e := range m.ProofSpecs {
			l = e.Size()
			n += 1 + l + sovTendermint(uint64(l))
		}
	}
	if len(m.UpgradePath) > 0 {
		for _, s := range m.UpgradePath {
			l = len(s)
			n += 1 + l + sovTendermint(uint64(l))
		}
	}
	if m.AllowUpdateAfterExpiry {
		n += 2
	}
	if m.AllowUpdateAfterMisbehaviour {
		n += 2
	}
	return n
}

func (m *ConsensusState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovTendermint(uint64(l))
	l = m.Root.Size()
	n += 1 + l + sovTendermint(uint64(l))
	l = len(m.NextValidatorsHash)
	if l > 0 {
		n += 1 + l + sovTendermint(uint64(l))
	}
	return n
}

func (m *Misbehaviour) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovTendermint(uint64(l))
	}
	if m.Header1 != nil {
		l = m.Header1.Size()
		n += 1 + l + sovTendermint(uint64(l))
	}
	if m.Header2 != nil {
		l = m.Header2.Size()
		n += 1 + l + sovTendermint(uint64(l))
	}
	return n
}

func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignedHeader != nil {
		l = m.SignedHeader.Size()
		n += 1 + l + sovTendermint(uint64(l))
	}
	if m.ValidatorSet != nil {
		l = m.ValidatorSet.Size()
		n += 1 + l + sovTendermint(uint64(l))
	}
	l = m.TrustedHeight.Size()
	n += 1 + l + sovTendermint(uint64(l))
	if m.TrustedValidators != nil {
		l = m.TrustedValidators.Size()
		n += 1 + l + sovTendermint(uint64(l))
	}
	return n
}

func (m *Fraction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Numerator != 0 {
		n += 1 + sovTendermint(uint64(m.Numerator))
	}
	if m.Denominator != 0 {
		n += 1 + sovTendermint(uint64(m.Denominator))
	}
	return n
}

func sovTendermint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTendermint(x uint64) (n int) {
	return sovTendermint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTendermint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClientState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustLevel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TrustLevel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustingPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.TrustingPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.UnbondingPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxClockDrift", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.MaxClockDrift, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrozenHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FrozenHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LatestHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofSpecs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProofSpecs = append(m.ProofSpecs, &_go.ProofSpec{})
			if err := m.ProofSpecs[len(m.ProofSpecs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpgradePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpgradePath = append(m.UpgradePath, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowUpdateAfterExpiry", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowUpdateAfterExpiry = bool(v != 0)
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowUpdateAfterMisbehaviour", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowUpdateAfterMisbehaviour = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTendermint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTendermint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConsensusState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTendermint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConsensusState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsensusState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Root.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextValidatorsHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextValidatorsHash = append(m.NextValidatorsHash[:0], dAtA[iNdEx:postIndex]...)
			if m.NextValidatorsHash == nil {
				m.NextValidatorsHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTendermint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTendermint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Misbehaviour) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTendermint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Misbehaviour: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Misbehaviour: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header1 == nil {
				m.Header1 = &Header{}
			}
			if err := m.Header1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header2 == nil {
				m.Header2 = &Header{}
			}
			if err := m.Header2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTendermint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTendermint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTendermint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SignedHeader == nil {
				m.SignedHeader = &types2.SignedHeader{}
			}
			if err := m.SignedHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ValidatorSet == nil {
				m.ValidatorSet = &types2.ValidatorSet{}
			}
			if err := m.ValidatorSet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TrustedHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedValidators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTendermint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTendermint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TrustedValidators == nil {
				m.TrustedValidators = &types2.ValidatorSet{}
			}
			if err := m.TrustedValidators.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTendermint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTendermint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Fraction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTendermint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Fraction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fraction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Numerator", wireType)
			}
			m.Numerator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Numerator |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denominator", wireType)
			}
			m.Denominator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Denominator |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTendermint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTendermint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTendermint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTendermint
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTendermint
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTendermint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTendermint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTendermint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTendermint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTendermint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTendermint = fmt.Errorf("proto: unexpected end of group")
)
