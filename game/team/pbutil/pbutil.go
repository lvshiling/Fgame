package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/core/utils"
	commonpbutil "fgame/fgame/game/common/pbutil"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/team/team"
)

func BuildSCTeamGet(teamData *team.TeamObject, login bool, playerId int64) *uipb.SCTeamGet {
	teamGet := &uipb.SCTeamGet{}
	teamId := teamData.GetTeamId()
	teamGet.TeamId = &teamId
	for pos, member := range teamData.GetMemberList() {
		isPrepare := false
		if login && member.GetPlayerId() == playerId {
			isPrepare = true
		} else {
			pl := player.GetOnlinePlayerManager().GetPlayerById(member.GetPlayerId())
			if pl != nil {
				s := pl.GetScene()
				if s != nil {
					if s.MapTemplate().IsWorld() {
						isPrepare = true
					}
				}
			}
		}
		teamGet.MemberList = append(teamGet.MemberList, buildTeamMember(int32(pos), member, isPrepare))
	}
	if login {
		loginAgain := int32(1)
		teamGet.Login = &loginAgain
	}
	match := teamData.IsMatch()
	autoReview := teamData.IsAutoReview()
	purpose := int32(teamData.GetTeamPurpose())
	teamGet.Match = &match
	teamGet.AutoReview = &autoReview
	teamGet.Purpose = &purpose
	return teamGet
}

func BuildSCTeamInvite(invitedId int64) *uipb.SCTeamInvite {
	teamInvite := &uipb.SCTeamInvite{}
	teamInvite.InvitedId = &invitedId
	return teamInvite
}

func BuildSCTeamInviteToInvited(typ int32, id int64, inviteName string, teamName string) *uipb.SCTeamInviteToInvited {
	teamInviteToInvited := &uipb.SCTeamInviteToInvited{}
	teamInviteToInvited.Typ = &typ
	teamInviteToInvited.Id = &id
	teamInviteToInvited.InviteName = &inviteName
	teamInviteToInvited.TeamName = &teamName
	return teamInviteToInvited
}

func BuildSCTeamInviteBroadcast(result int32, name string) *uipb.SCTeamInviteBroadcast {
	teamInviteBroadcast := &uipb.SCTeamInviteBroadcast{}
	teamInviteBroadcast.Result = &result
	teamInviteBroadcast.PlayerName = &name
	return teamInviteBroadcast
}

func BuildSCTeamBroadcast(typ int32, name string, teamData *team.TeamObject) *uipb.SCTeamBroadcast {
	teamBroadcast := &uipb.SCTeamBroadcast{}
	teamBroadcast.Typ = &typ
	teamBroadcast.Name = &name
	teamId := teamData.GetTeamId()
	teamBroadcast.TeamId = &teamId
	autoReview := teamData.IsAutoReview()
	teamBroadcast.AutoReview = &autoReview
	for pos, member := range teamData.GetMemberList() {
		isPrepare := false
		playerId := member.GetPlayerId()
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl != nil {
			s := pl.GetScene()
			if s != nil {
				if s.MapTemplate().IsWorld() {
					isPrepare = true
				}
			}
		}
		teamBroadcast.MemberList = append(teamBroadcast.MemberList, buildTeamMember(int32(pos), member, isPrepare))
	}
	purpose := int32(teamData.GetTeamPurpose())
	teamBroadcast.Purpose = &purpose
	return teamBroadcast
}

func BuildSCTeamLeave(name string) *uipb.SCTeamLeave {
	teamLeave := &uipb.SCTeamLeave{}
	teamLeave.Name = &name
	return teamLeave
}

func BuildSCTeamByLeavedToLeave(name string) *uipb.SCTeamByLeavedToLeave {
	teamByLeavedToLeave := &uipb.SCTeamByLeavedToLeave{}
	teamByLeavedToLeave.Name = &name
	return teamByLeavedToLeave
}

