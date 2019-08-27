package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/tower/tower"
)

func BuildSCTowerEnter(boss scene.NPC, remainTime int64, logList []*tower.TowerLogObject, floor int32) *uipb.SCTowerEnter {
	scMsg := &uipb.SCTowerEnter{}
	scMsg.RemainTime = &remainTime
	for _, log := range logList {
		scMsg.LogList = append(scMsg.LogList, buildTowerLog(log))
	}
	scMsg.Floor = &floor
	if boss != nil {
		deadTime := boss.GetDeadTime()
		scMsg.DeadTime = &deadTime
		isDead := boss.IsDead()
		scMsg.IsDead = &isDead
	}

	return scMsg
}

func BuildSCTowerTimeNotice(remainTime int64) *uipb.SCTowerTimeNotice {
	scMsg := &uipb.SCTowerTimeNotice{}
	scMsg.RemainTime = &remainTime
	return scMsg
}

func BuildSCTowerBossInfoReBornNotice(boss scene.NPC, bossName, mapName string, floor int32) *uipb.SCTowerBossInfoNotice {
	scMsg := &uipb.SCTowerBossInfoNotice{}
	scMsg.FloorInfo = buildFloorInfo(boss, floor)
	scMsg.BossName = &bossName
	scMsg.MapName = &mapName

	return scMsg
}

func BuildSCTowerBossInfoDeadNotice(boss scene.NPC, playerName, bossName string, floor int32) *uipb.SCTowerBossInfoNotice {
	scMsg := &uipb.SCTowerBossInfoNotice{}
	scMsg.FloorInfo = buildFloorInfo(boss, floor)
	scMsg.PlayerName = &playerName
	scMsg.BossName = &bossName

	return scMsg
}

func BuildSCTowerResultNotice(totalExp, remainTime int64, dropMap map[int32]int32) *uipb.SCTowerResultNotice {
	scMsg := &uipb.SCTowerResultNotice{}
	scMsg.TotalExp = &totalExp
	scMsg.RemainTime = &remainTime
	for itemId, num := range dropMap {
		scMsg.DropList = append(scMsg.DropList, buildDropInfo(itemId, num))
	}

	return scMsg
}

func BuildSCTowerFloorList(bossList map[int32]scene.NPC, time int64) *uipb.SCTowerFloorList {
	scMsg := &uipb.SCTowerFloorList{}

	for floor, boss := range bossList {
		scMsg.FloorList = append(scMsg.FloorList, buildFloorInfo(boss, floor))
	}
	scMsg.RemainTime = &time

	return scMsg
}

func BuildSCTowerLogIncr(logList []*tower.TowerLogObject) *uipb.SCTowerLogIncr {
	scMsg := &uipb.SCTowerLogIncr{}
	for _, log := range logList {
		scMsg.LogList = append(scMsg.LogList, buildTowerLog(log))
	}

	return scMsg
}

func BuildSCTowerDaBao(remainTime int64) *uipb.SCTowerDaBao {
	scMsg := &uipb.SCTowerDaBao{}
	scMsg.RemainTime = &remainTime
	return scMsg
}

func buildDropInfo(itemId, num int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num

	return dropInfo
}

func buildFloorInfo(boss scene.NPC, floor int32) *uipb.FloorInfo {
	floorInfo := &uipb.FloorInfo{}
	biologyId := int32(boss.GetBiologyTemplate().TemplateId())
	floorInfo.BiologyId = &biologyId
	deadTime := boss.GetDeadTime()
	floorInfo.DeadTime = &deadTime
	isDead := boss.IsDead()
	floorInfo.IsDead = &isDead
	floorInfo.Floor = &floor

	return floorInfo
}

func buildTowerLog(log *tower.TowerLogObject) *uipb.TowerLog {
	towerLog := &uipb.TowerLog{}
	createTime := log.GetCreateTime()
	towerLog.CreateTime = &createTime
	playerName := log.GetPlayerName()
	towerLog.PlayerName = &playerName
	biologyName := log.GetBiologyName()
	towerLog.BiologyName = &biologyName
	itemId := log.GetItemId()
	towerLog.ItemId = &itemId

	return towerLog
}
