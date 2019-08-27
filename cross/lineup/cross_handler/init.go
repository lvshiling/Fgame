package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	initCodec()
}

func initCodec() {

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LINEUP_ATTEND_TYPE), (*crosspb.ISLineupAttend)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LINEUP_ATTEND_TYPE), (*crosspb.SILineupAttend)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LINEUP_CANCEL_TYPE), (*crosspb.ISLineupCancle)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LINEUP_CANCEL_TYPE), (*crosspb.SILineupCancle)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LINEUP_SUCCESS_TYPE), (*crosspb.ISLineupSuccess)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LINEUP_SUCCESS_TYPE), (*crosspb.SILineupSuccess)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LINEUP_SCENE_FINISH_TO_CANCEL_TYPE), (*crosspb.ISLineupSceneFinishToCancel)(nil))
}
