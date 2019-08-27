package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	collecttypes "fgame/fgame/game/collect/types"
	droppbutil "fgame/fgame/game/drop/pbutil"
	droptemplate "fgame/fgame/game/drop/template"
)

func BuildSCSceneCollect(success bool, npcId int64, startTime int64) *uipb.SCSceneCollect {
	scSceneCollect := &uipb.SCSceneCollect{}
	scSceneCollect.NpcId = &npcId
	scSceneCollect.StartTime = &startTime
	scSceneCollect.Success = &success
	return scSceneCollect
}

func BuildSCSceneCollectStop(npcId int64) *uipb.SCSceneCollectStop {
	scSceneCollectStop := &uipb.SCSceneCollectStop{}
	scSceneCollectStop.NpcId = &npcId
	return scSceneCollectStop
}

func BuildSCSceneCollectFinish(npcId int64, biologyId int32, dropItemList []*droptemplate.DropItemData) *uipb.SCSceneCollectFinish {
	scSceneCollectFinish := &uipb.SCSceneCollectFinish{}
	scSceneCollectFinish.NpcId = &npcId
	scSceneCollectFinish.BiologyId = &biologyId
	scSceneCollectFinish.DropInfoList = droppbutil.BuildDropInfoList(dropItemList)
	return scSceneCollectFinish
}

func BuildSCSceneCollectMiZangFinish(npcId int64, biologyId int32, parentId int32) *uipb.SCSceneCollectFinish {
	scSceneCollectFinish := &uipb.SCSceneCollectFinish{}
	scSceneCollectFinish.NpcId = &npcId
	scSceneCollectFinish.BiologyId = &biologyId
	scSceneCollectFinish.BossId = &parentId
	return scSceneCollectFinish
}

func BuildSCSceneCollectChooseResult(npcId int64, typ collecttypes.CollectChooseFinishType, isSuccess bool) *uipb.SCSceneCollectChooseResult {
	scSceneCollectChooseResult := &uipb.SCSceneCollectChooseResult{}
	typInt := int32(typ)
	scSceneCollectChooseResult.NpcId = &npcId
	scSceneCollectChooseResult.Result = &typInt
	scSceneCollectChooseResult.Success = &isSuccess
	return scSceneCollectChooseResult
}

func BuildSCSceneCollectMiZangOpen(openType int32, npcId int64, dropItemList []*droptemplate.DropItemData) *uipb.SCSceneCollectMiZangOpen {
	scMsg := &uipb.SCSceneCollectMiZangOpen{}
	scMsg.Type = &openType
	scMsg.NpcId = &npcId
	scMsg.DropInfoList = droppbutil.BuildDropInfoList(dropItemList)
	return scMsg
}

func BuildSCSceneCollectMiZangGiveup(npcId int64) *uipb.SCSceneCollectMiZangGiveup {
	scMsg := &uipb.SCSceneCollectMiZangGiveup{}
	scMsg.NpcId = &npcId
	return scMsg
}
