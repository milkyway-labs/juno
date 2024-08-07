// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/channel/v1/channel.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/forbole/juno/v5/ibc-go/modules/core/02-client/types"
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

type Packet struct {
	// number corresponds to the order of sends and receives, where a Packet
	// with an earlier sequence number must be sent and received before a Packet
	// with a later sequence number.
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// identifies the port on the sending chain.
	SourcePort string `protobuf:"bytes,2,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty"`
	// identifies the channel end on the sending chain.
	SourceChannel string `protobuf:"bytes,3,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty"`
	// identifies the port on the receiving chain.
	DestinationPort string `protobuf:"bytes,4,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	// identifies the channel end on the receiving chain.
	DestinationChannel string `protobuf:"bytes,5,opt,name=destination_channel,json=destinationChannel,proto3" json:"destination_channel,omitempty"`
	// actual opaque bytes transferred directly to the application module
	Data []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// block height after which the packet times out
	TimeoutHeight types.Height `protobuf:"bytes,7,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height"`
	// block timestamp (in nanoseconds) after which the packet times out
	TimeoutTimestamp uint64 `protobuf:"varint,8,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a07336710636a0, []int{0}
}
func (m *Packet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return m.Size()
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func init() {
	// proto.RegisterType((*Packet)(nil), "ibc.core.channel.v1.Packet")
}

//func init() { proto.RegisterFile("ibc/core/channel/v1/channel.proto", fileDescriptor_c3a07336710636a0) }

var fileDescriptor_c3a07336710636a0 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0xcd, 0xce, 0x93, 0x40,
	0x14, 0x86, 0x99, 0x8a, 0xb5, 0x4e, 0x6d, 0xad, 0x53, 0x17, 0x84, 0x05, 0xa0, 0x89, 0x09, 0xc6,
	0x94, 0xb1, 0xfe, 0x6c, 0xdc, 0x98, 0xd4, 0x85, 0x2e, 0x1b, 0xe2, 0x46, 0x37, 0x0d, 0x4c, 0x47,
	0x18, 0x05, 0x0e, 0xc2, 0x40, 0xe2, 0x1d, 0xb8, 0xf4, 0x12, 0xbc, 0x9c, 0x2e, 0xbb, 0xec, 0xca,
	0x98, 0xf6, 0x46, 0xbe, 0x74, 0x06, 0xbe, 0x76, 0xc5, 0xcb, 0x7b, 0x1e, 0x9e, 0x90, 0x33, 0x83,
	0x9f, 0x88, 0x98, 0x51, 0x06, 0x15, 0xa7, 0x2c, 0x8d, 0x8a, 0x82, 0x67, 0xb4, 0x5d, 0xf6, 0x31,
	0x28, 0x2b, 0x90, 0x40, 0xe6, 0x22, 0x66, 0xc1, 0x19, 0x09, 0xfa, 0xbe, 0x5d, 0xda, 0x8f, 0x13,
	0x48, 0x40, 0xcd, 0xe9, 0x39, 0x69, 0xd4, 0x76, 0x2f, 0xb6, 0x4c, 0xf0, 0x42, 0x2a, 0x99, 0x4a,
	0x1a, 0x78, 0x7a, 0x18, 0xe0, 0xe1, 0x3a, 0x62, 0x3f, 0xb8, 0x24, 0x36, 0x1e, 0xd5, 0xfc, 0x67,
	0xc3, 0x0b, 0xc6, 0x2d, 0xe4, 0x21, 0xdf, 0x0c, 0x6f, 0xdf, 0x89, 0x8b, 0xc7, 0x35, 0x34, 0x15,
	0xe3, 0x9b, 0x12, 0x2a, 0x69, 0x0d, 0x3c, 0xe4, 0xdf, 0x0f, 0xb1, 0xae, 0xd6, 0x50, 0x49, 0xf2,
	0x0c, 0x4f, 0x3b, 0xa0, 0xfb, 0x27, 0xeb, 0x8e, 0x62, 0x26, 0xba, 0xfd, 0xa0, 0x4b, 0xf2, 0x1c,
	0xcf, 0xb6, 0xbc, 0x96, 0xa2, 0x88, 0xa4, 0x80, 0x42, 0xcb, 0x4c, 0x05, 0x3e, 0xbc, 0xea, 0x95,
	0x91, 0xe2, 0xf9, 0x35, 0xda, 0x6b, 0xef, 0x2a, 0x9a, 0x5c, 0x8d, 0x7a, 0x37, 0xc1, 0xe6, 0x36,
	0x92, 0x91, 0x35, 0xf4, 0x90, 0xff, 0x20, 0x54, 0x99, 0x7c, 0xc4, 0x53, 0x29, 0x72, 0x0e, 0x8d,
	0xdc, 0xa4, 0x5c, 0x24, 0xa9, 0xb4, 0xee, 0x79, 0xc8, 0x1f, 0xbf, 0xb2, 0x83, 0xcb, 0x0e, 0xf5,
	0x3a, 0xda, 0x65, 0xf0, 0x49, 0x11, 0x2b, 0x73, 0xf7, 0xcf, 0x35, 0xc2, 0x49, 0xf7, 0x9d, 0x2e,
	0xc9, 0x0b, 0xfc, 0xa8, 0x17, 0x9d, 0x9f, 0xb5, 0x8c, 0xf2, 0xd2, 0x1a, 0xa9, 0x2d, 0xcd, 0xba,
	0xc1, 0xe7, 0xbe, 0x7f, 0x67, 0xfe, 0xfe, 0xeb, 0x1a, 0xab, 0x2f, 0xbb, 0xa3, 0x83, 0xf6, 0x47,
	0x07, 0xfd, 0x3f, 0x3a, 0xe8, 0xcf, 0xc9, 0x31, 0xf6, 0x27, 0xc7, 0x38, 0x9c, 0x1c, 0xe3, 0xeb,
	0xfb, 0x44, 0xc8, 0xb4, 0x89, 0x03, 0x06, 0x39, 0xfd, 0x06, 0x55, 0x0c, 0x19, 0xa7, 0xdf, 0x9b,
	0x02, 0x68, 0xfb, 0x96, 0x8a, 0x98, 0x2d, 0x12, 0xa0, 0x39, 0x6c, 0x9b, 0x8c, 0xd7, 0xfa, 0xec,
	0x5e, 0xbe, 0x59, 0xf4, 0x97, 0x41, 0xfe, 0x2a, 0x79, 0x1d, 0x0f, 0xd5, 0xe1, 0xbd, 0xbe, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x9c, 0x41, 0xfd, 0x06, 0x2d, 0x02, 0x00, 0x00,
}

