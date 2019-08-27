package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertytypes "fgame/fgame/game/property/types"
	playersoulruins "fgame/fgame/game/soulruins/player"
)

func BuildSCSoulRuinsGet(pl player.Player) *uipb.SCSoulRuinsGet {
	soulruinsGet := &uipb.SCSoulRuinsGet{}
	manager := pl.GetPlayerDataManager(types.PlayerSoulRuinsDataManagerType).(*playersoulruins.PlayerSoulRuinsDataManager)
	numObj := manager.GetSoulRuinsNum()
	ruinsObjMap := manager.GetSoulRuinsMap()
	rewObjMap := manager.GetSoulRuinsRewChapterMap()
	soulruinsGet.NumInfo = buildNum(numObj)

	//soulruinsGet.RuinsList
	for _, ruinsTypeMap := range ruinsObjMap {
		for _, levelMap := range ruinsTypeMap {
			for _, ruinsObj := range levelMap {
				soulruinsGet.RuinsList = append(soulruinsGet.RuinsList, buildSoulRuins(ruinsObj))
			}
		}
	}

	//soulruinsGet.RewList
	for _, rewTypeMap := range rewObjMap {
		for _, rewObj := range rewTypeMap {
			chapter := rewObj.Chapter
			typ := int32(rewObj.Type)
			soulruinsGet.RewList = append(soulruinsGet.RewList, buildChapter(chapter, typ))
		}
	}
	return soulruinsGet
}

func BuildSCSoulRuinsChallenge(numObj *playersoulruins.PlayerSoulRuinsNumObject) *uipb.SCSoulRuinsChallenge {
	soulRuinsChallenge := &uipb.SCSoulRuinsChallenge{}
	soulRuinsChallenge.NumInfo = buildNum(numObj)
	return soulRuinsChallenge
}

func BuildSCSoulRuinsRewReceive(rewObj *playersoulruins.PlayerSoulRuinsRewChapterObject) *uipb.SCSoulRuinsRewReceive {
	soulRuinsRewReceive := &uipb.SCSoulRuinsRewReceive{}
	chapter := rewObj.Chapter
	typ := int32(rewObj.Type)
	soulRuinsRewReceive.ChapterInfo = buildChapter(chapter, typ)
	return soulRuinsRewReceive
}

func BuildSCSoulRuinsSweep(numObj *playersoulruins.PlayerSoulRuinsNumObject, rd *propertytypes.RewData, dropItemList [][]*droptemplate.DropItemData) *uipb.SCSoulRuinsSweep {
	soulRuinsSweep := &uipb.SCSoulRuinsSweep{}
	soulRuinsSweep.NumInfo = buildNum(numObj)
	soulRuinsSweep.RewInfo = buildRewProperty(rd)

	for i := int(0); i < len(dropItemList); i++ {
		soulRuinsSweep.SweepList = append(soulRuinsSweep.SweepList, buildSweepDrop(dropItemList[i]))
	}
	return soulRuinsSweep
}

func BuildSCSoulRuinsEvent(eventType int32) *uipb.SCSoulRuinsEvent {
	soulRuinsEvent := &uipb.SCSoulRuinsEvent{}
	soulRuinsEvent.EventType = &eventType
	return soulRuinsEvent
}

func BuildSCSoulRuinsDealEvent(eventType int32, accept bool) *uipb.SCSoulRuinsDealEvent {
	soulRuinsDealEvent := &uipb.SCSoulRuinsDealEvent{}
	soulRuinsDealEvent.EventType = &eventType
	soulRuinsDealEvent.Accept = &accept
	return soulRuinsDealEvent
}

func BuildSCSoulRuinsScene(state int32, eventType int32, starTime int64, chapter int32, typ int32, level int32) *uipb.SCSoulRuinsSceneInfo {
	soulRuinsScene := &uipb.SCSoulRuinsSceneInfo{}
	soulRuinsScene.State = &state
	soulRuinsScene.EventType = &eventType
	soulRuinsScene.CreateTime = &starTime
	soulRuinsScene.ChapterInfo = buildChapter(chapter, typ)
	soulRuinsScene.Level = &level
	return soulRuinsScene
}

