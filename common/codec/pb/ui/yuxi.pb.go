// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yuxi.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type YuXiCollect struct {
	NcpId            *int64    `protobuf:"varint,1,req,name=ncpId" json:"ncpId,omitempty"`
	Typ              *int32    `protobuf:"varint,2,req,name=typ" json:"typ,omitempty"`
	IsDead           *bool     `protobuf:"varint,3,req,name=isDead" json:"isDead,omitempty"`
	StatusTime       *int64    `protobuf:"varint,4,req,name=statusTime" json:"statusTime,omitempty"`
	Pos              *Position `protobuf:"bytes,5,opt,name=pos" json:"pos,omitempty"`
	BiologyId        *int32    `protobuf:"varint,6,req,name=biologyId" json:"biologyId,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *YuXiCollect) Reset()                    { *m = YuXiCollect{} }
func (m *YuXiCollect) String() string            { return proto.CompactTextString(m) }
func (*YuXiCollect) ProtoMessage()               {}
func (*YuXiCollect) Descriptor() ([]byte, []int) { return fileDescriptor141, []int{0} }

func (m *YuXiCollect) GetNcpId() int64 {
	if m != nil && m.NcpId != nil {
		return *m.NcpId
	}
	return 0
}

func (m *YuXiCollect) GetTyp() int32 {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return 0
}

func (m *YuXiCollect) GetIsDead() bool {
	if m != nil && m.IsDead != nil {
		return *m.IsDead
	}
	return false
}

func (m *YuXiCollect) GetStatusTime() int64 {
	if m != nil && m.StatusTime != nil {
		return *m.StatusTime
	}
	return 0
}

func (m *YuXiCollect) GetPos() *Position {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *YuXiCollect) GetBiologyId() int32 {
	if m != nil && m.BiologyId != nil {
		return *m.BiologyId
	}
	return 0
}

type YuXiOwner struct {
	KeepStartTime    *int64    `protobuf:"varint,1,req,name=keepStartTime" json:"keepStartTime,omitempty"`
	Pos              *Position `protobuf:"bytes,2,req,name=pos" json:"pos,omitempty"`
	PlayerName       *string   `protobuf:"bytes,3,req,name=playerName" json:"playerName,omitempty"`
	PlayerId         *int64    `protobuf:"varint,4,req,name=playerId" json:"playerId,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *YuXiOwner) Reset()                    { *m = YuXiOwner{} }
func (m *YuXiOwner) String() string            { return proto.CompactTextString(m) }
func (*YuXiOwner) ProtoMessage()               {}
func (*YuXiOwner) Descriptor() ([]byte, []int) { return fileDescriptor141, []int{1} }

func (m *YuXiOwner) GetKeepStartTime() int64 {
	if m != nil && m.KeepStartTime != nil {
		return *m.KeepStartTime
	}
	return 0
}

func (m *YuXiOwner) GetPos() *Position {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *YuXiOwner) GetPlayerName() string {
	if m != nil && m.PlayerName != nil {
		return *m.PlayerName
	}
	return ""
}

func (m *YuXiOwner) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

