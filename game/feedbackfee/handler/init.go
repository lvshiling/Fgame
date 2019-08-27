package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FEEDBACK_FEE_INFO_TYPE), (*uipb.CSFeedbackFeeInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FEEDBACK_FEE_INFO_TYPE), (*uipb.SCFeedbackFeeInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FEEDBACK_FEE_EXCHANGE_TYPE), (*uipb.CSFeedbackFeeExchange)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FEEDBACK_FEE_EXCHANGE_TYPE), (*uipb.SCFeedbackFeeExchange)(nil))

}
