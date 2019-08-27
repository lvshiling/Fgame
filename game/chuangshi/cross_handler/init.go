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
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_CHUANGSHI_ENTER_CITY_TYPE), (*crosspb.ISChuangShiEnterCity)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_CHUANGSHI_ENTER_CITY_TYPE), (*crosspb.SIChuangShiEnterCity)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_CHUANGSHI_SCENE_FINISH_TYPE), (*crosspb.SIChuangShiSceneFinish)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_CHUANGSHI_SCENE_FINISH_TYPE), (*crosspb.ISChuangShiSceneFinish)(nil))
}
