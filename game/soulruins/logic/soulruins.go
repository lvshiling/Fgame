package logic

import (
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/soulruins/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

type SoulRuinsSceneData struct {
	*scene.SceneDelegateBase
	s scene.Scene
	//玩家id
	ownerId int64
	//帝陵遗迹数据
	currentSoulRuinsTemplate *gametemplate.SoulRuinsTemplate
	//场景掉落物品
	dropMap map[int32]int32
	//阶段
	stage types.SoulRuinsStageType
	//事件类型
	eventType types.SoulRuinsEventType
	//事件接受情况 false拒绝 true接受
	eventAccept bool
}

func (sd *SoulRuinsSceneData) GetScene() scene.Scene {
	return sd.s
}

//场景开始
func (sd *SoulRuinsSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
}

//刷怪
func (sd *SoulRuinsSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {

}

//场景心跳
func (sd *SoulRuinsSceneData) OnSceneTick(s scene.Scene) {

}

//生物进入
func (sd *SoulRuinsSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {

}

func (sd *SoulRuinsSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}
func (sd *SoulRuinsSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *SoulRuinsSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("soulruins:帝陵遗迹应该是同一个场景"))
	}

}

//怪物死亡
func (sd *SoulRuinsSceneData) OnSceneBiologyAllDead(s scene.Scene) {
	//杀怪
	if sd.stage == types.SoulRuinsStageTypeKillMonster {
		//TODO: ylz:通过场景获取
		// pl := player.GetOnlinePlayerManager().GetPlayerById(sd.ownerId)
		// if pl == nil {
		// 	return
		// }
		p := s.GetPlayer(sd.ownerId)
		if p == nil {
			return
		}
		pl, ok := p.(player.Player)
		if !ok {
			return
		}
		//触发事件
		eventType := triggerSpecialEvent(pl, sd.currentSoulRuinsTemplate)
		if eventType != types.SoulRuinsEventTypeNot {
			onSpecialEventType(pl, eventType)
			sd.stage = types.SoulRuinsStageTypeEvent
			sd.eventType = eventType
		} else {
			//完成
			sd.s.Finish(true)
		}
	}
	//第二阶段
	if sd.stage == types.SoulRuinsStageTypeSecond {
		sd.stage = types.SoulRuinsStageTypeFinshed
		sd.s.Finish(true)
	}
}

//玩家重生
func (sd *SoulRuinsSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

//玩家死亡
func (sd *SoulRuinsSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("soulruins:帝陵遗迹应该是同一个场景"))
	}
}
func (sd *SoulRuinsSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *SoulRuinsSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("soulruins:帝陵遗迹应该是同一个场景"))
	}
	//TODO: ylz:通过场景获取
	// pl := player.GetOnlinePlayerManager().GetPlayerById(sd.ownerId)
	// if pl == nil {
	// 	return
	// }
	pl, ok := p.(player.Player)
	if !ok {
		return
	}
	//帝魂降临
	if sd.stage == types.SoulRuinsStageTypeSecond &&
		sd.eventType == types.SoulRuinsEventTypeSoul {
		sd.stage = types.SoulRuinsStageTypeFinshed
		s.Finish(true)
	}
	chapter := sd.currentSoulRuinsTemplate.Chapter
	typ := sd.currentSoulRuinsTemplate.Type
	level := sd.currentSoulRuinsTemplate.Level
	onPushSceneInfo(pl, sd.stage, sd.eventType, s.GetStartTime(), chapter, typ, level)
}

//玩家退出
func (sd *SoulRuinsSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("soulruins:帝陵遗迹应该是同一个场景"))
	}
	//主动退出 结束副本
	if active {
		sd.s.Stop(true, false)
	}
}

//场景完成
func (sd *SoulRuinsSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("soulruins:帝陵遗迹应该是同一个场景"))
	}
	//TODO: ylz:通过场景获取
	// p := player.GetOnlinePlayerManager().GetPlayerById(sd.ownerId)
	// if p == nil {
	// 	return
	// }
	sp := s.GetPlayer(sd.ownerId)
	if sp == nil {
		return
	}
	p, ok := sp.(player.Player)
	if !ok {
		return
	}

	sd.stage = types.SoulRuinsStageTypeFinshed
	err := onSoulRuinsFinish(p, sd.currentSoulRuinsTemplate, success, useTime, sd.dropMap)
	if err != nil {
		p.Close(err)
	}
}

//场景退出了
func (sd *SoulRuinsSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("soulruins:帝陵遗迹应该是同一个场景"))
	}
}

//场景获取物品
func (sd *SoulRuinsSceneData) OnScenePlayerGetItem(s scene.Scene, pl scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("soulruins:帝陵遗迹应该是同一个场景"))
	}
	itemId := itemData.ItemId
	itemNum := itemData.Num

	if itemNum <= 0 {
		panic(fmt.Errorf("soulruins: 帝陵遗迹应该是同一个场景 itemNum > 0"))
	}
	_, ok := sd.dropMap[itemId]
	if ok {
		sd.dropMap[itemId] += itemNum
	} else {
		sd.dropMap[itemId] = itemNum
	}
}

//玩家获得经验
func (sd *SoulRuinsSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("soulruins:帝陵遗迹应该是同一个场景"))
	}
}

func CreateSoulRuinsSceneData(ownerId int64, currentSoulRuinsTemplate *gametemplate.SoulRuinsTemplate, stage types.SoulRuinsStageType) *SoulRuinsSceneData {
	sd := &SoulRuinsSceneData{
		ownerId:                  ownerId,
		currentSoulRuinsTemplate: currentSoulRuinsTemplate,
		stage:                    stage,
		eventType:                types.SoulRuinsEventTypeNot,
		eventAccept:              false,
	}
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	sd.dropMap = make(map[int32]int32)
	return sd
}

func (sd *SoulRuinsSceneData) GetSoulRuinsTemplate() *gametemplate.SoulRuinsTemplate {
	return sd.currentSoulRuinsTemplate
}

func (sd *SoulRuinsSceneData) GetEventType() types.SoulRuinsEventType {
	return sd.eventType
}

func (sd *SoulRuinsSceneData) GetStageType() types.SoulRuinsStageType {
	return sd.stage
}

func (sd *SoulRuinsSceneData) GetAccept() bool {
	return sd.eventAccept
}

func (sd *SoulRuinsSceneData) getGroupId() (groupId int32) {
	groupId = 0
	eventType := sd.GetEventType()
	if eventType == types.SoulRuinsEventTypeNot {
		return
	}
	groupId = sd.currentSoulRuinsTemplate.GetGroupIdByEventType(eventType)
	return
}

func (sd *SoulRuinsSceneData) AcceptEvent(accept bool) (flag bool) {
	if sd.stage != types.SoulRuinsStageTypeEvent {
		return
	}
	sd.stage = types.SoulRuinsStageTypeSecond
	sd.eventAccept = accept
	if sd.eventType != types.SoulRuinsEventTypeSoul {
		groupId := sd.getGroupId()
		sd.s.RefreshBiology(groupId)
	}
	flag = true
	return
}
