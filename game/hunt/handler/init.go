package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_HUNT_XUNBAO_TYPE), (*uipb.CSHuntXunBao)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HUNT_XUNBAO_TYPE), (*uipb.SCHuntXunBao)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_HUNT_INFO_NOTICE_TYPE), (*uipb.SCHuntInfoNotice)(nil))
}
