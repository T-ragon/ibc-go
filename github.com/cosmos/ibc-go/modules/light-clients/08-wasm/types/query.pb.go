// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/lightclients/wasm/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryChecksumsRequest is the request type for the Query/Checksums RPC method.
type QueryChecksumsRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryChecksumsRequest) Reset()         { *m = QueryChecksumsRequest{} }
func (m *QueryChecksumsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryChecksumsRequest) ProtoMessage()    {}
func (*QueryChecksumsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3718a8cb915777, []int{0}
}
func (m *QueryChecksumsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryChecksumsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryChecksumsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryChecksumsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryChecksumsRequest.Merge(m, src)
}
func (m *QueryChecksumsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryChecksumsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryChecksumsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryChecksumsRequest proto.InternalMessageInfo

func (m *QueryChecksumsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryChecksumsResponse is the response type for the Query/Checksums RPC method.
type QueryChecksumsResponse struct {
	// checksums is a list of the hex encoded checksums of all wasm codes stored.
	Checksums []string `protobuf:"bytes,1,rep,name=checksums,proto3" json:"checksums,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryChecksumsResponse) Reset()         { *m = QueryChecksumsResponse{} }
func (m *QueryChecksumsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryChecksumsResponse) ProtoMessage()    {}
func (*QueryChecksumsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3718a8cb915777, []int{1}
}
func (m *QueryChecksumsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryChecksumsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryChecksumsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryChecksumsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryChecksumsResponse.Merge(m, src)
}
func (m *QueryChecksumsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryChecksumsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryChecksumsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryChecksumsResponse proto.InternalMessageInfo

func (m *QueryChecksumsResponse) GetChecksums() []string {
	if m != nil {
		return m.Checksums
	}
	return nil
}

func (m *QueryChecksumsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryCodeRequest is the request type for the Query/Code RPC method.
type QueryCodeRequest struct {
	// checksum is a hex encoded string of the code stored.
	Checksum string `protobuf:"bytes,1,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (m *QueryCodeRequest) Reset()         { *m = QueryCodeRequest{} }
func (m *QueryCodeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCodeRequest) ProtoMessage()    {}
func (*QueryCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3718a8cb915777, []int{2}
}
func (m *QueryCodeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCodeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCodeRequest.Merge(m, src)
}
func (m *QueryCodeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCodeRequest proto.InternalMessageInfo

func (m *QueryCodeRequest) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

// QueryCodeResponse is the response type for the Query/Code RPC method.
type QueryCodeResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *QueryCodeResponse) Reset()         { *m = QueryCodeResponse{} }
func (m *QueryCodeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCodeResponse) ProtoMessage()    {}
func (*QueryCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3718a8cb915777, []int{3}
}
func (m *QueryCodeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCodeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCodeResponse.Merge(m, src)
}
func (m *QueryCodeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCodeResponse proto.InternalMessageInfo

func (m *QueryCodeResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryChecksumsRequest)(nil), "ibc.lightclients.wasm.v1.QueryChecksumsRequest")
	proto.RegisterType((*QueryChecksumsResponse)(nil), "ibc.lightclients.wasm.v1.QueryChecksumsResponse")
	proto.RegisterType((*QueryCodeRequest)(nil), "ibc.lightclients.wasm.v1.QueryCodeRequest")
	proto.RegisterType((*QueryCodeResponse)(nil), "ibc.lightclients.wasm.v1.QueryCodeResponse")
}

func init() {
	proto.RegisterFile("ibc/lightclients/wasm/v1/query.proto", fileDescriptor_9e3718a8cb915777)
}

var fileDescriptor_9e3718a8cb915777 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x3b, 0xf5, 0x2a, 0x66, 0x74, 0xa1, 0x03, 0x4a, 0x09, 0x97, 0x70, 0x89, 0x7f, 0xee,
	0xe5, 0x5e, 0x3a, 0xd3, 0xb4, 0x08, 0x82, 0xae, 0x14, 0x74, 0xab, 0x59, 0xb8, 0x70, 0x23, 0x93,
	0xc9, 0x90, 0x0e, 0x26, 0x99, 0xb4, 0x33, 0xa9, 0x14, 0x11, 0xc1, 0x27, 0x10, 0x5c, 0xea, 0xe3,
	0xb8, 0x70, 0x59, 0x70, 0xe3, 0x52, 0x5a, 0x1f, 0x44, 0x32, 0x93, 0x34, 0xad, 0x58, 0xda, 0xdd,
	0xe4, 0xf0, 0x3b, 0xdf, 0xf7, 0x9d, 0x93, 0x03, 0xef, 0x8a, 0x88, 0x91, 0x54, 0x24, 0x63, 0xcd,
	0x52, 0xc1, 0x73, 0xad, 0xc8, 0x3b, 0xaa, 0x32, 0x32, 0x0b, 0xc8, 0xa4, 0xe4, 0xd3, 0x39, 0x2e,
	0xa6, 0x52, 0x4b, 0xd4, 0x13, 0x11, 0xc3, 0x9b, 0x14, 0xae, 0x28, 0x3c, 0x0b, 0xdc, 0xe3, 0x44,
	0xca, 0x24, 0xe5, 0x84, 0x16, 0x82, 0xd0, 0x3c, 0x97, 0x9a, 0x6a, 0x21, 0x73, 0x65, 0xfb, 0xdc,
	0x73, 0x26, 0x55, 0x26, 0x15, 0x89, 0xa8, 0xe2, 0x56, 0x90, 0xcc, 0x82, 0x88, 0x6b, 0x1a, 0x90,
	0x82, 0x26, 0x22, 0x37, 0xb0, 0x65, 0xfd, 0x37, 0xf0, 0xd6, 0xcb, 0x8a, 0x78, 0x3a, 0xe6, 0xec,
	0xad, 0x2a, 0x33, 0x15, 0xf2, 0x49, 0xc9, 0x95, 0x46, 0xcf, 0x20, 0x6c, 0xe1, 0x1e, 0x38, 0x01,
	0x67, 0xd7, 0x86, 0xf7, 0xb1, 0x55, 0xc6, 0x95, 0x32, 0xb6, 0x51, 0x6b, 0x65, 0xfc, 0x82, 0x26,
	0xbc, 0xee, 0x0d, 0x37, 0x3a, 0xfd, 0x8f, 0xf0, 0xf6, 0xbf, 0x06, 0xaa, 0x90, 0xb9, 0xe2, 0xe8,
	0x18, 0x3a, 0xac, 0x29, 0xf6, 0xc0, 0xc9, 0xa5, 0x33, 0x27, 0x6c, 0x0b, 0xe8, 0xf9, 0x96, 0x7f,
	0xd7, 0xf8, 0x9f, 0xee, 0xf5, 0xb7, 0xd2, 0x5b, 0x01, 0x30, 0xbc, 0x61, 0x03, 0xc8, 0xb8, 0x09,
	0x88, 0x5c, 0x78, 0xb5, 0x71, 0x32, 0xa3, 0x39, 0xe1, 0xfa, 0xdb, 0x3f, 0x85, 0x37, 0x37, 0xf8,
	0x3a, 0x2b, 0x82, 0x47, 0x31, 0xd5, 0xd4, 0xc0, 0xd7, 0x43, 0xf3, 0x1e, 0x7e, 0xef, 0xc2, 0xcb,
	0x86, 0x44, 0x5f, 0x01, 0x74, 0xd6, 0xf3, 0x21, 0x82, 0x77, 0xfd, 0x37, 0xfc, 0xdf, 0x55, 0xbb,
	0x83, 0xc3, 0x1b, 0x6c, 0x1c, 0xff, 0xe2, 0xd3, 0xcf, 0x3f, 0x5f, 0xba, 0xf7, 0xd0, 0x1d, 0xb2,
	0xf3, 0x90, 0xda, 0x4d, 0x7e, 0x03, 0xf0, 0xa8, 0x1a, 0x06, 0x9d, 0xef, 0xf3, 0x69, 0x37, 0xe4,
	0x5e, 0x1c, 0xc4, 0xd6, 0x71, 0x1e, 0x99, 0x38, 0x0f, 0xd0, 0xe8, 0x80, 0x38, 0xe4, 0x7d, 0xf3,
	0xfc, 0x40, 0x98, 0x8c, 0xf9, 0x93, 0x57, 0x3f, 0x96, 0x1e, 0x58, 0x2c, 0x3d, 0xf0, 0x7b, 0xe9,
	0x81, 0xcf, 0x2b, 0xaf, 0xb3, 0x58, 0x79, 0x9d, 0x5f, 0x2b, 0xaf, 0xf3, 0xfa, 0x71, 0x22, 0xf4,
	0xb8, 0x8c, 0x30, 0x93, 0x19, 0xa9, 0x4f, 0x5a, 0x44, 0xac, 0x9f, 0x48, 0x92, 0xc9, 0xb8, 0x4c,
	0xb9, 0xb2, 0x56, 0xfd, 0xc6, 0x6b, 0xf0, 0xb0, 0x6f, 0xec, 0xf4, 0xbc, 0xe0, 0x2a, 0xba, 0x62,
	0x0e, 0x7c, 0xf4, 0x37, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x82, 0x4a, 0xb3, 0x6c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Get all Wasm checksums
	Checksums(ctx context.Context, in *QueryChecksumsRequest, opts ...grpc.CallOption) (*QueryChecksumsResponse, error)
	// Get Wasm code for given checksum
	Code(ctx context.Context, in *QueryCodeRequest, opts ...grpc.CallOption) (*QueryCodeResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Checksums(ctx context.Context, in *QueryChecksumsRequest, opts ...grpc.CallOption) (*QueryChecksumsResponse, error) {
	out := new(QueryChecksumsResponse)
	err := c.cc.Invoke(ctx, "/ibc.lightclients.wasm.v1.Query/Checksums", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Code(ctx context.Context, in *QueryCodeRequest, opts ...grpc.CallOption) (*QueryCodeResponse, error) {
	out := new(QueryCodeResponse)
	err := c.cc.Invoke(ctx, "/ibc.lightclients.wasm.v1.Query/Code", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Get all Wasm checksums
	Checksums(context.Context, *QueryChecksumsRequest) (*QueryChecksumsResponse, error)
	// Get Wasm code for given checksum
	Code(context.Context, *QueryCodeRequest) (*QueryCodeResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Checksums(ctx context.Context, req *QueryChecksumsRequest) (*QueryChecksumsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checksums not implemented")
}
func (*UnimplementedQueryServer) Code(ctx context.Context, req *QueryCodeRequest) (*QueryCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Code not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Checksums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryChecksumsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Checksums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.lightclients.wasm.v1.Query/Checksums",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Checksums(ctx, req.(*QueryChecksumsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Code_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Code(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.lightclients.wasm.v1.Query/Code",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Code(ctx, req.(*QueryCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.lightclients.wasm.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Checksums",
			Handler:    _Query_Checksums_Handler,
		},
		{
			MethodName: "Code",
			Handler:    _Query_Code_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/lightclients/wasm/v1/query.proto",
}

func (m *QueryChecksumsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryChecksumsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryChecksumsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryChecksumsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryChecksumsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryChecksumsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Checksums) > 0 {
		for iNdEx := len(m.Checksums) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Checksums[iNdEx])
			copy(dAtA[i:], m.Checksums[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Checksums[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryCodeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCodeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCodeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Checksum) > 0 {
		i -= len(m.Checksum)
		copy(dAtA[i:], m.Checksum)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Checksum)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCodeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCodeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCodeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryChecksumsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryChecksumsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Checksums) > 0 {
		for _, s := range m.Checksums {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCodeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCodeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryChecksumsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryChecksumsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryChecksumsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryChecksumsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryChecksumsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryChecksumsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksums", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksums = append(m.Checksums, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCodeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCodeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCodeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCodeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCodeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCodeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
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
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)