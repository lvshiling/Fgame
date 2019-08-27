package godsiege

import (
	coretypes "fgame/fgame/core/types"
	gameevent "fgame/fgame/game/event"
	godsiegeeventtypes "fgame/fgame/game/godsiege/event/types"
	godsiegescene "fgame/fgame/game/godsiege/scene"
	godsiegetemplate "fgame/fgame/game/godsiege/template"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

type godSiegeData struct {
	godSiegeSceneData godsiegescene.GodSiegeSceneData
	//参加记录
	attendMap map[int64]struct{}
	//排队记录
	lineList []int64
	//获取复活统计
	bornMap map[godsiegetypes.GodSiegePosType]int32
	//场景人数
	num int32
	//总人数
	totalNum int32
	//地图id
	mapId int32
	//类型
	godType godsiegetypes.GodSiegeType
}

func (gd *godSiegeData) init() {
	gd.attendMap = make(map[int64]struct{})
	gd.lineList = make([]int64, 0, 8)
	gd.bornMap = make(map[godsiegetypes.GodSiegePosType]int32)
	gd.num = 0
	gd.mapId = 0
	gd.godType = 0
	gd.godSiegeSceneData = nil
}

func createGodSiegeData(mapId int32, godType godsiegetypes.GodSiegeType) (data *godSiegeData) {
	if !godType.Valid() {
		panic(fmt.Errorf("godsiege:神兽攻城godType应该是有效的"))
	}

	data = &godSiegeData{}
	data.init()
	constantTmeplate := godsiegetemplate.GetGodSiegeTemplateService().GetConstantTemplate()
	data.godType = godType
	data.mapId = constantTmeplate.GetMapId(godType)
	if godType != godsiegetypes.GodSiegeTypeDenseWat {
		data.totalNum = constantTmeplate.PlayerLimitCount
	} else {
		data.totalNum = constantTmeplate.MoneyLimitCount
	}

	return
}

func (gd *godSiegeData) godSiegeSceneFinish() {
	gd.init()
}

func (gd *godSiegeData) attend(playerId int64) (lineUpPos int32, lineUpFlag bool) {
	lineLen := int32(len(gd.lineList))
	if gd.num >= gd.totalNum {
		lineUpFlag = true
		gd.lineList = append(gd.lineList, playerId)
		lineUpPos = lineLen
		return
	}
	gd.attendMap[playerId] = struct{}{}
	gd.num++
	return 0, false
}

func (gd *godSiegeData) getHasLineUp(playerId int64) (lineUpPos int32, flag bool) {
	for index, curPlayerId := range gd.lineList {
		if playerId == curPlayerId {
			lineUpPos = int32(index)
			flag = true
			break
		}
	}
	return
}

func (gd *godSiegeData) createGodSiegeScene(mapId int32, endTime int64, godType godsiegetypes.GodSiegeType) (se scene.Scene) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeCrossGodSiege &&
		mapTemplate.GetMapType() != scenetypes.SceneTypeCrossDenseWat {
		return nil
	}

	gd.godSiegeSceneData = godsiegescene.CreateGodSiegeSceneData(godType)
	se = scene.CreateScene(mapTemplate, endTime, gd.godSiegeSceneData)
	return
}

func (gd *godSiegeData) getGodSiegeScene() (se scene.Scene) {
	if gd.godSiegeSceneData == nil {
		return
	}
	se = gd.godSiegeSceneData.GetScene()
	return
}

func (gd *godSiegeData) getRebornPos(playerId int64) (pos coretypes.Position, flag bool) {
	_, exist := gd.attendMap[playerId]
	if !exist {
		return
	}

	initPos := godsiegetypes.GodSiegePosTypeMin
	initPosCount := gd.bornMap[godsiegetypes.GodSiegePosTypeMin]
	for i := godsiegetypes.GodSiegePosTypeMin; i <= godsiegetypes.GodSiegePosTypeMax; i++ {
		curPosCount := gd.bornMap[i]
		if curPosCount >= initPosCount {
			continue
		}
		initPos = i
		initPosCount = curPosCount
	}
	gd.bornMap[initPos] += 1

	posTemplate := godsiegetemplate.GetGodSiegeTemplateService().GetPlayerPosTemplate(gd.mapId, initPos)
	if posTemplate != nil {
		pos = posTemplate.GetPos()
		flag = true
	}
	return
}

func (gd *godSiegeData) cancleLineUp(playerId int64) (flag bool) {

	index, flag := gd.getHasLineUp(playerId)
	if !flag {
		return
	}
	if index == 0 && len(gd.lineList) <= 1 {
		gd.lineList = make([]int64, 0, 8)
		return true
	}
	gd.lineList = append(gd.lineList[:index], gd.lineList[index+1:]...)

	eventData := godsiegeeventtypes.CreateGodSiegeCancleLineUpEventData(int32(index), gd.godType)
	gameevent.Emit(godsiegeeventtypes.EventTypeGodSiegeCancleLineUp, gd.lineList, eventData)
	flag = true
	return
}

func (gd *godSiegeData) syncSceneNum(scenePlayerNum int32) {
	if scenePlayerNum < 0 || scenePlayerNum > gd.totalNum {
		return
	}
	gd.num = scenePlayerNum
}

func (gd *godSiegeData) getAllLineUpList() (lineUpList []int64) {
	return gd.lineList
}

func (gd *godSiegeData) isGodSiegeActivityTime() (flag bool) {
	return gd.godSiegeSceneData != nil
}

func (gd *godSiegeData) removeFirstLineUpPlayer(scenePlayerNum int32) {
	if scenePlayerNum < 0 || scenePlayerNum > gd.totalNum {
		return
	}

	playerId := int64(0)
	gd.num = scenePlayerNum
	lineLen := int32(len(gd.lineList))
	if gd.num < gd.totalNum && lineLen != 0 {
		playerId = gd.lineList[0]
		if lineLen == 1 {
			gd.lineList = make([]int64, 0, 8)
		} else {
			gd.lineList = gd.lineList[1:]
		}
	}

	if playerId != 0 {
		gd.attendMap[playerId] = struct{}{}
		eventData := godsiegeeventtypes.CreateGodSiegeFinishLineUpEventData(playerId, gd.godType)
		gameevent.Emit(godsiegeeventtypes.EventTypeGodSiegePlayerLineUpFinish, gd.lineList, eventData)
	}
	return
}
