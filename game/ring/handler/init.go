package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RING_INFO_GET_TYPE), (*uipb.CSRingInfoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RING_INFO_GET_TYPE), (*uipb.SCRingInfoGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RING_ADVANCE_TYPE), (*uipb.CSRingAdvance)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RING_ADVANCE_TYPE), (*uipb.SCRingAdvance)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RING_STRENGTHEN_TYPE), (*uipb.CSRingStrengthen)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RING_STRENGTHEN_TYPE), (*uipb.SCRingStrengthen)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RING_JINGLING_TYPE), (*uipb.CSRingJingLing)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RING_JINGLING_TYPE), (*uipb.SCRingJingLing)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RING_FUSE_TYPE), (*uipb.CSRingFuse)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RING_FUSE_TYPE), (*uipb.SCRingFuse)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RING_EQUIP_TYPE), (*uipb.CSRingEquip)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RING_EQUIP_TYPE), (*uipb.SCRingEquip)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RING_UNLOAD_TYPE), (*uipb.CSRingUnload)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RING_UNLOAD_TYPE), (*uipb.SCRingUnload)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RING_SLOT_CHANGE_TYPE), (*uipb.SCRingSlotChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RING_BAOKU_INFO_TYPE), (*uipb.CSRingBaoKuInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RING_BAOKU_INFO_TYPE), (*uipb.SCRingBaoKuInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RING_BAOKU_ATTEND_TYPE), (*uipb.CSRingBaoKuAttend)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RING_BAOKU_ATTEND_TYPE), (*uipb.SCRingBaoKuAttend)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RING_LUCKY_POINTS_CHANGE_TYPE), (*uipb.SCRingLuckyPointsChange)(nil))
}
