package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
	"fgame/fgame/game/processor"
)

func init() {
	initAlliance()
	initAllianceScene()
	initProxy()
	initAllianceBoss()
}

func initAllianceScene() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_SCENE_INFO_TYPE), (*uipb.SCAllianceSceneInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_SCENE_CALL_TYPE), (*uipb.CSAllianceSceneCall)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_SCENE_CALL_TYPE), (*uipb.SCAllianceSceneCall)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_SCENE_CALLED_GUARD_LIST_TYPE), (*uipb.CSAllianceSceneCalledGuardList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_SCENE_CALLED_GUARD_LIST_TYPE), (*uipb.SCAllianceSceneCalledGuardList)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_SCENE_DOOR_BROKE_TYPE), (*uipb.SCAllianceSceneDoorBroke)(nil))
	// gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_SCENE_OCCUPYING_TYPE), (*uipb.SCAllianceSceneOccupying)(nil))
	// gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_SCENE_OCCUPY_STOP_TYPE), (*uipb.SCAllianceSceneOccupyStop)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLAINCE_SCENE_OCCUPY_TYPE), (*uipb.SCAllianceSceneOccupy)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLAINCE_SCENE_FINISH_TYPE), (*uipb.SCAllianceSceneFinish)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_SCENE_DEFEND_HU_FU_CHANGED_TYPE), (*uipb.SCAllianceSceneDefendHuFuChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_SCENE_GET_REWARD_TYPE), (*uipb.CSAllianceSceneGetReward)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_SCENE_GET_REWARD_TYPE), (*uipb.SCAllianceSceneGetReward)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_HEGEMON_INFO_TYPE), (*uipb.CSAllianceHegemonInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_HEGEMON_INFO_TYPE), (*uipb.SCAllianceHegemonInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_SCENE_RELIVE_OCCUPYING_TYPE), (*uipb.CSAllianceSceneReliveOccupying)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_SCENE_RELIVE_OCCUPYING_TYPE), (*uipb.SCAllianceSceneReliveOccupying)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_SCENE_RELIVE_OCCUPY_STOP_TYPE), (*uipb.SCAllianceSceneReliveOccupyStop)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_SCENE_RELIVE_OCCUPY_TYPE), (*uipb.SCAllianceSceneReliveOccupy)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_AGREE_JOIN_APPLY_BATCH_TYPE), (*uipb.CSAllianceAgreeJoinApplyBatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_AGREE_JOIN_APPLY_BATCH_TYPE), (*uipb.SCAllianceAgreeJoinApplyBatch)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_SCENE_WAR_POINT_CHANGED_TYPE), (*uipb.SCAllianceSceneWarPointChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_SCENE_YUXI_BROADCAST_TYPE), (*uipb.SCAllianceSceneYuXiBroadcast)(nil))
}

