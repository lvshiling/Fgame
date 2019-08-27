package scene

import (
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	guidereplica "fgame/fgame/game/guidereplica/guidereplica"
	"fgame/fgame/game/guidereplica/pbutil"
	guidereplicatypes "fgame/fgame/game/guidereplica/types"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

func init() {
	guidereplica.RegisterGuideReplicaSdHandler(guidereplicatypes.GuideReplicaTypeMoJian, guidereplica.GuideReplicaSdHandlerFunc(createMoJianSceneData))
}

type moJianSceneData struct {
	*scene.SceneDelegateBase
	s                scene.Scene
	ownerId          int64
	questId          int32
	guidereplicaTemp *gametemplate.GuideReplicaTemplate
	itemList         []*droptemplate.DropItemData
}

func (sd *moJianSceneData) GetScene() scene.Scene {
	return sd.s
}

//场景开始
func (sd *moJianSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
}

//刷怪
func (sd *moJianSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {
}

//生物进入
func (sd *moJianSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {
}
func (sd *moJianSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *moJianSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("mojian:魔剑副本应该是同一个场景"))
	}

	if sd.s.GetNumOfNPC() == 2 {
		//是否只剩下boss了
		moJianTemp := sd.guidereplicaTemp.GetMoJianGuideTemp()
		bossNpcList := sd.s.GetNPCListByBiology(moJianTemp.BossId)
		if len(bossNpcList) == int(1) && npc.GetBiologyTemplate().TemplateId() != int(moJianTemp.BossId) {
			bossNpc := bossNpcList[0]
			if bossNpc.GetBuff(moJianTemp.GetWuDiBuffTemplate().Group) != nil {
				scenelogic.RemoveBuff(bossNpc, moJianTemp.BuffId)
			}
		}
		return
	} else if sd.s.GetNumOfNPC() == 1 {
		//判断是否剩余最后一只npc
		sd.s.Finish(true)
		return
	}
}

func (sd *moJianSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *moJianSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}
func (sd *moJianSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *moJianSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("mojian:魔剑副本应该是同一个场景"))
	}
	startTime := s.GetStartTime()
	pl := p.(player.Player)
	onPushMoJianSceneInfo(pl, startTime, int32(sd.s.MapTemplate().TemplateId()), int32(sd.guidereplicaTemp.GetGuideType()), sd.questId)
}

//玩家重生
func (sd *moJianSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

//玩家死亡
func (sd *moJianSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("mojian:魔剑副本应该是同一个场景"))
	}
}

//玩家退出
func (sd *moJianSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("mojian:魔剑副本应该是同一个场景"))
	}

	//主动退出 结束副本
	sd.s.Stop(true, false)
}

//场景完成
func (sd *moJianSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("mojian:魔剑副本应该是同一个场景"))
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
	onMoJianFinish(pl, newItemList, success)
}

//场景退出了
func (sd *moJianSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("mojian:魔剑副本应该是同一个场景"))
	}
}

//场景获取物品
func (sd *moJianSceneData) OnScenePlayerGetItem(s scene.Scene, p scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("mojian:魔剑副本应该是同一个场景"))
	}

	sd.itemList = append(sd.itemList, itemData)
}

//玩家获得经验
func (sd *moJianSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("mojian:魔剑副本应该是同一个场景"))
	}
}

func createMoJianSceneData(pl player.Player, temp *gametemplate.GuideReplicaTemplate, questId int32) scene.SceneDelegate {
	csd := &moJianSceneData{
		ownerId:          pl.GetId(),
		guidereplicaTemp: temp,
		questId:          questId,
	}
	csd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return csd
}

//挑战结束
func onMoJianFinish(p player.Player, itemList []*droptemplate.DropItemData, isSuccess bool) {
	scMsg := pbutil.BuildSCGuideReplicaChallengeResult(isSuccess, itemList)
	p.SendMsg(scMsg)
}

//下发场景信息
func onPushMoJianSceneInfo(p player.Player, startTime int64, mapId int32, guideType int32, questId int32) {
	scMsg := pbutil.BuildSCGuideReplicaSceneInfo(startTime, mapId, guideType, questId)
	p.SendMsg(scMsg)
}
