// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/privval/types.proto

package privval

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	keys "github.com/okex/exchain/libs/tendermint/proto/crypto/keys"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// FilePVKey stores the immutable part of PrivValidator.
type FilePVKey struct {
	Address              []byte          `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey               keys.PublicKey  `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key"`
	PrivKey              keys.PrivateKey `protobuf:"bytes,3,opt,name=priv_key,json=privKey,proto3" json:"priv_key"`
	FilePath             string          `protobuf:"bytes,4,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FilePVKey) Reset()         { *m = FilePVKey{} }
func (m *FilePVKey) String() string { return proto.CompactTextString(m) }
func (*FilePVKey) ProtoMessage()    {}
func (*FilePVKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d74c406df3ad93, []int{0}
}
func (m *FilePVKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilePVKey.Unmarshal(m, b)
}
func (m *FilePVKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilePVKey.Marshal(b, m, deterministic)
}
func (m *FilePVKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilePVKey.Merge(m, src)
}
func (m *FilePVKey) XXX_Size() int {
	return xxx_messageInfo_FilePVKey.Size(m)
}
func (m *FilePVKey) XXX_DiscardUnknown() {
	xxx_messageInfo_FilePVKey.DiscardUnknown(m)
}

var xxx_messageInfo_FilePVKey proto.InternalMessageInfo

func (m *FilePVKey) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *FilePVKey) GetPubKey() keys.PublicKey {
	if m != nil {
		return m.PubKey
	}
	return keys.PublicKey{}
}

func (m *FilePVKey) GetPrivKey() keys.PrivateKey {
	if m != nil {
		return m.PrivKey
	}
	return keys.PrivateKey{}
}

func (m *FilePVKey) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

// FilePVLastSignState stores the mutable part of PrivValidator.
type FilePVLastSignState struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Round                int64    `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Step                 int32    `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	SignBytes            []byte   `protobuf:"bytes,5,opt,name=sign_bytes,json=signBytes,proto3" json:"sign_bytes,omitempty"`
	FilePath             string   `protobuf:"bytes,6,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilePVLastSignState) Reset()         { *m = FilePVLastSignState{} }
func (m *FilePVLastSignState) String() string { return proto.CompactTextString(m) }
func (*FilePVLastSignState) ProtoMessage()    {}
func (*FilePVLastSignState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d74c406df3ad93, []int{1}
}
func (m *FilePVLastSignState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilePVLastSignState.Unmarshal(m, b)
}
func (m *FilePVLastSignState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilePVLastSignState.Marshal(b, m, deterministic)
}
func (m *FilePVLastSignState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilePVLastSignState.Merge(m, src)
}
func (m *FilePVLastSignState) XXX_Size() int {
	return xxx_messageInfo_FilePVLastSignState.Size(m)
}
func (m *FilePVLastSignState) XXX_DiscardUnknown() {
	xxx_messageInfo_FilePVLastSignState.DiscardUnknown(m)
}

var xxx_messageInfo_FilePVLastSignState proto.InternalMessageInfo

func (m *FilePVLastSignState) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *FilePVLastSignState) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *FilePVLastSignState) GetStep() int32 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *FilePVLastSignState) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *FilePVLastSignState) GetSignBytes() []byte {
	if m != nil {
		return m.SignBytes
	}
	return nil
}

func (m *FilePVLastSignState) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func init() {
	proto.RegisterType((*FilePVKey)(nil), "tendermint.proto.privval.FilePVKey")
	proto.RegisterType((*FilePVLastSignState)(nil), "tendermint.proto.privval.FilePVLastSignState")
}

func init() { proto.RegisterFile("proto/privval/types.proto", fileDescriptor_a9d74c406df3ad93) }

