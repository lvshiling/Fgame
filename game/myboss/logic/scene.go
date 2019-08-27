package logic

import (
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

type myBossSceneData struct {
	*scene.SceneDelegateBase
	s        scene.Scene
	ownerId  int64
	bossId   int32
	itemList []*droptemplate.DropItemData
}

func (sd *myBossSceneData) GetScene() scene.Scene {
	return sd.s
}

//场景开始
func (sd *myBossSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
}

//刷怪
func (sd *myBossSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {
}

//场景心跳
func (sd *myBossSceneData) OnSceneTick(s scene.Scene) {
}

//生物进入
func (sd *myBossSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {
}
func (sd *myBossSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *myBossSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("myboss:个人BOSS应该是同一个场景"))
	}

	biologyId := int32(npc.GetBiologyTemplate().TemplateId())
	if sd.bossId == biologyId {
		sd.s.Finish(true)
		return
	}
}

func (sd *myBossSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *myBossSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}
func (sd *myBossSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *myBossSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("myboss:个人BOSS应该是同一个场景"))
	}
	startTime := s.GetStartTime()
	pl := p.(player.Player)
	onPushSceneInfo(pl, startTime, sd.bossId)
}

//玩家重生
func (sd *myBossSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

//玩家死亡
func (sd *myBossSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("myboss:个人BOSS应该是同一个场景"))
	}
}

//玩家退出
func (sd *myBossSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("myboss:个人BOSS应该是同一个场景"))
	}

	//主动退出 结束副本
	if active {
		sd.s.Stop(true, false)
	}
}

//场景完成
func (sd *myBossSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("myboss:个人BOSS应该是同一个场景"))
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
func (sd *myBossSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("myboss:个人BOSS应该是同一个场景"))
	}
}

//场景获取物品
func (sd *myBossSceneData) OnScenePlayerGetItem(s scene.Scene, p scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("myboss:个人BOSS应该是同一个场景"))
	}

	sd.itemList = append(sd.itemList, itemData)
}

//玩家获得经验
func (sd *myBossSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("myboss:个人BOSS应该是同一个场景"))
	}

}

func createMyBossSceneData(ownerId int64, bossId int32) *myBossSceneData {
	sd := &myBossSceneData{
		ownerId: ownerId,
		bossId:  bossId,
	}
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return sd
}
