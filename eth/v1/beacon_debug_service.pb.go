// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eth/v1/beacon_debug_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ForkChoiceHeadsResponse struct {
	Data                 []*ForkChoiceHead `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ForkChoiceHeadsResponse) Reset()         { *m = ForkChoiceHeadsResponse{} }
func (m *ForkChoiceHeadsResponse) String() string { return proto.CompactTextString(m) }
func (*ForkChoiceHeadsResponse) ProtoMessage()    {}
func (*ForkChoiceHeadsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c4891f2a5dc2c4a, []int{0}
}
func (m *ForkChoiceHeadsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ForkChoiceHeadsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ForkChoiceHeadsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ForkChoiceHeadsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForkChoiceHeadsResponse.Merge(m, src)
}
func (m *ForkChoiceHeadsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ForkChoiceHeadsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ForkChoiceHeadsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ForkChoiceHeadsResponse proto.InternalMessageInfo

func (m *ForkChoiceHeadsResponse) GetData() []*ForkChoiceHead {
	if m != nil {
		return m.Data
	}
	return nil
}

type ForkChoiceHead struct {
	Root                 []byte   `protobuf:"bytes,1,opt,name=root,proto3" json:"root,omitempty" ssz-size:"32"`
	Slot                 uint64   `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForkChoiceHead) Reset()         { *m = ForkChoiceHead{} }
func (m *ForkChoiceHead) String() string { return proto.CompactTextString(m) }
func (*ForkChoiceHead) ProtoMessage()    {}
func (*ForkChoiceHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c4891f2a5dc2c4a, []int{1}
}
func (m *ForkChoiceHead) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ForkChoiceHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ForkChoiceHead.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ForkChoiceHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForkChoiceHead.Merge(m, src)
}
func (m *ForkChoiceHead) XXX_Size() int {
	return m.Size()
}
func (m *ForkChoiceHead) XXX_DiscardUnknown() {
	xxx_messageInfo_ForkChoiceHead.DiscardUnknown(m)
}

var xxx_messageInfo_ForkChoiceHead proto.InternalMessageInfo

func (m *ForkChoiceHead) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *ForkChoiceHead) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

