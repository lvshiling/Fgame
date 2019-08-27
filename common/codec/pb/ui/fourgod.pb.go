// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fourgod.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FourGodBio struct {
	NcpId            *int64    `protobuf:"varint,1,req,name=ncpId" json:"ncpId,omitempty"`
	Typ              *int32    `protobuf:"varint,2,req,name=typ" json:"typ,omitempty"`
	Status           *int32    `protobuf:"varint,3,req,name=status" json:"status,omitempty"`
	StatusTime       *int64    `protobuf:"varint,4,req,name=statusTime" json:"statusTime,omitempty"`
	Pos              *Position `protobuf:"bytes,5,opt,name=pos" json:"pos,omitempty"`
	BiologyId        *int32    `protobuf:"varint,6,req,name=biologyId" json:"biologyId,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *FourGodBio) Reset()                    { *m = FourGodBio{} }
func (m *FourGodBio) String() string            { return proto.CompactTextString(m) }
func (*FourGodBio) ProtoMessage()               {}
func (*FourGodBio) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{0} }

func (m *FourGodBio) GetNcpId() int64 {
	if m != nil && m.NcpId != nil {
		return *m.NcpId
	}
	return 0
}

func (m *FourGodBio) GetTyp() int32 {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return 0
}

func (m *FourGodBio) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *FourGodBio) GetStatusTime() int64 {
	if m != nil && m.StatusTime != nil {
		return *m.StatusTime
	}
	return 0
}

func (m *FourGodBio) GetPos() *Position {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *FourGodBio) GetBiologyId() int32 {
	if m != nil && m.BiologyId != nil {
		return *m.BiologyId
	}
	return 0
}

type SCFourGodGet struct {
	KeyNum           *int32        `protobuf:"varint,1,req,name=keyNum" json:"keyNum,omitempty"`
	BioList          []*FourGodBio `protobuf:"bytes,2,rep,name=bioList" json:"bioList,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *SCFourGodGet) Reset()                    { *m = SCFourGodGet{} }
func (m *SCFourGodGet) String() string            { return proto.CompactTextString(m) }
func (*SCFourGodGet) ProtoMessage()               {}
func (*SCFourGodGet) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{1} }

func (m *SCFourGodGet) GetKeyNum() int32 {
	if m != nil && m.KeyNum != nil {
		return *m.KeyNum
	}
	return 0
}

func (m *SCFourGodGet) GetBioList() []*FourGodBio {
	if m != nil {
		return m.BioList
	}
	return nil
}

type CSFourGodOpenBox struct {
	NpcId            *int64 `protobuf:"varint,1,req,name=npcId" json:"npcId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSFourGodOpenBox) Reset()                    { *m = CSFourGodOpenBox{} }
func (m *CSFourGodOpenBox) String() string            { return proto.CompactTextString(m) }
func (*CSFourGodOpenBox) ProtoMessage()               {}
func (*CSFourGodOpenBox) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{2} }

func (m *CSFourGodOpenBox) GetNpcId() int64 {
	if m != nil && m.NpcId != nil {
		return *m.NpcId
	}
	return 0
}

type SCFourGodOpenBox struct {
	NpcId            *int64 `protobuf:"varint,1,req,name=npcId" json:"npcId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCFourGodOpenBox) Reset()                    { *m = SCFourGodOpenBox{} }
func (m *SCFourGodOpenBox) String() string            { return proto.CompactTextString(m) }
func (*SCFourGodOpenBox) ProtoMessage()               {}
func (*SCFourGodOpenBox) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{3} }

func (m *SCFourGodOpenBox) GetNpcId() int64 {
	if m != nil && m.NpcId != nil {
		return *m.NpcId
	}
	return 0
}

type SCFourGodOpenBoxStop struct {
	NpcId            *int64 `protobuf:"varint,1,req,name=npcId" json:"npcId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCFourGodOpenBoxStop) Reset()                    { *m = SCFourGodOpenBoxStop{} }
func (m *SCFourGodOpenBoxStop) String() string            { return proto.CompactTextString(m) }
func (*SCFourGodOpenBoxStop) ProtoMessage()               {}
func (*SCFourGodOpenBoxStop) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{4} }

func (m *SCFourGodOpenBoxStop) GetNpcId() int64 {
	if m != nil && m.NpcId != nil {
		return *m.NpcId
	}
	return 0
}

type SCFourGodOpenBoxFinish struct {
	NpcId            *int64 `protobuf:"varint,1,req,name=npcId" json:"npcId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCFourGodOpenBoxFinish) Reset()                    { *m = SCFourGodOpenBoxFinish{} }
func (m *SCFourGodOpenBoxFinish) String() string            { return proto.CompactTextString(m) }
func (*SCFourGodOpenBoxFinish) ProtoMessage()               {}
func (*SCFourGodOpenBoxFinish) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{5} }

