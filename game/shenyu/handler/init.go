package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENYU_SCENE_INFO_NOTICE_TYPE), (*uipb.SCShenYuSceneInfoNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENYU_KEY_NUM_CHANGED_NOTICE_TYPE), (*uipb.SCShenYuKeyNumChangeNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENYU_FINISH_REW_TYPE), (*uipb.SCShenYuFinishRew)(nil))

}
