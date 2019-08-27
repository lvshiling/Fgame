package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MAJOR_INVITE_TYPE), (*uipb.CSMajorInvite)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MAJOR_INVITE_TYPE), (*uipb.SCMajorInvite)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MAJOR_INVITE_CANCLE_TYPE), (*uipb.CSMajorInviteCancle)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MAJOR_INVITE_CANCLE_TYPE), (*uipb.SCMajorInviteCancle)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MAJOR_INVITE_PUSH_SPOUSE_TYPE), (*uipb.SCMajorInvitePushSpouse)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MAJOR_INVITE_PUSH_CANCLE_TYPE), (*uipb.SCMajorInvitePushCancle)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MAJOR_INVITE_DEAL_TYPE), (*uipb.CSMajorInviteDeal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MAJOR_INVITE_DEAL_TYPE), (*uipb.SCMajorInviteDeal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MAJOR_SPOUSE_REFUSED_TYPE), (*uipb.SCMajorSpouseRefused)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MAJOR_RESULT_TYPE), (*uipb.SCMajorResult)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MAJOR_SCENE_INFO_TYPE), (*uipb.SCMajorSceneInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MAJOR_NUM_TYPE), (*uipb.CSMajorNum)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MAJOR_NUM_TYPE), (*uipb.SCMajorNum)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MAJOR_NUM_NOTICE_TYPE), (*uipb.SCMajorNumNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MAJOR_SAODANG_TYPE), (*uipb.CSMajorSaoDang)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MAJOR_SAODANG_TYPE), (*uipb.SCMajorSaoDang)(nil))
}