func BuildSCTeamNearGet(mapId int32, teamList []*team.TeamObject) *uipb.SCTeamNearGet {
	teamNearGet := &uipb.SCTeamNearGet{}
	for _, teamData := range teamList {
		teamNearGet.TeamNearList = append(teamNearGet.TeamNearList, buildTeamObject(mapId, teamData))
	}
	return teamNearGet
}

func BuildSCTeamNearJoin(result int32, teamId int64) *uipb.SCTeamNearJoin {
	teamNearJoin := &uipb.SCTeamNearJoin{}
	teamNearJoin.TeamId = &teamId
	teamNearJoin.Result = &result
	return teamNearJoin
}

func BuildSCTeamNearJoinResultToApply(name string) *uipb.SCTeamNearJoinResultToApply {
	teamNearJoinResultToApply := &uipb.SCTeamNearJoinResultToApply{}
	teamNearJoinResultToApply.Name = &name
	return teamNearJoinResultToApply
}

func BuildSCTeamNearJoinResult(applyId int64) *uipb.SCTeamNearJoinResult {
	tamNearJoinResult := &uipb.SCTeamNearJoinResult{}
	tamNearJoinResult.ApplyId = &applyId
	return tamNearJoinResult
}

func BuildSCTeamNearPlayerGet(playerList []player.Player) *uipb.SCTeamNearPlayerGet {
	teamNearPlayerGet := &uipb.SCTeamNearPlayerGet{}
	for _, pl := range playerList {
		if pl == nil {
			continue
		}
		teamNearPlayerGet.PlayerNearList = append(teamNearPlayerGet.PlayerNearList, buildNearPlayer(pl))
	}
	return teamNearPlayerGet
}

func BuildSCTeamApplyGet(applyList []*team.TeamApplyData) *uipb.SCTeamApplyGet {
	teamApplyGet := &uipb.SCTeamApplyGet{}

	for _, applyData := range applyList {
		teamApplyGet.ApplyList = append(teamApplyGet.ApplyList, buildApply(applyData))
	}
	return teamApplyGet
}

func BuildSCTeamMemberLogin(playerId int64, hp int64, maxHp int64) *uipb.SCTeamDataChanged {
	teamDataChanged := &uipb.SCTeamDataChanged{}
	teamDataChanged.PlayerId = &playerId
	teamDataChanged.Hp = &hp
	teamDataChanged.MaxHp = &maxHp
	online := int32(1)
	teamDataChanged.Online = &online
	return teamDataChanged
}

func BuildSCTeamForceChange(playerId int64, force int64) *uipb.SCTeamDataChanged {
	teamDataChanged := &uipb.SCTeamDataChanged{}
	teamDataChanged.PlayerId = &playerId
	teamDataChanged.Force = &force
	return teamDataChanged
}

func BuildSCTeamHpChange(playerId int64, hp int64) *uipb.SCTeamDataChanged {
	teamDataChanged := &uipb.SCTeamDataChanged{}
	teamDataChanged.PlayerId = &playerId
	teamDataChanged.Hp = &hp
	return teamDataChanged
}

func BuildSCTeamMaxHpChange(playerId int64, maxHp int64) *uipb.SCTeamDataChanged {
	teamDataChanged := &uipb.SCTeamDataChanged{}
	teamDataChanged.PlayerId = &playerId
	teamDataChanged.MaxHp = &maxHp
	return teamDataChanged
}

func BuildSCTeamLevelChange(playerId int64, level int32) *uipb.SCTeamDataChanged {
	teamDataChanged := &uipb.SCTeamDataChanged{}
	teamDataChanged.PlayerId = &playerId
	teamDataChanged.Level = &level
	return teamDataChanged
}

func BuildSCTeamFashionChange(playerId int64, fashionId int32) *uipb.SCTeamDataChanged {
	teamDataChanged := &uipb.SCTeamDataChanged{}
	teamDataChanged.PlayerId = &playerId
	teamDataChanged.FashionId = &fashionId
	return teamDataChanged
}

func BuildSCTeamZhuanShengChange(playerId int64, zhuanSheng int32) *uipb.SCTeamDataChanged {
	teamDataChanged := &uipb.SCTeamDataChanged{}
	teamDataChanged.PlayerId = &playerId
	teamDataChanged.ZhuanSheng = &zhuanSheng
	return teamDataChanged
}

