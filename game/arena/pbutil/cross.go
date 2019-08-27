package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	crosspbutil "fgame/fgame/game/cross/pbutil"
	"fgame/fgame/game/player"
)

func BuildSIArenaMatch(pList []player.Player) *crosspb.SIArenaMatch {
	siArenaMatch := &crosspb.SIArenaMatch{}

	for _, p := range pList {
		arenaPlayer := BuildArenaPlayer(p)
		siArenaMatch.PlayerList = append(siArenaMatch.PlayerList, arenaPlayer)
	}

	return siArenaMatch
}

func BuildArenaPlayer(p player.Player) *crosspb.ArenaPlayer {
	arenaPlayer := &crosspb.ArenaPlayer{}
	playerId := p.GetId()
	fashionId := p.GetFashionId()
	force := p.GetForce()
	level := p.GetLevel()
	name := p.GetOriginName()
	online := true
	serverId := p.GetServerId()
	role := int32(p.GetRole())
	sex := int32(p.GetSex())
	winCount := p.GetArenaWinTime()

	arenaPlayer.ServerId = &serverId
	arenaPlayer.PlayerId = &playerId
	arenaPlayer.FashionId = &fashionId
	arenaPlayer.Force = &force
	arenaPlayer.Level = &level
	arenaPlayer.Name = &name
	arenaPlayer.Online = &online
	arenaPlayer.Role = &role
	arenaPlayer.Sex = &sex
	arenaPlayer.BattlePropertyData = crosspbutil.BuildBattlePropertyData(p.GetAllSystemBattleProperties())
	arenaPlayer.SkillList = crosspbutil.BuildPlayerSkillDataList(p)
	arenaPlayer.WinCount = &winCount
	return arenaPlayer
}

func BuildSIArenaStopMatch() *crosspb.SIArenaStopMatch {
	siArenaStopMatch := &crosspb.SIArenaStopMatch{}

	return siArenaStopMatch
}

func BuildSIArenaWin() *crosspb.SIArenaWin {
	siArenaWin := &crosspb.SIArenaWin{}
	return siArenaWin
}

func BuildSIArenaRelive(flag bool) *crosspb.SIArenaRelive {
	siArenaRelive := &crosspb.SIArenaRelive{}
	siArenaRelive.Success = &flag
	return siArenaRelive
}

var (
	siArenaCollectExpTree = &crosspb.SIArenaCollectExpTree{}
)

func BuildSIArenaCollectExpTree() *crosspb.SIArenaCollectExpTree {

	return siArenaCollectExpTree
}

var (
	siArenaCollectBox = &crosspb.SIArenaCollectBox{}
)

func BuildSIArenaCollectBox() *crosspb.SIArenaCollectBox {
	return siArenaCollectBox
}

func BuildSIArenaGiveUp() *crosspb.SIArenaGiveUp {
	siMsg := &crosspb.SIArenaGiveUp{}
	return siMsg
}
