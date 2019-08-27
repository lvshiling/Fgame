package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
	"fgame/fgame/game/processor"
)

func init() {
	initCodec()
	initProxy()
}

func initCodec() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_GET_TYPE), (*uipb.CSTeamGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_GET_TYPE), (*uipb.SCTeamGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_CREATE_BY_PLAYER_TYPE), (*uipb.CSTeamCreateByPlayer)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_INVITE_TYPE), (*uipb.CSTeamInvite)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_INVITE_TYPE), (*uipb.SCTeamInvite)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_INVITE_TO_INVITED_TYPE), (*uipb.SCTeamInviteToInvited)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_INVITE_RESULT_TYPE), (*uipb.CSTeamInviteResult)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_INVITE_BROADCAST_TYPE), (*uipb.SCTeamInviteBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_BROADCAST_TYPE), (*uipb.SCTeamBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_TRANSFER_CAPTAIN_TYPE), (*uipb.CSTeamTransferCaptain)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_LEAVE_TYPE), (*uipb.CSTeamLeave)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_LEAVE_TYPE), (*uipb.SCTeamLeave)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_BY_LEAVED_TYPE), (*uipb.CSTeamByLeaved)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_BY_LEAVED_TO_LEAVE_TYPE), (*uipb.SCTeamByLeavedToLeave)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_NEAR_GET_TYPE), (*uipb.CSTeamNearGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_NEAR_GET_TYPE), (*uipb.SCTeamNearGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_NEAR_JOIN_TYPE), (*uipb.CSTeamNearJoin)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_NEAR_JOIN_TYPE), (*uipb.SCTeamNearJoin)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_NEAR_JOIN_RESULT_TYPE), (*uipb.CSTeamNearJoinResult)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_NEAR_JOIN_RESULT_TYPE), (*uipb.SCTeamNearJoinResult)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_NEAR_JOIN_RESULT_TO_APPLY_TYPE), (*uipb.SCTeamNearJoinResultToApply)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_NEAR_PLAYER_GET_TYPE), (*uipb.CSTeamNearPlayerGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_NEAR_PLAYER_GET_TYPE), (*uipb.SCTeamNearPlayerGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_APPLY_GET_TYPE), (*uipb.CSTeamApplyGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_APPLY_GET_TYPE), (*uipb.SCTeamApplyGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_CLEAR_APPLY_TYPE), (*uipb.CSTeamClearApply)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_DATA_CHANGE_TYPE), (*uipb.SCTeamDataChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_NEAR_JOIN_TO_CAPTAIN_TYPE), (*uipb.SCTeamNearJoinToCaptain)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_APPLY_ALL_TYPE), (*uipb.CSTeamApplyAll)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_APPLY_ALL_TYPE), (*uipb.SCTeamApplyAll)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_INVITE_ALL_TYPE), (*uipb.CSTeamInviteAll)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_INVITE_ALL_TYPE), (*uipb.SCTeamInviteAll)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_MEMBER_POS_TYPE), (*uipb.CSTeamMemberPos)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_MEMBER_POS_TYPE), (*uipb.SCTeamMemberPos)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_AUTO_REVIEW_TYPE), (*uipb.CSTeamAutoReview)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_AUTO_REVIEW_TYPE), (*uipb.SCTeamAutoReview)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_HOUSES_GET_TYPE), (*uipb.CSTeamHousesGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_HOUSES_GET_TYPE), (*uipb.SCTeamHousesGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_CREATE_HOUSE_TYPE), (*uipb.CSTeamCreateHouse)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_CREATE_HOUSE_TYPE), (*uipb.SCTeamCreateHouse)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_MATCH_CONDTION_FAILED_TYPE), (*uipb.SCTeamMatchCondtionFailed)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_MATCH_CONDTION_FAILED_DEAL_TYPE), (*uipb.CSTeamMatchCondtionFailedDeal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_MATCH_CONDTION_FAILED_DEAL_TYPE), (*uipb.SCTeamMatchCondtionFailedDeal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_MATCH_CONDTION_FAILED_BROADCAST_TYPE), (*uipb.SCTeamMatchCondtionFailedBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_MATCH_CONDTION_FAILED_TO_PREPARE_TYPE), (*uipb.SCTeamMatchCondtionFailedToPrepare)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_MATCH_CONDTION_PREPARE_DEAL_TYPE), (*uipb.CSTeamMatchCondtionPrepareDeal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_MATCH_CONDTION_PREPARE_DEAL_TYPE), (*uipb.SCTeamMatchCondtionPrepareDeal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_MATCH_CONDTION_PREPARE_BROADCAST_TYPE), (*uipb.SCTeamMatchCondtionPrepareBroadcast)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TEAM_MATCH_RUSH_START_TYPE), (*uipb.CSTeamMatchRushStart)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_MATCH_RUSH_START_TYPE), (*uipb.SCTeamMatchRushStart)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TEAM_MATCH_RUSH_TO_CAPTAIN_TYPE), (*uipb.SCTeamMatchRushToCaptain)(nil))
}

//TODO
func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_TEAM_MATCH_CONDTION_PREPARE_DEAL_TYPE))
}