func BuildSCSoulRuinsResult(chapter int32, typ int32, level int32, usedTime int32, dropMap map[int32]int32, sucess bool) *uipb.SCSoulRuinsResult {
	soulRuinsResult := &uipb.SCSoulRuinsResult{}
	soulRuinsResult.Result = buildResult(chapter, typ, level, usedTime, dropMap, sucess)
	return soulRuinsResult
}

func BuildSCSoulRuinsFirstPass(numObj *playersoulruins.PlayerSoulRuinsNumObject, chapter int32, typ int32, level int32, usedTime int32, dropMap map[int32]int32, sucess bool) *uipb.SCSoulRuinsFirstPass {
	soulRuinsFirstPass := &uipb.SCSoulRuinsFirstPass{}
	soulRuinsFirstPass.NumInfo = buildNum(numObj)
	soulRuinsFirstPass.Result = buildResult(chapter, typ, level, usedTime, dropMap, sucess)
	return soulRuinsFirstPass
}

func BuildSCSoulRuinsBuyNum(numObj *playersoulruins.PlayerSoulRuinsNumObject) *uipb.SCSoulRuinsBuyNum {
	soulRuinsBuyNum := &uipb.SCSoulRuinsBuyNum{}
	soulRuinsBuyNum.NumInfo = buildNum(numObj)
	return soulRuinsBuyNum
}

func buildNum(numObj *playersoulruins.PlayerSoulRuinsNumObject) *uipb.SoulRuinsNum {
	numInfo := &uipb.SoulRuinsNum{}
	num := numObj.Num
	buyNum := numObj.BuyNum
	extraNum := numObj.RewNum + numObj.ExtraBuyNum
	numInfo.Num = &num
	numInfo.BuyNum = &buyNum
	numInfo.ExtraNum = &extraNum
	return numInfo
}

func buildSoulRuins(ruinsObj *playersoulruins.PlayerSoulRuinsObject) *uipb.SoulRuins {
	ruinsInfo := &uipb.SoulRuins{}
	chapter := ruinsObj.Chapter
	typ := int32(ruinsObj.Type)
	level := ruinsObj.Level
	star := ruinsObj.Star
	ruinsInfo.ChapterInfo = buildChapter(chapter, typ)
	ruinsInfo.Level = &level
	ruinsInfo.Star = &star
	return ruinsInfo
}

func buildResult(chapter int32, typ int32, level int32, usedTime int32, dropMap map[int32]int32, sucess bool) *uipb.SoulRuinsResult {
	soulRuinsResult := &uipb.SoulRuinsResult{}
	soulRuinsResult.ChapterInfo = buildChapter(chapter, typ)
	soulRuinsResult.Level = &level
	soulRuinsResult.UsedTime = &usedTime
	soulRuinsResult.Sucess = &sucess
	for itemId, num := range dropMap {
		soulRuinsResult.DropList = append(soulRuinsResult.DropList, buildDrop(itemId, num, 0))
	}
	return soulRuinsResult
}

func buildChapter(chapter int32, typ int32) *uipb.SoulRuinsChapter {
	chapterInfo := &uipb.SoulRuinsChapter{}
	chapterInfo.Chapter = &chapter
	chapterInfo.Typ = &typ
	return chapterInfo
}

func buildRewProperty(rd *propertytypes.RewData) *uipb.RewProperty {
	rewProperty := &uipb.RewProperty{}
	rewExp := rd.GetRewExp()
	rewExpPoint := rd.GetRewExpPoint()
	rewGold := rd.GetRewGold()
	rewBindGold := rd.GetRewBindGold()
	rewSilver := rd.GetRewSilver()

	rewProperty.Exp = &rewExp
	rewProperty.ExpPoint = &rewExpPoint
	rewProperty.Silver = &rewSilver
	rewProperty.Gold = &rewGold
	rewProperty.BindGold = &rewBindGold

	return rewProperty
}

func buildSweepDrop(dropItemMap []*droptemplate.DropItemData) *uipb.SoulRuinsSweepDrop {
	sweepDrp := &uipb.SoulRuinsSweepDrop{}
	for _, itemData := range dropItemMap {
		itemId := itemData.ItemId
		num := itemData.Num
		level := itemData.Level
		sweepDrp.DropList = append(sweepDrp.DropList, buildDrop(itemId, num, level))
	}
	return sweepDrp
}

func buildDrop(itemId int32, num int32, level int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	dropInfo.Level = &level
	return dropInfo
}
