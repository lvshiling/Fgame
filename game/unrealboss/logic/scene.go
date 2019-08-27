package logic

import (
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

type unrealBossSceneData struct {
	*scene.SceneDelegateBase
	s        scene.Scene
	ownerId  int64
	bossId   int32
	itemList []*droptemplate.DropItemData
}

func (sd *unrealBossSceneData) GetScene() scene.Scene {
	return sd.s
}

//场景开始
func (sd *unrealBossSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
}

//刷怪
func (sd *unrealBossSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {
}

//场景心跳
func (sd *unrealBossSceneData) OnSceneTick(s scene.Scene) {
}

//生物进入
func (sd *unrealBossSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {
}
func (sd *unrealBossSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *unrealBossSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("unrealboss:个人BOSS应该是同一个场景"))
	}

	biologyId := int32(npc.GetBiologyTemplate().TemplateId())
	if sd.bossId == biologyId {
		sd.s.Finish(true)
		return
	}
}

func (sd *unrealBossSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *unrealBossSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}
func (sd *unrealBossSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *unrealBossSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("unrealboss:个人BOSS应该是同一个场景"))
	}
	startTime := s.GetStartTime()
	pl := p.(player.Player)
	onPushSceneInfo(pl, startTime, sd.bossId)
}

//玩家重生
func (sd *unrealBossSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

//玩家死亡
func (sd *unrealBossSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("unrealboss:个人BOSS应该是同一个场景"))
	}
}

//玩家退出
func (sd *unrealBossSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("unrealboss:个人BOSS应该是同一个场景"))
	}

	//主动退出 结束副本
	if active {
		sd.s.Stop(true, false)
	}
}

//场景完成
func (sd *unrealBossSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("unrealboss:个人BOSS应该是同一个场景"))
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
func (sd *unrealBossSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("unrealboss:个人BOSS应该是同一个场景"))
	}
}

//场景获取物品
func (sd *unrealBossSceneData) OnScenePlayerGetItem(s scene.Scene, p scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("unrealboss:个人BOSS应该是同一个场景"))
	}

	sd.itemList = append(sd.itemList, itemData)
}

//玩家获得经验
func (sd *unrealBossSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("unrealboss:个人BOSS应该是同一个场景"))
	}

}

func createMyBossSceneData(ownerId int64, bossId int32) *unrealBossSceneData {
	d := &unrealBossSceneData{
		ownerId: ownerId,
		bossId:  bossId,
	}
	d.SceneDelegateBase = scene.NewSceneDelegateBase()
	return d
}
