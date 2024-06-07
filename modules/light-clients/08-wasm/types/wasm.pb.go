// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/lightclients/wasm/v1/wasm.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/T-ragon/ibc-go/v9/modules/core/02-client/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Wasm light client's Client state
type ClientState struct {
	// bytes encoding the client state of the underlying light client
	// implemented as a Wasm contract.
	Data         []byte       `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Checksum     []byte       `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
	LatestHeight types.Height `protobuf:"bytes,3,opt,name=latest_height,json=latestHeight,proto3" json:"latest_height"`
}

func (m *ClientState) Reset()         { *m = ClientState{} }
func (m *ClientState) String() string { return proto.CompactTextString(m) }
func (*ClientState) ProtoMessage()    {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_678928ebbdee1807, []int{0}
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

// Wasm light client's ConsensusState
type ConsensusState struct {
	// bytes encoding the consensus state of the underlying light client
	// implemented as a Wasm contract.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ConsensusState) Reset()         { *m = ConsensusState{} }
func (m *ConsensusState) String() string { return proto.CompactTextString(m) }
func (*ConsensusState) ProtoMessage()    {}
func (*ConsensusState) Descriptor() ([]byte, []int) {
	return fileDescriptor_678928ebbdee1807, []int{1}
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

// Wasm light client message (either header(s) or misbehaviour)
type ClientMessage struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ClientMessage) Reset()         { *m = ClientMessage{} }
func (m *ClientMessage) String() string { return proto.CompactTextString(m) }
func (*ClientMessage) ProtoMessage()    {}
func (*ClientMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_678928ebbdee1807, []int{2}
}
func (m *ClientMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientMessage.Merge(m, src)
}
func (m *ClientMessage) XXX_Size() int {
	return m.Size()
}
func (m *ClientMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ClientMessage proto.InternalMessageInfo

// Checksums defines a list of all checksums that are stored
//
// Deprecated: This message is deprecated in favor of storing the checksums
// using a Collections.KeySet.
//
// Deprecated: Do not use.
type Checksums struct {
	Checksums [][]byte `protobuf:"bytes,1,rep,name=checksums,proto3" json:"checksums,omitempty"`
}

func (m *Checksums) Reset()         { *m = Checksums{} }
func (m *Checksums) String() string { return proto.CompactTextString(m) }
func (*Checksums) ProtoMessage()    {}
func (*Checksums) Descriptor() ([]byte, []int) {
	return fileDescriptor_678928ebbdee1807, []int{3}
}
func (m *Checksums) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Checksums) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Checksums.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Checksums) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checksums.Merge(m, src)
}
func (m *Checksums) XXX_Size() int {
	return m.Size()
}
func (m *Checksums) XXX_DiscardUnknown() {
	xxx_messageInfo_Checksums.DiscardUnknown(m)
}

var xxx_messageInfo_Checksums proto.InternalMessageInfo

func (m *Checksums) GetChecksums() [][]byte {
	if m != nil {
		return m.Checksums
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientState)(nil), "ibc.lightclients.wasm.v1.ClientState")
	proto.RegisterType((*ConsensusState)(nil), "ibc.lightclients.wasm.v1.ConsensusState")
	proto.RegisterType((*ClientMessage)(nil), "ibc.lightclients.wasm.v1.ClientMessage")
	proto.RegisterType((*Checksums)(nil), "ibc.lightclients.wasm.v1.Checksums")
}

func init() {
	proto.RegisterFile("ibc/lightclients/wasm/v1/wasm.proto", fileDescriptor_678928ebbdee1807)
}

