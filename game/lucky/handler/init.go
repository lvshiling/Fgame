package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LUCKY_INFO_CHANGED_TYPE), (*uipb.SCLuckyInfoChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LUCKY_INFO_NOTICE_TYPE), (*uipb.SCLuckyInfoNotice)(nil))
}
