package scene

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	fourgodeventtypes "fgame/fgame/game/fourgod/event/types"
	fourgodtemplate "fgame/fgame/game/fourgod/template"
	"fgame/fgame/game/global"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/merge/merge"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

//四神遗迹主战场场景
func createFourGodWarScene(mapId int32, endTime int64, sh scene.SceneDelegate) (s scene.Scene) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeFourGodWar {
		return nil
	}
	s = scene.CreateScene(mapTemplate, endTime, sh)
	return s
}

//采集宝箱信息
type collectBoxInfo struct {
	playerId        int64 //采集宝箱玩家
	collectStarTime int64 //采集开始时间
	boxId           int64 //宝箱npcid
}

func newCollectBoxInfo(npcId int64) *collectBoxInfo {
	collectBoxInfo := &collectBoxInfo{
		playerId:        0,
		collectStarTime: 0,
		boxId:           npcId,
	}
	return collectBoxInfo
}

func (cb *collectBoxInfo) GetPlayerId() int64 {
	return cb.playerId
}

type FourGodWarSceneData interface {
	FourGodSubSceneData
	//获取所有生物
	GetAllNPCS() map[int64]scene.NPC
	//获取副本生物npc
	GetNpc(curNpcId int64) (curNpc scene.NPC)
	//增加特殊怪
	AddSpecial()
	//获取特殊怪数量
	GetSpecialNum() int32
	//获取特殊怪出生点
	GetSpecialPos() coretypes.Position

	//获取采集宝箱
	GetCollectBox() map[int64]*collectBoxInfo
	//能否采集
	IfCanCollectBox(npcId int64) (*collectBoxInfo, bool)
	//玩家采集宝箱
	CollectBox(npcId int64, playerId int64) bool
	//玩家是否有采集宝箱
	HasCollectBox(playerId int64) (npcId int64, has bool)
	//采集被打断
	CollectBoxInterrupt(npcId int64)
	//采集完成
	FinishCollectBox(npc scene.NPC)
}

//四神遗迹主战场数据
type fourGodWarSceneData struct {
	*scene.SceneDelegateBase
	s scene.Scene
	//开始时间
	starTime int64
	//副本生物状态
	npcMap map[int64]scene.NPC
	//最后一次普通怪被杀的地点
	pos coretypes.Position
	//特殊怪数量
	specialNum int32
	//采集宝箱
	collectBoxMap map[int64]*collectBoxInfo
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

func createFourGodWarSceneData(defaultPos coretypes.Position) FourGodWarSceneData {
	csd := &fourGodWarSceneData{
		npcMap:        make(map[int64]scene.NPC),
		pos:           defaultPos,
		specialNum:    0,
		collectBoxMap: make(map[int64]*collectBoxInfo),
	}
	csd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return csd
}

func (sd *fourGodWarSceneData) GetScene() (s scene.Scene) {
	return sd.s
}

//场景开始
func (sd *fourGodWarSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeFourGod)
	activityTimeTemplate, _ := activityTemplate.GetActivityTimeTemplate(now, openTime, mergeTime)
	sd.starTime, _ = activityTimeTemplate.GetBeginTime(now)
	fourGodTemplate := fourgodtemplate.GetFourGodTemplateService().GetFourGodConstTemplate()
	//心跳任务
	sd.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	sd.heartbeatRunner.AddTask(CreateFourGodWarTask(sd, sd.starTime, fourGodTemplate))
}