var fileDescriptor_678928ebbdee1807 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe3, 0xb6, 0x42, 0xd4, 0x6d, 0x19, 0x22, 0x86, 0x28, 0x42, 0x69, 0x55, 0x96, 0x82,
	0x14, 0x9b, 0xc2, 0x82, 0x2a, 0xa6, 0x56, 0x48, 0x2c, 0x2c, 0x45, 0x62, 0x60, 0x41, 0x89, 0x6b,
	0x25, 0x16, 0x49, 0x5d, 0xf5, 0x9c, 0x22, 0xde, 0x00, 0x31, 0xf1, 0x08, 0x3c, 0x4e, 0xc7, 0x8e,
	0x4c, 0x08, 0xb5, 0x2f, 0x82, 0x6c, 0xa7, 0xc0, 0x02, 0x53, 0x2e, 0xbf, 0x3f, 0xdf, 0x7d, 0xf2,
	0xe1, 0x43, 0x11, 0x33, 0x9a, 0x89, 0x24, 0x55, 0x2c, 0x13, 0x7c, 0xaa, 0x80, 0x3e, 0x46, 0x90,
	0xd3, 0x45, 0xdf, 0x7c, 0xc9, 0x6c, 0x2e, 0x95, 0x74, 0x3d, 0x11, 0x33, 0xf2, 0x1b, 0x22, 0xe6,
	0x70, 0xd1, 0xf7, 0xf7, 0x13, 0x99, 0x48, 0x03, 0x51, 0x5d, 0x59, 0xde, 0x6f, 0xeb, 0xa6, 0x4c,
	0xce, 0x39, 0xb5, 0xbc, 0x6e, 0x67, 0x2b, 0x0b, 0x74, 0x5f, 0x10, 0x6e, 0x8c, 0x4c, 0x70, 0xa3,
	0x22, 0xc5, 0x5d, 0x17, 0xd7, 0x26, 0x91, 0x8a, 0x3c, 0xd4, 0x41, 0xbd, 0xe6, 0xd8, 0xd4, 0xae,
	0x8f, 0x77, 0x59, 0xca, 0xd9, 0x03, 0x14, 0xb9, 0x57, 0x31, 0xf9, 0xf7, 0xbf, 0x7b, 0x89, 0x5b,
	0x59, 0xa4, 0x38, 0xa8, 0xfb, 0x94, 0x6b, 0x2d, 0xaf, 0xda, 0x41, 0xbd, 0xc6, 0xa9, 0x4f, 0xb4,
	0xa8, 0x1e, 0x4c, 0xca, 0x71, 0x8b, 0x3e, 0xb9, 0x32, 0xc4, 0xb0, 0xb6, 0xfc, 0x68, 0x3b, 0xe3,
	0xa6, 0xbd, 0x66, 0xb3, 0x41, 0xed, 0xf9, 0xad, 0xed, 0x74, 0x8f, 0xf1, 0xde, 0x48, 0x4e, 0x81,
	0x4f, 0xa1, 0x80, 0x3f, 0x75, 0x4a, 0xf6, 0x08, 0xb7, 0xac, 0xf7, 0x35, 0x07, 0x88, 0x92, 0xff,
	0xd0, 0x10, 0xd7, 0x47, 0xa5, 0x2f, 0xb8, 0x07, 0xb8, 0xbe, 0x95, 0x07, 0x0f, 0x75, 0xaa, 0xbd,
	0xe6, 0xf8, 0x27, 0x18, 0x54, 0x3c, 0x34, 0xbc, 0x5d, 0xae, 0x03, 0xb4, 0x5a, 0x07, 0xe8, 0x73,
	0x1d, 0xa0, 0xd7, 0x4d, 0xe0, 0xac, 0x36, 0x81, 0xf3, 0xbe, 0x09, 0x9c, 0xbb, 0x8b, 0x44, 0xa8,
	0xb4, 0x88, 0x09, 0x93, 0x39, 0x65, 0x12, 0x72, 0x09, 0x54, 0xc4, 0x2c, 0x4c, 0x24, 0xcd, 0xe5,
	0xa4, 0xc8, 0x38, 0xd8, 0xfd, 0x85, 0xdb, 0x05, 0x9e, 0x9c, 0x87, 0x66, 0x87, 0xea, 0x69, 0xc6,
	0x21, 0xde, 0x31, 0x2f, 0x7e, 0xf6, 0x15, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x78, 0x4d, 0x63, 0xe9,
	0x01, 0x00, 0x00,
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
	{
		size, err := m.LatestHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintWasm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Checksum) > 0 {
		i -= len(m.Checksum)
		copy(dAtA[i:], m.Checksum)
		i = encodeVarintWasm(dAtA, i, uint64(len(m.Checksum)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintWasm(dAtA, i, uint64(len(m.Data)))
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
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintWasm(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClientMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintWasm(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Checksums) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Checksums) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Checksums) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Checksums) > 0 {
		for iNdEx := len(m.Checksums) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Checksums[iNdEx])
			copy(dAtA[i:], m.Checksums[iNdEx])
			i = encodeVarintWasm(dAtA, i, uint64(len(m.Checksums[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintWasm(dAtA []byte, offset int, v uint64) int {
	offset -= sovWasm(v)
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
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovWasm(uint64(l))
	}
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovWasm(uint64(l))
	}
	l = m.LatestHeight.Size()
	n += 1 + l + sovWasm(uint64(l))
	return n
}

func (m *ConsensusState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovWasm(uint64(l))
	}
	return n
}

func (m *ClientMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovWasm(uint64(l))
	}
	return n
}

func (m *Checksums) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Checksums) > 0 {
		for _, b := range m.Checksums {
			l = len(b)
			n += 1 + l + sovWasm(uint64(l))
		}
	}
	return n
}

func sovWasm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWasm(x uint64) (n int) {
	return sovWasm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWasm
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
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
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
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
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
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = append(m.Checksum[:0], dAtA[iNdEx:postIndex]...)
			if m.Checksum == nil {
				m.Checksum = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
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
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LatestHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWasm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWasm
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
				return ErrIntOverflowWasm
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
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
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
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWasm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWasm
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
func (m *ClientMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWasm
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
			return fmt.Errorf("proto: ClientMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
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
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWasm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWasm
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
func (m *Checksums) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWasm
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
			return fmt.Errorf("proto: Checksums: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Checksums: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksums", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
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
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksums = append(m.Checksums, make([]byte, postIndex-iNdEx))
			copy(m.Checksums[len(m.Checksums)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWasm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWasm
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
func skipWasm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWasm
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
					return 0, ErrIntOverflowWasm
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
					return 0, ErrIntOverflowWasm
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
				return 0, ErrInvalidLengthWasm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWasm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWasm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWasm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWasm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWasm = fmt.Errorf("proto: unexpected end of group")
)
