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
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_SHENMO_ATTEND_TYPE), (*crosspb.SIShenMoAttend)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_SHENMO_ATTEND_TYPE), (*crosspb.ISShenMoAttend)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_SHENMO_LINEUP_SUCCESS_TYPE), (*crosspb.ISShenMoLineUpSuccess)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_SHENMO_LINEUP_SUCCESS_TYPE), (*crosspb.SIShenMoLineUpSuccess)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_SHENMO_CANCLE_LINEUP_TYPE), (*crosspb.SIShenMoCancleLineUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_SHENMO_CANCLE_LINEUP_TYPE), (*crosspb.ISShenMoCancleLineUp)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_SHENMO_FINISH_LINEUP_CANCLE_TYPE), (*crosspb.ISShenMoFinishLineUpCancle)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_SHENMO_FINISH_LINEUP_CANCLE_TYPE), (*crosspb.SIShenMoFinishLineUpCancle)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_GONGXUN_ADD_TYPE), (*crosspb.ISPlayerGongXunAdd)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAEYR_GONGXUN_ADD_TYPE), (*crosspb.SIPlayerGongXunAdd)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_GONGXUN_SUB_TYPE), (*crosspb.ISPlayerGongXunSub)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAEYR_GONGXUN_SUB_TYPE), (*crosspb.SIPlayerGongXunSub)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_GONGXUN_CHANGED_TYPE), (*crosspb.SIPlayerGongXunChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_GONGXUN_CHANGED_TYPE), (*crosspb.ISPlayerGongXunChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_SHENMO_KILLNUM_CHANGED_TYPE), (*crosspb.ISShenMoKillNumChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_SHENMO_KILLNUM_CHANGED_TYPE), (*crosspb.SIShenMoKillNumChanged)(nil))

}