func initAlliance() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_CREATE_TYPE), (*uipb.CSAllianceCreate)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_CREATE_TYPE), (*uipb.SCAllianceCreate)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_JOIN_APPLY_TYPE), (*uipb.CSAllianceJoinApply)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_JOIN_APPLY_TYPE), (*uipb.SCAllianceJoinApply)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_JOIN_APPLY_BROADCAST_TYPE), (*uipb.SCAllianceJoinApplyBroadcast)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_LIST_TYPE), (*uipb.CSAllianceList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_LIST_TYPE), (*uipb.SCAllianceList)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_DISMISS_TYPE), (*uipb.CSAllianceDismiss)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_DISMISS_TYPE), (*uipb.SCAllianceDismiss)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_DISMISS_BROADCAST_TYPE), (*uipb.SCAllianceDismissBroadcast)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_AGREE_JOIN_APPLY_TYPE), (*uipb.CSAllianceAgreeJoinApply)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_AGREE_JOIN_APPLY_TYPE), (*uipb.SCAllianceAgreeJoinApply)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_AGREE_JOIN_APPLY_TO_APPLY_TYPE), (*uipb.SCAllianceAgreeJoinApplyToApply)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_AGREE_JOIN_APPLY_TO_MANAGER_TYPE), (*uipb.SCAllianceAgreeJoinApplyToManager)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_INFO_TYPE), (*uipb.SCAllianceInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_JOIN_APPLY_LIST_TYPE), (*uipb.CSAllianceJoinApplyList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_JOIN_APPLY_LIST_TYPE), (*uipb.SCAllianceJoinApplyList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_EXIT_TYPE), (*uipb.CSAllianceExit)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_EXIT_TYPE), (*uipb.SCAllianceExit)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_CHARM_TYPE), (*uipb.CSAllianceCharm)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_CHARM_TYPE), (*uipb.SCAllianceCharm)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_KICK_TYPE), (*uipb.CSAllianceKick)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_KICK_TYPE), (*uipb.SCAllianceKick)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_KICK_NOTICE_TYPE), (*uipb.SCAllianceKickNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_COMMIT_TYPE), (*uipb.CSAllianceCommit)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_COMMIT_TYPE), (*uipb.SCAllianceCommit)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_COMMIT_NOTICE_TYPE), (*uipb.SCAllianceCommitNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_TRANSFER_TYPE), (*uipb.CSAllianceTransfer)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_TRANSFER_TYPE), (*uipb.SCAllianceTransfer)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_TRANSFER_BROADCAST_TYPE), (*uipb.SCAllianceTransferBroadcast)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_IMPEACH_TYPE), (*uipb.CSAllianceImpeach)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_IMPEACH_TYPE), (*uipb.SCAllianceImpeach)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_IMPEACH_BROADCAST_TYPE), (*uipb.SCAllianceImpeachBroadcast)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_MEMBER_LIST_TYPE), (*uipb.CSAllianceMemberList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_MEMBER_LIST_TYPE), (*uipb.SCAllianceMemberList)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_INVITATION_TYPE), (*uipb.CSAllianceInvitation)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_INVITATION_TYPE), (*uipb.SCAllianceInvitation)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_INVITATION_NOTICE_TYPE), (*uipb.SCAllianceInvitationNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_AGREE_INVITATION_TYPE), (*uipb.CSAllianceAgreeInvitation)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_AGREE_INVITATION_TYPE), (*uipb.SCAllianceAgreeInvitation)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_AGREE_INVITATION_NOTICE_TYPE), (*uipb.SCAllianceAgreeInvitationNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_SKILL_UPGRADE_TYPE), (*uipb.CSAllianceSkillUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_SKILL_UPGRADE_TYPE), (*uipb.SCAllianceSkillUpgrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_DONATE_TYPE), (*uipb.CSAllianceDonate)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_DONATE_TYPE), (*uipb.SCAllianceDonate)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_DONATE_HUFU_TYPE), (*uipb.CSAllianceDonateHuFu)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_DONATE_HUFU_TYPE), (*uipb.SCAllianceDonateHuFu)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_DONATE_HUFU_BROADCAST_TYPE), (*uipb.SCAllianceDonateHuFuBroadcast)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_BRIEF_INFO_TYPE), (*uipb.CSAlliance)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCALLIANCE_BRIEF_INFO_TYPE), (*uipb.SCAlliance)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_NOTICE_CHANGE_TYPE), (*uipb.CSAllianceNoticeChange)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_NOTICE_CHANGE_TYPE), (*uipb.SCAllianceNoticeChange)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCALLIANCE_NOTICE_BROADCAST_TYPE), (*uipb.SCAllianceNoticeBroadcast)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_LOG_TYPE), (*uipb.CSAllianceLog)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_LOG_TYPE), (*uipb.SCAllianceLog)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_PLAYER_INFO_TYPE), (*uipb.CSAlliancePlayerInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_PLAYER_INFO_TYPE), (*uipb.SCAlliancePlayerInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_PLAYER_YAO_PAI_CHANGED), (*uipb.SCAlliancePlayerYaoPaiChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_YAO_PAI_CONVERT_TYPE), (*uipb.CSYaoPaiConvert)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_YAO_PAI_CONVERT_TYPE), (*uipb.SCYaoPaiConvert)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_DOU_SHEN_LIST_TYPE), (*uipb.CSAllianceDouShenList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_DOU_SHEN_LIST_TYPE), (*uipb.SCAllianceDouShenList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_DOU_SHEN_LINGYU_CHANGED_BROADCAST_TYPE), (*uipb.SCAllianceDouShenLingyuChangedBroadcast)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_SCENE_RELIVE_TIME_CHANGE_TYPE), (*uipb.SCAllianceSceneReliveTimeChange)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_MEMBER_CHANGED_TYPE), (*uipb.SCAllianceMemberChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SAVE_IN_ALLIANCE_DEPOT_TYPE), (*uipb.CSSaveInAllianceDepot)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SAVE_IN_ALLIANCE_DEPOT_TYPE), (*uipb.SCSaveInAllianceDepot)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TAKE_OUT_ALLIANCE_DEPOT_TYPE), (*uipb.CSTakeOutAllianceDepot)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TAKE_OUT_ALLIANCE_DEPOT_TYPE), (*uipb.SCTakeOutAllianceDepot)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MERGE_ALLIANCE_DEPOT_TYPE), (*uipb.CSAllianceDepotMerge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MERGE_ALLIANCE_DEPOT_TYPE), (*uipb.SCAllianceDepotMerge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_DEPOT_CHANGED_NOTICE_TYPE), (*uipb.SCAllianceDepotChangedNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_DEPOT_MERGE_NOTICE_TYPE), (*uipb.SCAllianceDepotMergeNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_MEMBER_CALL_TYPE), (*uipb.CSAllianceMemberCall)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_MEMBER_CALL_TYPE), (*uipb.SCAllianceMemberCall)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_MEMBER_CALL_BROADCAST_TYPE), (*uipb.SCAllianceMemberCallBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_MEMBER_RESCUE_TYPE), (*uipb.CSAllianceMemberRescue)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_MEMBER_RESCUE_TYPE), (*uipb.SCAllianceMemberRescue)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_MEMBER_POS_TYPE), (*uipb.CSAllianceMemberPos)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_MEMBER_POS_TYPE), (*uipb.SCAllianceMemberPos)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_AUTO_AGREE_JOIN_TYPE), (*uipb.CSAllianceAutoAgreeJoin)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_AUTO_AGREE_JOIN_TYPE), (*uipb.SCAllianceAutoAgreeJoin)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_JOIN_APPLAY_BATCH_TYPE), (*uipb.CSAllianceJoinApplyBatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_JOIN_APPLAY_BATCH_TYPE), (*uipb.SCAllianceJoinApplyBatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_DEPOT_AUTO_REMOVE_TYPE), (*uipb.CSAllianceDepotAutoRemove)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_DEPOT_AUTO_REMOVE_TYPE), (*uipb.SCAllianceDepotAutoRemove)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_DEPOT_SETTING_NOTICE_TYPE), (*uipb.SCAllianceDepotSettingNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_MENGZHU_INFO_NOTICE_TYPE), (*uipb.SCAllianceMengZhuInfoNotice)(nil))
}