func (m *SCFourGodOpenBoxFinish) GetNpcId() int64 {
	if m != nil && m.NpcId != nil {
		return *m.NpcId
	}
	return 0
}

type SCFourGodKeyNumChange struct {
	KeyNum           *int32 `protobuf:"varint,1,req,name=keyNum" json:"keyNum,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCFourGodKeyNumChange) Reset()                    { *m = SCFourGodKeyNumChange{} }
func (m *SCFourGodKeyNumChange) String() string            { return proto.CompactTextString(m) }
func (*SCFourGodKeyNumChange) ProtoMessage()               {}
func (*SCFourGodKeyNumChange) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{6} }

func (m *SCFourGodKeyNumChange) GetKeyNum() int32 {
	if m != nil && m.KeyNum != nil {
		return *m.KeyNum
	}
	return 0
}

type SCFourGodBioBroadcast struct {
	Bio              *FourGodBio `protobuf:"bytes,1,req,name=bio" json:"bio,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SCFourGodBioBroadcast) Reset()                    { *m = SCFourGodBioBroadcast{} }
func (m *SCFourGodBioBroadcast) String() string            { return proto.CompactTextString(m) }
func (*SCFourGodBioBroadcast) ProtoMessage()               {}
func (*SCFourGodBioBroadcast) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{7} }

func (m *SCFourGodBioBroadcast) GetBio() *FourGodBio {
	if m != nil {
		return m.Bio
	}
	return nil
}

type SCFourGodTotal struct {
	Exp              *int64      `protobuf:"varint,1,req,name=exp" json:"exp,omitempty"`
	ItemList         []*ItemInfo `protobuf:"bytes,2,rep,name=itemList" json:"itemList,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SCFourGodTotal) Reset()                    { *m = SCFourGodTotal{} }
func (m *SCFourGodTotal) String() string            { return proto.CompactTextString(m) }
func (*SCFourGodTotal) ProtoMessage()               {}
func (*SCFourGodTotal) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{8} }

func (m *SCFourGodTotal) GetExp() int64 {
	if m != nil && m.Exp != nil {
		return *m.Exp
	}
	return 0
}

func (m *SCFourGodTotal) GetItemList() []*ItemInfo {
	if m != nil {
		return m.ItemList
	}
	return nil
}

type CSFourGodUseMasked struct {
	Result           *bool  `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSFourGodUseMasked) Reset()                    { *m = CSFourGodUseMasked{} }
func (m *CSFourGodUseMasked) String() string            { return proto.CompactTextString(m) }
func (*CSFourGodUseMasked) ProtoMessage()               {}
func (*CSFourGodUseMasked) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{9} }

func (m *CSFourGodUseMasked) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

