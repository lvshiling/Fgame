package pbuitl

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerdan "fgame/fgame/game/dan/player"
)

func BuildSCDanGet(danInfo *playerdan.PlayerDanObject) *uipb.SCDanGet {
	danGet := &uipb.SCDanGet{}
	levelId := danInfo.LevelId
	danGet.Level = &levelId
	danGet.DanList = buildDanList(danInfo.DanInfoMap)
	return danGet
}

func BuildSCDanUse(dans map[int]int32) *uipb.SCDanUse {
	danUse := &uipb.SCDanUse{}
	danUse.DanList = buildDanList(dans)
	return danUse
}

func BuildSCDanUpgrade(level int32) *uipb.SCDanUpgrade {
	danUpgrade := &uipb.SCDanUpgrade{}
	danUpgrade.Level = &level
	return danUpgrade
}

func BuildSCAlchemyGet(alchemyList map[int]*playerdan.PlayerAlchemyObject) *uipb.SCAlchemyGet {
	alchemyGet := &uipb.SCAlchemyGet{}
	alchemyGet.AlchemyList = buildAlchemyList(alchemyList)
	return alchemyGet
}

func BuildSCAlchemyStart(alchemyList map[int]*playerdan.PlayerAlchemyObject, kindId int) *uipb.SCAlchemyStart {
	alchemyStart := &uipb.SCAlchemyStart{}
	alchemyItem := alchemyList[kindId]
	alchemyStart.AlchemyInfo = buildAlchemy(alchemyItem)
	return alchemyStart
}

func BuildSCAlchemyAccelerate(kindId int32) *uipb.SCAlchemyAccelerate {
	alchemyAccelerate := &uipb.SCAlchemyAccelerate{}
	alchemyAccelerate.KindId = &kindId
	return alchemyAccelerate
}

func BuildSCAlchemyReceive(num int32, kindId int) *uipb.SCAlchemyReceive {
	alchemyReceive := &uipb.SCAlchemyReceive{}
	tKindId := int32(kindId)
	alchemyReceive.KindId = &tKindId
	alchemyReceive.Num = &num
	return alchemyReceive
}

func buildDanList(dans map[int]int32) (danList []*uipb.DanInfo) {
	for id, num := range dans {
		danList = append(danList, buildDan(id, num))
	}
	return danList
}

func buildDan(id int, num int32) *uipb.DanInfo {
	danInfo := &uipb.DanInfo{}
	tId := int32(id)
	danInfo.Id = &tId
	danInfo.Num = &num
	return danInfo
}

func buildAlchemyList(alchemyInfo map[int]*playerdan.PlayerAlchemyObject) (alchemyList []*uipb.AlchemyInfo) {

	for _, alchemyItem := range alchemyInfo {
		alchemyList = append(alchemyList, buildAlchemy(alchemyItem))
	}
	return alchemyList
}

func buildAlchemy(alchemyItem *playerdan.PlayerAlchemyObject) *uipb.AlchemyInfo {
	alchemyInfo := &uipb.AlchemyInfo{}
	kindId := int32(alchemyItem.KindId)
	num := int32(alchemyItem.Num)
	startTime := alchemyItem.StartTime
	state := int32(alchemyItem.State)

	alchemyInfo.KindId = &kindId
	alchemyInfo.Num = &num
	alchemyInfo.StartTime = &startTime
	alchemyInfo.State = &state

	return alchemyInfo
}
