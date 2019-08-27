package scene

import (
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	shenmoeventtypes "fgame/fgame/game/shenmo/event/types"
	"fmt"
)

//神魔战场场景
func CreateShenMoScene(mapId int32, endTime int64, sh scene.SceneDelegate) (s scene.Scene) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeCrossShenMo {
		return nil
	}
	s = scene.CreateScene(mapTemplate, endTime, sh)
	return s
}

type ShenMoSceneData interface {
	scene.SceneDelegate
	//玩家人数
	GetScenePlayerNum() int32
}

//神魔战场数据
type shenMoSceneData struct {
	*scene.SceneDelegateBase
	s scene.Scene
	//当前人数
	num int32
}

func CreateShenMoSceneData() ShenMoSceneData {
	csd := &shenMoSceneData{}
	csd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return csd
}

func (sd *shenMoSceneData) GetScene() (s scene.Scene) {
	return sd.s
}

//场景开始
func (sd *shenMoSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
}

//刷怪
func (sd *shenMoSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {

}

//场景心跳
func (sd *shenMoSceneData) OnSceneTick(s scene.Scene) {
	sd.Heartbeat()
}

//怪物死亡
func (sd *shenMoSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}

//生物进入
func (sd *shenMoSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {

}

func (sd *shenMoSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *shenMoSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("shenmo:神魔战场应该是同一个场景"))
	}
}

//生物重生
func (sd *shenMoSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("shenmo:神魔战场应该是同一个场景"))
	}
}

//玩家复活
func (sd *shenMoSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("shenmo:神魔战场应该是同一个场景"))
	}

}

//玩家死亡
func (sd *shenMoSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("shenmo:神魔战场应该是同一个场景"))
	}
}
func (sd *shenMoSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *shenMoSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("shenmo:神魔战场应该是同一个场景"))
	}
	if p.GetAllianceId() != 0 {
		p.SwitchPkState(pktypes.PkStateBangPai, pktypes.PkCommonCampDefault)
	} else {
		p.SwitchPkState(pktypes.PkStateAll, pktypes.PkCommonCampDefault)
	}
	sd.num++
	//发送事件
	gameevent.Emit(shenmoeventtypes.EventTypeShenMoPlayerEnter, sd, p)

}

//玩家退出
func (sd *shenMoSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("shenmo:神魔战场应该是同一个场景"))
	}
	sd.num--
	gameevent.Emit(shenmoeventtypes.EventTypeShenMoPlayerExit, sd, nil)
}

//场景完成
func (sd *shenMoSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("shenmo:神魔战场应该是同一个场景"))
	}

	gameevent.Emit(shenmoeventtypes.EventTypeShenMoSceneFinish, sd, nil)
}

//场景退出了
func (sd *shenMoSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("shenmo:神魔战场应该是同一个场景"))
	}
}

//场景获取物品
func (sd *shenMoSceneData) OnScenePlayerGetItem(s scene.Scene, pl scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("shenmo:神魔战场应该是同一个场景"))
	}
}

//玩家获得经验
func (sd *shenMoSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("shenmo:神魔战场应该是同一个场景"))
	}
}

//心跳
func (sd *shenMoSceneData) Heartbeat() {
}

//玩家人数
func (sd *shenMoSceneData) GetScenePlayerNum() int32 {
	return sd.num
}
