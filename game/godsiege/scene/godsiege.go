package scene

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/core/utils"
	activitytemplate "fgame/fgame/game/activity/template"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	godsiegeeventtypes "fgame/fgame/game/godsiege/event/types"
	godsiegetemplate "fgame/fgame/game/godsiege/template"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

type GodSiegeSceneData interface {
	scene.SceneDelegate
	//玩家人数
	GetScenePlayerNum() int32
	//获取boss
	GetBoss() *godSiegeBoss
	//获取物品
	GetItemMap() map[int64]map[int32]int32
	//获取场景类型
	GetGodType() godsiegetypes.GodSiegeType
	//获取物品
	GetItemMapByPlayer(p scene.Player) map[int32]int32
	//获取采集物列表
	GetCollectNpcList() map[int64]scene.NPC
}

type godSiegeBoss struct {
	npc scene.NPC
	//boss状态
	bossStatus godsiegetypes.GodSiegeBossStatusType
}

func newGodSiegeBoss() *godSiegeBoss {
	d := &godSiegeBoss{
		bossStatus: godsiegetypes.GodSiegeBossStatusTypeInit,
	}
	return d
}

func (l *godSiegeBoss) GetBossStatus() godsiegetypes.GodSiegeBossStatusType {
	return l.bossStatus
}

func (l *godSiegeBoss) GetNpc() scene.NPC {
	return l.npc
}

//神兽攻城战场数据
type godSiegeSceneData struct {
	*scene.SceneDelegateBase
	s scene.Scene
	//类型
	godType godsiegetypes.GodSiegeType
	//神兽攻城活动开始时间
	starTime int64
	//玩家获得的物品
	itemInfoMap map[int64]map[int32]int32
	//当前人数
	num int32
	//boss
	boss *godSiegeBoss
	//金银密窟超级采集物npc
	collectNpcMap map[int64]scene.NPC
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

func CreateGodSiegeSceneData(godType godsiegetypes.GodSiegeType) GodSiegeSceneData {
	csd := &godSiegeSceneData{
		num:           0,
		godType:       godType,
		itemInfoMap:   make(map[int64]map[int32]int32),
		collectNpcMap: make(map[int64]scene.NPC),
	}
	csd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return csd
}

func (sd *godSiegeSceneData) GetScene() (s scene.Scene) {
	return sd.s
}

//场景开始
func (sd *godSiegeSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
	sd.boss = newGodSiegeBoss()
	now := global.GetGame().GetTimeService().Now()
	activityTemplate := sd.GetActivityTemplate(sd.godType)
	activityTimeTemplate, _ := activityTemplate.GetActivityTimeTemplate(now, 0, 0)
	sd.starTime, _ = activityTimeTemplate.GetBeginTime(now)

	//心跳任务
	mapId := sd.GetScene().MapId()
	if sd.godType != godsiegetypes.GodSiegeTypeDenseWat {
		sd.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
		sd.heartbeatRunner.AddTask(CreateGodSiegeTask(sd, sd.starTime, mapId))
	}
}

//刷怪
func (sd *godSiegeSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {

}

//场景心跳
func (sd *godSiegeSceneData) OnSceneTick(s scene.Scene) {
	sd.Heartbeat()
}

//怪物死亡
func (sd *godSiegeSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}

//生物进入
func (sd *godSiegeSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {
	//Boss刷新
	npcType := npc.GetBiologyTemplate().GetBiologyScriptType()
	if npcType == scenetypes.BiologyScriptTypeGodSiegeBoss {
		sd.boss.npc = npc
		sd.boss.bossStatus = godsiegetypes.GodSiegeBossStatusTypeLive
		//发送事件
		gameevent.Emit(godsiegeeventtypes.EventTypeGodSiegeBossStatusRefresh, sd, sd.GetScene().MapId())
	}

	constantTemplate := godsiegetemplate.GetGodSiegeTemplateService().GetConstantTemplate()
	npcBilogyId := int32(npc.GetBiologyTemplate().Id)
	if utils.ContainInt32(constantTemplate.GetDensewatCollectList(), npcBilogyId) {
		sd.collectNpcMap[npc.GetId()] = npc
	}
}

func (sd *godSiegeSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *godSiegeSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("godsiege:神兽攻城战场应该是同一个场景"))
	}
	//Boss死亡
	npcType := npc.GetBiologyTemplate().GetBiologyScriptType()
	if npcType == scenetypes.BiologyScriptTypeGodSiegeBoss {
		sd.boss.bossStatus = godsiegetypes.GodSiegeBossStatusTypeDead
		sd.boss.npc = nil
		//发送事件
		gameevent.Emit(godsiegeeventtypes.EventTypeGodSiegeBossStatusRefresh, sd, sd.GetScene().MapId())
	}
	if npcType == scenetypes.BiologyScriptTypeGeneralCollect {
		_, exist := sd.collectNpcMap[npc.GetId()]
		if !exist {
			//	panic(fmt.Errorf("fourgod:金银密窟超级资源应该存在的"))
			return
		}
		//发送事件
		gameevent.Emit(godsiegeeventtypes.EventTypeGodSiegeCollectNpcChanged, sd, npc)
	}
}

//生物重生
func (sd *godSiegeSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("godsiege:神兽攻城战场应该是同一个场景"))
	}

	if npc.GetBiologyTemplate().GetBiologyScriptType() == scenetypes.BiologyScriptTypeGeneralCollect {
		_, exist := sd.collectNpcMap[npc.GetId()]
		if !exist {
			//	panic(fmt.Errorf("fourgod:金银密窟超级资源应该存在的"))
			return
		}
		//发送事件
		gameevent.Emit(godsiegeeventtypes.EventTypeGodSiegeCollectNpcChanged, sd, npc)
	}
}

