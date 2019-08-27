package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIANFU_GET_TYPE), (*uipb.CSXianfuGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANFU_GET_TYPE), (*uipb.SCXianfuGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIANFU_UPGRADE_TYPE), (*uipb.CSXianfuUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANFU_UPGRADE_TYPE), (*uipb.SCXianfuUpgrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIANFU_ACCELERATE_TYPE), (*uipb.CSXianfuAccelerate)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANFU_ACCELERATE_TYPE), (*uipb.SCXianfuAccelerate)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIANFU_SAODANG_TYPE), (*uipb.CSXianfuSaoDang)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANFU_SAODANG_TYPE), (*uipb.SCXianfuSaoDang)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIANFU_CHALLENGE_TYPE), (*uipb.CSXianfuChallenge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANFU_CHALLENGE_TYPE), (*uipb.SCXianfuChallenge)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANFU_SCENEINFO_TYPE), (*uipb.SCXianfuSceneInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANFU_CHALLENGE_RESULT_TYPE), (*uipb.SCXianfuChallengeResult)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANFU_REFRESH_BIOLOGY_TYPE), (*uipb.SCXianfuRefreshBiology)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_XIANFU_FINISH_ALL_TYPE), (*uipb.CSXianfuFinishAll)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANFU_FINISH_ALL_TYPE), (*uipb.SCXianfuFinishAll)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANFU_REW_NOTICE_TYPE), (*uipb.SCXianfuRewNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANFU_KILL_NUM_NOTICE_TYPE), (*uipb.SCXianfuKillNumNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_XIANFU_BOSS_HP_CHANGED_NOTICE_TYPE), (*uipb.SCXianfuBossHpChangedNotice)(nil))
} 
