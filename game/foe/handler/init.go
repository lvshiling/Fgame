package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOE_GET_TYPE), (*uipb.SCFoesGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FOE_REMOVE_TYPE), (*uipb.CSFoeRemove)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOE_REMOVE_TYPE), (*uipb.SCFoeRemove)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOE_ADD_TYPE), (*uipb.SCFoeAdd)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FOE_VIEW_POS_TYPE), (*uipb.CSFoeViewPos)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOE_VIEW_POS_TYPE), (*uipb.SCFoeViewPos)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FOE_TRANSFER_TYPE), (*uipb.CSFoeTransfer)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOE_TRANSFER_TYPE), (*uipb.SCFoeTransfer)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOE_NOTICE_TYPE), (*uipb.SCFoeNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOE_FEEDBACK_NOTICE_TYPE), (*uipb.SCFoeFeedbackNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FOE_FEEDBACK_TYPE), (*uipb.CSFoeFeedback)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOE_FEEDBACK_TYPE), (*uipb.SCFoeFeedback)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FOE_FEEDBACK_BUY_PROTECT_TYPE), (*uipb.CSFoeFeedbackBuyProtect)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOE_FEEDBACK_BUY_PROTECT_TYPE), (*uipb.SCFoeFeedbackBuyProtect)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOE_FEEDBACK_INFO_TYPE), (*uipb.SCFoeFeedbackInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FOE_FEEDBACK_READ_TYPE), (*uipb.CSFoeFeedbackRead)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOE_FEEDBACK_READ_TYPE), (*uipb.SCFoeFeedbackRead)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FOE_KILL_NOTICE_TYPE), (*uipb.SCFoeKillNotice)(nil))
}
