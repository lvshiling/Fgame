package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUIZ_ANSWER_TYPE), (*uipb.CSQuizAnswer)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUIZ_ANSWER_TYPE), (*uipb.SCQuizAnswer)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUIZ_ASSIGN_INFO_TYPE), (*uipb.SCQuizAssignInfo)(nil))
}
