package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GEM_MINE_GET_TYPE), (*uipb.CSGemMineGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GEM_MINE_GET_TYPE), (*uipb.SCGemMineGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GEM_MINE_ACTIVE_TYPE), (*uipb.CSGemMineActive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GEM_MINE_ACTIVE_TYPE), (*uipb.SCGemMineActive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GEM_MINE_RECEIVE_TYPE), (*uipb.CSGemMineReceive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GEM_MINE_RECEIVE_TYPE), (*uipb.SCGemMineReceive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GEM_GAMBLE_TYPE), (*uipb.CSGemGamble)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GEM_GAMBLE_TYPE), (*uipb.SCGemGamble)(nil))
}
