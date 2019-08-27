package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DRAGON_GET_TYPE), (*uipb.CSDragonGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DRAGON_GET_TYPE), (*uipb.SCDragonGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DRAGON_FEED_TYPE), (*uipb.CSDragonFeed)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DRAGON_FEED_TYPE), (*uipb.SCDragonFeed)(nil))

}
