// Code generated by protoc-gen-go. DO NOT EDIT.
// source: item.proto

package cross

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ItemInfo struct {
	ItemId           *int32 `protobuf:"varint,1,req,name=itemId" json:"itemId,omitempty"`
	Num              *int32 `protobuf:"varint,2,req,name=num" json:"num,omitempty"`
	Level            *int32 `protobuf:"varint,3,opt,name=level" json:"level,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ItemInfo) Reset()                    { *m = ItemInfo{} }
func (m *ItemInfo) String() string            { return proto.CompactTextString(m) }
func (*ItemInfo) ProtoMessage()               {}
func (*ItemInfo) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *ItemInfo) GetItemId() int32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *ItemInfo) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *ItemInfo) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func init() {
	proto.RegisterType((*ItemInfo)(nil), "cross.ItemInfo")
}

func init() { proto.RegisterFile("item.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x2c, 0x49, 0xcd,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2e, 0xca, 0x2f, 0x2e, 0x56, 0x32, 0xe3,
	0xe2, 0xf0, 0x2c, 0x49, 0xcd, 0xf5, 0xcc, 0x4b, 0xcb, 0x17, 0xe2, 0xe3, 0x62, 0x03, 0x29, 0xf0,
	0x4c, 0x91, 0x60, 0x54, 0x60, 0xd2, 0x60, 0x15, 0xe2, 0xe6, 0x62, 0xce, 0x2b, 0xcd, 0x95, 0x60,
	0x02, 0x73, 0x78, 0xb9, 0x58, 0x73, 0x52, 0xcb, 0x52, 0x73, 0x24, 0x98, 0x15, 0x18, 0x35, 0x58,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x7e, 0x00, 0xde, 0x4b, 0x00, 0x00, 0x00,
}