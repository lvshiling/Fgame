package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ACTIVITY_ATTEND_TYPE), (*uipb.CSActivityAttend)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ACTIVITY_ATTEND_TYPE), (*uipb.SCActivityAttend)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ACTIVITY_COLLECT_INFO_NOTICE_TYPE), (*uipb.SCActivityCollectInfoNotice)(nil))
}
