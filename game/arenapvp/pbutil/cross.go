package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildSIArenapvpAttend() *crosspb.SIArenapvpAttend {
	siMsg := &crosspb.SIArenapvpAttend{}
	return siMsg
}

func BuildSIArenapvpRelive(success bool) *crosspb.SIArenapvpRelive {
	siMsg := &crosspb.SIArenapvpRelive{}
	siMsg.Success = &success
	return siMsg
}

// func BuildArenapvpPlayer(p player.Player) *crosspb.ArenapvpPlayer {
// 	arenapvpPlayer := &crosspb.ArenapvpPlayer{}
// 	playerId := p.GetId()
// 	fashionId := p.GetFashionId()
// 	force := p.GetForce()
// 	level := p.GetLevel()
// 	name := p.GetOriginName()
// 	online := true
// 	serverId := p.GetServerId()
// 	role := int32(p.GetRole())
// 	sex := int32(p.GetSex())
// 	winCount := p.GetArenapvpWinTime()

// 	arenapvpPlayer.ServerId = &serverId
// 	arenapvpPlayer.PlayerId = &playerId
// 	arenapvpPlayer.FashionId = &fashionId
// 	arenapvpPlayer.Force = &force
// 	arenapvpPlayer.Level = &level
// 	arenapvpPlayer.Name = &name
// 	arenapvpPlayer.Online = &online
// 	arenapvpPlayer.Role = &role
// 	arenapvpPlayer.Sex = &sex
// 	arenapvpPlayer.BattlePropertyData = crosspbutil.BuildBattlePropertyData(p.GetAllSystemBattleProperties())
// 	arenapvpPlayer.SkillList = crosspbutil.BuildPlayerSkillDataList(p)
// 	arenapvpPlayer.WinCount = &winCount
// 	return arenapvpPlayer
// }
