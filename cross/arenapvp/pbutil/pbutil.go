package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	arenapvpdata "fgame/fgame/game/arenapvp/data"
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	scenepbutil "fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

func BuildSCArenapvpSceneData(reliveTimes int32, s scene.Scene, pvpType int32, battlePl1, battlePl2 *arenapvpdata.PvpPlayerInfo) *uipb.SCArenapvpSceneData {
	scMsg := &uipb.SCArenapvpSceneData{}
	if battlePl1 != nil {
		pl := s.GetPlayer(battlePl1.PlayerId)
		scMsg.BattleList = append(scMsg.BattleList, buildArenapvpPlayerAllData(pl, battlePl1))
	}
	if battlePl2 != nil {
		pl := s.GetPlayer(battlePl2.PlayerId)
		scMsg.BattleList = append(scMsg.BattleList, buildArenapvpPlayerAllData(pl, battlePl2))
	}
	scMsg.PvpType = &pvpType
	return scMsg
}

func buildBattlePlayerBasicInfo(battlePl *arenapvpdata.PvpPlayerInfo) *uipb.BattlePlayerBasicInfo {
	info := &uipb.BattlePlayerBasicInfo{}
	info.Platform = &battlePl.Platform
	info.PlayerId = &battlePl.PlayerId
	info.PlayerName = &battlePl.PlayerName
	info.ServerId = &battlePl.ServerId
	info.Role = &battlePl.Role
	info.Sex = &battlePl.Sex
	info.WeaponId = &battlePl.WeaponId
	info.WingId = &battlePl.WingId
	info.FashionId = &battlePl.FashionId

	return info
}

func buildBattlePlayerShowData(pl scene.Player, battlePl *arenapvpdata.PvpPlayerInfo) *uipb.ArenapvpPlayerShowData {
	playerId := battlePl.PlayerId
	state := int32(battlePl.GetState())

	playerShowData := &uipb.ArenapvpPlayerShowData{}
	playerShowData.PlayerId = &playerId
	playerShowData.State = &state

	if pl == nil {
		online := int32(arenapvpdata.BattlePlayerStatusOffline)
		playerShowData.Online = &online
		return playerShowData
	}

	remainReliveTime := pl.GetArenapvpReliveTimes()
	maxHP := pl.GetMaxHP()
	hp := pl.GetHP()
	online := int32(arenapvpdata.BattlePlayerStatusOnline)

	playerShowData.RemainReliveTime = &remainReliveTime
	playerShowData.MaxHp = &maxHP
	playerShowData.Hp = &hp
	playerShowData.Online = &online

	if pl.IsDead() {
		isDead := int32(1)
		deadTime := pl.GetDeadTime()
		playerShowData.IsDead = &isDead
		playerShowData.DeadTime = &deadTime
	} else {
		isDead := int32(0)
		playerShowData.IsDead = &isDead
	}
	return playerShowData
}

func BuildSCArenapvpBattleStart() *uipb.SCArenapvpBattleStart {
	scMsg := &uipb.SCArenapvpBattleStart{}
	return scMsg
}

func buildArenapvpPlayerAllData(pl scene.Player, battlePl *arenapvpdata.PvpPlayerInfo) *uipb.ArenapvpPlayerAllData {
	playerAllData := &uipb.ArenapvpPlayerAllData{}
	playerAllData.BasicInfo = buildBattlePlayerBasicInfo(battlePl)
	playerAllData.ShowData = buildBattlePlayerShowData(pl, battlePl)

	return playerAllData
}

func BuildSCArenapvpBattleEnd(winnerId int64, pvpType arenapvptypes.ArenapvpType) *uipb.SCArenapvpBattleEnd {
	typeInt := int32(pvpType)
	scMsg := &uipb.SCArenapvpBattleEnd{}
	scMsg.WinnerPlayerId = &winnerId
	scMsg.PvpType = &typeInt

	return scMsg
}

func buildArenapvpBattlePlayer(battlePl *arenapvpdata.PvpPlayerInfo, pvpType arenapvptypes.ArenapvpType) *uipb.ArenapvpBattlePlayer {
	info := &uipb.ArenapvpBattlePlayer{}
	result := battlePl.GetBattleData(pvpType)
	if result == nil {
		return info
	}
	typeInt := int32(pvpType)
	info.Type = &typeInt
	info.Index = &result.Index
	info.WinnerId = &result.WinnerId
	info.BasicInfo = buildBattlePlayerBasicInfo(battlePl)
	return info
}

func BuildSCArenapvpElectionFailedNotice() *uipb.SCArenapvpElectionFailedNotice {
	scMsg := &uipb.SCArenapvpElectionFailedNotice{}
	return scMsg
}

func BuildSCArenapvpElectionSceneData(pl scene.Player, rankMap map[scenetypes.SceneRankType]*scene.SceneRank) *uipb.SCArenapvpElectionSceneData {
	reliveTimes := pl.GetArenapvpReliveTimes()

	scMsg := &uipb.SCArenapvpElectionSceneData{}
	scMsg.ReliveTimes = &reliveTimes
	for _, r := range rankMap {
		scMsg.RankInfoList = append(scMsg.RankInfoList, scenepbutil.BuildSceneRankInfo(pl, r))
	}
	return scMsg
}