//玩家复活
func (sd *godSiegeSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("godsiege:神兽攻城战场应该是同一个场景"))
	}

}

//玩家死亡
func (sd *godSiegeSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("godsiege:神兽攻城战场应该是同一个场景"))
	}
}

func (sd *godSiegeSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *godSiegeSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("godsiege:神兽攻城战场应该是同一个场景"))
	}
	sd.num++
	//发送事件
	gameevent.Emit(godsiegeeventtypes.EventTypeGodSiegePlayerEnter, sd, p)

}

//玩家退出
func (sd *godSiegeSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("godsiege:神兽攻城战场应该是同一个场景"))
	}
	sd.num--
	gameevent.Emit(godsiegeeventtypes.EventTypeGodSiegePlayerExit, sd, nil)
}

//场景完成
func (sd *godSiegeSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("godsiege:神兽攻城战场应该是同一个场景"))
	}
	gameevent.Emit(godsiegeeventtypes.EventTypeGodSiegeSceneFinish, sd, nil)
}

//场景退出了
func (sd *godSiegeSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("godsiege:神兽攻城战场应该是同一个场景"))
	}
}

//场景获取物品
func (sd *godSiegeSceneData) OnScenePlayerGetItem(s scene.Scene, pl scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("godsiege:神兽攻城战场应该是同一个场景"))
	}
	itemId := itemData.ItemId
	itemNum := itemData.Num
	playerId := pl.GetId()
	itemMap, exist := sd.itemInfoMap[playerId]
	if !exist {
		itemMap = make(map[int32]int32)
		sd.itemInfoMap[playerId] = itemMap
	}
	itemMap[itemId] += itemNum

	if sd.godType == godsiegetypes.GodSiegeTypeDenseWat {
		gameevent.Emit(godsiegeeventtypes.EventTypeDenseWatItemChanged, pl, sd)
	}
}

//玩家获得经验
func (sd *godSiegeSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("godsiege:神兽攻城战场应该是同一个场景"))
	}
}

//心跳
func (sd *godSiegeSceneData) Heartbeat() {
	if sd.godType != godsiegetypes.GodSiegeTypeDenseWat {
		sd.heartbeatRunner.Heartbeat()
	}
}

//玩家人数
func (sd *godSiegeSceneData) GetScenePlayerNum() int32 {
	return sd.num
}

//获取boss
func (sd *godSiegeSceneData) GetBoss() *godSiegeBoss {
	return sd.boss
}

//获取物品
func (sd *godSiegeSceneData) GetItemMap() map[int64]map[int32]int32 {
	return sd.itemInfoMap
}

//获取场景类型
func (sd *godSiegeSceneData) GetGodType() godsiegetypes.GodSiegeType {
	return sd.godType
}

//获取物品
func (sd *godSiegeSceneData) GetItemMapByPlayer(p scene.Player) map[int32]int32 {
	return sd.itemInfoMap[p.GetId()]
}

//获取物品
func (sd *godSiegeSceneData) GetCollectNpcList() map[int64]scene.NPC {
	return sd.collectNpcMap
}

//获取活动模板
func (sd *godSiegeSceneData) GetActivityTemplate(godType godsiegetypes.GodSiegeType) (activityTemplate *gametemplate.ActivityTemplate) {
	if !godType.Valid() {
		panic(fmt.Errorf("godsiege:神兽攻城战场godType应该是有效的"))
	}
	activityTypes, _ := godType.GetActivityType()
	activityTemplate = activitytemplate.GetActivityTemplateService().GetActiveByType(activityTypes)
	return
}