type SCYuXiCollectInfoBroadcast struct {
	CollectInfo      *YuXiCollect `protobuf:"bytes,1,req,name=collectInfo" json:"collectInfo,omitempty"`
	OwnerInfo        *YuXiOwner   `protobuf:"bytes,2,opt,name=ownerInfo" json:"ownerInfo,omitempty"`
	RebornType       *int32       `protobuf:"varint,3,req,name=rebornType" json:"rebornType,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SCYuXiCollectInfoBroadcast) Reset()                    { *m = SCYuXiCollectInfoBroadcast{} }
func (m *SCYuXiCollectInfoBroadcast) String() string            { return proto.CompactTextString(m) }
func (*SCYuXiCollectInfoBroadcast) ProtoMessage()               {}
func (*SCYuXiCollectInfoBroadcast) Descriptor() ([]byte, []int) { return fileDescriptor141, []int{2} }

func (m *SCYuXiCollectInfoBroadcast) GetCollectInfo() *YuXiCollect {
	if m != nil {
		return m.CollectInfo
	}
	return nil
}

func (m *SCYuXiCollectInfoBroadcast) GetOwnerInfo() *YuXiOwner {
	if m != nil {
		return m.OwnerInfo
	}
	return nil
}

func (m *SCYuXiCollectInfoBroadcast) GetRebornType() int32 {
	if m != nil && m.RebornType != nil {
		return *m.RebornType
	}
	return 0
}

type SCYuXiPosBroadcast struct {
	OwnerInfo        *YuXiOwner `protobuf:"bytes,1,req,name=ownerInfo" json:"ownerInfo,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *SCYuXiPosBroadcast) Reset()                    { *m = SCYuXiPosBroadcast{} }
func (m *SCYuXiPosBroadcast) String() string            { return proto.CompactTextString(m) }
func (*SCYuXiPosBroadcast) ProtoMessage()               {}
func (*SCYuXiPosBroadcast) Descriptor() ([]byte, []int) { return fileDescriptor141, []int{3} }

func (m *SCYuXiPosBroadcast) GetOwnerInfo() *YuXiOwner {
	if m != nil {
		return m.OwnerInfo
	}
	return nil
}

type CSYuXiGetInfo struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSYuXiGetInfo) Reset()                    { *m = CSYuXiGetInfo{} }
func (m *CSYuXiGetInfo) String() string            { return proto.CompactTextString(m) }
func (*CSYuXiGetInfo) ProtoMessage()               {}
func (*CSYuXiGetInfo) Descriptor() ([]byte, []int) { return fileDescriptor141, []int{4} }

