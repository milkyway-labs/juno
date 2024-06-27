// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/authz/v1beta1/tx.proto

package authz

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	any "github.com/cosmos/gogoproto/types/any"
	_ "github.com/forbole/juno/v5/cosmos-sdk/types/msgservice"
	_ "github.com/forbole/juno/v5/cosmos-sdk/types/tx/amino"
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

// MsgExec attempts to execute the provided messages using
// authorizations granted to the grantee. Each message should have only
// one signer corresponding to the granter of the authorization.
type MsgExec struct {
	Grantee string `protobuf:"bytes,1,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// Execute Msg.
	// The x/authz will try to find a grant matching (msg.signers[0], grantee,
	// MsgTypeURL(msg)) triple and validate it.
	Msgs []*any.Any `protobuf:"bytes,2,rep,name=msgs,proto3" json:"msgs,omitempty"`
}

func (m *MsgExec) Reset()         { *m = MsgExec{} }
func (m *MsgExec) String() string { return proto.CompactTextString(m) }
func (*MsgExec) ProtoMessage()    {}
func (*MsgExec) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ceddab7d8589ad1, []int{0}
}
func (m *MsgExec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExec.Merge(m, src)
}
func (m *MsgExec) XXX_Size() int {
	return m.Size()
}
func (m *MsgExec) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExec.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExec proto.InternalMessageInfo

// MsgExecResponse defines the Msg/MsgExecResponse response type.
type MsgExecResponse struct {
	Results [][]byte `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (m *MsgExecResponse) Reset()         { *m = MsgExecResponse{} }
func (m *MsgExecResponse) String() string { return proto.CompactTextString(m) }
func (*MsgExecResponse) ProtoMessage()    {}
func (*MsgExecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ceddab7d8589ad1, []int{1}
}
func (m *MsgExecResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecResponse.Merge(m, src)
}
func (m *MsgExecResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgExec)(nil), "cosmos.authz.v1beta1.MsgExec")
	proto.RegisterType((*MsgExecResponse)(nil), "cosmos.authz.v1beta1.MsgExecResponse")
}

func init() { proto.RegisterFile("cosmos/authz/v1beta1/tx.proto", fileDescriptor_3ceddab7d8589ad1) }

var fileDescriptor_3ceddab7d8589ad1 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x31, 0x4f, 0xc2, 0x40,
	0x14, 0xc7, 0x5b, 0x31, 0x12, 0xab, 0x89, 0xb1, 0x21, 0xb1, 0x60, 0x6c, 0x08, 0x2e, 0x04, 0xd3,
	0xbb, 0x80, 0x71, 0x71, 0x83, 0x84, 0xc9, 0xb0, 0xd4, 0xcd, 0xc5, 0xb4, 0x70, 0x1c, 0x28, 0xed,
	0x23, 0x7d, 0x57, 0x02, 0x8e, 0x8e, 0x4e, 0x7e, 0x0d, 0x37, 0x06, 0x3e, 0x04, 0x71, 0x22, 0x4e,
	0x8e, 0x0a, 0x03, 0x5f, 0xc3, 0x70, 0x77, 0x4d, 0x5c, 0xee, 0xde, 0xfb, 0xff, 0xff, 0x77, 0xf7,
	0x7b, 0x39, 0xeb, 0xa2, 0x0b, 0x18, 0x01, 0xd2, 0x20, 0x15, 0x83, 0x17, 0x3a, 0xa9, 0x87, 0x4c,
	0x04, 0x75, 0x2a, 0xa6, 0x64, 0x9c, 0x80, 0x00, 0xbb, 0xa0, 0x6c, 0x22, 0x6d, 0xa2, 0xed, 0x52,
	0x51, 0xa9, 0x8f, 0x32, 0x43, 0x75, 0x44, 0x36, 0xa5, 0x02, 0x07, 0x0e, 0x4a, 0xdf, 0x55, 0x5a,
	0x2d, 0x72, 0x00, 0x3e, 0x62, 0x54, 0x76, 0x61, 0xda, 0xa7, 0x41, 0x3c, 0xd3, 0xd6, 0x99, 0x06,
	0x88, 0x90, 0xd3, 0x49, 0x7d, 0xb7, 0x69, 0xe3, 0x34, 0x88, 0x86, 0x31, 0x50, 0xb9, 0x2a, 0xa9,
	0xf2, 0x61, 0x5a, 0xf9, 0x0e, 0xf2, 0xf6, 0x94, 0x75, 0xed, 0x86, 0x95, 0xe7, 0x49, 0x10, 0x0b,
	0xc6, 0x1c, 0xb3, 0x6c, 0x56, 0x0f, 0x5b, 0xce, 0xd7, 0xc2, 0xcb, 0x70, 0x9b, 0xbd, 0x5e, 0xc2,
	0x10, 0xef, 0x45, 0x32, 0x8c, 0xb9, 0x9f, 0x05, 0xed, 0xb6, 0xb5, 0x1f, 0x21, 0x47, 0x67, 0xaf,
	0x9c, 0xab, 0x1e, 0x35, 0x0a, 0x44, 0x51, 0x91, 0x8c, 0x8a, 0x34, 0xe3, 0x59, 0xeb, 0xfc, 0x73,
	0xe1, 0x69, 0x26, 0x12, 0x06, 0xc8, 0xb2, 0xa1, 0x49, 0x07, 0xb9, 0x2f, 0x8f, 0xdf, 0x5e, 0xbe,
	0x6e, 0xe7, 0xb5, 0xec, 0xd2, 0xb7, 0xed, 0xbc, 0x66, 0xab, 0xbc, 0x87, 0xbd, 0x67, 0xaa, 0xf9,
	0x2a, 0x57, 0xd6, 0x89, 0x2e, 0x7d, 0x86, 0x63, 0x88, 0x91, 0xd9, 0x8e, 0x95, 0x4f, 0x18, 0xa6,
	0x23, 0x81, 0x8e, 0x59, 0xce, 0x55, 0x8f, 0xfd, 0xac, 0x6d, 0xdd, 0x2d, 0x7f, 0x5d, 0x63, 0xb9,
	0x76, 0xcd, 0xd5, 0xda, 0x35, 0x7f, 0xd6, 0xae, 0xf9, 0xbe, 0x71, 0x8d, 0xd5, 0xc6, 0x35, 0xbe,
	0x37, 0xae, 0xf1, 0xe0, 0xf1, 0xa1, 0x18, 0xa4, 0x21, 0xe9, 0x42, 0x44, 0xfb, 0x90, 0x84, 0x30,
	0x62, 0xf4, 0x29, 0x8d, 0x81, 0x4e, 0x6e, 0xe8, 0xbf, 0x97, 0xa7, 0xea, 0x13, 0xc3, 0x03, 0x39,
	0xcf, 0xf5, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x70, 0x9d, 0x30, 0xdb, 0x01, 0x00, 0x00,
}

func (m *MsgExec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for iNdEx := len(m.Msgs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Msgs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgExecResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Results) > 0 {
		for iNdEx := len(m.Results) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Results[iNdEx])
			copy(dAtA[i:], m.Results[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Results[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgExec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgExecResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, b := range m.Results {
			l = len(b)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgExec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgExec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msgs = append(m.Msgs, &any.Any{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgExecResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgExecResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, make([]byte, postIndex-iNdEx))
			copy(m.Results[len(m.Results)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
