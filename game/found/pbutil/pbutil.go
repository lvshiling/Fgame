package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droppbutil "fgame/fgame/game/drop/pbutil"
	playerfound "fgame/fgame/game/found/player"
	foundtemplate "fgame/fgame/game/found/template"
)

func BuildSCFoundResouceList(resList []*playerfound.PlayerFoundBackObject) *uipb.SCFoundResouceList {
	scFoundResouceList := &uipb.SCFoundResouceList{}

	for _, res := range resList {
		temp := foundtemplate.GetFoundTemplateService().GetFoundTemplateByType(res.GetResType(), res.GetResLevel())
		if temp == nil {
			continue
		}

		resId := int32(temp.TemplateId())
		num := res.GetFoundTimes()
		status := int32(res.GetFoundStatus())
		group := res.GetGroup()
		scFoundResouceList.FoundList = append(scFoundResouceList.FoundList, buildFoundBriefInfo(resId, num, status, group))
	}

	return scFoundResouceList
}

func buildFoundBriefInfo(resId, num, isReceive, group int32) *uipb.FoundBriefInfo {

	briefInfo := &uipb.FoundBriefInfo{}
	briefInfo.ResId = &resId
	briefInfo.Num = &num
	briefInfo.IsReceive = &isReceive
	briefInfo.Group = &group

	return briefInfo
}

func BuildSCFound(resType int32, itemMap map[int32]int32) *uipb.SCFound {
	scFound := &uipb.SCFound{}
	scFound.ResType = &resType
	scFound.DropInfoList = droppbutil.BuildSimpleDropInfoList(itemMap)
	return scFound
}

func BuildSCFoundBatch(itemMap map[int32]int32) *uipb.SCFoundBatch {
	scFoundBatch := &uipb.SCFoundBatch{}
	scFoundBatch.DropInfoList = droppbutil.BuildSimpleDropInfoList(itemMap)
	return scFoundBatch
}
