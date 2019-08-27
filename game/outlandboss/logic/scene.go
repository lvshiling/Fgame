package logic

import (
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

type outlandBossSceneData struct {
	*scene.SceneDelegateBase
	s        scene.Scene
	ownerId  int64
	bossId   int32
	itemList []*droptemplate.DropItemData
}

func (sd *outlandBossSceneData) GetScene() scene.Scene {
	return sd.s
}

//场景开始
func (sd *outlandBossSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
}

//刷怪
func (sd *outlandBossSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {
}

//场景心跳
func (sd *outlandBossSceneData) OnSceneTick(s scene.Scene) {
}

//生物进入
func (sd *outlandBossSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {
}
func (sd *outlandBossSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *outlandBossSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("outlandboss:个人BOSS应该是同一个场景"))
	}

	biologyId := int32(npc.GetBiologyTemplate().TemplateId())
	if sd.bossId == biologyId {
		sd.s.Finish(true)
		return
	}
}

func (sd *outlandBossSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *outlandBossSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}
func (sd *outlandBossSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *outlandBossSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("outlandboss:个人BOSS应该是同一个场景"))
	}
	startTime := s.GetStartTime()
	pl := p.(player.Player)
	onPushSceneInfo(pl, startTime, sd.bossId)
}

//玩家重生
func (sd *outlandBossSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

//玩家死亡
func (sd *outlandBossSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("outlandboss:个人BOSS应该是同一个场景"))
	}
}

//玩家退出
func (sd *outlandBossSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("outlandboss:个人BOSS应该是同一个场景"))
	}

	//主动退出 结束副本
	if active {
		sd.s.Stop(true, false)
	}
}

//场景完成
func (sd *outlandBossSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("outlandboss:个人BOSS应该是同一个场景"))
	}
	p := sd.s.GetPlayer(sd.ownerId)
	if p == nil {
		return
	}

	pl, ok := p.(player.Player)
	if !ok {
		return
	}

	//捡起所有东西
	scenelogic.FuBenGetAllItems(p)

	newItemList := droplogic.MergeItemLevel(sd.itemList)
	onMyBossFinish(pl, newItemList, success)
}

//场景退出了
func (sd *outlandBossSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("outlandboss:个人BOSS应该是同一个场景"))
	}
}

//场景获取物品
func (sd *outlandBossSceneData) OnScenePlayerGetItem(s scene.Scene, p scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("outlandboss:个人BOSS应该是同一个场景"))
	}

	sd.itemList = append(sd.itemList, itemData)
}

//玩家获得经验
func (sd *outlandBossSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("outlandboss:个人BOSS应该是同一个场景"))
	}

}

func createMyBossSceneData(ownerId int64, bossId int32) *outlandBossSceneData {
	d := &outlandBossSceneData{
		ownerId: ownerId,
		bossId:  bossId,
	}
	d.SceneDelegateBase = scene.NewSceneDelegateBase()
	return d
}
