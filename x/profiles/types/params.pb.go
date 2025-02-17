// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Params contains the parameters for the profiles module
type Params struct {
	MonikerParams MonikerParams                          `protobuf:"bytes,1,opt,name=moniker_params,json=monikerParams,proto3" json:"moniker_params" yaml:"moniker_params"`
	DtagParams    DTagParams                             `protobuf:"bytes,2,opt,name=dtag_params,json=dtagParams,proto3" json:"dtag_params" yaml:"dtag_params"`
	MaxBioLength  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=max_bio_length,json=maxBioLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_bio_length" yaml:"max_bio_length"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_821862d20041ec2d, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMonikerParams() MonikerParams {
	if m != nil {
		return m.MonikerParams
	}
	return MonikerParams{}
}

func (m *Params) GetDtagParams() DTagParams {
	if m != nil {
		return m.DtagParams
	}
	return DTagParams{}
}

// MonikerParams defines the parameters related to the profiles monikers
type MonikerParams struct {
	MinMonikerLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=min_moniker_length,json=minMonikerLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_moniker_length" yaml:"min_moniker_length"`
	MaxMonikerLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=max_moniker_length,json=maxMonikerLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_moniker_length" yaml:"max_moniker_length"`
}

func (m *MonikerParams) Reset()         { *m = MonikerParams{} }
func (m *MonikerParams) String() string { return proto.CompactTextString(m) }
func (*MonikerParams) ProtoMessage()    {}
func (*MonikerParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_821862d20041ec2d, []int{1}
}
func (m *MonikerParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MonikerParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MonikerParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MonikerParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonikerParams.Merge(m, src)
}
func (m *MonikerParams) XXX_Size() int {
	return m.Size()
}
func (m *MonikerParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MonikerParams.DiscardUnknown(m)
}

var xxx_messageInfo_MonikerParams proto.InternalMessageInfo

// DTagParams defines the parameters related to profile DTags
type DTagParams struct {
	RegEx         string                                 `protobuf:"bytes,1,opt,name=reg_ex,json=regEx,proto3" json:"reg_ex,omitempty" yaml:"reg_ex"`
	MinDtagLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=min_dtag_length,json=minDtagLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_dtag_length" yaml:"min_dtag_length"`
	MaxDtagLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=max_dtag_length,json=maxDtagLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_dtag_length" yaml:"max_dtag_length"`
}

func (m *DTagParams) Reset()         { *m = DTagParams{} }
func (m *DTagParams) String() string { return proto.CompactTextString(m) }
func (*DTagParams) ProtoMessage()    {}
func (*DTagParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_821862d20041ec2d, []int{2}
}
func (m *DTagParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DTagParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DTagParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DTagParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DTagParams.Merge(m, src)
}
func (m *DTagParams) XXX_Size() int {
	return m.Size()
}
func (m *DTagParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DTagParams.DiscardUnknown(m)
}

var xxx_messageInfo_DTagParams proto.InternalMessageInfo

func (m *DTagParams) GetRegEx() string {
	if m != nil {
		return m.RegEx
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "desmos.profiles.v1beta1.Params")
	proto.RegisterType((*MonikerParams)(nil), "desmos.profiles.v1beta1.MonikerParams")
	proto.RegisterType((*DTagParams)(nil), "desmos.profiles.v1beta1.DTagParams")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/params.proto", fileDescriptor_821862d20041ec2d)
}