type SCYuXiGetInfo struct {
	IsReceive        *int32  `protobuf:"varint,1,req,name=isReceive" json:"isReceive,omitempty"`
	WinAllianceId    *int64  `protobuf:"varint,2,req,name=winAllianceId" json:"winAllianceId,omitempty"`
	WinAllianceName  *string `protobuf:"bytes,3,req,name=winAllianceName" json:"winAllianceName,omitempty"`
	WinMengName      *string `protobuf:"bytes,4,req,name=winMengName" json:"winMengName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SCYuXiGetInfo) Reset()                    { *m = SCYuXiGetInfo{} }
func (m *SCYuXiGetInfo) String() string            { return proto.CompactTextString(m) }
func (*SCYuXiGetInfo) ProtoMessage()               {}
func (*SCYuXiGetInfo) Descriptor() ([]byte, []int) { return fileDescriptor141, []int{5} }

func (m *SCYuXiGetInfo) GetIsReceive() int32 {
	if m != nil && m.IsReceive != nil {
		return *m.IsReceive
	}
	return 0
}

func (m *SCYuXiGetInfo) GetWinAllianceId() int64 {
	if m != nil && m.WinAllianceId != nil {
		return *m.WinAllianceId
	}
	return 0
}

func (m *SCYuXiGetInfo) GetWinAllianceName() string {
	if m != nil && m.WinAllianceName != nil {
		return *m.WinAllianceName
	}
	return ""
}

func (m *SCYuXiGetInfo) GetWinMengName() string {
	if m != nil && m.WinMengName != nil {
		return *m.WinMengName
	}
	return ""
}

type CSYuXiReceiveDayRew struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSYuXiReceiveDayRew) Reset()                    { *m = CSYuXiReceiveDayRew{} }
func (m *CSYuXiReceiveDayRew) String() string            { return proto.CompactTextString(m) }
func (*CSYuXiReceiveDayRew) ProtoMessage()               {}
func (*CSYuXiReceiveDayRew) Descriptor() ([]byte, []int) { return fileDescriptor141, []int{6} }

type SCYuXiReceiveDayRew struct {
	DropInfo         []*DropInfo `protobuf:"bytes,1,rep,name=dropInfo" json:"dropInfo,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SCYuXiReceiveDayRew) Reset()                    { *m = SCYuXiReceiveDayRew{} }
func (m *SCYuXiReceiveDayRew) String() string            { return proto.CompactTextString(m) }
func (*SCYuXiReceiveDayRew) ProtoMessage()               {}
func (*SCYuXiReceiveDayRew) Descriptor() ([]byte, []int) { return fileDescriptor141, []int{7} }

func (m *SCYuXiReceiveDayRew) GetDropInfo() []*DropInfo {
	if m != nil {
		return m.DropInfo
	}
	return nil
}

type SCYuXiWinnerBroadcast struct {
	WinAllianceId    *int64 `protobuf:"varint,1,req,name=winAllianceId" json:"winAllianceId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCYuXiWinnerBroadcast) Reset()                    { *m = SCYuXiWinnerBroadcast{} }
func (m *SCYuXiWinnerBroadcast) String() string            { return proto.CompactTextString(m) }
func (*SCYuXiWinnerBroadcast) ProtoMessage()               {}
func (*SCYuXiWinnerBroadcast) Descriptor() ([]byte, []int) { return fileDescriptor141, []int{8} }

func (m *SCYuXiWinnerBroadcast) GetWinAllianceId() int64 {
	if m != nil && m.WinAllianceId != nil {
		return *m.WinAllianceId
	}
	return 0
}

func init() {
	proto.RegisterType((*YuXiCollect)(nil), "ui.YuXiCollect")
	proto.RegisterType((*YuXiOwner)(nil), "ui.YuXiOwner")
	proto.RegisterType((*SCYuXiCollectInfoBroadcast)(nil), "ui.SCYuXiCollectInfoBroadcast")
	proto.RegisterType((*SCYuXiPosBroadcast)(nil), "ui.SCYuXiPosBroadcast")
	proto.RegisterType((*CSYuXiGetInfo)(nil), "ui.CSYuXiGetInfo")
	proto.RegisterType((*SCYuXiGetInfo)(nil), "ui.SCYuXiGetInfo")
	proto.RegisterType((*CSYuXiReceiveDayRew)(nil), "ui.CSYuXiReceiveDayRew")
	proto.RegisterType((*SCYuXiReceiveDayRew)(nil), "ui.SCYuXiReceiveDayRew")
	proto.RegisterType((*SCYuXiWinnerBroadcast)(nil), "ui.SCYuXiWinnerBroadcast")
}

func init() { proto.RegisterFile("yuxi.proto", fileDescriptor141) }

var fileDescriptor141 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0xd5, 0x84, 0x4c, 0xcd, 0xcb, 0x42, 0x99, 0xab, 0x8a, 0xd0, 0x03, 0x8a, 0x2c, 0x0e,
	0x3d, 0xf5, 0x30, 0x09, 0xee, 0xd0, 0x4a, 0x28, 0x07, 0x60, 0x5a, 0x27, 0x01, 0x47, 0x2f, 0x79,
	0x54, 0x16, 0x89, 0x6d, 0x39, 0x0e, 0x59, 0xc4, 0x3f, 0x8f, 0x6c, 0xa7, 0x6a, 0xca, 0x8e, 0xef,
	0x93, 0xfd, 0x7d, 0xdf, 0xfb, 0xd9, 0x00, 0x43, 0xf7, 0xc4, 0xb7, 0x4a, 0x4b, 0x23, 0x49, 0xd0,
	0xf1, 0x35, 0x54, 0x5a, 0x2a, 0x3f, 0xaf, 0x6f, 0x34, 0xf6, 0x4a, 0x4b, 0x85, 0xda, 0x0c, 0xa3,
	0x74, 0x5d, 0xca, 0xa6, 0x91, 0xc2, 0x4f, 0xf4, 0x2f, 0x24, 0x3f, 0xbb, 0x1f, 0x7c, 0x27, 0xeb,
	0x1a, 0x4b, 0x43, 0x52, 0x88, 0x44, 0xa9, 0x8a, 0x2a, 0x9b, 0xe5, 0xc1, 0x26, 0x24, 0x09, 0x84,
	0x66, 0x50, 0x59, 0x90, 0x07, 0x9b, 0x88, 0xbc, 0x84, 0x2b, 0xde, 0xee, 0x91, 0x55, 0x59, 0x98,
	0x07, 0x9b, 0x39, 0x21, 0x00, 0xad, 0x61, 0xa6, 0x6b, 0x1f, 0x78, 0x83, 0xd9, 0x0b, 0x77, 0xe1,
	0x0d, 0x84, 0x4a, 0xb6, 0x59, 0x94, 0xcf, 0x36, 0xc9, 0xed, 0xf5, 0xb6, 0xe3, 0xdb, 0x3b, 0xd9,
	0x72, 0xc3, 0xa5, 0x20, 0x37, 0x10, 0x3f, 0x72, 0x59, 0xcb, 0xe3, 0x50, 0x54, 0xd9, 0x95, 0x75,
	0xa4, 0x25, 0xc4, 0x36, 0xfc, 0x5b, 0x2f, 0x50, 0x93, 0x15, 0xa4, 0xbf, 0x11, 0xd5, 0xc1, 0x30,
	0x6d, 0x9c, 0xe3, 0x6c, 0xea, 0x68, 0x2b, 0xfc, 0xef, 0x48, 0x00, 0x54, 0xcd, 0x06, 0xd4, 0x5f,
	0x59, 0x83, 0xae, 0x54, 0x4c, 0x5e, 0xc1, 0xdc, 0x6b, 0x45, 0xe5, 0x2b, 0xd1, 0x27, 0x58, 0x1f,
	0x76, 0x93, 0x1d, 0x0b, 0xf1, 0x4b, 0x7e, 0xd2, 0x92, 0x55, 0x25, 0x6b, 0x0d, 0x79, 0x07, 0x49,
	0x79, 0xd6, 0x5d, 0x66, 0x72, 0xbb, 0xb0, 0x31, 0x53, 0x2c, 0x39, 0xc4, 0xd2, 0x96, 0x74, 0x67,
	0x02, 0xb7, 0x5c, 0x7a, 0x3a, 0xe3, 0xdb, 0x13, 0x00, 0x8d, 0x8f, 0x52, 0x8b, 0x87, 0x41, 0xf9,
	0x2e, 0x11, 0xfd, 0x00, 0xc4, 0x27, 0xdf, 0xc9, 0xf6, 0x9c, 0x78, 0xe1, 0xe5, 0xf3, 0x2e, 0xbd,
	0xe8, 0x02, 0xd2, 0xdd, 0xc1, 0x8e, 0x9f, 0xd1, 0xb5, 0xa2, 0x47, 0x48, 0xbd, 0xd1, 0x28, 0x58,
	0x96, 0xbc, 0xbd, 0xc7, 0x12, 0xf9, 0x1f, 0xcf, 0x29, 0xb2, 0xf8, 0x7a, 0x2e, 0x3e, 0xd6, 0x35,
	0x67, 0xa2, 0xc4, 0xa2, 0x72, 0xc4, 0x42, 0xf2, 0x1a, 0x16, 0x13, 0x79, 0x02, 0x6a, 0x09, 0x49,
	0xcf, 0xc5, 0x17, 0x14, 0x47, 0x27, 0x5a, 0x56, 0x31, 0x5d, 0xc1, 0xd2, 0x27, 0x8f, 0xde, 0x7b,
	0x36, 0xdc, 0x63, 0x4f, 0xdf, 0xc3, 0xd2, 0xe7, 0x5f, 0xc8, 0xe4, 0x2d, 0xcc, 0xed, 0x57, 0x1b,
	0x17, 0x09, 0x4f, 0xef, 0xb3, 0x1f, 0x35, 0xba, 0x85, 0x95, 0xbf, 0xf6, 0x9d, 0x0b, 0x81, 0xfa,
	0x8c, 0xe0, 0x59, 0x57, 0xf7, 0xd4, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x1d, 0x57, 0x59,
	0xc9, 0x02, 0x00, 0x00,
}