// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/authz/v1beta1/authz.proto

package authz

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	any "github.com/cosmos/gogoproto/types/any"
	_ "github.com/forbole/juno/v5/cosmos-sdk/types/tx/amino"
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

// GenericAuthorization gives the grantee unrestricted permissions to execute
// the provided method on behalf of the granter's account.
type GenericAuthorization struct {
	// Msg, identified by it's type URL, to grant unrestricted permissions to
	// execute
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *GenericAuthorization) Reset()         { *m = GenericAuthorization{} }
func (m *GenericAuthorization) String() string { return proto.CompactTextString(m) }
func (*GenericAuthorization) ProtoMessage()    {}
func (*GenericAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_544dc2e84b61c637, []int{0}
}
func (m *GenericAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenericAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenericAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenericAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericAuthorization.Merge(m, src)
}
func (m *GenericAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *GenericAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_GenericAuthorization proto.InternalMessageInfo

// Grant gives permissions to execute
// the provide method with expiration time.
type Grant struct {
	Authorization *any.Any `protobuf:"bytes,1,opt,name=authorization,proto3" json:"authorization,omitempty"`
	// time when the grant will expire and will be pruned. If null, then the grant
	// doesn't have a time expiration (other conditions  in `authorization`
	// may apply to invalidate the grant)
	Expiration *time.Time `protobuf:"bytes,2,opt,name=expiration,proto3,stdtime" json:"expiration,omitempty"`
}

func (m *Grant) Reset()         { *m = Grant{} }
func (m *Grant) String() string { return proto.CompactTextString(m) }
func (*Grant) ProtoMessage()    {}
func (*Grant) Descriptor() ([]byte, []int) {
	return fileDescriptor_544dc2e84b61c637, []int{1}
}
func (m *Grant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Grant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Grant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Grant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grant.Merge(m, src)
}
func (m *Grant) XXX_Size() int {
	return m.Size()
}
func (m *Grant) XXX_DiscardUnknown() {
	xxx_messageInfo_Grant.DiscardUnknown(m)
}

var xxx_messageInfo_Grant proto.InternalMessageInfo

// GrantAuthorization extends a grant with both the addresses of the grantee and
// granter. It is used in genesis.proto and query.proto
type GrantAuthorization struct {
	Granter       string     `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee       string     `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	Authorization *any.Any   `protobuf:"bytes,3,opt,name=authorization,proto3" json:"authorization,omitempty"`
	Expiration    *time.Time `protobuf:"bytes,4,opt,name=expiration,proto3,stdtime" json:"expiration,omitempty"`
}

func (m *GrantAuthorization) Reset()         { *m = GrantAuthorization{} }
func (m *GrantAuthorization) String() string { return proto.CompactTextString(m) }
func (*GrantAuthorization) ProtoMessage()    {}
func (*GrantAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_544dc2e84b61c637, []int{2}
}
func (m *GrantAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrantAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrantAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GrantAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrantAuthorization.Merge(m, src)
}
func (m *GrantAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *GrantAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_GrantAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_GrantAuthorization proto.InternalMessageInfo

// GrantQueueItem contains the list of TypeURL of a sdk.Msg.
type GrantQueueItem struct {
	// msg_type_urls contains the list of TypeURL of a sdk.Msg.
	MsgTypeUrls []string `protobuf:"bytes,1,rep,name=msg_type_urls,json=msgTypeUrls,proto3" json:"msg_type_urls,omitempty"`
}

func (m *GrantQueueItem) Reset()         { *m = GrantQueueItem{} }
func (m *GrantQueueItem) String() string { return proto.CompactTextString(m) }
func (*GrantQueueItem) ProtoMessage()    {}
func (*GrantQueueItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_544dc2e84b61c637, []int{3}
}
func (m *GrantQueueItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrantQueueItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrantQueueItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GrantQueueItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrantQueueItem.Merge(m, src)
}
func (m *GrantQueueItem) XXX_Size() int {
	return m.Size()
}
func (m *GrantQueueItem) XXX_DiscardUnknown() {
	xxx_messageInfo_GrantQueueItem.DiscardUnknown(m)
}

var xxx_messageInfo_GrantQueueItem proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenericAuthorization)(nil), "cosmos.authz.v1beta1.GenericAuthorization")
	proto.RegisterType((*Grant)(nil), "cosmos.authz.v1beta1.Grant")
	proto.RegisterType((*GrantAuthorization)(nil), "cosmos.authz.v1beta1.GrantAuthorization")
	proto.RegisterType((*GrantQueueItem)(nil), "cosmos.authz.v1beta1.GrantQueueItem")
}