var fileDescriptor_821862d20041ec2d = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x9b, 0x4c, 0x54, 0xc2, 0x5b, 0x07, 0x58, 0xfc, 0x19, 0x93, 0x48, 0x27, 0x83, 0xa6,
	0x5e, 0xe6, 0x68, 0x70, 0xe3, 0x18, 0x0d, 0x31, 0x04, 0x48, 0x28, 0xe2, 0xc4, 0x25, 0x38, 0xab,
	0xf1, 0xac, 0xc5, 0x71, 0x94, 0x18, 0xe4, 0x7e, 0x00, 0xee, 0x9c, 0xf9, 0x42, 0xec, 0xb8, 0x23,
	0xe2, 0x50, 0xa1, 0xf6, 0x1b, 0xec, 0x13, 0x20, 0xff, 0xa9, 0x92, 0x30, 0xed, 0x30, 0xf5, 0x54,
	0xbb, 0x7a, 0xde, 0xe7, 0xf7, 0xbc, 0xaf, 0xf3, 0x82, 0x67, 0x53, 0xda, 0x08, 0xd9, 0xc4, 0x55,
	0x2d, 0xbf, 0xf0, 0x82, 0x36, 0xf1, 0xb7, 0xc3, 0x9c, 0x2a, 0x72, 0x18, 0x57, 0xa4, 0x26, 0xa2,
	0xc1, 0x55, 0x2d, 0x95, 0x84, 0x8f, 0x9c, 0x0a, 0xaf, 0x54, 0xd8, 0xab, 0x76, 0xef, 0x33, 0xc9,
	0xa4, 0xd5, 0xc4, 0xe6, 0xe4, 0xe4, 0xe8, 0x57, 0x08, 0x86, 0x1f, 0x6c, 0x3d, 0x2c, 0xc0, 0xb6,
	0x90, 0x25, 0x3f, 0xa3, 0x75, 0xe6, 0x1c, 0x77, 0x82, 0xbd, 0x60, 0xb2, 0xf9, 0x7c, 0x1f, 0x5f,
	0x63, 0x89, 0xdf, 0x3b, 0xb9, 0xab, 0x4f, 0x9e, 0x9c, 0xcf, 0xc7, 0x83, 0xcb, 0xf9, 0xf8, 0xc1,
	0x8c, 0x88, 0xe2, 0x25, 0xea, 0x7b, 0xa1, 0x74, 0x24, 0xba, 0x6a, 0xf8, 0x19, 0x6c, 0x4e, 0x15,
	0x61, 0x2b, 0x54, 0x68, 0x51, 0x4f, 0xaf, 0x45, 0x1d, 0x7d, 0x24, 0xcc, 0x73, 0x76, 0x3d, 0x07,
	0x3a, 0x4e, 0xc7, 0x05, 0xa5, 0xc0, 0xdc, 0x3c, 0x41, 0x80, 0x6d, 0x41, 0x74, 0x96, 0x73, 0x99,
	0x15, 0xb4, 0x64, 0xea, 0x74, 0x67, 0x63, 0x2f, 0x98, 0x6c, 0x25, 0xaf, 0x4d, 0xfd, 0x9f, 0xf9,
	0x78, 0x9f, 0x71, 0x75, 0xfa, 0x35, 0xc7, 0x27, 0x52, 0xc4, 0x27, 0xd2, 0x8e, 0xd6, 0xfd, 0x1c,
	0x34, 0xd3, 0xb3, 0x58, 0xcd, 0x2a, 0xda, 0xe0, 0x37, 0xa5, 0xea, 0x74, 0xd4, 0x73, 0x43, 0xe9,
	0x96, 0x20, 0x3a, 0xe1, 0xf2, 0x9d, 0xbb, 0x7e, 0x0f, 0xc1, 0xa8, 0x37, 0x10, 0x38, 0x03, 0x50,
	0xf0, 0x32, 0x5b, 0x0d, 0xc2, 0x87, 0x08, 0x6c, 0x88, 0xb7, 0x37, 0x0e, 0xf1, 0xd8, 0x87, 0xb8,
	0xe2, 0x88, 0xd2, 0xbb, 0x82, 0x97, 0x1e, 0xed, 0xc2, 0x58, 0x34, 0xd1, 0xff, 0xa3, 0xc3, 0x35,
	0xd1, 0x57, 0x1c, 0x0d, 0x9a, 0xe8, 0x1e, 0x1a, 0xfd, 0x0c, 0x01, 0x68, 0x5f, 0x0b, 0x4e, 0xc0,
	0xb0, 0xa6, 0x2c, 0xa3, 0xda, 0x36, 0x7e, 0x3b, 0xb9, 0x77, 0x39, 0x1f, 0x8f, 0x9c, 0x9f, 0xfb,
	0x1f, 0xa5, 0xb7, 0x6a, 0xca, 0x5e, 0x69, 0x58, 0x81, 0x3b, 0xa6, 0x39, 0xfb, 0x9e, 0xbd, 0xc0,
	0xc7, 0x37, 0x0e, 0xfc, 0xb0, 0x9d, 0x55, 0xc7, 0xce, 0x7c, 0x83, 0xbc, 0x3c, 0x52, 0x84, 0xf9,
	0x29, 0x19, 0x22, 0xd1, 0x3d, 0xe2, 0xc6, 0x9a, 0xc4, 0xbe, 0x9d, 0x21, 0x12, 0xdd, 0x12, 0x93,
	0xe3, 0xf3, 0x45, 0x14, 0x5c, 0x2c, 0xa2, 0xe0, 0xef, 0x22, 0x0a, 0x7e, 0x2c, 0xa3, 0xc1, 0xc5,
	0x32, 0x1a, 0xfc, 0x5e, 0x46, 0x83, 0x4f, 0xb8, 0x83, 0x72, 0x4b, 0x70, 0x50, 0x90, 0xbc, 0xf1,
	0xe7, 0x58, 0xb7, 0x6b, 0x6f, 0xb1, 0xf9, 0xd0, 0xee, 0xef, 0x8b, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xce, 0x1d, 0x51, 0xff, 0x16, 0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxBioLength.Size()
		i -= size
		if _, err := m.MaxBioLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.DtagParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.MonikerParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MonikerParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MonikerParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MonikerParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxMonikerLength.Size()
		i -= size
		if _, err := m.MaxMonikerLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinMonikerLength.Size()
		i -= size
		if _, err := m.MinMonikerLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *DTagParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DTagParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DTagParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxDtagLength.Size()
		i -= size
		if _, err := m.MaxDtagLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinDtagLength.Size()
		i -= size
		if _, err := m.MinDtagLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.RegEx) > 0 {
		i -= len(m.RegEx)
		copy(dAtA[i:], m.RegEx)
		i = encodeVarintParams(dAtA, i, uint64(len(m.RegEx)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MonikerParams.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.DtagParams.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxBioLength.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *MonikerParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinMonikerLength.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxMonikerLength.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *DTagParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RegEx)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.MinDtagLength.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxDtagLength.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonikerParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MonikerParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DtagParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DtagParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBioLength", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxBioLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *MonikerParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: MonikerParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MonikerParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinMonikerLength", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinMonikerLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMonikerLength", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxMonikerLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *DTagParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: DTagParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DTagParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegEx", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegEx = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDtagLength", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinDtagLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDtagLength", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxDtagLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
