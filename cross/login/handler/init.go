package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LOGIN_TYPE), (*crosspb.ISLogin)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LOGIN_TYPE), (*crosspb.SILogin)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_DATA_TYPE), (*crosspb.SIPlayerData)(nil))
}
