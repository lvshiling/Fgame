package handler

import (
	clientcodec "fgame/fgame/client/codec"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
)

func init() {
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_PROPERTY_TYPE), (*uipb.SCPlayerProperty)(nil))
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_INFO_TYPE), (*uipb.SCPlayerInfo)(nil))
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GUA_JI_TYPE), (*uipb.CSGuaJi)(nil))
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GUA_JI_TYPE), (*uipb.SCGuaJi)(nil))
}
