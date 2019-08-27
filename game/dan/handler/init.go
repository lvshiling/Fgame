package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DAN_GET_TYPE), (*uipb.CSDanGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DAN_GET_TYPE), (*uipb.SCDanGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DAN_USE_TYPE), (*uipb.CSDanUse)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DAN_USE_TYPE), (*uipb.SCDanUse)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DAN_UPGRADE_TYPE), (*uipb.CSDanUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DAN_UPGRADE_TYPE), (*uipb.SCDanUpgrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DAN_ALCHEMYGET_TYPE), (*uipb.CSAlchemyGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DAN_ALCHEMYGET_TYPE), (*uipb.SCAlchemyGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DAN_ALCHEMYSTART_TYPE), (*uipb.CSAlchemyStart)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DAN_ALCHEMYSTART_TYPE), (*uipb.SCAlchemyStart)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DAN_ALCHEMYACCELERATE_TYPE), (*uipb.CSAlchemyAccelerate)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DAN_ALCHEMYACCELERATE_TYPE), (*uipb.SCAlchemyAccelerate)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DAN_ALCHEMYRECEIVE_TYPE), (*uipb.CSAlchemyReceive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DAN_ALCHEMYRECEIVE_TYPE), (*uipb.SCAlchemyReceive)(nil))
}