func BuildSCArenapvpElectionSceneDataChanged(reliveTimes int32) *uipb.SCArenapvpElectionSceneDataChanged {
	scMsg := &uipb.SCArenapvpElectionSceneDataChanged{}
	scMsg.ReliveTimes = &reliveTimes
	return scMsg
}

func BuildSCArenapvpPlayerShowDataReliveTimeChanged(pl scene.Player) *uipb.SCArenapvpPlayerShowDataChanged {
	playerId := pl.GetId()
	reliveTimes := pl.GetArenapvpReliveTimes()
	playerShowData := &uipb.ArenapvpPlayerShowData{}
	playerShowData.PlayerId = &playerId
	playerShowData.RemainReliveTime = &reliveTimes

	scMsg := &uipb.SCArenapvpPlayerShowDataChanged{}
	scMsg.PlayerShowData = playerShowData
	return scMsg
}

func BuildSCArenapvpPlayerShowDataDeadChanged(pl scene.Player) *uipb.SCArenapvpPlayerShowDataChanged {
	scMsg := &uipb.SCArenapvpPlayerShowDataChanged{}
	playerShowData := &uipb.ArenapvpPlayerShowData{}
	scMsg.PlayerShowData = playerShowData
	playerId := pl.GetId()
	playerShowData.PlayerId = &playerId
	if pl.IsDead() {
		isDead := int32(1)
		deadTime := pl.GetDeadTime()
		playerShowData.IsDead = &isDead
		playerShowData.DeadTime = &deadTime
	} else {
		isDead := int32(0)
		playerShowData.IsDead = &isDead
	}
	return scMsg
}

func BuildSCArenapvpPlayerShowDataHpChanged(pl scene.Player) *uipb.SCArenapvpPlayerShowDataChanged {
	scMsg := &uipb.SCArenapvpPlayerShowDataChanged{}
	playerShowData := &uipb.ArenapvpPlayerShowData{}
	scMsg.PlayerShowData = playerShowData
	playerId := pl.GetId()
	playerShowData.PlayerId = &playerId
	hp := pl.GetHP()
	playerShowData.Hp = &hp
	return scMsg
}

func BuildSCArenapvpPlayerShowDataMaxHpChanged(pl scene.Player) *uipb.SCArenapvpPlayerShowDataChanged {
	scMsg := &uipb.SCArenapvpPlayerShowDataChanged{}
	playerShowData := &uipb.ArenapvpPlayerShowData{}
	scMsg.PlayerShowData = playerShowData
	playerId := pl.GetId()
	playerShowData.PlayerId = &playerId
	hp := pl.GetHP()
	playerShowData.Hp = &hp
	maxHp := pl.GetMaxHP()
	playerShowData.MaxHp = &maxHp
	return scMsg
}

func BuildSCArenapvpPlayerShowDataMaxOnlineChanged(pl scene.Player) *uipb.SCArenapvpPlayerShowDataChanged {
	scMsg := &uipb.SCArenapvpPlayerShowDataChanged{}
	playerShowData := &uipb.ArenapvpPlayerShowData{}
	scMsg.PlayerShowData = playerShowData
	playerId := pl.GetId()
	playerShowData.PlayerId = &playerId
	hp := pl.GetHP()
	playerShowData.Hp = &hp
	maxHp := pl.GetMaxHP()
	playerShowData.MaxHp = &maxHp
	return scMsg
}

func BuildSCArenapvpPlayerOnlineChanged(pl scene.Player) *uipb.SCArenapvpPlayerShowDataChanged {
	scMsg := &uipb.SCArenapvpPlayerShowDataChanged{}
	playerShowData := &uipb.ArenapvpPlayerShowData{}
	scMsg.PlayerShowData = playerShowData
	playerId := pl.GetId()
	playerShowData.PlayerId = &playerId
	online := int32(1)
	playerShowData.Online = &online
	return scMsg
}

func BuildSCArenapvpPlayerOfflineChanged(pl scene.Player) *uipb.SCArenapvpPlayerShowDataChanged {
	playerId := pl.GetId()
	online := int32(0)
	playerShowData := &uipb.ArenapvpPlayerShowData{}
	playerShowData.PlayerId = &playerId
	playerShowData.Online = &online

	scMsg := &uipb.SCArenapvpPlayerShowDataChanged{}
	scMsg.PlayerShowData = playerShowData
	return scMsg
}

func BuildSCArenapvpPlayerStateChanged(battlePl *arenapvpdata.PvpPlayerInfo) *uipb.SCArenapvpPlayerShowDataChanged {
	scMsg := &uipb.SCArenapvpPlayerShowDataChanged{}
	playerShowData := &uipb.ArenapvpPlayerShowData{}
	scMsg.PlayerShowData = playerShowData
	playerId := battlePl.PlayerId
	playerShowData.PlayerId = &playerId

	state := int32(battlePl.GetState())
	fmt.Println("状态", state)
	playerShowData.State = &state
	return scMsg
}

func BuildSCArenapvpElectionEnd(isWin bool) *uipb.SCArenapvpElectionEnd {
	scMsg := &uipb.SCArenapvpElectionEnd{}
	scMsg.IsWin = &isWin
	return scMsg
}
