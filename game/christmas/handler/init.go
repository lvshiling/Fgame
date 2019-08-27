package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHRISTMAS_COLLECT_REFRESH_BROADCAST_TYPE), (*uipb.SCChristmasCollectRefreshBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHRISTMAS_COLLECT_COLLECT_NUM_NOTICE), (*uipb.SCChristmasCollectNumNotice)(nil))
}
