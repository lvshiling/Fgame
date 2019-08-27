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
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LIANYU_ATTEND_TYPE), (*crosspb.SILianYuAttend)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LIANYU_ATTEND_TYPE), (*crosspb.ISLianYuAttend)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LIANYU_LINEUP_SUCCESS_TYPE), (*crosspb.ISLianYuLineUpSuccess)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LIANYU_LINEUP_SUCCESS_TYPE), (*crosspb.SILianYuLineUpSuccess)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LIANYU_CANCLE_LINEUP_TYPE), (*crosspb.SILianYuCancleLineUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LIANYU_CANCLE_LINEUP_TYPE), (*crosspb.ISLianYuCancleLineUp)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LIANYU_FINISH_LINEUP_CANCLE_TYPE), (*crosspb.ISLianYuFinishLineUpCancle)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LIANYU_FINISH_LINEUP_CANCLE_TYPE), (*crosspb.SILianYuFinishLineUpCancle)(nil))
}