func BuildSCTeamOnlineChange(playerId int64, online bool) *uipb.SCTeamDataChanged {
	teamDataChanged := &uipb.SCTeamDataChanged{}
	teamDataChanged.PlayerId = &playerId
	on := int32(0)
	if online {
		on = 1
	}
	teamDataChanged.Online = &on
	return teamDataChanged
}

func BuildSCTeamSexChange(playerId int64, sexType playertypes.SexType) *uipb.SCTeamDataChanged {
	teamDataChanged := &uipb.SCTeamDataChanged{}
	teamDataChanged.PlayerId = &playerId
	sex := int32(sexType)
	teamDataChanged.Sex = &sex
	return teamDataChanged
}

func BuildSCTeamNameChange(playerId int64, name string) *uipb.SCTeamDataChanged {
	teamDataChanged := &uipb.SCTeamDataChanged{}
	teamDataChanged.PlayerId = &playerId
	teamDataChanged.Name = &name
	return teamDataChanged
}

func BuildSCTeamPurposeChange(playerId int64, purpose int32) *uipb.SCTeamDataChanged {
	teamDataChanged := &uipb.SCTeamDataChanged{}
	teamDataChanged.PlayerId = &playerId
	teamDataChanged.Purpose = &purpose
	return teamDataChanged
}

func BuildSCTeamNearJoinToCaptain(applyId int64) *uipb.SCTeamNearJoinToCaptain {
	teamNearJoinToCaptain := &uipb.SCTeamNearJoinToCaptain{}
	teamNearJoinToCaptain.ApplyId = &applyId
	return teamNearJoinToCaptain
}

func BuildSCTeamApplyAll() *uipb.SCTeamApplyAll {
	teamApplyAll := &uipb.SCTeamApplyAll{}
	return teamApplyAll
}

func BuildSCTeamInviteAll() *uipb.SCTeamInviteAll {
	teamInviteAll := &uipb.SCTeamInviteAll{}
	return teamInviteAll
}

func BuildSCTeamMemberPos(playerId int64, teamData *team.TeamObject) *uipb.SCTeamMemberPos {
	teamMemberPos := &uipb.SCTeamMemberPos{}
	for _, memberObj := range teamData.GetMemberList() {
		memberId := memberObj.GetPlayerId()
		if memberId == playerId {
			continue
		}
		pl := player.GetOnlinePlayerManager().GetPlayerById(memberId)
		if pl == nil {
			continue
		}
		s := pl.GetScene()
		if s == nil {
			continue
		}
		teamMemberPos.MemberPosList = append(teamMemberPos.MemberPosList, buildMemberPos(pl))
	}
	return teamMemberPos
}

func BuildSCTeamAutoReview(autoReview bool) *uipb.SCTeamAutoReview {
	teamAutoReview := &uipb.SCTeamAutoReview{}
	teamAutoReview.AutoReview = &autoReview
	return teamAutoReview
}

func BuildSCTeamHousesGet(purpose int32, teamList []*team.TeamObject) *uipb.SCTeamHousesGet {
	teamHousesGet := &uipb.SCTeamHousesGet{}
	teamHousesGet.Purpose = &purpose
	for _, teamData := range teamList {
		teamHousesGet.TeamInfoList = append(teamHousesGet.TeamInfoList, buildTeamInfo(teamData))
	}
	return teamHousesGet
}

func BuildSCTeamCreateHouse(purpose int32, teamData *team.TeamObject) *uipb.SCTeamCreateHouse {
	teamCreateHouse := &uipb.SCTeamCreateHouse{}
	teamCreateHouse.Purpose = &purpose
	teamCreateHouse.TeamInfo = buildTeamInfo(teamData)
	return teamCreateHouse
}

