package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	commonpbutil "fgame/fgame/game/common/pbutil"
	droppbutil "fgame/fgame/game/drop/pbutil"
	"fgame/fgame/game/scene/scene"
)

func BuildSCYuXiPosBroadcast(pl scene.Player, keepStartTime int64) *uipb.SCYuXiPosBroadcast {
	scMsg := &uipb.SCYuXiPosBroadcast{}
	scMsg.OwnerInfo = buildYuXiOwner(pl, keepStartTime)
	return scMsg
}

func BuildSCYuXiCollectInfoBroadcast(npc scene.NPC, pl scene.Player, keepStartTime int64, rebornTypeInt int32) *uipb.SCYuXiCollectInfoBroadcast {
	scMsg := &uipb.SCYuXiCollectInfoBroadcast{}
	scMsg.CollectInfo = buildYuXiCollect(npc)
	if pl != nil {
		scMsg.OwnerInfo = buildYuXiOwner(pl, keepStartTime)
	}
	scMsg.RebornType = &rebornTypeInt
	return scMsg
}

func BuildSCYuXiReceiveDayRew(itemMap map[int32]int32) *uipb.SCYuXiReceiveDayRew {
	scMsg := &uipb.SCYuXiReceiveDayRew{}
	scMsg.DropInfo = droppbutil.BuildSimpleDropInfoList(itemMap)
	return scMsg
}

func BuildSCYuXiGetInfo(isReceive int32, winAllianceId int64, winAllianceName, winMengZhuName string) *uipb.SCYuXiGetInfo {
	scMsg := &uipb.SCYuXiGetInfo{}
	scMsg.IsReceive = &isReceive
	scMsg.WinAllianceId = &winAllianceId
	scMsg.WinAllianceName = &winAllianceName
	scMsg.WinMengName = &winMengZhuName
	return scMsg
}

func BuildSCYuXiWinnerBroadcast(winAllianceId int64) *uipb.SCYuXiWinnerBroadcast {
	scMsg := &uipb.SCYuXiWinnerBroadcast{}
	scMsg.WinAllianceId = &winAllianceId
	return scMsg
}

func buildYuXiOwner(pl scene.Player, keepStartTime int64) *uipb.YuXiOwner {
	playerId := pl.GetId()
	playerName := pl.GetName()

	info := &uipb.YuXiOwner{}
	info.KeepStartTime = &keepStartTime
	info.Pos = commonpbutil.BuildPos(pl.GetPos())
	info.PlayerId = &playerId
	info.PlayerName = &playerName
	return info
}

func buildYuXiCollect(npc scene.NPC) *uipb.YuXiCollect {
	npcId := npc.GetId()
	typ := int32(npc.GetBiologyTemplate().GetBiologyScriptType())
	status := npc.IsDead()
	statusTime := int64(0)
	if status {
		statusTime = npc.GetDeadTime()
	}
	pos := npc.GetPosition()
	biologyId := int32(npc.GetBiologyTemplate().Id)

	bio := &uipb.YuXiCollect{}
	bio.NcpId = &npcId
	bio.Typ = &typ
	bio.IsDead = &status
	bio.StatusTime = &statusTime
	bio.Pos = commonpbutil.BuildPos(pos)
	bio.BiologyId = &biologyId
	return bio
}
