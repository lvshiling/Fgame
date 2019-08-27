package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOUL_GET_TYPE), (*uipb.CSSoulGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOUL_GET_TYPE), (*uipb.SCSoulGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOUL_ACTIVE_TYPE), (*uipb.CSSoulActive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOUL_ACTIVE_TYPE), (*uipb.SCSoulActive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOUL_EMBED_TYPE), (*uipb.CSSoulEmbed)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOUL_EMBED_TYPE), (*uipb.SCSoulEmbed)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOUL_FEED_TYPE), (*uipb.CSSoulFeed)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOUL_FEED_TYPE), (*uipb.SCSoulFeed)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOUL_AWAKEN_TYPE), (*uipb.CSSoulAwaken)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOUL_AWAKEN_TYPE), (*uipb.SCSoulAwaken)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOUL_STRENGTHEN_TYPE), (*uipb.CSSoulStrengthen)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOUL_STRENGTHEN_TYPE), (*uipb.SCSoulStrengthen)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOUL_UPGRADE_TYPE), (*uipb.CSSoulUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOUL_UPGRADE_TYPE), (*uipb.SCSoulUpgrade)(nil))

}