func init() { proto.RegisterFile("cosmos/authz/v1beta1/authz.proto", fileDescriptor_544dc2e84b61c637) }

var fileDescriptor_544dc2e84b61c637 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xbb, 0x6e, 0x13, 0x41,
	0x14, 0x86, 0x3d, 0x71, 0xb8, 0x64, 0xa2, 0x20, 0x58, 0xb9, 0x30, 0x2e, 0xd6, 0xd6, 0x16, 0x28,
	0x42, 0xf2, 0x8e, 0x12, 0xa0, 0xa1, 0xc2, 0x16, 0x52, 0x04, 0x54, 0x2c, 0xa1, 0xa1, 0xb1, 0x66,
	0xed, 0x93, 0xf1, 0x80, 0x67, 0x66, 0x35, 0x97, 0x28, 0xce, 0x23, 0x50, 0xe5, 0x19, 0x78, 0x02,
	0x90, 0xf2, 0x10, 0x16, 0x55, 0x44, 0x45, 0xc5, 0xc5, 0x2e, 0x78, 0x0d, 0xe4, 0x19, 0xaf, 0xb0,
	0x71, 0x24, 0x52, 0xd0, 0x58, 0x73, 0xe6, 0xfc, 0xff, 0x39, 0xbf, 0xbf, 0xdd, 0xc5, 0xad, 0xbe,
	0x32, 0x42, 0x19, 0x42, 0x9d, 0x1d, 0x9e, 0x92, 0xe3, 0xbd, 0x1c, 0x2c, 0xdd, 0x0b, 0x55, 0x5a,
	0x68, 0x65, 0x55, 0x54, 0x0b, 0x8a, 0x34, 0xdc, 0x2d, 0x14, 0x8d, 0x3b, 0x54, 0x70, 0xa9, 0x88,
	0xff, 0x0d, 0xc2, 0xc6, 0xdd, 0x20, 0xec, 0xf9, 0x8a, 0x2c, 0x5c, 0xa1, 0xd5, 0x64, 0x4a, 0xb1,
	0x11, 0x10, 0x5f, 0xe5, 0xee, 0x88, 0x58, 0x2e, 0xc0, 0x58, 0x2a, 0x8a, 0x85, 0xa0, 0xc6, 0x14,
	0x53, 0xc1, 0x38, 0x3f, 0x95, 0x13, 0xff, 0xb6, 0x51, 0x39, 0x0e, 0xad, 0xc4, 0xe2, 0xda, 0x01,
	0x48, 0xd0, 0xbc, 0xdf, 0x71, 0x76, 0xa8, 0x34, 0x3f, 0xa5, 0x96, 0x2b, 0x19, 0xdd, 0xc6, 0x55,
	0x61, 0x58, 0x1d, 0xb5, 0xd0, 0xee, 0x56, 0x36, 0x3f, 0x3e, 0x7e, 0xfe, 0xf9, 0xbc, 0x9d, 0x5c,
	0xf6, 0x1f, 0xd2, 0x15, 0xe7, 0xfb, 0x5f, 0x1f, 0xef, 0x37, 0x83, 0xac, 0x6d, 0x06, 0xef, 0xc8,
	0x65, 0xd3, 0x93, 0x4f, 0x08, 0x5f, 0x3b, 0xd0, 0x54, 0xda, 0x28, 0xc7, 0x3b, 0x74, 0xb9, 0xe5,
	0x37, 0x6e, 0xef, 0xd7, 0xd2, 0x10, 0x39, 0x2d, 0x23, 0xa7, 0x1d, 0x39, 0xee, 0xde, 0xbb, 0x5a,
	0x84, 0x6c, 0x75, 0x64, 0xf4, 0x14, 0x63, 0x38, 0x29, 0xb8, 0x0e, 0x0b, 0x36, 0xfc, 0x82, 0xc6,
	0xda, 0x82, 0xc3, 0x12, 0x65, 0xf7, 0xe6, 0xe4, 0x5b, 0x13, 0x9d, 0x7d, 0x6f, 0xa2, 0x6c, 0xc9,
	0x97, 0x7c, 0xd8, 0xc0, 0x91, 0xcf, 0xbc, 0x0a, 0x6a, 0x1f, 0xdf, 0x60, 0xf3, 0x5b, 0xd0, 0x01,
	0x56, 0xb7, 0xfe, 0xe5, 0xbc, 0x5d, 0x3e, 0xeb, 0xce, 0x60, 0xa0, 0xc1, 0x98, 0x57, 0x56, 0x73,
	0xc9, 0xb2, 0x52, 0xf8, 0xc7, 0x03, 0x3e, 0xcd, 0x15, 0x3c, 0xb0, 0x0e, 0xaa, 0xfa, 0xff, 0x41,
	0x3d, 0x59, 0x01, 0xb5, 0xf9, 0x4f, 0x50, 0x9b, 0x6b, 0x90, 0x1e, 0xe2, 0x5b, 0x9e, 0xd1, 0x4b,
	0x07, 0x0e, 0x9e, 0x59, 0x10, 0x51, 0x82, 0x77, 0x84, 0x61, 0x3d, 0x3b, 0x2e, 0xa0, 0xe7, 0xf4,
	0xc8, 0xd4, 0x51, 0xab, 0xba, 0xbb, 0x95, 0x6d, 0x0b, 0xc3, 0x0e, 0xc7, 0x05, 0xbc, 0xd6, 0x23,
	0xd3, 0x7d, 0x31, 0xf9, 0x19, 0x57, 0x26, 0xd3, 0x18, 0x5d, 0x4c, 0x63, 0xf4, 0x63, 0x1a, 0xa3,
	0xb3, 0x59, 0x5c, 0xb9, 0x98, 0xc5, 0x95, 0xaf, 0xb3, 0xb8, 0xf2, 0xa6, 0xcd, 0xb8, 0x1d, 0xba,
	0x3c, 0xed, 0x2b, 0x41, 0x8e, 0x94, 0xce, 0xd5, 0x08, 0xc8, 0x5b, 0x27, 0x15, 0x39, 0x7e, 0x44,
	0x96, 0x5e, 0xb4, 0x93, 0xf0, 0xb5, 0xe5, 0xd7, 0x7d, 0xd0, 0x07, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x3c, 0x1c, 0xa0, 0x8f, 0x92, 0x03, 0x00, 0x00,
}