//刷怪
func (sd *fourGodWarSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {

}

//场景心跳
func (sd *fourGodWarSceneData) OnSceneTick(s scene.Scene) {
	sd.Heartbeat()
}

//生物进入
func (sd *fourGodWarSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {
	npcType := npc.GetBiologyTemplate().GetBiologyScriptType()
	switch npcType {
	case scenetypes.BiologyScriptTypeGeneralCollect:
		{
			sd.npcMap[npc.GetId()] = npc
			collectBoxInfo := newCollectBoxInfo(npc.GetId())
			sd.collectBoxMap[npc.GetId()] = collectBoxInfo
			break
		}
	case scenetypes.BiologyScriptTypeFourGodBoss:
		{
			sd.npcMap[npc.GetId()] = npc
			break
		}
	}
	//Boss刷新
	if npcType == scenetypes.BiologyScriptTypeFourGodBoss {
		//发送事件
		gameevent.Emit(fourgodeventtypes.EventTypeFourGodBioChange, sd, npc)
	}
}

func (sd *fourGodWarSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *fourGodWarSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("fourgod:四神遗迹主战场应该是同一个场景"))
	}

	npcType := npc.GetBiologyTemplate().GetBiologyScriptType()
	switch npcType {
	//普通怪死亡普通怪死亡
	case scenetypes.BiologyScriptTypeMonster:
		{
			sd.pos = npc.GetPosition()
			break
		}
	//特殊怪死亡
	case scenetypes.BiologyScriptTypeFourGodSpecial:
		{
			sd.specialNum--
			break
		}
	//boss死亡
	case scenetypes.BiologyScriptTypeFourGodBoss:
		{
			//发送事件
			gameevent.Emit(fourgodeventtypes.EventTypeFourGodBioChange, sd, npc)
			break
		}
	}

}

//生物重生
func (sd *fourGodWarSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("fourgod:四神遗迹主战场应该是同一个场景"))
	}
	if npc.GetBiologyTemplate().GetBiologyScriptType() == scenetypes.BiologyScriptTypeGeneralCollect {
		sd.npcMap[npc.GetId()] = npc
		collectBoxInfo, exist := sd.collectBoxMap[npc.GetId()]
		if !exist {
			panic(fmt.Errorf("fourgod:四神遗迹宝箱重生应该存在的"))
		}
		collectBoxInfo.playerId = 0
		collectBoxInfo.collectStarTime = 0
		//发送事件
		gameevent.Emit(fourgodeventtypes.EventTypeFourGodBioChange, sd, npc)
	}
}

//怪物死亡
func (sd *fourGodWarSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}

//玩家复活
func (sd *fourGodWarSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("fourgod:四神遗迹主战场应该是同一个场景"))
	}

}

//玩家死亡
func (sd *fourGodWarSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("fourgod:四神遗迹主战场应该是同一个场景"))
	}
}
func (sd *fourGodWarSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *fourGodWarSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("fourgod:四神遗迹主战场应该是同一个场景"))
	}
	allianceId := p.GetAllianceId()
	if allianceId == 0 {
		p.SwitchPkState(pktypes.PkStateAll, pktypes.PkCommonCampDefault)
	} else {
		p.SwitchPkState(pktypes.PkStateBangPai, pktypes.PkCommonCampDefault)
	}
	//发送事件
	gameevent.Emit(fourgodeventtypes.EventTypeFourGodPlayerEnter, p, sd)

}

//玩家退出
func (sd *fourGodWarSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("fourgod:四神遗迹主战场应该是同一个场景"))
	}
	//清空采集宝箱
	npcId, hasCollect := sd.HasCollectBox(p.GetId())
	if hasCollect {
		sd.CollectBoxInterrupt(npcId)
	}
	if active {
		gameevent.Emit(fourgodeventtypes.EventTypeFourGodPlayerExit, p, sd)
	}
}

//场景完成
func (sd *fourGodWarSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("fourgod:四神遗迹主战场应该是同一个场景"))
	}

	gameevent.Emit(fourgodeventtypes.EventTypeFourGodSceneFinish, sd, nil)
}

//场景退出了
func (sd *fourGodWarSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("fourgod:四神遗迹主战场应该是同一个场景"))
	}
}

//场景获取物品
func (sd *fourGodWarSceneData) OnScenePlayerGetItem(s scene.Scene, pl scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("fourgod:四神遗迹主战场应该是同一个场景"))
	}
	itemId := itemData.ItemId
	itemNum := itemData.Num
	level := itemData.Level
	if itemId == fourgodtemplate.GetFourGodTemplateService().GetFourGodConstTemplate().ItemId {
		return
	}
	to := template.GetTemplateService().Get(int(itemId), (*gametemplate.ItemTemplate)(nil))
	if to == nil {
		return
	}
	itemTemplate := to.(*gametemplate.ItemTemplate)
	if itemTemplate.GetItemSubType() == itemtypes.ItemAutoUseResSubTypeKey {
		return
	}
	//发送事件
	eventData := fourgodeventtypes.CreateFourGodItemGetEventData(itemId, itemNum, level)
	gameevent.Emit(fourgodeventtypes.EventTypeFourGodGetItem, pl, eventData)

}

