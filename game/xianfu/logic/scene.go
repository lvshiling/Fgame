package logic

import (
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	xianfutemplate "fgame/fgame/game/xianfu/template"
	xianfutypes "fgame/fgame/game/xianfu/types"
	"fmt"
)

type XianFuSceneData interface {
	GetCurTemplate() xianfutemplate.XianFuTemplate
	GetOwnerId() int64
}

type xianFuSceneData struct {
	*scene.SceneDelegateBase
	s                     scene.Scene
	ownerId               int64
	currentXianFuTempalte xianfutemplate.XianFuTemplate
	resource              int64
	killNum               int32
	protectedNPC          scene.NPC
}

func (sd *xianFuSceneData) GetScene() scene.Scene {
	return sd.s
}

//场景开始
func (sd *xianFuSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
}

//刷怪
func (sd *xianFuSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {
	if currentGroup == 0 {
		switch sd.currentXianFuTempalte.GetXianFuType() {
		case xianfutypes.XianfuTypeExp:
			{
				//查找保护的人
				for _, n := range s.GetAllNPCS() {
					if sd.currentXianFuTempalte.GetBossId() == int32(n.GetBiologyTemplate().TemplateId()) {
						sd.protectedNPC = n
						break
					}
				}
			}
		}
	}

	//设置目标
	for _, n := range s.GetAllNPCS() {
		if sd.protectedNPC != nil && n != sd.protectedNPC {
			n.SetDefaultAttackTarget(sd.protectedNPC)
		}
	}

}

//场景心跳
func (sd *xianFuSceneData) OnSceneTick(s scene.Scene) {

}

//生物进入
func (sd *xianFuSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {

}
func (sd *xianFuSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *xianFuSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("xianfu:仙府应该是同一个场景"))
	}
	biologyId := int32(npc.GetBiologyTemplate().TemplateId())
	xianfuType := sd.currentXianFuTempalte.GetXianFuType()
	switch xianfuType {
	case xianfutypes.XianfuTypeExp:
		{
			if sd.currentXianFuTempalte.GetBossId() == biologyId {
				sd.s.Finish(false)
				return
			}
			//怪物死亡数量
			sd.killNum += 1
			onPushKillNum(sd.ownerId, sd.killNum, s)

			//TODO 特殊处理剩余npc和1只怪
			//判断是否剩余最后一只npc
			if sd.s.GetNumOfNPC() != 2 {
				return
			}

			//判断是否是最后一波怪
			maxGroupIndex := sd.currentXianFuTempalte.GetMapTemplate().GetNumGroup() - 1
			currentGroupIndex := sd.s.GetCurrentGroup()
			if currentGroupIndex == maxGroupIndex {

				sd.s.Finish(true)
				return
			} else {
				//刷新怪
				nextGroupIndex := currentGroupIndex + 1
				sd.s.RefreshBiology(nextGroupIndex)
				onPushBiologyGroupInfo(sd.ownerId, nextGroupIndex, s)
			}
		}
	case xianfutypes.XianfuTypeSilver:
		{
			if sd.s.GetNumOfNPC() == 1 {
				sd.s.Finish(true)
			}
		}
	}

}

func (sd *xianFuSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *xianFuSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}
func (sd *xianFuSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *xianFuSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("xianfu:仙府应该是同一个场景"))
	}
	onPushSceneInfo(sd.ownerId, sd.currentXianFuTempalte, sd.killNum, sd.resource, sd.s)
}

//玩家重生
func (sd *xianFuSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

//玩家死亡
func (sd *xianFuSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("xianfu:仙府应该是同一个场景"))
	}

	sd.s.Finish(false)
}

//玩家退出
func (sd *xianFuSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("xianfu:仙府应该是同一个场景"))
	}

	//主动退出 结束副本
	if active {
		sd.s.Stop(true, false)
	}
}

//场景完成
func (sd *xianFuSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("xianfu:仙府应该是同一个场景"))
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

	onXianFuFinish(pl, sd.currentXianFuTempalte, success, useTime, sd.resource, sd.s.GetCurrentGroup()+1)
}

//场景退出了
func (sd *xianFuSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("xianfu:仙府应该是同一个场景"))
	}
}

//场景获取物品
func (sd *xianFuSceneData) OnScenePlayerGetItem(s scene.Scene, p scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("xianfu:仙府应该是同一个场景"))
	}
	itemId := itemData.ItemId
	itemNum := itemData.Num

	to := item.GetItemService().GetItem(int(itemId))
	typ := to.GetItemType()
	subType := to.GetItemSubType()
	if typ == itemtypes.ItemTypeAutoUseRes && subType == itemtypes.ItemAutoUseResSubTypeSilver {
		sd.resource += int64(itemNum)
	}
	pl := p.(player.Player)
	onPushResourceInfo(pl, sd.resource)

}

//玩家获得经验
func (sd *xianFuSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("xianfu:仙府应该是同一个场景"))
	}

	if sd.currentXianFuTempalte.GetXianFuType() == xianfutypes.XianfuTypeExp {
		pl := p.(player.Player)
		sd.resource += num
		onPushResourceInfo(pl, sd.resource)
	}

}

func (sd *xianFuSceneData) GetCurTemplate() xianfutemplate.XianFuTemplate {
	return sd.currentXianFuTempalte
}
func (sd *xianFuSceneData) GetOwnerId() int64 {
	return sd.ownerId
}

func createXianFuSceneData(ownerId int64, xianFuTemplate xianfutemplate.XianFuTemplate) *xianFuSceneData {
	d := &xianFuSceneData{
		ownerId:               ownerId,
		currentXianFuTempalte: xianFuTemplate,
	}
	d.SceneDelegateBase = scene.NewSceneDelegateBase()
	return d
}