func (m *GenericAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenericAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenericAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintAuthz(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Grant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Grant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Grant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Expiration != nil {
		n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.Expiration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Expiration):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintAuthz(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x12
	}
	if m.Authorization != nil {
		{
			size, err := m.Authorization.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuthz(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GrantAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrantAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GrantAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Expiration != nil {
		n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.Expiration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Expiration):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintAuthz(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0x22
	}
	if m.Authorization != nil {
		{
			size, err := m.Authorization.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuthz(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintAuthz(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintAuthz(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GrantQueueItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrantQueueItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GrantQueueItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgTypeUrls) > 0 {
		for iNdEx := len(m.MsgTypeUrls) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.MsgTypeUrls[iNdEx])
			copy(dAtA[i:], m.MsgTypeUrls[iNdEx])
			i = encodeVarintAuthz(dAtA, i, uint64(len(m.MsgTypeUrls[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthz(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthz(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenericAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovAuthz(uint64(l))
	}
	return n
}

func (m *Grant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Authorization != nil {
		l = m.Authorization.Size()
		n += 1 + l + sovAuthz(uint64(l))
	}
	if m.Expiration != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Expiration)
		n += 1 + l + sovAuthz(uint64(l))
	}
	return n
}

func (m *GrantAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovAuthz(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovAuthz(uint64(l))
	}
	if m.Authorization != nil {
		l = m.Authorization.Size()
		n += 1 + l + sovAuthz(uint64(l))
	}
	if m.Expiration != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Expiration)
		n += 1 + l + sovAuthz(uint64(l))
	}
	return n
}

func (m *GrantQueueItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MsgTypeUrls) > 0 {
		for _, s := range m.MsgTypeUrls {
			l = len(s)
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	return n
}

func sovAuthz(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthz(x uint64) (n int) {
	return sovAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenericAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: GenericAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenericAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func (m *Grant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: Grant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Grant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authorization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authorization == nil {
				m.Authorization = &any.Any{}
			}
			if err := m.Authorization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expiration == nil {
				m.Expiration = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.Expiration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func (m *GrantAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: GrantAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrantAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authorization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authorization == nil {
				m.Authorization = &any.Any{}
			}
			if err := m.Authorization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expiration == nil {
				m.Expiration = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.Expiration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func (m *GrantQueueItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: GrantQueueItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrantQueueItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrls", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrls = append(m.MsgTypeUrls, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func skipAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
				return 0, ErrInvalidLengthAuthz
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthz
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthz
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthz        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthz          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthz = fmt.Errorf("proto: unexpected end of group")
)
