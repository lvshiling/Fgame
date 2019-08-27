package scene

import (
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/merge/merge"
	onearenaeventtypes "fgame/fgame/game/onearena/event/types"
	onearenatypes "fgame/fgame/game/onearena/types"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/robot/robot"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

type OneArenaSceneData interface {
	scene.SceneDelegate
	GetOneArenaRobot() scene.Player
}

type oneArenaSceneData struct {
	*scene.SceneDelegateBase
	s                scene.Scene
	pl               player.Player //抢夺者
	ownerId          int64         //灵池拥有者id
	ownerName        string        //灵池拥有者名字
	npc              scene.NPC     //灵池守卫者
	spl              scene.Player  //灵池镜像
	isFinish         bool
	oneArenaTemplate *gametemplate.OneArenaTemplate
}

func CreateOneArenaSceneData(pl player.Player, ownerId int64, ownerName string, oneArenaTemplate *gametemplate.OneArenaTemplate) OneArenaSceneData {
	oasd := &oneArenaSceneData{
		pl:               pl,
		ownerId:          ownerId,
		ownerName:        ownerName,
		oneArenaTemplate: oneArenaTemplate,
	}
	oasd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return oasd
}

func (sd *oneArenaSceneData) GetScene() scene.Scene {
	return sd.s
}

//场景开始
func (sd *oneArenaSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
	pos := sd.oneArenaTemplate.GetPos()
	//刷守卫者
	if sd.ownerId == 0 {
		biologyTemplate := sd.oneArenaTemplate.GetBiologyTemplate()
		sd.npc = scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, int64(0), int32(0), biologyTemplate, pos, 0, 0)
		if sd.npc != nil {
			//设置场景
			sd.s.AddSceneObject(sd.npc)
		}
	} else { //刷镜像
		playerInfo, err := player.GetPlayerService().GetPlayerInfo(sd.ownerId)
		if err != nil {
			return
		}
		robotProperties := make(map[propertytypes.BattlePropertyType]int64)

		for typ, val := range playerInfo.BattleProperty {
			robotProperties[propertytypes.BattlePropertyType(typ)] = val
		}
		power := propertylogic.CulculateAllForce(robotProperties)
		showServerId := merge.GetMergeService().GetMergeTime() != 0
		sd.spl = robot.GetRobotService().CreateOneArenaRobot(playerInfo, power, showServerId)
		sd.spl.SetEnterPos(pos)

		if sd.spl != nil {
			sd.s.AddSceneObject(sd.spl)
		}
	}
}

//刷怪
func (sd *oneArenaSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {

}

//场景心跳
func (sd *oneArenaSceneData) OnSceneTick(s scene.Scene) {

}

//生物进入
func (sd *oneArenaSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {

}

func (sd *oneArenaSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

func (sd *oneArenaSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *oneArenaSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("onearena:灵池争夺应该是同一个场景"))
	}
	sd.s.Finish(true)
}

//怪物死亡
func (sd *oneArenaSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}

//玩家重生
func (sd *oneArenaSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

//玩家死亡
func (sd *oneArenaSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("onearena:灵池争夺应该是同一个场景"))
	}

	if p.GetId() == sd.pl.GetId() {
		//挑战者死亡
		sd.s.Finish(false)
	} else {
		//镜像死亡
		sd.s.Finish(true)
	}
}
func (sd *oneArenaSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *oneArenaSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("onearena:灵池争夺应该是同一个场景"))
	}

}

//玩家退出
func (sd *oneArenaSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("onearena:灵池争夺应该是同一个场景"))
	}
	if !p.IsRobot() {
		if !sd.isFinish {
			sd.endSendEmit(false)
		}
		sd.s.Stop(true, false)
	}
}

//场景完成
func (sd *oneArenaSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("onearena:灵池争夺应该是同一个场景"))
	}
	sd.isFinish = true
	//时间到 血量判断输赢
	if !success && sd.ownerId != 0 {
		hp := sd.pl.GetHP()
		peerHp := sd.peerHp()
		if hp > peerHp {
			success = true
		}
	}
	sd.endSendEmit(success)
}

//场景退出了
func (sd *oneArenaSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("onearena:灵池争夺应该是同一个场景"))
	}
}

//场景获取物品
func (sd *oneArenaSceneData) OnScenePlayerGetItem(s scene.Scene, pl scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("onearena:灵池争夺应该是同一个场景"))
	}
}

//玩家获得经验
func (sd *oneArenaSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("onearena:灵池争夺应该是同一个场景"))
	}
}

func (sd *oneArenaSceneData) peerHp() int64 {
	return sd.spl.GetHP()
}

func (sd *oneArenaSceneData) GetOneArenaRobot() scene.Player {
	return sd.spl
}

func (sd *oneArenaSceneData) endSendEmit(sucess bool) {
	levelType := onearenatypes.OneArenaLevelType(sd.oneArenaTemplate.Level)
	pos := sd.oneArenaTemplate.PosId
	eventData := onearenaeventtypes.CreateOneArenaRobEndEventData(sucess, levelType, pos, sd.ownerName)
	gameevent.Emit(onearenaeventtypes.EventTypeOneArenaRobEnd, sd.pl, eventData)
}
