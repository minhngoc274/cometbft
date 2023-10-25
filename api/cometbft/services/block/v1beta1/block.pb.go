// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cometbft/services/block/v1beta1/block.proto

package v1beta1

import (
	fmt "fmt"
	v1beta1 "github.com/cometbft/cometbft/api/cometbft/types/v1beta1"
	v1beta3 "github.com/cometbft/cometbft/api/cometbft/types/v1beta3"
	proto "github.com/cosmos/gogoproto/proto"
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

type GetByHeightRequest struct {
	// The height of the block requested. If set to 0, the latest height will be returned.
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *GetByHeightRequest) Reset()         { *m = GetByHeightRequest{} }
func (m *GetByHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetByHeightRequest) ProtoMessage()    {}
func (*GetByHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_260532424c9a31db, []int{0}
}
func (m *GetByHeightRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetByHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetByHeightRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetByHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByHeightRequest.Merge(m, src)
}
func (m *GetByHeightRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetByHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByHeightRequest proto.InternalMessageInfo

func (m *GetByHeightRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetByHeightResponse struct {
	BlockId *v1beta1.BlockID `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Block   *v1beta3.Block   `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (m *GetByHeightResponse) Reset()         { *m = GetByHeightResponse{} }
func (m *GetByHeightResponse) String() string { return proto.CompactTextString(m) }
func (*GetByHeightResponse) ProtoMessage()    {}
func (*GetByHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_260532424c9a31db, []int{1}
}
func (m *GetByHeightResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetByHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetByHeightResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetByHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByHeightResponse.Merge(m, src)
}
func (m *GetByHeightResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetByHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetByHeightResponse proto.InternalMessageInfo

func (m *GetByHeightResponse) GetBlockId() *v1beta1.BlockID {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *GetByHeightResponse) GetBlock() *v1beta3.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type GetLatestRequest struct {
}

func (m *GetLatestRequest) Reset()         { *m = GetLatestRequest{} }
func (m *GetLatestRequest) String() string { return proto.CompactTextString(m) }
func (*GetLatestRequest) ProtoMessage()    {}
func (*GetLatestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_260532424c9a31db, []int{2}
}
func (m *GetLatestRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLatestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLatestRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLatestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestRequest.Merge(m, src)
}
func (m *GetLatestRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetLatestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestRequest proto.InternalMessageInfo

type GetLatestResponse struct {
	BlockId *v1beta1.BlockID `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Block   *v1beta3.Block   `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (m *GetLatestResponse) Reset()         { *m = GetLatestResponse{} }
func (m *GetLatestResponse) String() string { return proto.CompactTextString(m) }
func (*GetLatestResponse) ProtoMessage()    {}
func (*GetLatestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_260532424c9a31db, []int{3}
}
func (m *GetLatestResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLatestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLatestResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLatestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestResponse.Merge(m, src)
}
func (m *GetLatestResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetLatestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestResponse proto.InternalMessageInfo

func (m *GetLatestResponse) GetBlockId() *v1beta1.BlockID {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *GetLatestResponse) GetBlock() *v1beta3.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

// GetLatestHeightRequest - empty message since no parameter is required
type GetLatestHeightRequest struct {
}

func (m *GetLatestHeightRequest) Reset()         { *m = GetLatestHeightRequest{} }
func (m *GetLatestHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetLatestHeightRequest) ProtoMessage()    {}
func (*GetLatestHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_260532424c9a31db, []int{4}
}
func (m *GetLatestHeightRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLatestHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLatestHeightRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLatestHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestHeightRequest.Merge(m, src)
}
func (m *GetLatestHeightRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetLatestHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestHeightRequest proto.InternalMessageInfo

// GetLatestHeightResponse provides the height of the latest committed block.
type GetLatestHeightResponse struct {
	// The height of the latest committed block. Will be 0 if no data has been
	// committed yet.
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *GetLatestHeightResponse) Reset()         { *m = GetLatestHeightResponse{} }
func (m *GetLatestHeightResponse) String() string { return proto.CompactTextString(m) }
func (*GetLatestHeightResponse) ProtoMessage()    {}
func (*GetLatestHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_260532424c9a31db, []int{5}
}
func (m *GetLatestHeightResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLatestHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLatestHeightResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLatestHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestHeightResponse.Merge(m, src)
}
func (m *GetLatestHeightResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetLatestHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestHeightResponse proto.InternalMessageInfo

func (m *GetLatestHeightResponse) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*GetByHeightRequest)(nil), "cometbft.services.block.v1beta1.GetByHeightRequest")
	proto.RegisterType((*GetByHeightResponse)(nil), "cometbft.services.block.v1beta1.GetByHeightResponse")
	proto.RegisterType((*GetLatestRequest)(nil), "cometbft.services.block.v1beta1.GetLatestRequest")
	proto.RegisterType((*GetLatestResponse)(nil), "cometbft.services.block.v1beta1.GetLatestResponse")
	proto.RegisterType((*GetLatestHeightRequest)(nil), "cometbft.services.block.v1beta1.GetLatestHeightRequest")
	proto.RegisterType((*GetLatestHeightResponse)(nil), "cometbft.services.block.v1beta1.GetLatestHeightResponse")
}

func init() {
	proto.RegisterFile("cometbft/services/block/v1beta1/block.proto", fileDescriptor_260532424c9a31db)
}

var fileDescriptor_260532424c9a31db = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0xce, 0xcf, 0x4d,
	0x2d, 0x49, 0x4a, 0x2b, 0xd1, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x4f, 0xca,
	0xc9, 0x4f, 0xce, 0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0x84, 0xf0, 0xf4, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xe4, 0x61, 0x8a, 0xf5, 0x60, 0x8a, 0xf5, 0x20, 0xd2, 0x50, 0xc5, 0x52,
	0x4a, 0x70, 0xd3, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0xe1, 0x86, 0x80, 0x79, 0x10, 0x43, 0x70, 0xa8,
	0x31, 0x46, 0xb6, 0x48, 0x49, 0x87, 0x4b, 0xc8, 0x3d, 0xb5, 0xc4, 0xa9, 0xd2, 0x23, 0x35, 0x33,
	0x3d, 0xa3, 0x24, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x8c, 0x8b, 0x2d, 0x03, 0x2c,
	0x20, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1c, 0x04, 0xe5, 0x29, 0xb5, 0x31, 0x72, 0x09, 0xa3, 0x28,
	0x2f, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xb2, 0xe2, 0xe2, 0x00, 0x1b, 0x1a, 0x9f, 0x99, 0x02,
	0xd6, 0xc1, 0x6d, 0x24, 0xaf, 0x07, 0xf7, 0x01, 0xc4, 0x49, 0x50, 0x07, 0xea, 0x39, 0x81, 0xd4,
	0x79, 0xba, 0x04, 0xb1, 0x83, 0x35, 0x78, 0xa6, 0x08, 0x19, 0x73, 0xb1, 0x82, 0x99, 0x12, 0x4c,
	0x60, 0x8d, 0xb2, 0xd8, 0x35, 0x1a, 0x43, 0x34, 0x06, 0x41, 0xd4, 0x2a, 0x09, 0x71, 0x09, 0xb8,
	0xa7, 0x96, 0xf8, 0x24, 0x96, 0xa4, 0x16, 0xc3, 0x1c, 0xad, 0xd4, 0xc2, 0xc8, 0x25, 0x88, 0x24,
	0x38, 0x50, 0x4e, 0x93, 0xe0, 0x12, 0x83, 0xbb, 0x02, 0x25, 0x54, 0x95, 0x0c, 0xb9, 0xc4, 0x31,
	0x64, 0xa0, 0xae, 0xc4, 0x11, 0xe0, 0x4e, 0x31, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7,
	0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c,
	0xc7, 0x10, 0xe5, 0x94, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0x04, 0x72, 0x92, 0x3e, 0x3c, 0x9e, 0xe1,
	0x0c, 0x70, 0xe4, 0xea, 0x13, 0x48, 0x71, 0x49, 0x6c, 0x60, 0x65, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x2a, 0x9d, 0xa0, 0x5f, 0x9b, 0x02, 0x00, 0x00,
}

func (m *GetByHeightRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetByHeightRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetByHeightRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetByHeightResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetByHeightResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetByHeightResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BlockId != nil {
		{
			size, err := m.BlockId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetLatestRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLatestRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLatestRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetLatestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLatestResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLatestResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BlockId != nil {
		{
			size, err := m.BlockId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetLatestHeightRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLatestHeightRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLatestHeightRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetLatestHeightResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLatestHeightResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLatestHeightResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetByHeightRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovBlock(uint64(m.Height))
	}
	return n
}

func (m *GetByHeightResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockId != nil {
		l = m.BlockId.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	return n
}

func (m *GetLatestRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetLatestResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockId != nil {
		l = m.BlockId.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	return n
}

func (m *GetLatestHeightRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetLatestHeightResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovBlock(uint64(m.Height))
	}
	return n
}

func sovBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetByHeightRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: GetByHeightRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetByHeightRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func (m *GetByHeightResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: GetByHeightResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetByHeightResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockId == nil {
				m.BlockId = &v1beta1.BlockID{}
			}
			if err := m.BlockId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Block == nil {
				m.Block = &v1beta3.Block{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func (m *GetLatestRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: GetLatestRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLatestRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func (m *GetLatestResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: GetLatestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLatestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockId == nil {
				m.BlockId = &v1beta1.BlockID{}
			}
			if err := m.BlockId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Block == nil {
				m.Block = &v1beta3.Block{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func (m *GetLatestHeightRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: GetLatestHeightRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLatestHeightRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func (m *GetLatestHeightResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: GetLatestHeightResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLatestHeightResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
				return 0, ErrInvalidLengthBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlock = fmt.Errorf("proto: unexpected end of group")
)