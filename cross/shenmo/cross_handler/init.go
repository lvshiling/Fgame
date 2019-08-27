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
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_SHENMO_ATTEND_TYPE), (*crosspb.SIShenMoAttend)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_SHENMO_ATTEND_TYPE), (*crosspb.ISShenMoAttend)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_SHENMO_LINEUP_SUCCESS_TYPE), (*crosspb.ISShenMoLineUpSuccess)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_SHENMO_LINEUP_SUCCESS_TYPE), (*crosspb.SIShenMoLineUpSuccess)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_SHENMO_CANCLE_LINEUP_TYPE), (*crosspb.SIShenMoCancleLineUp)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_SHENMO_CANCLE_LINEUP_TYPE), (*crosspb.ISShenMoCancleLineUp)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_SHENMO_FINISH_LINEUP_CANCLE_TYPE), (*crosspb.ISShenMoFinishLineUpCancle)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_SHENMO_FINISH_LINEUP_CANCLE_TYPE), (*crosspb.SIShenMoFinishLineUpCancle)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_GONGXUN_ADD_TYPE), (*crosspb.ISPlayerGongXunAdd)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAEYR_GONGXUN_ADD_TYPE), (*crosspb.SIPlayerGongXunAdd)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_GONGXUN_SUB_TYPE), (*crosspb.ISPlayerGongXunSub)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAEYR_GONGXUN_SUB_TYPE), (*crosspb.SIPlayerGongXunSub)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_GONGXUN_CHANGED_TYPE), (*crosspb.SIPlayerGongXunChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_GONGXUN_CHANGED_TYPE), (*crosspb.ISPlayerGongXunChanged)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_SHENMO_KILLNUM_CHANGED_TYPE), (*crosspb.ISShenMoKillNumChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_SHENMO_KILLNUM_CHANGED_TYPE), (*crosspb.SIShenMoKillNumChanged)(nil))
}
