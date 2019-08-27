package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HONGBAO_GET_TYPE), (*uipb.CSHongbaoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HONGBAO_GET_TYPE), (*uipb.SCHongbaoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HONGBAO_SEND_TYPE), (*uipb.CSHongbaoSend)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HONGBAO_SEND_TYPE), (*uipb.SCHongbaoSend)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HONGBAO_SNATCH_TYPE), (*uipb.CSHongbaoSnatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HONGBAO_SNATCH_TYPE), (*uipb.SCHongbaoSnatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HONGBAO_SNATCH_GET_TYPE), (*uipb.CSHongbaoSnatchGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HONGBAO_SNATCH_GET_TYPE), (*uipb.SCHongbaoSnatchGet)(nil))
}
