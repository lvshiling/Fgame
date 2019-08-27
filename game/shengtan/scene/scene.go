package scene

import (
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	shengtaneventtypes "fgame/fgame/game/shengtan/event/types"
	shengtantemplate "fgame/fgame/game/shengtan/template"
)

type ShengTanSceneData interface {
	//排行信息
	//活动奖励
	scene.SceneDelegate
	AddJiuNiang(jiuLiang int32, jiuNiangPercent int32) bool
	GetJiuNiangNum() (int32, int32)
	GetBossHp() (int64, int64)
	GetAllianceId() int64
}

type shengTanSceneData struct {
	*scene.SceneDelegateBase
	s scene.Scene
	//酒量
	currentJiuNiang int32
	//酒酿加成
	currentJiuNiangExpPercent int32
	//仙盟id
	allianceId int64
	//圣坛
	protectedNPC scene.NPC
	//当前组
	currentGroup int32
	//上次刷新时间
	lastRefreshTime int64
}

func (sd *shengTanSceneData) GetScene() scene.Scene {
	return sd.s
}

//场景开始
func (sd *shengTanSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
}

//刷怪
func (sd *shengTanSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {
	now := global.GetGame().GetTimeService().Now()
	sd.lastRefreshTime = now
	sd.currentGroup = currentGroup
	shengTanTemplate := shengtantemplate.GetShengTanTemplateService().GetShengTanTemplate()
	if currentGroup == 0 {
		//查找保护的人
		for _, n := range s.GetAllNPCS() {
			if shengTanTemplate.ShengtanId == int32(n.GetBiologyTemplate().TemplateId()) {
				sd.protectedNPC = n
				break
			}
		}
	}

	//设置目标
	for _, n := range s.GetAllNPCS() {
		if n.GetForeverAttackTarget() != nil {
			continue
		}
		if sd.protectedNPC != nil && n != sd.protectedNPC {
			n.SetForeverAttackTarget(sd.protectedNPC)
		}
	}

}

//场景心跳
func (sd *shengTanSceneData) OnSceneTick(s scene.Scene) {
	//不刷怪
	if sd.protectedNPC == nil || sd.protectedNPC.IsDead() {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	elapse := now - sd.lastRefreshTime
	shengTanTemplate := shengtantemplate.GetShengTanTemplateService().GetShengTanTemplate()
	if sd.currentGroup == 0 {
		//刷新小怪
		if elapse >= int64(shengTanTemplate.ShuaguaiBeginTime) {
			s.RefreshBiology(sd.currentGroup + 1)
		}
	} else {
		//刷新小怪
		if elapse >= int64(shengTanTemplate.XiaoguaiTime) {
			s.RefreshBiology(sd.currentGroup + 1)
		}
	}
}

//生物进入
func (sd *shengTanSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {

}

func (sd *shengTanSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *shengTanSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {

}

func (sd *shengTanSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *shengTanSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}
func (sd *shengTanSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *shengTanSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	//推送场景信息
}

//玩家重生
func (sd *shengTanSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

//玩家死亡
func (sd *shengTanSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {

}

//玩家退出
func (sd *shengTanSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {

}

//场景完成
func (sd *shengTanSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	//推送结算

	//发送场景事件
	gameevent.Emit(shengtaneventtypes.EventTypeShengTanSceneEnd, sd, sd.allianceId)
}

//场景退出了
func (sd *shengTanSceneData) OnSceneStop(s scene.Scene) {

}

//场景获取物品
func (sd *shengTanSceneData) OnScenePlayerGetItem(s scene.Scene, p scene.Player, itemData *droptemplate.DropItemData) {

}

//玩家获得经验
func (sd *shengTanSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {

}

//添加酒量
func (sd *shengTanSceneData) AddJiuNiang(jiuLiang int32, jiuNiangPercent int32) bool {
	sd.currentJiuNiang += jiuLiang
	sd.currentJiuNiangExpPercent += jiuNiangPercent
	gameevent.Emit(shengtaneventtypes.EventTypeShengTanSceneJiuNiangChanged, sd, nil)
	return true
}

//获取当前酒量
func (sd *shengTanSceneData) GetJiuNiangNum() (num int32, percent int32) {
	return sd.currentJiuNiang, sd.currentJiuNiangExpPercent
}

//获取当前酒量
func (sd *shengTanSceneData) GetBossHp() (int64, int64) {
	if sd.protectedNPC == nil {
		return 0, 0
	}
	hp := sd.protectedNPC.GetHP()
	maxHp := sd.protectedNPC.GetMaxHP()
	return hp, maxHp
}

func (sd *shengTanSceneData) GetAllianceId() int64 {
	return sd.allianceId
}

func (sd *shengTanSceneData) GetExpAddPercent(p scene.Player) int32 {
	return sd.currentJiuNiangExpPercent
}

func createShengTanSceneData(allianceId int64) ShengTanSceneData {
	sd := &shengTanSceneData{}
	sd.allianceId = allianceId
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return sd
}

func CreateShengTanSceneData(allianceId int64, mapId int32, endTime int64) ShengTanSceneData {

	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeAllianceShengTan {
		return nil
	}
	sh := createShengTanSceneData(allianceId)

	scene.CreateScene(mapTemplate, endTime, sh)

	return sh
}
