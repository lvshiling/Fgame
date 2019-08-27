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
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LINEUP_ATTEND_TYPE), (*crosspb.ISLineupAttend)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LINEUP_ATTEND_TYPE), (*crosspb.SILineupAttend)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LINEUP_CANCEL_TYPE), (*crosspb.ISLineupCancle)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LINEUP_CANCEL_TYPE), (*crosspb.SILineupCancle)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LINEUP_SUCCESS_TYPE), (*crosspb.ISLineupSuccess)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_LINEUP_SUCCESS_TYPE), (*crosspb.SILineupSuccess)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_LINEUP_SCENE_FINISH_TO_CANCEL_TYPE), (*crosspb.ISLineupSceneFinishToCancel)(nil))
}
