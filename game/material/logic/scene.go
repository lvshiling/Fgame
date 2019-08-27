package logic

import (
	droptemplate "fgame/fgame/game/drop/template"
	materialtypes "fgame/fgame/game/material/types"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

type MaterialSceneData interface {
	GetMaterialType() materialtypes.MaterialType
}

type materialSceneData struct {
	*scene.SceneDelegateBase
	s                       scene.Scene
	ownerId                 int64
	currentMaterialTempalte *gametemplate.MaterialTemplate
	itemMap                 map[int32]int32
}

func (sd *materialSceneData) GetScene() scene.Scene {
	return sd.s
}

//场景开始
func (sd *materialSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
}

//刷怪
func (sd *materialSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {

}

//场景心跳
func (sd *materialSceneData) OnSceneTick(s scene.Scene) {

}

//生物进入
func (sd *materialSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {

}
func (sd *materialSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *materialSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("material:材料副本应该是同一个场景"))
	}

	//判断是否剩余最后一只npc
	if sd.s.GetNumOfNPC() != 1 {
		return
	}

	//判断是否是最后一波怪
	maxGroupIndex := sd.currentMaterialTempalte.GetMapTemplate().GetNumGroup() - 1
	curGroupIndex := sd.s.GetCurrentGroup()
	if curGroupIndex == maxGroupIndex {
		sd.s.Finish(true)
		return
	} else {
		//刷新怪
		nextGroupIndex := curGroupIndex + 1
		sd.s.RefreshBiology(nextGroupIndex)

		spl := s.GetPlayer(sd.ownerId)
		if spl == nil {
			return
		}
		pl := spl.(player.Player)
		onPushBiologyGroupInfo(pl, nextGroupIndex, sd.currentMaterialTempalte)
	}

}

func (sd *materialSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *materialSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}
func (sd *materialSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *materialSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("material:材料副本应该是同一个场景"))
	}
	onPushSceneInfo(sd.ownerId, sd.currentMaterialTempalte, sd.s)
}

//玩家重生
func (sd *materialSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

//玩家死亡
func (sd *materialSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("material:材料副本应该是同一个场景"))
	}

	sd.s.Finish(false)
}

//玩家退出
func (sd *materialSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("material:材料副本应该是同一个场景"))
	}

	//主动退出 结束副本
	if active {
		sd.s.Stop(true, false)
	}
}

//场景完成
func (sd *materialSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("material:材料副本应该是同一个场景"))
	}
	p := s.GetPlayer(sd.ownerId)
	if p == nil {
		return
	}

	pl, ok := p.(player.Player)
	if !ok {
		return
	}

	//捡起所有东西
	scenelogic.FuBenGetAllItems(p)

	showGroup := sd.s.GetCurrentGroup() + 1
	onMaterialFinish(pl, sd.currentMaterialTempalte, sd.itemMap, success, useTime, showGroup)
}

//场景退出了
func (sd *materialSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("material:材料副本应该是同一个场景"))
	}
}

//场景获取物品
func (sd *materialSceneData) OnScenePlayerGetItem(s scene.Scene, p scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("material:材料副本应该是同一个场景"))
	}
	itemId := itemData.ItemId
	itemNum := itemData.Num

	_, ok := sd.itemMap[itemId]
	if !ok {
		sd.itemMap[itemId] = itemNum
	} else {
		sd.itemMap[itemId] += itemNum
	}

}

//玩家获得经验
func (sd *materialSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("material:材料副本应该是同一个场景"))
	}
}

func (sd *materialSceneData) GetMaterialType() materialtypes.MaterialType {
	return sd.currentMaterialTempalte.GetMaterialType()
}

func createMaterialSceneData(ownerId int64, materialTemplate *gametemplate.MaterialTemplate) *materialSceneData {
	msd := &materialSceneData{
		ownerId:                 ownerId,
		currentMaterialTempalte: materialTemplate,
		itemMap:                 map[int32]int32{},
	}
	msd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return msd
}