func initAllianceBoss() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_BOSS_SUMMON_TYPE), (*uipb.CSAllianceBossSummon)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_BOSS_SUMMON_TYPE), (*uipb.SCAllianceBossSummon)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_BOSS_ENTER_TYPE), (*uipb.CSAllianceBossEnter)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_BOSS_ENTER_TYPE), (*uipb.SCAllianceBossEnter)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_BOSS_CHANGED_TYPE), (*uipb.SCAllianceBossChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_BOSS_RANK_CHANGED_TYPE), (*uipb.SCAllianceBossRank)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_BOSS_END_TYPE), (*uipb.SCAllianceBossEnd)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_BOSS_TYPE), (*uipb.CSAllianceBoss)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_BOSS_TYPE), (*uipb.SCAllianceBoss)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_BOSS_SUMMON_SUCESS_TYPE), (*uipb.SCAllianceBossSummonSucess)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_INVITE_MERGE_TYPE), (*uipb.CSAllianceInviteMerge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_INVITE_MERGE_TYPE), (*uipb.SCAllianceInviteMerge)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_INVITE_MERGE_NOTICE_TYPE), (*uipb.SCAllianceInviteMergeNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_INVITE_MERGE_FEEDBACK_TYPE), (*uipb.CSAllianceInviteMergeFeedback)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_INVITE_MERGE_FEEDBACK_TYPE), (*uipb.SCAllianceInviteMergeFeedback)(nil))
}

func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_ALLIANCE_MEMBER_POS_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_ALLIANCE_MEMBER_POS_TYPE))
}