//玩家获得经验
func (sd *fourGodWarSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("fourgod:四神遗迹主战场应该是同一个场景"))
	}

	// pl := p.(player.Player)
	// if pl != nil {
	// 	//添加经验
	// 	manager := pl.GetPlayerDataManager(types.PlayerFourGodDataManagerType).(*playerfourgod.PlayerFourGodDataManager)
	// 	manager.AddExp(num)
	// }
}

//心跳
func (sd *fourGodWarSceneData) Heartbeat() {
	sd.heartbeatRunner.Heartbeat()
}

//获取副本生物
func (sd *fourGodWarSceneData) GetNpc(curNpcId int64) (curNpc scene.NPC) {
	curNpc, exist := sd.npcMap[curNpcId]
	if !exist {
		return nil
	}
	return
}

//获取副本生物
func (sd *fourGodWarSceneData) GetAllNPCS() map[int64]scene.NPC {
	return sd.npcMap
}

//增加特殊怪
func (sd *fourGodWarSceneData) AddSpecial() {
	sd.specialNum++
}

//获取特殊怪数量
func (sd *fourGodWarSceneData) GetSpecialNum() int32 {
	return sd.specialNum
}

//获取特殊怪出生点
func (sd *fourGodWarSceneData) GetSpecialPos() coretypes.Position {
	return sd.pos
}

//宝箱是否正在采集
func (sd *fourGodWarSceneData) IfCanCollectBox(npcId int64) (*collectBoxInfo, bool) {
	collectBoxInfo, exist := sd.collectBoxMap[npcId]
	if !exist {
		return nil, false
	}
	if collectBoxInfo.playerId != 0 {
		return collectBoxInfo, false
	}
	return collectBoxInfo, true
}

//玩家是否有采集
func (sd *fourGodWarSceneData) HasCollectBox(playerId int64) (npcId int64, has bool) {
	npcId = 0
	has = false
	for _, collectBoxInfo := range sd.collectBoxMap {
		if collectBoxInfo.playerId == playerId {
			has = true
			npcId = collectBoxInfo.boxId
			return
		}
	}
	return
}

//获取采集宝箱
func (sd *fourGodWarSceneData) GetCollectBox() map[int64]*collectBoxInfo {
	return sd.collectBoxMap
}

//玩家开宝箱
func (sd *fourGodWarSceneData) CollectBox(npcId int64, playerId int64) bool {
	now := global.GetGame().GetTimeService().Now()
	collectBoxInfo, flag := sd.IfCanCollectBox(npcId)
	if !flag {
		return false
	}
	collectBoxInfo.playerId = playerId
	collectBoxInfo.collectStarTime = now
	return true
}

//采集被打断
func (sd *fourGodWarSceneData) CollectBoxInterrupt(npcId int64) {
	collectBoxInfo, exist := sd.collectBoxMap[npcId]
	if !exist {
		return
	}
	collectBoxInfo.playerId = 0
	collectBoxInfo.collectStarTime = 0
}

//采集完成
func (sd *fourGodWarSceneData) FinishCollectBox(npc scene.NPC) {
	npcId := npc.GetId()
	collectBoxInfo, exist := sd.collectBoxMap[npcId]
	if !exist {
		return
	}
	// curHp := npc.GetHP()
	// flag := npc.CostHP(curHp, collectBoxInfo.playerId)
	flag := npc.Recycle(collectBoxInfo.playerId)
	if !flag {
		panic(fmt.Errorf("fourgod:  CostHP should be ok"))
	}
	collectBoxInfo.playerId = 0
	collectBoxInfo.collectStarTime = 0
}
