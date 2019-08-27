package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_RELIVE_TYPE), (*crosspb.SIPlayerRelive)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_RELIVE_TYPE), (*crosspb.ISPlayerRelive)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_RELIVE_SYNC_TYPE), (*crosspb.SIPlayerReliveSync)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_RELIVE_SYNC_TYPE), (*crosspb.ISPlayerReliveSync)(nil))

}
