package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droppbutil "fgame/fgame/game/drop/pbutil"
	droptemplate "fgame/fgame/game/drop/template"
)

func BuildSCSceneCollect(npcId int64, startTime int64) *uipb.SCSceneCollect {
	scSceneCollect := &uipb.SCSceneCollect{}
	scSceneCollect.NpcId = &npcId
	scSceneCollect.StartTime = &startTime
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
