package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_RELIVE_TYPE), (*crosspb.ISPlayerRelive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_RELIVE_TYPE), (*crosspb.SIPlayerRelive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_RELIVE_SYNC_TYPE), (*crosspb.ISPlayerReliveSync)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_RELIVE_SYNC_TYPE), (*crosspb.SIPlayerReliveSync)(nil))
}