func BuildSCTeamMatchConditionFail(teamData *team.TeamObject, memberIdList []int64) *uipb.SCTeamMatchCondtionFailed {
	scTeamMatchCondtionFailed := &uipb.SCTeamMatchCondtionFailed{}
	for index, member := range teamData.GetMemberList() {
		flag := utils.ContainInt64(memberIdList, member.GetPlayerId())
		if !flag {
			continue
		}
		scTeamMatchCondtionFailed.MemberList = append(scTeamMatchCondtionFailed.MemberList, buildTeamMember(int32(index), member, false))
	}
	return scTeamMatchCondtionFailed
}

func BuildSCTeamMatchCondtionFailedDeal(result bool) *uipb.SCTeamMatchCondtionFailedDeal {
	scTeamMatchCondtionFailedDeal := &uipb.SCTeamMatchCondtionFailedDeal{}
	scTeamMatchCondtionFailedDeal.Result = &result
	return scTeamMatchCondtionFailedDeal
}

func BuildSCTeamMatchCondtionFailedToPrepare() *uipb.SCTeamMatchCondtionFailedToPrepare {
	scTeamMatchCondtionFailedToPrepare := &uipb.SCTeamMatchCondtionFailedToPrepare{}
	return scTeamMatchCondtionFailedToPrepare
}

func BuildSCTeamMatchCondtionFailedBroadcast(memberIdList []int64) *uipb.SCTeamMatchCondtionFailedBroadcast {
	scTeamMatchCondtionFailedBroadcast := &uipb.SCTeamMatchCondtionFailedBroadcast{}
	for _, playerId := range memberIdList {
		scTeamMatchCondtionFailedBroadcast.PlayerIdList = append(scTeamMatchCondtionFailedBroadcast.PlayerIdList, playerId)
	}
	return scTeamMatchCondtionFailedBroadcast
}

func BuildSCTeamMatchCondtionPrepareDeal(result bool) *uipb.SCTeamMatchCondtionPrepareDeal {
	scTeamMatchCondtionPrepareDeal := &uipb.SCTeamMatchCondtionPrepareDeal{}
	scTeamMatchCondtionPrepareDeal.Result = &result
	return scTeamMatchCondtionPrepareDeal
}

func BuildSCTeamMatchCondtionPrepareBroadcast(playerId int64) *uipb.SCTeamMatchCondtionPrepareBroadcast {
	scTeamMatchCondtionPrepareBroadcast := &uipb.SCTeamMatchCondtionPrepareBroadcast{}
	scTeamMatchCondtionPrepareBroadcast.PlayerId = &playerId
	return scTeamMatchCondtionPrepareBroadcast
}

func BuildSCTeamMatchRushStart() *uipb.SCTeamMatchRushStart {
	scTeamMatchRushStart := &uipb.SCTeamMatchRushStart{}
	return scTeamMatchRushStart
}

func BuildSCTeamMatchRushToCaptain() *uipb.SCTeamMatchRushToCaptain {
	scTeamMatchRushToCaptain := &uipb.SCTeamMatchRushToCaptain{}
	return scTeamMatchRushToCaptain
}

func buildTeamInfo(teamData *team.TeamObject) *uipb.TeamInfo {
	teamInfo := &uipb.TeamInfo{}
	teamId := teamData.GetTeamId()
	match := teamData.IsMatch()
	autoReview := teamData.IsAutoReview()
	purpose := int32(teamData.GetTeamPurpose())
	teamInfo.TeamId = &teamId
	teamInfo.Match = &match
	teamInfo.AutoReview = &autoReview
	teamInfo.Purpose = &purpose

	for pos, member := range teamData.GetMemberList() {
		playerId := member.GetPlayerId()
		isPrepare := false
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl != nil {
			s := pl.GetScene()
			if s != nil {
				if s.MapTemplate().IsWorld() {
					isPrepare = true
				}
			}
		}
		teamInfo.MemberList = append(teamInfo.MemberList, buildTeamMember(int32(pos), member, isPrepare))
	}
	return teamInfo
}
func buildMemberPos(pl player.Player) *uipb.TeamMemberPos {
	teamMemberPos := &uipb.TeamMemberPos{}
	playerId := pl.GetId()
	mapId := pl.GetMapId()
	pos := pl.GetPos()
	teamMemberPos.PlayerId = &playerId
	teamMemberPos.MapId = &mapId
	teamMemberPos.Pos = commonpbutil.BuildPos(pos)
	return teamMemberPos
}

