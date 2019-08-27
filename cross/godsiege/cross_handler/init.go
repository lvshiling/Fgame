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
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_GODSIEGE_ATTEND_TYPE), (*crosspb.SIGodSiegeAttend)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_GODSIEGE_ATTEND_TYPE), (*crosspb.ISGodSiegeAttend)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_GODSIEGE_LINEUP_SUCCESS_TYPE), (*crosspb.ISGodSiegeLineUpSuccess)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_GODSIEGE_LINEUP_SUCCESS_TYPE), (*crosspb.SIGodSiegeLineUpSuccess)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_GODSIEGE_CANCLE_LINEUP_TYPE), (*crosspb.SIGodSiegeCancleLineUp)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_GODSIEGE_CANCLE_LINEUP_TYPE), (*crosspb.ISGodSiegeCancleLineUp)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_GODSIEGE_FINISH_LINEUP_CANCLE_TYPE), (*crosspb.ISGodSiegeFinishLineUpCancle)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_GODSIEGE_FINISH_LINEUP_CANCLE_TYPE), (*crosspb.SIGodSiegeFinishLineUpCancle)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_DENSEWAT_SYNC_TYPE), (*crosspb.SIDenseWatSync)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_DENSEWAT_SYNC_TYPE), (*crosspb.ISDenseWatSync)(nil))
}
