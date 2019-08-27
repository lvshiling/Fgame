package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	initCodec()
}

func initCodec() {
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_KILL_BIOLOGY_TYPE), (*crosspb.ISPlayerKillBiology)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLYAER_KILL_BIOLOGY_TYPE), (*crosspb.SIPlayerKillBiology)(nil))
}
