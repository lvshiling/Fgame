package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANTAO_GET_TYPE), (*uipb.SCXiantaoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIANTAO_PEACH_COMMIT_TYPE), (*uipb.CSXiantaoPeachCommit)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANTAO_PEACH_COMMIT_TYPE), (*uipb.SCXiantaoPeachCommit)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANTAO_PLAYER_ATTEND_CHANGE_TYPE), (*uipb.SCXiantaoPlayerAttendChange)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANTAO_PEACH_POINT_CHANGE_TYPE), (*uipb.SCXiantaoPeachPointChange)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANTAO_RESULT_TYPE), (*uipb.SCXiantaoResult)(nil))
}
