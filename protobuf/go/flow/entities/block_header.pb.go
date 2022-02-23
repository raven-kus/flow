// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow/entities/block_header.proto

package entities

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BlockHeader struct {
	Id                   []byte                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId             []byte                 `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Height               uint64                 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp            *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	View                 uint64                 `protobuf:"varint,5,opt,name=view,proto3" json:"view,omitempty"`
	ParentVoterIds       [][]byte               `protobuf:"bytes,6,rep,name=parent_voter_ids,json=parentVoterIds,proto3" json:"parent_voter_ids,omitempty"`
	ParentVoterSigData   []byte                 `protobuf:"bytes,7,opt,name=parent_voter_sig_data,json=parentVoterSigData,proto3" json:"parent_voter_sig_data,omitempty"`
	ProposerId           []byte                 `protobuf:"bytes,8,opt,name=proposer_id,json=proposerId,proto3" json:"proposer_id,omitempty"`
	ProposerSigData      []byte                 `protobuf:"bytes,9,opt,name=proposer_sig_data,json=proposerSigData,proto3" json:"proposer_sig_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9d8363da0276a74, []int{0}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *BlockHeader) GetParentId() []byte {
	if m != nil {
		return m.ParentId
	}
	return nil
}

func (m *BlockHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeader) GetTimestamp() *timestamppb.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *BlockHeader) GetView() uint64 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *BlockHeader) GetParentVoterIds() [][]byte {
	if m != nil {
		return m.ParentVoterIds
	}
	return nil
}

func (m *BlockHeader) GetParentVoterSigData() []byte {
	if m != nil {
		return m.ParentVoterSigData
	}
	return nil
}

func (m *BlockHeader) GetProposerId() []byte {
	if m != nil {
		return m.ProposerId
	}
	return nil
}

func (m *BlockHeader) GetProposerSigData() []byte {
	if m != nil {
		return m.ProposerSigData
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "flow.entities.BlockHeader")
}

func init() { proto.RegisterFile("flow/entities/block_header.proto", fileDescriptor_b9d8363da0276a74) }

var fileDescriptor_b9d8363da0276a74 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x3f, 0x4f, 0xf3, 0x30,
	0x18, 0xc4, 0x95, 0xb4, 0x6f, 0xdf, 0xf6, 0x69, 0x29, 0x60, 0x09, 0x64, 0x15, 0xa4, 0x46, 0x4c,
	0x11, 0x83, 0xcd, 0x9f, 0x85, 0xb9, 0x62, 0xa0, 0x1b, 0x0a, 0x88, 0x81, 0x25, 0x72, 0x6a, 0xd7,
	0xb1, 0x68, 0xfb, 0x44, 0x89, 0xdb, 0x7e, 0x47, 0x3e, 0x15, 0xaa, 0x13, 0x07, 0xba, 0x44, 0xf1,
	0xdd, 0x4f, 0xe7, 0xb3, 0x0e, 0xa2, 0xe5, 0x0a, 0xf7, 0x5c, 0x6d, 0xac, 0xb1, 0x46, 0x55, 0x3c,
	0x5b, 0xe1, 0xe2, 0x2b, 0xcd, 0x95, 0x90, 0xaa, 0x64, 0x45, 0x89, 0x16, 0xc9, 0xc9, 0x81, 0x60,
	0x9e, 0x98, 0x4c, 0x35, 0xa2, 0x5e, 0x29, 0xee, 0xcc, 0x6c, 0xbb, 0xe4, 0xd6, 0xac, 0x55, 0x65,
	0xc5, 0xba, 0xa8, 0xf9, 0x9b, 0xef, 0x10, 0x86, 0xb3, 0x43, 0xcc, 0x8b, 0x4b, 0x21, 0x63, 0x08,
	0x8d, 0xa4, 0x41, 0x14, 0xc4, 0xa3, 0x24, 0x34, 0x92, 0x5c, 0xc1, 0xa0, 0x10, 0xa5, 0xda, 0xd8,
	0xd4, 0x48, 0x1a, 0x3a, 0xb9, 0x5f, 0x0b, 0x73, 0x49, 0x2e, 0xa1, 0x97, 0x2b, 0xa3, 0x73, 0x4b,
	0x3b, 0x51, 0x10, 0x77, 0x93, 0xe6, 0x44, 0x9e, 0x60, 0xd0, 0xde, 0x43, 0xbb, 0x51, 0x10, 0x0f,
	0x1f, 0x26, 0xac, 0x6e, 0xc2, 0x7c, 0x13, 0xf6, 0xee, 0x89, 0xe4, 0x17, 0x26, 0x04, 0xba, 0x3b,
	0xa3, 0xf6, 0xf4, 0x9f, 0xcb, 0x73, 0xff, 0x24, 0x86, 0xb3, 0xa6, 0xc2, 0x0e, 0xad, 0x2a, 0x53,
	0x23, 0x2b, 0xda, 0x8b, 0x3a, 0xf1, 0x28, 0x19, 0xd7, 0xfa, 0xc7, 0x41, 0x9e, 0xcb, 0x8a, 0xdc,
	0xc3, 0xc5, 0x11, 0x59, 0x19, 0x9d, 0x4a, 0x61, 0x05, 0xfd, 0xef, 0x8a, 0x93, 0x3f, 0xf8, 0x9b,
	0xd1, 0xcf, 0xc2, 0x0a, 0x32, 0x85, 0x61, 0x51, 0x62, 0x81, 0x95, 0x0b, 0xa6, 0x7d, 0x07, 0x82,
	0x97, 0xe6, 0x92, 0xdc, 0xc2, 0x79, 0x0b, 0xb4, 0x79, 0x03, 0x87, 0x9d, 0x7a, 0xa3, 0x09, 0x9b,
	0xbd, 0xc2, 0x35, 0x96, 0x9a, 0xe1, 0xc6, 0x8d, 0xd0, 0xbe, 0xd4, 0xaf, 0xf1, 0x79, 0xa7, 0x8d,
	0xcd, 0xb7, 0x19, 0x5b, 0xe0, 0x9a, 0xd7, 0x10, 0x77, 0x9f, 0x76, 0x1d, 0x8d, 0xfc, 0x68, 0xe1,
	0xac, 0xe7, 0xac, 0xc7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x20, 0x2a, 0xc8, 0x71, 0xf9, 0x01,
	0x00, 0x00,
}