func buildApply(applyData *team.TeamApplyData) *uipb.TeamNearPlayer {
	teamNearPlayer := &uipb.TeamNearPlayer{}
	playerId := applyData.GetApplyId()
	role := int32(0)
	sex := int32(0)
	name := ""
	force := int64(0)
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl != nil {
		role = int32(pl.GetRole())
		sex = int32(pl.GetSex())
		name = pl.GetName()
		force = pl.GetForce()
	} else {
		playerInfo, _ := player.GetPlayerService().GetPlayerInfo(playerId)
		name = playerInfo.Name
		sex = int32(playerInfo.Sex)
		role = int32(playerInfo.Role)
		force = playerInfo.Force
	}

	teamNearPlayer.PlayerId = &playerId
	teamNearPlayer.Role = &role
	teamNearPlayer.Sex = &sex
	teamNearPlayer.Name = &name
	teamNearPlayer.Force = &force
	return teamNearPlayer

}

func buildNearPlayer(pl player.Player) *uipb.TeamNearPlayer {
	teamNearPlayer := &uipb.TeamNearPlayer{}
	playerId := pl.GetId()
	role := int32(pl.GetRole())
	sex := int32(pl.GetSex())
	name := pl.GetName()
	force := pl.GetForce()

	teamNearPlayer.Name = &name
	teamNearPlayer.PlayerId = &playerId
	teamNearPlayer.Role = &role
	teamNearPlayer.Sex = &sex
	teamNearPlayer.Force = &force
	return teamNearPlayer
}

//TODO 添加角色
func buildTeamObject(mapId int32, teamData *team.TeamObject) *uipb.TeamNearInfo {
	teamNearInfo := &uipb.TeamNearInfo{}
	teamNearInfo.MapId = &mapId
	teamId := teamData.GetTeamId()
	teamNearInfo.TeamId = &teamId
	for pos, member := range teamData.GetMemberList() {
		if pos == 0 {
			sex := int32(member.GetSex())
			name := member.GetName()
			teamNearInfo.Sex = &sex
			teamNearInfo.TeamName = &name
		}
		teamNearInfo.MemberList = append(teamNearInfo.MemberList, buildTeamNearMember(member))
	}
	return teamNearInfo
}

func buildTeamNearMember(member *team.TeamMemberObject) *uipb.TeamNearMember {
	teamNearMember := &uipb.TeamNearMember{}
	level := member.GetLevel()
	role := int32(member.GetRole())
	teamNearMember.Level = &level
	teamNearMember.Role = &role
	return teamNearMember
}

func buildTeamMember(pos int32, member *team.TeamMemberObject, isPrepare bool) *uipb.TeamMember {
	teamMember := &uipb.TeamMember{}
	playerId := member.GetPlayerId()
	name := member.GetName()
	level := member.GetLevel()
	role := int32(member.GetRole())
	sex := int32(member.GetSex())
	force := member.GetForce()
	fashionId := member.GetFashionId()
	zhuanSheng := member.GetZhuanSheng()
	online := int32(0)
	if member.GetOnline() {
		online = 1
	}

	hp := int64(0)
	maxHp := int64(0)
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl != nil {
		hp = pl.GetHP()
		maxHp = pl.GetMaxHP()
	}

	teamMember.Pos = &pos
	teamMember.PlayerId = &playerId
	teamMember.Name = &name
	teamMember.Level = &level
	teamMember.Role = &role
	teamMember.Sex = &sex
	teamMember.Hp = &hp
	teamMember.MaxHp = &maxHp
	teamMember.Force = &force
	teamMember.FashionId = &fashionId
	teamMember.Online = &online
	teamMember.ZhuanSheng = &zhuanSheng
	teamMember.IsPrepare = &isPrepare
	return teamMember
}