type BeaconStateResponse struct {
	Data                 *BeaconState `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BeaconStateResponse) Reset()         { *m = BeaconStateResponse{} }
func (m *BeaconStateResponse) String() string { return proto.CompactTextString(m) }
func (*BeaconStateResponse) ProtoMessage()    {}
func (*BeaconStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c4891f2a5dc2c4a, []int{2}
}
func (m *BeaconStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BeaconStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BeaconStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BeaconStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconStateResponse.Merge(m, src)
}
func (m *BeaconStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *BeaconStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconStateResponse proto.InternalMessageInfo

func (m *BeaconStateResponse) GetData() *BeaconState {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ForkChoiceHeadsResponse)(nil), "ethereum.eth.v1.ForkChoiceHeadsResponse")
	proto.RegisterType((*ForkChoiceHead)(nil), "ethereum.eth.v1.ForkChoiceHead")
	proto.RegisterType((*BeaconStateResponse)(nil), "ethereum.eth.v1.BeaconStateResponse")
}

func init() { proto.RegisterFile("eth/v1/beacon_debug_service.proto", fileDescriptor_5c4891f2a5dc2c4a) }

var fileDescriptor_5c4891f2a5dc2c4a = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x8b, 0x13, 0x31,
	0x14, 0xc7, 0x49, 0x2d, 0x1e, 0x52, 0xdd, 0xd5, 0x14, 0xb4, 0x8e, 0xb5, 0x5b, 0x07, 0x95, 0x39,
	0xb8, 0x89, 0x6d, 0x6f, 0x1e, 0xab, 0x75, 0x05, 0x45, 0x96, 0x7a, 0x93, 0x85, 0x92, 0x99, 0x3e,
	0x67, 0x82, 0xed, 0x64, 0x9c, 0xbc, 0x29, 0xec, 0x2e, 0x5e, 0x04, 0xf1, 0xae, 0xff, 0x85, 0x7f,
	0x89, 0x78, 0x12, 0xbc, 0x8b, 0x14, 0xff, 0x02, 0xff, 0x02, 0x99, 0x34, 0xad, 0x6d, 0x57, 0xf4,
	0x94, 0xe4, 0xfd, 0xfc, 0xbc, 0xf7, 0x0d, 0xbd, 0x09, 0x98, 0x88, 0x59, 0x47, 0x84, 0x20, 0x23,
	0x9d, 0x8e, 0xc6, 0x10, 0x16, 0xf1, 0xc8, 0x40, 0x3e, 0x53, 0x11, 0xf0, 0x2c, 0xd7, 0xa8, 0xd9,
	0x2e, 0x60, 0x02, 0x39, 0x14, 0x53, 0x0e, 0x98, 0xf0, 0x59, 0xc7, 0xdb, 0x8f, 0x15, 0x26, 0x45,
	0xc8, 0x23, 0x3d, 0x15, 0xb1, 0x8e, 0xb5, 0xb0, 0x71, 0x61, 0xf1, 0xd2, 0xbe, 0xec, 0xc3, 0xde,
	0x16, 0xf9, 0x5e, 0x33, 0xd6, 0x3a, 0x9e, 0x80, 0x90, 0x99, 0x12, 0x32, 0x4d, 0x35, 0x4a, 0x54,
	0x3a, 0x35, 0xce, 0x7b, 0xdd, 0x79, 0x57, 0x35, 0x60, 0x9a, 0xe1, 0xb1, 0x73, 0x5e, 0xdb, 0xa4,
	0x33, 0x28, 0xd1, 0x51, 0x79, 0x5b, 0xe0, 0x51, 0x22, 0x55, 0xba, 0x09, 0xee, 0x3f, 0xa3, 0x57,
	0x1f, 0xe9, 0xfc, 0xd5, 0x83, 0x44, 0xab, 0x08, 0x1e, 0x83, 0x1c, 0x9b, 0x21, 0x98, 0x4c, 0xa7,
	0x06, 0x58, 0x8f, 0x56, 0xc7, 0x12, 0x65, 0x83, 0xb4, 0xcf, 0x05, 0xb5, 0xee, 0x1e, 0xdf, 0x1a,
	0x91, 0x6f, 0xe6, 0x0d, 0x6d, 0xb0, 0xff, 0x84, 0xee, 0x6c, 0xda, 0xd9, 0x6d, 0x5a, 0xcd, 0xb5,
	0xc6, 0x06, 0x69, 0x93, 0xe0, 0x42, 0xff, 0xf2, 0xaf, 0xef, 0x7b, 0x17, 0x8d, 0x39, 0xd9, 0x37,
	0xea, 0x04, 0xee, 0xfb, 0xbd, 0xae, 0x3f, 0xb4, 0x6e, 0xc6, 0x68, 0xd5, 0x4c, 0x34, 0x36, 0x2a,
	0x6d, 0x12, 0x54, 0x87, 0xf6, 0xee, 0x1f, 0xd0, 0x7a, 0xdf, 0xa2, 0x3f, 0x2f, 0x87, 0x5a, 0x81,
	0xdd, 0x5b, 0x81, 0x91, 0xa0, 0xd6, 0x6d, 0x9e, 0x01, 0x5b, 0xcf, 0xb1, 0x91, 0xdd, 0x0f, 0x15,
	0x5a, 0x5b, 0x58, 0x1f, 0x96, 0xe2, 0xb1, 0x77, 0x84, 0xee, 0x1c, 0x00, 0xae, 0x05, 0xb2, 0x1b,
	0x67, 0xca, 0xb8, 0xa6, 0xaf, 0x0b, 0x30, 0xe8, 0xdd, 0xfa, 0x67, 0x17, 0x47, 0xe6, 0xf3, 0xb7,
	0xdf, 0x7e, 0x7e, 0xac, 0x04, 0xec, 0x8e, 0x70, 0x9b, 0xb7, 0x7f, 0xc5, 0xed, 0x5f, 0x58, 0x69,
	0x8c, 0x38, 0xb5, 0xe7, 0x48, 0x8d, 0xdf, 0xb0, 0x53, 0x5a, 0x7f, 0xaa, 0x0c, 0x6e, 0x29, 0xc0,
	0xae, 0xf0, 0x85, 0xe0, 0x7c, 0x29, 0x38, 0x1f, 0x94, 0x82, 0x7b, 0xc1, 0x7f, 0x34, 0x58, 0x69,
	0xe7, 0xfb, 0x16, 0xa4, 0xc9, 0xbc, 0xbf, 0x82, 0x24, 0x65, 0x6c, 0xff, 0x3d, 0xf9, 0x3c, 0x6f,
	0x91, 0xaf, 0xf3, 0x16, 0xf9, 0x31, 0x6f, 0x11, 0x5a, 0xd7, 0x79, 0xbc, 0x5d, 0xbf, 0x7f, 0x69,
	0x6d, 0x6b, 0x87, 0x25, 0xcc, 0x21, 0x79, 0x71, 0x77, 0xed, 0x6b, 0x67, 0xf9, 0xb1, 0x99, 0x4a,
	0x54, 0xd1, 0x44, 0x86, 0x46, 0x2c, 0xb3, 0x65, 0xa6, 0x8c, 0x6b, 0xfc, 0xa9, 0xb2, 0x3b, 0x58,
	0xd6, 0x1c, 0xd8, 0x9a, 0x5f, 0xfe, 0x58, 0x8e, 0x06, 0x98, 0x1c, 0xcd, 0x3a, 0xe1, 0x79, 0x3b,
	0x67, 0xef, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x49, 0x5d, 0xbd, 0x78, 0x69, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BeaconDebugClient is the client API for BeaconDebug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BeaconDebugClient interface {
	GetBeaconState(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*BeaconStateResponse, error)
	ListForkChoiceHeads(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*ForkChoiceHeadsResponse, error)
}

type beaconDebugClient struct {
	cc *grpc.ClientConn
}

func NewBeaconDebugClient(cc *grpc.ClientConn) BeaconDebugClient {
	return &beaconDebugClient{cc}
}

func (c *beaconDebugClient) GetBeaconState(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*BeaconStateResponse, error) {
	out := new(BeaconStateResponse)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1.BeaconDebug/GetBeaconState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconDebugClient) ListForkChoiceHeads(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*ForkChoiceHeadsResponse, error) {
	out := new(ForkChoiceHeadsResponse)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1.BeaconDebug/ListForkChoiceHeads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeaconDebugServer is the server API for BeaconDebug service.
type BeaconDebugServer interface {
	GetBeaconState(context.Context, *StateRequest) (*BeaconStateResponse, error)
	ListForkChoiceHeads(context.Context, *types.Empty) (*ForkChoiceHeadsResponse, error)
}

// UnimplementedBeaconDebugServer can be embedded to have forward compatible implementations.
type UnimplementedBeaconDebugServer struct {
}

func (*UnimplementedBeaconDebugServer) GetBeaconState(ctx context.Context, req *StateRequest) (*BeaconStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBeaconState not implemented")
}
func (*UnimplementedBeaconDebugServer) ListForkChoiceHeads(ctx context.Context, req *types.Empty) (*ForkChoiceHeadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListForkChoiceHeads not implemented")
}

func RegisterBeaconDebugServer(s *grpc.Server, srv BeaconDebugServer) {
	s.RegisterService(&_BeaconDebug_serviceDesc, srv)
}

func _BeaconDebug_GetBeaconState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconDebugServer).GetBeaconState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1.BeaconDebug/GetBeaconState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconDebugServer).GetBeaconState(ctx, req.(*StateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconDebug_ListForkChoiceHeads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconDebugServer).ListForkChoiceHeads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1.BeaconDebug/ListForkChoiceHeads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconDebugServer).ListForkChoiceHeads(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _BeaconDebug_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethereum.eth.v1.BeaconDebug",
	HandlerType: (*BeaconDebugServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBeaconState",
			Handler:    _BeaconDebug_GetBeaconState_Handler,
		},
		{
			MethodName: "ListForkChoiceHeads",
			Handler:    _BeaconDebug_ListForkChoiceHeads_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eth/v1/beacon_debug_service.proto",
}

func (m *ForkChoiceHeadsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ForkChoiceHeadsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ForkChoiceHeadsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBeaconDebugService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ForkChoiceHead) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ForkChoiceHead) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ForkChoiceHead) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Slot != 0 {
		i = encodeVarintBeaconDebugService(dAtA, i, uint64(m.Slot))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Root) > 0 {
		i -= len(m.Root)
		copy(dAtA[i:], m.Root)
		i = encodeVarintBeaconDebugService(dAtA, i, uint64(len(m.Root)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BeaconStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BeaconStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BeaconStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBeaconDebugService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBeaconDebugService(dAtA []byte, offset int, v uint64) int {
	offset -= sovBeaconDebugService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ForkChoiceHeadsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovBeaconDebugService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ForkChoiceHead) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Root)
	if l > 0 {
		n += 1 + l + sovBeaconDebugService(uint64(l))
	}
	if m.Slot != 0 {
		n += 1 + sovBeaconDebugService(uint64(m.Slot))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BeaconStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovBeaconDebugService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBeaconDebugService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBeaconDebugService(x uint64) (n int) {
	return sovBeaconDebugService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ForkChoiceHeadsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBeaconDebugService
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
			return fmt.Errorf("proto: ForkChoiceHeadsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForkChoiceHeadsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBeaconDebugService
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
				return ErrInvalidLengthBeaconDebugService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBeaconDebugService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &ForkChoiceHead{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBeaconDebugService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBeaconDebugService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBeaconDebugService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ForkChoiceHead) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBeaconDebugService
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
			return fmt.Errorf("proto: ForkChoiceHead: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForkChoiceHead: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBeaconDebugService
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
				return ErrInvalidLengthBeaconDebugService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBeaconDebugService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = append(m.Root[:0], dAtA[iNdEx:postIndex]...)
			if m.Root == nil {
				m.Root = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slot", wireType)
			}
			m.Slot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBeaconDebugService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Slot |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBeaconDebugService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBeaconDebugService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBeaconDebugService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BeaconStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBeaconDebugService
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
			return fmt.Errorf("proto: BeaconStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BeaconStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBeaconDebugService
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
				return ErrInvalidLengthBeaconDebugService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBeaconDebugService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &BeaconState{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBeaconDebugService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBeaconDebugService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBeaconDebugService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBeaconDebugService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBeaconDebugService
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
					return 0, ErrIntOverflowBeaconDebugService
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
					return 0, ErrIntOverflowBeaconDebugService
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
				return 0, ErrInvalidLengthBeaconDebugService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBeaconDebugService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBeaconDebugService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBeaconDebugService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBeaconDebugService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBeaconDebugService = fmt.Errorf("proto: unexpected end of group")
)
