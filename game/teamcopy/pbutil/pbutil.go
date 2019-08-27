package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	teamtypes "fgame/fgame/game/team/types"
	playerteam "fgame/fgame/game/teamcopy/player"
)

func BuildSCTeamCopyAllGet(copyMap map[teamtypes.TeamPurposeType]*playerteam.PlayerTeamCopyObject) *uipb.SCTeamCopyAllGet {
	teamCopyAllGet := &uipb.SCTeamCopyAllGet{}
	for _, teamCopyObj := range copyMap {
		teamCopyAllGet.TeamCopyList = append(teamCopyAllGet.TeamCopyList, buildTeamCopyInfo(teamCopyObj))
	}
	return teamCopyAllGet
}

func BuildSCTeamCopyStartBattle() *uipb.SCTeamCopyStartBattle {
	teamCopyStartBattle := &uipb.SCTeamCopyStartBattle{}
	return teamCopyStartBattle
}

func BuildSCTeamCopyStartBattleBroadcast(purpose int32) *uipb.SCTeamCopyStartBattleBroadcast {
	teamCopyStartBattleBroadcast := &uipb.SCTeamCopyStartBattleBroadcast{}
	teamCopyStartBattleBroadcast.Purpose = &purpose
	return teamCopyStartBattleBroadcast
}

func BuildSCTeamCopyStartBattleResultBroadcast(purpose int32, sucess bool) *uipb.SCTeamCopyStartBattleResultBroadcast {
	resultBroadcast := &uipb.SCTeamCopyStartBattleResultBroadcast{}
	resultBroadcast.Purpose = &purpose
	resultBroadcast.Sucess = &sucess
	return resultBroadcast
}

func BuildSCTeamCopyResult(obj *playerteam.PlayerTeamCopyObject, success bool, isRew bool) *uipb.SCTeamCopyResult {
	teamCopyResult := &uipb.SCTeamCopyResult{}
	purpose := int32(obj.GetPurpose())
	teamCopyResult.Sucess = &success
	teamCopyResult.TeamCopyInfo = buildTeamCopyInfo(obj)
	teamCopyResult.IsRew = &isRew
	teamCopyResult.Purpose = &purpose
	return teamCopyResult
}

func buildTeamCopyInfo(obj *playerteam.PlayerTeamCopyObject) *uipb.TeamCopyInfo {
	teamCopyInfo := &uipb.TeamCopyInfo{}
	purpose := int32(obj.GetPurpose())
	num := int32(obj.GetNum())
	teamCopyInfo.Purpose = &purpose
	teamCopyInfo.Num = &num
	return teamCopyInfo
}
