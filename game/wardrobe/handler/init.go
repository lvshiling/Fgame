package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WARDROBE_GET_TYPE), (*uipb.CSWardrobeGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WARDROBE_GET_TYPE), (*uipb.SCWardrobeGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WARDROBE_ACTIVE_TYPE), (*uipb.SCWardrobeActive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WARDROBE_ROMOVE_TYPE), (*uipb.SCWardrobeRemove)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_WARDROBE_PEIYANG_TYPE), (*uipb.CSWardrobePeiYang)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_WARDROBE_PEIYANG_TYPE), (*uipb.SCWardrobePeiYang)(nil))
}
