// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cometbft/types/v1beta1/canonical.proto

package v1beta1

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

type CanonicalBlockID struct {
	Hash          []byte                 `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	PartSetHeader CanonicalPartSetHeader `protobuf:"bytes,2,opt,name=part_set_header,json=partSetHeader,proto3" json:"part_set_header"`
}

func (m *CanonicalBlockID) Reset()         { *m = CanonicalBlockID{} }
func (m *CanonicalBlockID) String() string { return proto.CompactTextString(m) }
func (*CanonicalBlockID) ProtoMessage()    {}
func (*CanonicalBlockID) Descriptor() ([]byte, []int) {
	return fileDescriptor_476991b7998505b4, []int{0}
}
func (m *CanonicalBlockID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CanonicalBlockID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CanonicalBlockID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CanonicalBlockID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CanonicalBlockID.Merge(m, src)
}
func (m *CanonicalBlockID) XXX_Size() int {
	return m.Size()
}
func (m *CanonicalBlockID) XXX_DiscardUnknown() {
	xxx_messageInfo_CanonicalBlockID.DiscardUnknown(m)
}

var xxx_messageInfo_CanonicalBlockID proto.InternalMessageInfo

func (m *CanonicalBlockID) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *CanonicalBlockID) GetPartSetHeader() CanonicalPartSetHeader {
	if m != nil {
		return m.PartSetHeader
	}
	return CanonicalPartSetHeader{}
}

type CanonicalPartSetHeader struct {
	Total uint32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Hash  []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *CanonicalPartSetHeader) Reset()         { *m = CanonicalPartSetHeader{} }
func (m *CanonicalPartSetHeader) String() string { return proto.CompactTextString(m) }
func (*CanonicalPartSetHeader) ProtoMessage()    {}
func (*CanonicalPartSetHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_476991b7998505b4, []int{1}
}
func (m *CanonicalPartSetHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CanonicalPartSetHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CanonicalPartSetHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CanonicalPartSetHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CanonicalPartSetHeader.Merge(m, src)
}
func (m *CanonicalPartSetHeader) XXX_Size() int {
	return m.Size()
}
func (m *CanonicalPartSetHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_CanonicalPartSetHeader.DiscardUnknown(m)
}

var xxx_messageInfo_CanonicalPartSetHeader proto.InternalMessageInfo

func (m *CanonicalPartSetHeader) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *CanonicalPartSetHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type CanonicalProposal struct {
	Type      SignedMsgType     `protobuf:"varint,1,opt,name=type,proto3,enum=cometbft.types.v1beta1.SignedMsgType" json:"type,omitempty"`
	Height    int64             `protobuf:"fixed64,2,opt,name=height,proto3" json:"height,omitempty"`
	Round     int64             `protobuf:"fixed64,3,opt,name=round,proto3" json:"round,omitempty"`
	POLRound  int64             `protobuf:"varint,4,opt,name=pol_round,json=polRound,proto3" json:"pol_round,omitempty"`
	BlockID   *CanonicalBlockID `protobuf:"bytes,5,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Timestamp time.Time         `protobuf:"bytes,6,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	ChainID   string            `protobuf:"bytes,7,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *CanonicalProposal) Reset()         { *m = CanonicalProposal{} }
func (m *CanonicalProposal) String() string { return proto.CompactTextString(m) }
func (*CanonicalProposal) ProtoMessage()    {}
func (*CanonicalProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_476991b7998505b4, []int{2}
}
func (m *CanonicalProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CanonicalProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CanonicalProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CanonicalProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CanonicalProposal.Merge(m, src)
}
func (m *CanonicalProposal) XXX_Size() int {
	return m.Size()
}
func (m *CanonicalProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CanonicalProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CanonicalProposal proto.InternalMessageInfo

func (m *CanonicalProposal) GetType() SignedMsgType {
	if m != nil {
		return m.Type
	}
	return UnknownType
}

func (m *CanonicalProposal) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *CanonicalProposal) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *CanonicalProposal) GetPOLRound() int64 {
	if m != nil {
		return m.POLRound
	}
	return 0
}

func (m *CanonicalProposal) GetBlockID() *CanonicalBlockID {
	if m != nil {
		return m.BlockID
	}
	return nil
}

func (m *CanonicalProposal) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *CanonicalProposal) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

type CanonicalVote struct {
	Type      SignedMsgType     `protobuf:"varint,1,opt,name=type,proto3,enum=cometbft.types.v1beta1.SignedMsgType" json:"type,omitempty"`
	Height    int64             `protobuf:"fixed64,2,opt,name=height,proto3" json:"height,omitempty"`
	Round     int64             `protobuf:"fixed64,3,opt,name=round,proto3" json:"round,omitempty"`
	BlockID   *CanonicalBlockID `protobuf:"bytes,4,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Timestamp time.Time         `protobuf:"bytes,5,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	ChainID   string            `protobuf:"bytes,6,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *CanonicalVote) Reset()         { *m = CanonicalVote{} }
func (m *CanonicalVote) String() string { return proto.CompactTextString(m) }
func (*CanonicalVote) ProtoMessage()    {}
func (*CanonicalVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_476991b7998505b4, []int{3}
}
func (m *CanonicalVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CanonicalVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CanonicalVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CanonicalVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CanonicalVote.Merge(m, src)
}
func (m *CanonicalVote) XXX_Size() int {
	return m.Size()
}
func (m *CanonicalVote) XXX_DiscardUnknown() {
	xxx_messageInfo_CanonicalVote.DiscardUnknown(m)
}

var xxx_messageInfo_CanonicalVote proto.InternalMessageInfo

func (m *CanonicalVote) GetType() SignedMsgType {
	if m != nil {
		return m.Type
	}
	return UnknownType
}

func (m *CanonicalVote) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *CanonicalVote) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *CanonicalVote) GetBlockID() *CanonicalBlockID {
	if m != nil {
		return m.BlockID
	}
	return nil
}

func (m *CanonicalVote) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *CanonicalVote) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func init() {
	proto.RegisterType((*CanonicalBlockID)(nil), "cometbft.types.v1beta1.CanonicalBlockID")
	proto.RegisterType((*CanonicalPartSetHeader)(nil), "cometbft.types.v1beta1.CanonicalPartSetHeader")
	proto.RegisterType((*CanonicalProposal)(nil), "cometbft.types.v1beta1.CanonicalProposal")
	proto.RegisterType((*CanonicalVote)(nil), "cometbft.types.v1beta1.CanonicalVote")
}

func init() {
	proto.RegisterFile("cometbft/types/v1beta1/canonical.proto", fileDescriptor_476991b7998505b4)
}

var fileDescriptor_476991b7998505b4 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb3, 0xa9, 0x93, 0x38, 0xdb, 0x06, 0xca, 0xaa, 0x8a, 0xac, 0x1c, 0xec, 0xc8, 0x12,
	0x95, 0xb9, 0xd8, 0x6a, 0x39, 0x20, 0xae, 0x2e, 0x07, 0x22, 0x81, 0x08, 0x6e, 0xc5, 0x01, 0x21,
	0x45, 0x6b, 0x7b, 0x6b, 0x5b, 0x38, 0xde, 0x95, 0xbd, 0x41, 0xea, 0x9d, 0x07, 0xe8, 0xd3, 0xf0,
	0x0c, 0x3d, 0xf6, 0xc8, 0x29, 0x20, 0x47, 0xbc, 0x07, 0xda, 0xdd, 0xc4, 0xe9, 0x21, 0x11, 0x07,
	0x10, 0xb7, 0x99, 0xf1, 0x3f, 0x33, 0x9f, 0xff, 0xd1, 0xc2, 0xd3, 0x88, 0xce, 0x09, 0x0f, 0xaf,
	0xb9, 0xc7, 0x6f, 0x18, 0xa9, 0xbc, 0x2f, 0x67, 0x21, 0xe1, 0xf8, 0xcc, 0x8b, 0x70, 0x41, 0x8b,
	0x2c, 0xc2, 0xb9, 0xcb, 0x4a, 0xca, 0x29, 0x1a, 0x6e, 0x74, 0xae, 0xd4, 0xb9, 0x6b, 0xdd, 0xe8,
	0x24, 0xa1, 0x09, 0x95, 0x12, 0x4f, 0x44, 0x4a, 0x3d, 0xb2, 0xf7, 0x4c, 0x55, 0xbd, 0x4a, 0x63,
	0x25, 0x94, 0x26, 0x39, 0xf1, 0x64, 0x16, 0x2e, 0xae, 0x3d, 0x9e, 0xcd, 0x49, 0xc5, 0xf1, 0x9c,
	0x29, 0x81, 0xfd, 0x15, 0xc0, 0xe3, 0x8b, 0x0d, 0x86, 0x9f, 0xd3, 0xe8, 0xf3, 0xe4, 0x15, 0x42,
	0x50, 0x4b, 0x71, 0x95, 0x1a, 0x60, 0x0c, 0x9c, 0xa3, 0x40, 0xc6, 0xe8, 0x13, 0x7c, 0xcc, 0x70,
	0xc9, 0x67, 0x15, 0xe1, 0xb3, 0x94, 0xe0, 0x98, 0x94, 0x46, 0x7b, 0x0c, 0x9c, 0xc3, 0x73, 0xd7,
	0xdd, 0x4d, 0xed, 0x36, 0x63, 0xa7, 0xb8, 0xe4, 0x97, 0x84, 0xbf, 0x96, 0x5d, 0xbe, 0x76, 0xb7,
	0xb4, 0x5a, 0xc1, 0x80, 0x3d, 0x2c, 0xda, 0x3e, 0x1c, 0xee, 0x96, 0xa3, 0x13, 0xd8, 0xe1, 0x94,
	0xe3, 0x5c, 0xc2, 0x0c, 0x02, 0x95, 0x34, 0x84, 0xed, 0x2d, 0xa1, 0xfd, 0xab, 0x0d, 0x9f, 0x6c,
	0x87, 0x94, 0x94, 0xd1, 0x0a, 0xe7, 0xe8, 0x25, 0xd4, 0x04, 0x96, 0x6c, 0x7f, 0x74, 0xfe, 0x74,
	0x1f, 0xec, 0x65, 0x96, 0x14, 0x24, 0x7e, 0x5b, 0x25, 0x57, 0x37, 0x8c, 0x04, 0xb2, 0x05, 0x0d,
	0x61, 0x37, 0x25, 0x59, 0x92, 0x72, 0xb9, 0xe6, 0x38, 0x58, 0x67, 0x02, 0xa9, 0xa4, 0x8b, 0x22,
	0x36, 0x0e, 0x64, 0x59, 0x25, 0xe8, 0x19, 0xec, 0x33, 0x9a, 0xcf, 0xd4, 0x17, 0x6d, 0x0c, 0x9c,
	0x03, 0xff, 0xa8, 0x5e, 0x5a, 0xfa, 0xf4, 0xdd, 0x9b, 0x40, 0xd4, 0x02, 0x9d, 0xd1, 0x5c, 0x46,
	0x68, 0x0a, 0xf5, 0x50, 0x58, 0x3d, 0xcb, 0x62, 0xa3, 0x23, 0x4d, 0x74, 0xfe, 0x68, 0xe2, 0xfa,
	0x36, 0xfe, 0x61, 0xbd, 0xb4, 0x7a, 0xeb, 0x24, 0xe8, 0xc9, 0x31, 0x93, 0x18, 0xf9, 0xb0, 0xdf,
	0x5c, 0xd6, 0xe8, 0xca, 0x91, 0x23, 0x57, 0xdd, 0xde, 0xdd, 0xdc, 0xde, 0xbd, 0xda, 0x28, 0x7c,
	0x5d, 0xdc, 0xe0, 0xf6, 0x87, 0x05, 0x82, 0x6d, 0x1b, 0x3a, 0x85, 0x7a, 0x94, 0xe2, 0xac, 0x10,
	0x54, 0xbd, 0x31, 0x70, 0xfa, 0x6a, 0xd7, 0x85, 0xa8, 0x89, 0x5d, 0xf2, 0xe3, 0x24, 0xb6, 0xbf,
	0xb5, 0xe1, 0xa0, 0xc1, 0xfa, 0x40, 0x39, 0xf9, 0x7f, 0x1e, 0x3f, 0x34, 0x4e, 0xfb, 0xf7, 0xc6,
	0x75, 0xfe, 0xde, 0xb8, 0xee, 0x7e, 0xe3, 0xfc, 0xf7, 0x77, 0xb5, 0x09, 0xee, 0x6b, 0x13, 0xfc,
	0xac, 0x4d, 0x70, 0xbb, 0x32, 0x5b, 0xf7, 0x2b, 0xb3, 0xf5, 0x7d, 0x65, 0xb6, 0x3e, 0xbe, 0x48,
	0x32, 0x9e, 0x2e, 0x42, 0xf1, 0x2f, 0x5e, 0xf3, 0xaa, 0x9b, 0x00, 0xb3, 0xcc, 0xdb, 0xfd, 0xd6,
	0xc3, 0xae, 0x64, 0x7c, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x10, 0x33, 0xb1, 0x97, 0x62, 0x04,
	0x00, 0x00,
}

func (m *CanonicalBlockID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CanonicalBlockID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CanonicalBlockID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PartSetHeader.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCanonical(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintCanonical(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CanonicalPartSetHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CanonicalPartSetHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CanonicalPartSetHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintCanonical(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Total != 0 {
		i = encodeVarintCanonical(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CanonicalProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CanonicalProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CanonicalProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintCanonical(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x3a
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintCanonical(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x32
	if m.BlockID != nil {
		{
			size, err := m.BlockID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCanonical(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.POLRound != 0 {
		i = encodeVarintCanonical(dAtA, i, uint64(m.POLRound))
		i--
		dAtA[i] = 0x20
	}
	if m.Round != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Round))
		i--
		dAtA[i] = 0x19
	}
	if m.Height != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Height))
		i--
		dAtA[i] = 0x11
	}
	if m.Type != 0 {
		i = encodeVarintCanonical(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CanonicalVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CanonicalVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CanonicalVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintCanonical(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x32
	}
	n4, err4 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintCanonical(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x2a
	if m.BlockID != nil {
		{
			size, err := m.BlockID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCanonical(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Round != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Round))
		i--
		dAtA[i] = 0x19
	}
	if m.Height != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Height))
		i--
		dAtA[i] = 0x11
	}
	if m.Type != 0 {
		i = encodeVarintCanonical(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCanonical(dAtA []byte, offset int, v uint64) int {
	offset -= sovCanonical(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CanonicalBlockID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovCanonical(uint64(l))
	}
	l = m.PartSetHeader.Size()
	n += 1 + l + sovCanonical(uint64(l))
	return n
}

func (m *CanonicalPartSetHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Total != 0 {
		n += 1 + sovCanonical(uint64(m.Total))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovCanonical(uint64(l))
	}
	return n
}

func (m *CanonicalProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCanonical(uint64(m.Type))
	}
	if m.Height != 0 {
		n += 9
	}
	if m.Round != 0 {
		n += 9
	}
	if m.POLRound != 0 {
		n += 1 + sovCanonical(uint64(m.POLRound))
	}
	if m.BlockID != nil {
		l = m.BlockID.Size()
		n += 1 + l + sovCanonical(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovCanonical(uint64(l))
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovCanonical(uint64(l))
	}
	return n
}

func (m *CanonicalVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCanonical(uint64(m.Type))
	}
	if m.Height != 0 {
		n += 9
	}
	if m.Round != 0 {
		n += 9
	}
	if m.BlockID != nil {
		l = m.BlockID.Size()
		n += 1 + l + sovCanonical(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovCanonical(uint64(l))
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovCanonical(uint64(l))
	}
	return n
}

func sovCanonical(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCanonical(x uint64) (n int) {
	return sovCanonical(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CanonicalBlockID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCanonical
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
			return fmt.Errorf("proto: CanonicalBlockID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CanonicalBlockID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartSetHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PartSetHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCanonical(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCanonical
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
func (m *CanonicalPartSetHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCanonical
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
			return fmt.Errorf("proto: CanonicalPartSetHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CanonicalPartSetHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCanonical(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCanonical
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
func (m *CanonicalProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCanonical
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
			return fmt.Errorf("proto: CanonicalProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CanonicalProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= SignedMsgType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Height = int64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Round = int64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field POLRound", wireType)
			}
			m.POLRound = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.POLRound |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockID == nil {
				m.BlockID = &CanonicalBlockID{}
			}
			if err := m.BlockID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCanonical(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCanonical
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
func (m *CanonicalVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCanonical
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
			return fmt.Errorf("proto: CanonicalVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CanonicalVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= SignedMsgType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Height = int64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Round = int64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockID == nil {
				m.BlockID = &CanonicalBlockID{}
			}
			if err := m.BlockID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonical
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
				return ErrInvalidLengthCanonical
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCanonical
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCanonical(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCanonical
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
func skipCanonical(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCanonical
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
					return 0, ErrIntOverflowCanonical
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
					return 0, ErrIntOverflowCanonical
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
				return 0, ErrInvalidLengthCanonical
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCanonical
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCanonical
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCanonical        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCanonical          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCanonical = fmt.Errorf("proto: unexpected end of group")
)