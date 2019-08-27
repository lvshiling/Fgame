// Code generated by protoc-gen-go. DO NOT EDIT.
// source: buff.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BuffData struct {
	BuffId           *int32   `protobuf:"varint,3,req,name=buffId" json:"buffId,omitempty"`
	BuffTime         *float32 `protobuf:"fixed32,4,req,name=buffTime" json:"buffTime,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BuffData) Reset()                    { *m = BuffData{} }
func (m *BuffData) String() string            { return proto.CompactTextString(m) }
func (*BuffData) ProtoMessage()               {}
func (*BuffData) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *BuffData) GetBuffId() int32 {
	if m != nil && m.BuffId != nil {
		return *m.BuffId
	}
	return 0
}

func (m *BuffData) GetBuffTime() float32 {
	if m != nil && m.BuffTime != nil {
		return *m.BuffTime
	}
	return 0
}

type CSBuffList struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSBuffList) Reset()                    { *m = CSBuffList{} }
func (m *CSBuffList) String() string            { return proto.CompactTextString(m) }
func (*CSBuffList) ProtoMessage()               {}
func (*CSBuffList) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

type SCBuffList struct {
	BuffList         []*BuffData `protobuf:"bytes,1,rep,name=buffList" json:"buffList,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SCBuffList) Reset()                    { *m = SCBuffList{} }
func (m *SCBuffList) String() string            { return proto.CompactTextString(m) }
func (*SCBuffList) ProtoMessage()               {}
func (*SCBuffList) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

func (m *SCBuffList) GetBuffList() []*BuffData {
	if m != nil {
		return m.BuffList
	}
	return nil
}

type CSBuffSearch struct {
	BuffId           *int32 `protobuf:"varint,1,req,name=buffId" json:"buffId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSBuffSearch) Reset()                    { *m = CSBuffSearch{} }
func (m *CSBuffSearch) String() string            { return proto.CompactTextString(m) }
func (*CSBuffSearch) ProtoMessage()               {}
func (*CSBuffSearch) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{3} }

func (m *CSBuffSearch) GetBuffId() int32 {
	if m != nil && m.BuffId != nil {
		return *m.BuffId
	}
	return 0
}

type SCBuffSearch struct {
	BuffId           *int32 `protobuf:"varint,1,req,name=buffId" json:"buffId,omitempty"`
	Result           *bool  `protobuf:"varint,2,req,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCBuffSearch) Reset()                    { *m = SCBuffSearch{} }
func (m *SCBuffSearch) String() string            { return proto.CompactTextString(m) }
func (*SCBuffSearch) ProtoMessage()               {}
func (*SCBuffSearch) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{4} }

func (m *SCBuffSearch) GetBuffId() int32 {
	if m != nil && m.BuffId != nil {
		return *m.BuffId
	}
	return 0
}

func (m *SCBuffSearch) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*BuffData)(nil), "ui.BuffData")
	proto.RegisterType((*CSBuffList)(nil), "ui.CSBuffList")
	proto.RegisterType((*SCBuffList)(nil), "ui.SCBuffList")
	proto.RegisterType((*CSBuffSearch)(nil), "ui.CSBuffSearch")
	proto.RegisterType((*SCBuffSearch)(nil), "ui.SCBuffSearch")
}

func init() { proto.RegisterFile("buff.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2a, 0x4d, 0x4b,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0xcd, 0x54, 0xd2, 0xe1, 0xe2, 0x70, 0x2a,
	0x4d, 0x4b, 0x73, 0x49, 0x2c, 0x49, 0x14, 0xe2, 0xe3, 0x62, 0x03, 0xc9, 0x7a, 0xa6, 0x48, 0x30,
	0x2b, 0x30, 0x69, 0xb0, 0x0a, 0x09, 0x70, 0x71, 0x80, 0xf8, 0x21, 0x99, 0xb9, 0xa9, 0x12, 0x2c,
	0x0a, 0x4c, 0x1a, 0x4c, 0x4a, 0x3c, 0x5c, 0x5c, 0xce, 0xc1, 0x20, 0xf5, 0x3e, 0x99, 0xc5, 0x25,
	0x4a, 0x3a, 0x5c, 0x5c, 0xc1, 0xce, 0x30, 0x9e, 0x90, 0x1c, 0x44, 0x35, 0x88, 0x2d, 0xc1, 0xa8,
	0xc0, 0xac, 0xc1, 0x6d, 0xc4, 0xa3, 0x57, 0x9a, 0xa9, 0x07, 0x33, 0x5d, 0x49, 0x8e, 0x8b, 0x07,
	0xa2, 0x37, 0x38, 0x35, 0xb1, 0x28, 0x39, 0x03, 0xc9, 0x36, 0x46, 0x90, 0x6d, 0x4a, 0x7a, 0x5c,
	0x3c, 0x10, 0xd3, 0xb0, 0xcb, 0x83, 0xf8, 0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x12, 0x4c, 0x0a,
	0x4c, 0x1a, 0x1c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x97, 0x8a, 0x6c, 0xca, 0x00, 0x00,
	0x00,
}