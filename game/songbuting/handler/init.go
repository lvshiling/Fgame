package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SONGBUTING_GET_TYPE), (*uipb.CSSongBuTingGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SONGBUTING_RECEIVE_TYPE), (*uipb.CSSongBuTingReceive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SONGBUTING_CHANGED_TYPE), (*uipb.SCSongBuTingChanged)(nil))
}