type SCFourGodUseMasked struct {
	Result           *bool  `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCFourGodUseMasked) Reset()                    { *m = SCFourGodUseMasked{} }
func (m *SCFourGodUseMasked) String() string            { return proto.CompactTextString(m) }
func (*SCFourGodUseMasked) ProtoMessage()               {}
func (*SCFourGodUseMasked) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{10} }

func (m *SCFourGodUseMasked) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*FourGodBio)(nil), "ui.FourGodBio")
	proto.RegisterType((*SCFourGodGet)(nil), "ui.SCFourGodGet")
	proto.RegisterType((*CSFourGodOpenBox)(nil), "ui.CSFourGodOpenBox")
	proto.RegisterType((*SCFourGodOpenBox)(nil), "ui.SCFourGodOpenBox")
	proto.RegisterType((*SCFourGodOpenBoxStop)(nil), "ui.SCFourGodOpenBoxStop")
	proto.RegisterType((*SCFourGodOpenBoxFinish)(nil), "ui.SCFourGodOpenBoxFinish")
	proto.RegisterType((*SCFourGodKeyNumChange)(nil), "ui.SCFourGodKeyNumChange")
	proto.RegisterType((*SCFourGodBioBroadcast)(nil), "ui.SCFourGodBioBroadcast")
	proto.RegisterType((*SCFourGodTotal)(nil), "ui.SCFourGodTotal")
	proto.RegisterType((*CSFourGodUseMasked)(nil), "ui.CSFourGodUseMasked")
	proto.RegisterType((*SCFourGodUseMasked)(nil), "ui.SCFourGodUseMasked")
}

func init() { proto.RegisterFile("fourgod.proto", fileDescriptor38) }

var fileDescriptor38 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x5f, 0x6b, 0xea, 0x40,
	0x10, 0xc5, 0x31, 0xb9, 0xf1, 0x7a, 0xc7, 0x3f, 0xdc, 0x2e, 0x6d, 0x49, 0x2d, 0xb4, 0x69, 0x68,
	0x69, 0x9e, 0x7c, 0x90, 0xbe, 0x96, 0x42, 0x04, 0x25, 0xf4, 0x2f, 0x68, 0x3f, 0x40, 0x4c, 0x56,
	0x5d, 0x34, 0x99, 0x25, 0xbb, 0x0b, 0xda, 0x4f, 0x5f, 0x76, 0xb5, 0xa9, 0x44, 0x7c, 0x3c, 0xf0,
	0x9b, 0x39, 0x73, 0xe6, 0x40, 0x7b, 0x86, 0xaa, 0x98, 0x63, 0xda, 0xe3, 0x05, 0x4a, 0x24, 0x96,
	0x62, 0x5d, 0x60, 0x92, 0x66, 0x5b, 0xdd, 0x6d, 0x25, 0x98, 0x65, 0x98, 0x6f, 0x95, 0xff, 0x05,
	0x30, 0x44, 0x55, 0x8c, 0x30, 0x0d, 0x19, 0x92, 0x36, 0x38, 0x79, 0xc2, 0xa3, 0xd4, 0xad, 0x79,
	0x56, 0x60, 0x93, 0x26, 0xd8, 0x72, 0xc3, 0x5d, 0xcb, 0xb3, 0x02, 0x87, 0x74, 0xa0, 0x2e, 0x64,
	0x2c, 0x95, 0x70, 0x6d, 0xa3, 0x09, 0xc0, 0x56, 0x4f, 0x58, 0x46, 0xdd, 0x3f, 0x66, 0xe0, 0x02,
	0x6c, 0x8e, 0xc2, 0x75, 0xbc, 0x5a, 0xd0, 0xec, 0xb7, 0x7a, 0x8a, 0xf5, 0x3e, 0x50, 0x30, 0xc9,
	0x30, 0x27, 0x27, 0xf0, 0x6f, 0xca, 0x70, 0x85, 0xf3, 0x4d, 0x94, 0xba, 0x75, 0xbd, 0xc1, 0x7f,
	0x82, 0xd6, 0x78, 0xb0, 0x73, 0x1f, 0x51, 0xa9, 0x1d, 0x96, 0x74, 0xf3, 0xa6, 0x32, 0x63, 0xef,
	0x90, 0x6b, 0xf8, 0x3b, 0x65, 0xf8, 0xc2, 0x84, 0x74, 0x2d, 0xcf, 0x0e, 0x9a, 0xfd, 0x8e, 0xde,
	0xf8, 0x7b, 0xae, 0x7f, 0x03, 0xff, 0x07, 0xe3, 0x9d, 0x7e, 0xe7, 0x34, 0x0f, 0x71, 0x6d, 0x22,
	0xf0, 0xe4, 0x27, 0x82, 0x46, 0x4a, 0x8f, 0x23, 0xc8, 0x1d, 0x9c, 0x56, 0x91, 0xb1, 0x44, 0x5e,
	0xc5, 0xee, 0xe1, 0xbc, 0x8a, 0x0d, 0x59, 0xce, 0xc4, 0xe2, 0x10, 0x3c, 0x2b, 0xc1, 0x67, 0x93,
	0x67, 0xb0, 0x88, 0xf3, 0x39, 0xad, 0xe6, 0xf3, 0x1f, 0xf6, 0xc0, 0x90, 0x61, 0x58, 0x60, 0x9c,
	0x26, 0xb1, 0x90, 0xe4, 0x12, 0xec, 0x29, 0x43, 0x43, 0x1d, 0x86, 0x7e, 0x84, 0x4e, 0x39, 0x35,
	0x41, 0x19, 0xaf, 0x74, 0x4d, 0x74, 0xcd, 0x77, 0x9d, 0x5d, 0x41, 0x43, 0x97, 0xbd, 0xf7, 0x35,
	0xd3, 0x43, 0x24, 0x69, 0x16, 0xe5, 0x33, 0xf4, 0x6f, 0x81, 0x94, 0x3f, 0xfb, 0x14, 0xf4, 0x35,
	0x16, 0x4b, 0x9a, 0xea, 0xd3, 0x0a, 0x2a, 0xd4, 0x4a, 0x9a, 0x2d, 0x0d, 0x4d, 0x95, 0x26, 0x47,
	0xa9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xce, 0xfe, 0xef, 0x6a, 0x02, 0x00, 0x00,
}