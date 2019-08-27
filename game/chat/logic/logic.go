package logic

import (
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/chat/chat"
	"fgame/fgame/game/chat/pbutil"
	chattemplate "fgame/fgame/game/chat/template"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	teamlogic "fgame/fgame/game/team/logic"
	"fgame/fgame/game/team/team"
	"fgame/fgame/pkg/timeutils"
)

func SystemBroadcastAllianceId(allianceId int64, msgType chattypes.MsgType, content []byte) {
	al := alliance.GetAllianceService().GetAlliance(allianceId)
	if al == nil {
		return
	}
	chatRecv := pbutil.BuildSCChatRecv(int64(0), chattypes.ChannelTypeBangPai, allianceId, msgType, content)

	for _, mem := range al.GetMemberList() {
		pl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
		if pl == nil {
			continue
		}
		pl.SendMsg(chatRecv)
	}
}

//系统帮派广播
func SystemBroadcastAlliance(al *alliance.Alliance, msgType chattypes.MsgType, content []byte) {
	if al == nil {
		return
	}
	allianceId := al.GetAllianceId()
	chatRecv := pbutil.BuildSCChatRecv(int64(0), chattypes.ChannelTypeBangPai, allianceId, msgType, content)

	for _, mem := range al.GetMemberList() {
		pl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
		if pl == nil {
			continue
		}
		pl.SendMsg(chatRecv)
	}
}

//帮派广播
func BroadcastAlliance(allianceId int64, sendId int64, sendName string, msgType chattypes.MsgType, content []byte, args string) {
	BroadcastAllianceToChatType(allianceId, sendId, sendName, chattypes.ChannelTypeBangPai, msgType, content, args)
}

//帮派广播
func BroadcastAllianceSystem(allianceId int64, sendId int64, sendName string, msgType chattypes.MsgType, content []byte, args string) {
	BroadcastAllianceToChatType(allianceId, sendId, sendName, chattypes.ChannelTypeSystem, msgType, content, args)
}

//帮派广播
func BroadcastAllianceToChatType(allianceId int64, sendId int64, sendName string, chatType chattypes.ChannelType, msgType chattypes.MsgType, content []byte, args string) {
	chatRecv := pbutil.BuildSCChatRecvWithCliArgs(sendId, sendName, chatType, allianceId, msgType, content, args)
	al := alliance.GetAllianceService().GetAlliance(allianceId)
	if al == nil {
		return
	}
	for _, mem := range al.GetMemberList() {
		pl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
		if pl == nil {
			continue
		}
		pl.SendMsg(chatRecv)
	}
}

func BroadcastAllianceExcludeSelf(allianceId int64, sendId int64, sendName string, msgType chattypes.MsgType, content []byte, args string) {
	chatRecv := pbutil.BuildSCChatRecvWithCliArgs(sendId, sendName, chattypes.ChannelTypeBangPai, allianceId, msgType, content, args)
	al := alliance.GetAllianceService().GetAlliance(allianceId)
	if al == nil {
		return
	}
	for _, mem := range al.GetMemberList() {
		pl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
		if pl == nil {
			continue
		}
		if pl.GetId() == sendId {
			continue
		}
		pl.SendMsg(chatRecv)
	}
}

//系统场景广播
func BroadcastScene(s scene.Scene, msgType chattypes.MsgType, content []byte) {
	if s == nil {
		return
	}
	chatRecv := pbutil.BuildSCChatRecv(int64(0), chattypes.ChannelTypeSystem, int64(0), msgType, content)
	s.BroadcastMsg(chatRecv)
}

//系统全服广播
func SystemBroadcast(msgType chattypes.MsgType, content []byte) {
	chatRecv := pbutil.BuildSCChatRecv(int64(0), chattypes.ChannelTypeSystem, int64(0), msgType, content)
	player.GetOnlinePlayerManager().BroadcastMsg(chatRecv)
	chat.GetChatService().AddSystemChat(content, msgType)
}

//队伍广播
func BroadcastTeam(teamId int64, sendId int64, msgType chattypes.MsgType, content []byte, args string) {
	chatRecv := pbutil.BuildSCChatRecvWithCliArgs(sendId, "", chattypes.ChannelTypeTeam, teamId, msgType, content, args)
	t := team.GetTeamService().GetTeam(teamId)
	if t == nil {
		return
	}
	teamlogic.BroadcastMsg(t, chatRecv)
}

//队伍广播
func BroadcastTeamExcludeSelf(teamId int64, sendId int64, sendName string, msgType chattypes.MsgType, content []byte, args string) {
	chatRecv := pbutil.BuildSCChatRecvWithCliArgs(sendId, sendName, chattypes.ChannelTypeTeam, teamId, msgType, content, args)
	t := team.GetTeamService().GetTeam(teamId)
	if t == nil {
		return
	}
	teamlogic.BroadcastPlayerMsg(t, sendId, chatRecv)
}

//答题频道广播
func BroadcastQuiz(sendId int64, msgType chattypes.MsgType, content []byte) {
	chatRecv := pbutil.BuildSCChatRecv(sendId, chattypes.ChannelTypeQuiz, int64(0), msgType, content)
	player.GetOnlinePlayerManager().BroadcastMsg(chatRecv)
}

//红包发送广播
func BroadcastHongBao(sendId int64, sendName string, content []byte, args string) {
	chatRecv := pbutil.BuildSCChatRecvWithCliArgs(sendId, sendName, chattypes.ChannelTypeWorld, int64(0), chattypes.MsgTypeHongBao, content, args)
	player.GetOnlinePlayerManager().BroadcastMsg(chatRecv)
}

//红包感谢广播
func BroadcastHongBaoThanks(pl player.Player, channelType chattypes.ChannelType, content []byte, args string) {
	switch channelType {
	case chattypes.ChannelTypeWorld:
		chatRecv := pbutil.BuildSCChatRecvWithCliArgs(pl.GetId(), pl.GetName(), chattypes.ChannelTypeWorld, int64(0), chattypes.MsgTypeText, content, args)
		player.GetOnlinePlayerManager().BroadcastMsg(chatRecv)
		break
	case chattypes.ChannelTypeBangPai:
		allianceId := pl.GetAllianceId()
		if allianceId == 0 {
			return
		}
		BroadcastAlliance(allianceId, pl.GetId(), pl.GetName(), chattypes.MsgTypeText, content, args)
		break
	}
}

//是否禁言时间
func IsForbiddenTime() bool {
	now := global.GetGame().GetTimeService().Now()
	startTimeStr := chattemplate.GetChatConstantService().GetChatConstant(chattypes.ChatConstantTypeStopChatStartTime)
	endTimeStr := chattemplate.GetChatConstantService().GetChatConstant(chattypes.ChatConstantTypeStopChatEndTime)
	beginDay, _ := timeutils.BeginOfNow(now)
	startTime, _ := timeutils.ParseDayOfHHMM(startTimeStr)
	startTime = beginDay + startTime
	endTime, _ := timeutils.ParseDayOfHHMM(endTimeStr)
	endTime = beginDay + endTime

	if now < startTime || now > endTime {
		return false
	}

	return true

}