var fileDescriptor_a9d74c406df3ad93 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6e, 0xe2, 0x30,
	0x14, 0x9e, 0x0c, 0x10, 0x88, 0x87, 0x95, 0x67, 0x34, 0xca, 0x30, 0x45, 0x45, 0x2c, 0xda, 0xac,
	0x92, 0xaa, 0xbd, 0x01, 0x0b, 0xa4, 0x8a, 0x2e, 0x50, 0x90, 0xba, 0xe8, 0x26, 0x72, 0xc8, 0x6b,
	0x62, 0x11, 0x12, 0xcb, 0x7e, 0x41, 0xf2, 0xb1, 0x7a, 0x8b, 0x5e, 0xa0, 0xdb, 0x9e, 0xa5, 0xb2,
	0x43, 0x15, 0x50, 0x17, 0xdd, 0xbd, 0xef, 0xf3, 0xf3, 0xf7, 0x63, 0x99, 0xfc, 0x13, 0xb2, 0xc6,
	0x3a, 0x12, 0x92, 0x1f, 0x0e, 0xac, 0x8c, 0x50, 0x0b, 0x50, 0xa1, 0xe5, 0xa8, 0x8f, 0x50, 0x65,
	0x20, 0xf7, 0xbc, 0xc2, 0x96, 0x09, 0x8f, 0x5b, 0x93, 0x2b, 0x2c, 0xb8, 0xcc, 0x12, 0xc1, 0x24,
	0xea, 0xa8, 0x15, 0xc8, 0xeb, 0xbc, 0xee, 0xa6, 0x76, 0x7f, 0x32, 0x6d, 0x99, 0xad, 0xd4, 0x02,
	0xeb, 0x68, 0x07, 0x5a, 0x9d, 0x1a, 0xcc, 0xdf, 0x1c, 0xe2, 0x2d, 0x79, 0x09, 0xeb, 0xc7, 0x15,
	0x68, 0xea, 0x93, 0x21, 0xcb, 0x32, 0x09, 0x4a, 0xf9, 0xce, 0xcc, 0x09, 0xc6, 0xf1, 0x27, 0xa4,
	0x4b, 0x32, 0x14, 0x4d, 0x9a, 0xec, 0x40, 0xfb, 0x3f, 0x67, 0x4e, 0xf0, 0xeb, 0xf6, 0x3a, 0xfc,
	0x12, 0xad, 0xf5, 0x08, 0x8d, 0x47, 0xb8, 0x6e, 0xd2, 0x92, 0x6f, 0x57, 0xa0, 0x17, 0xfd, 0xd7,
	0xf7, 0xcb, 0x1f, 0xb1, 0x2b, 0x9a, 0xd4, 0x38, 0xdc, 0x93, 0x91, 0x69, 0x60, 0x85, 0x7a, 0x56,
	0x28, 0xf8, 0x46, 0x48, 0xf2, 0x03, 0x43, 0xe8, 0x94, 0x86, 0xe6, 0xbe, 0x91, 0xfa, 0x4f, 0xbc,
	0x67, 0x5e, 0x42, 0x22, 0x18, 0x16, 0x7e, 0x7f, 0xe6, 0x04, 0x5e, 0x3c, 0x32, 0xc4, 0x9a, 0x61,
	0x31, 0x7f, 0x71, 0xc8, 0xef, 0xb6, 0xd7, 0x03, 0x53, 0xb8, 0xe1, 0x79, 0xb5, 0x41, 0x86, 0x40,
	0xff, 0x12, 0xb7, 0x00, 0x9e, 0x17, 0x68, 0x0b, 0xf6, 0xe2, 0x23, 0xa2, 0x7f, 0xc8, 0x40, 0xd6,
	0x4d, 0x95, 0xd9, 0x76, 0xbd, 0xb8, 0x05, 0x94, 0x92, 0xbe, 0x42, 0x10, 0x36, 0xe9, 0x20, 0xb6,
	0x33, 0xbd, 0x20, 0x9e, 0xe2, 0x79, 0xc5, 0xb0, 0x91, 0x60, 0x6d, 0xc7, 0x71, 0x47, 0xd0, 0x29,
	0x21, 0x06, 0x24, 0xa9, 0x46, 0x50, 0xfe, 0xa0, 0x3b, 0x5e, 0x18, 0xe2, 0x3c, 0xb3, 0x7b, 0x9e,
	0x79, 0x71, 0xf3, 0x14, 0xe6, 0x1c, 0x8b, 0x26, 0x0d, 0xb7, 0xf5, 0x3e, 0xea, 0x5e, 0xe5, 0x74,
	0x3c, 0xfb, 0x2a, 0xa9, 0x6b, 0xe1, 0xdd, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xa3, 0x78,
	0xe9, 0x42, 0x02, 0x00, 0x00,
}