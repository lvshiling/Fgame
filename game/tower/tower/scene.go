package tower

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/core/template"
	coretemplate "fgame/fgame/core/template"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	dummytemplate "fgame/fgame/game/dummy/template"
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

type TowerSceneData interface {
	scene.SceneDelegate
	//添加日志
	AddLog(plName, biologyName string, itemId, num int32)
	//获取日志
	GetLogByTime(time int64) []*TowerLogObject
	//清空日志
	GMClearLog()
	//获取打宝塔模板
	GetTowerTemplate() *gametemplate.TowerTemplate
}

const (
	maxLogList = 20
)

// 打宝塔场景数据
type towerSceneData struct {
	*scene.SceneDelegateBase
	s                   scene.Scene
	towerLogList        []*TowerLogObject //打宝日志列表
	lastAddDummyLogLime int64             //上次系统插入日志时间
	hbRunner            heartbeat.HeartbeatTaskRunner
	towerTemplate       *gametemplate.TowerTemplate
}

func (sd *towerSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
	sd.hbRunner.AddTask(CreateTowerDummyLogTask(sd))

	//初始化BOSS
	bossId := sd.towerTemplate.BossId
	if bossId > 0 {
		to := coretemplate.GetTemplateService().Get(int(bossId), (*gametemplate.BiologyTemplate)(nil))
		bossTemplate := to.(*gametemplate.BiologyTemplate)
		bornPos := sd.towerTemplate.GetBossBornPos()
		n := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, int64(0), int32(0), bossTemplate, bornPos, 0, 0)
		if n != nil {
			sd.s.AddSceneObject(n)
		}
	}
}

func (sd *towerSceneData) GetScene() scene.Scene {
	return sd.s
}

func (sd *towerSceneData) OnSceneTick(s scene.Scene) {
	sd.hbRunner.Heartbeat()
}
func (sd *towerSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {}
func (sd *towerSceneData) OnSceneStop(s scene.Scene)                                {}
func (sd *towerSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC)         {}
func (sd *towerSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC)          {}
func (sd *towerSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {

}
func (sd *towerSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC)      {}
func (sd *towerSceneData) OnSceneBiologyAllDead(s scene.Scene)                    {}
func (sd *towerSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player)      {}
func (sd *towerSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player)        {}
func (sd *towerSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {}

func (sd *towerSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("tower:打宝塔应该是同一个场景"))
	}
}

func (sd *towerSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("tower:打宝塔应该是同一个场景"))
	}
}

func (sd *towerSceneData) OnScenePlayerGetItem(s scene.Scene, p scene.Player, itemData *droptemplate.DropItemData) {
}
func (sd *towerSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {}
func (sd *towerSceneData) OnSceneRefreshGroup(s scene.Scene, group int32)               {}

func (sd *towerSceneData) AddLog(plName, biologyName string, itemId, num int32) {
	sd.appendLog(plName, biologyName, itemId, num)
}

func (sd *towerSceneData) GetLogByTime(time int64) []*TowerLogObject {

	for index, log := range sd.towerLogList {
		if time < log.createTime {
			return sd.towerLogList[index:]
		}
	}

	return nil
}

func (sd *towerSceneData) GMClearLog() {
	sd.towerLogList = []*TowerLogObject{}
}

func (sd *towerSceneData) GetTowerTemplate() *gametemplate.TowerTemplate {
	return sd.towerTemplate
}

func (sd *towerSceneData) appendLog(playerName, biologyName string, itemId, itemNum int32) {
	obj := sd.createLogObj(playerName, biologyName, itemId, itemNum)
	sd.towerLogList = append(sd.towerLogList, obj)
}

func (sd *towerSceneData) createLogObj(playerName, biologyName string, itemId, itemNum int32) *TowerLogObject {
	var obj *TowerLogObject
	if len(sd.towerLogList) >= maxLogList {
		obj = sd.towerLogList[0]
		sd.towerLogList = sd.towerLogList[1:]
	} else {
		obj = NewTowerLogObject()
	}

	now := global.GetGame().GetTimeService().Now()
	obj.playerName = playerName
	obj.biologyName = biologyName
	obj.itemId = itemId
	obj.itemNum = itemNum
	obj.createTime = now
	return obj
}

//生成系统假日志 紫色品质以上
func (sd *towerSceneData) addDummyLog() (err error) {
	minSecond := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTowerDummyLogMinTime))
	maxSecond := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTowerDummyLogMaxTime))
	minRandomTime := int(common.SECOND) * minSecond
	maxRandomTime := int(common.SECOND) * maxSecond
	now := global.GetGame().GetTimeService().Now()
	lastTime := sd.lastAddDummyLogLime
	diffTime := now - lastTime
	randTime := int64(mathutils.RandomRange(minRandomTime, maxRandomTime))
	if diffTime < randTime {
		return
	}

	name := dummytemplate.GetDummyTemplateService().GetGameRandomDummyName()
	biologyId, itemId := sd.towerTemplate.GetRandomBiologyDropItemId()
	to := template.GetTemplateService().Get(int(biologyId), (*gametemplate.BiologyTemplate)(nil))
	if to == nil {
		return
	}
	itemTemp := item.GetItemService().GetItem(int(itemId))
	if itemTemp == nil {
		return
	}
	biologyTemp := to.(*gametemplate.BiologyTemplate)
	sd.AddLog(name, biologyTemp.Name, itemId, 1)

	sd.lastAddDummyLogLime = now
	return
}

func CreateTowerSceneData(temp *gametemplate.TowerTemplate) TowerSceneData {
	sd := &towerSceneData{}
	sd.towerTemplate = temp
	sd.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return sd
}
