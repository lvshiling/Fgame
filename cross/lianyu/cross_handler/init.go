package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	initCodec()
}

func initCodec() {
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LIANYU_ATTEND_TYPE), (*crosspb.SILianYuAttend)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LIANYU_ATTEND_TYPE), (*crosspb.ISLianYuAttend)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LIANYU_LINEUP_SUCCESS_TYPE), (*crosspb.ISLianYuLineUpSuccess)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LIANYU_LINEUP_SUCCESS_TYPE), (*crosspb.SILianYuLineUpSuccess)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LIANYU_CANCLE_LINEUP_TYPE), (*crosspb.SILianYuCancleLineUp)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LIANYU_CANCLE_LINEUP_TYPE), (*crosspb.ISLianYuCancleLineUp)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LIANYU_FINISH_LINEUP_CANCLE_TYPE), (*crosspb.ISLianYuFinishLineUpCancle)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LIANYU_FINISH_LINEUP_CANCLE_TYPE), (*crosspb.SILianYuFinishLineUpCancle)(nil))

}
