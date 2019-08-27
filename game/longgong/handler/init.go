package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LONGGONG_GET_TYPE), (*uipb.SCLonggongGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LONGGONG_PLAYER_VAL_CHANGE_TYPE), (*uipb.SCLonggongPlayerValChange)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LONGGONG_SCENE_VAL_BROADCAST_TYPE), (*uipb.SCLonggongSceneValBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LONGGONG_RESULT_TYPE), (*uipb.SCLonggongResult)(nil))
}