func (m *Packet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Packet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Packet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintChannel(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x40
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintChannel(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintChannel(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DestinationChannel) > 0 {
		i -= len(m.DestinationChannel)
		copy(dAtA[i:], m.DestinationChannel)
		i = encodeVarintChannel(dAtA, i, uint64(len(m.DestinationChannel)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DestinationPort) > 0 {
		i -= len(m.DestinationPort)
		copy(dAtA[i:], m.DestinationPort)
		i = encodeVarintChannel(dAtA, i, uint64(len(m.DestinationPort)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintChannel(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintChannel(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0x12
	}
	if m.Sequence != 0 {
		i = encodeVarintChannel(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintChannel(dAtA []byte, offset int, v uint64) int {
	offset -= sovChannel(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Packet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovChannel(uint64(m.Sequence))
	}
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovChannel(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovChannel(uint64(l))
	}
	l = len(m.DestinationPort)
	if l > 0 {
		n += 1 + l + sovChannel(uint64(l))
	}
	l = len(m.DestinationChannel)
	if l > 0 {
		n += 1 + l + sovChannel(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovChannel(uint64(l))
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovChannel(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovChannel(uint64(m.TimeoutTimestamp))
	}
	return n
}

func sovChannel(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChannel(x uint64) (n int) {
	return sovChannel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Packet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChannel
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
			return fmt.Errorf("proto: Packet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Packet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChannel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChannel
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
				return ErrInvalidLengthChannel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChannel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChannel
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
				return ErrInvalidLengthChannel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChannel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationPort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChannel
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
				return ErrInvalidLengthChannel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChannel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationPort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChannel
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
				return ErrInvalidLengthChannel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChannel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChannel
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
				return ErrInvalidLengthChannel
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChannel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChannel
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
				return ErrInvalidLengthChannel
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChannel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChannel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChannel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChannel
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
func skipChannel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChannel
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
					return 0, ErrIntOverflowChannel
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
					return 0, ErrIntOverflowChannel
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
				return 0, ErrInvalidLengthChannel
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChannel
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChannel
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChannel        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChannel          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChannel = fmt.Errorf("proto: unexpected end of group")
)
