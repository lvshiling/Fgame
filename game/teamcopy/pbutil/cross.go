package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	crosspbutil "fgame/fgame/game/cross/pbutil"
	"fgame/fgame/game/player"
)

func BuildSITeamCopyStartBattle(pList []player.Player) *crosspb.SITeamCopyStartBattle {
	siTeamCopyStartBattle := &crosspb.SITeamCopyStartBattle{}
	for _, p := range pList {
		teamPlayer := BuildTeamPlayer(p)
		siTeamCopyStartBattle.PlayerList = append(siTeamCopyStartBattle.PlayerList, teamPlayer)
	}
	return siTeamCopyStartBattle
}

func BuildSITeamCopyBattleResult() *crosspb.SITeamCopyBattleResult {
	siTeamCopyBattleResult := &crosspb.SITeamCopyBattleResult{}
	return siTeamCopyBattleResult
}

func BuildTeamPlayer(p player.Player) *crosspb.TeamPlayer {
	teamPlayer := &crosspb.TeamPlayer{}
	playerId := p.GetId()
	fashionId := p.GetFashionId()
	force := p.GetForce()
	level := p.GetLevel()
	name := p.GetOriginName()
	online := true
	serverId := p.GetServerId()
	role := int32(p.GetRole())
	sex := int32(p.GetSex())
	purpose := int32(p.GetTeamPurpose())
	teamPlayer.ServerId = &serverId
	teamPlayer.PlayerId = &playerId
	teamPlayer.FashionId = &fashionId
	teamPlayer.Force = &force
	teamPlayer.Level = &level
	teamPlayer.Name = &name
	teamPlayer.Online = &online
	teamPlayer.Role = &role
	teamPlayer.Sex = &sex
	teamPlayer.Purpose = &purpose
	teamPlayer.BattlePropertyData = crosspbutil.BuildBattlePropertyData(p.GetAllSystemBattleProperties())
	teamPlayer.SkillList = crosspbutil.BuildPlayerSkillDataList(p)
	return teamPlayer
}
