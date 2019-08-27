package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ACTIVITY_NOTICE_TYPE), (*uipb.SCActivityNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ACTIVITY_DATA_NOTICE_TYPE), (*uipb.SCActivityDataNotice)(nil))
}
