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
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_GODSIEGE_ATTEND_TYPE), (*crosspb.SIGodSiegeAttend)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_GODSIEGE_ATTEND_TYPE), (*crosspb.ISGodSiegeAttend)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_GODSIEGE_LINEUP_SUCCESS_TYPE), (*crosspb.ISGodSiegeLineUpSuccess)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_GODSIEGE_LINEUP_SUCCESS_TYPE), (*crosspb.SIGodSiegeLineUpSuccess)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_GODSIEGE_CANCLE_LINEUP_TYPE), (*crosspb.SIGodSiegeCancleLineUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_GODSIEGE_CANCLE_LINEUP_TYPE), (*crosspb.ISGodSiegeCancleLineUp)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_GODSIEGE_FINISH_LINEUP_CANCLE_TYPE), (*crosspb.ISGodSiegeFinishLineUpCancle)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_GODSIEGE_FINISH_LINEUP_CANCLE_TYPE), (*crosspb.SIGodSiegeFinishLineUpCancle)(nil))
}
