package scene

import (
	"fgame/fgame/game/common/common"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/global"
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
	guidereplica.RegisterGuideReplicaSdHandler(guidereplicatypes.GuideReplicaTypeDefault, guidereplica.GuideReplicaSdHandlerFunc(createGuideReplicaSceneData))
}

const (
	guideReplicaAfterFinishTime = 10 * common.SECOND
)

type guideReplicaSceneData struct {
	*scene.SceneDelegateBase
	s                scene.Scene
	ownerId          int64
	questId          int32
	guidereplicaTemp *gametemplate.GuideReplicaTemplate
	itemList         []*droptemplate.DropItemData
}

func (sd *guideReplicaSceneData) GetScene() scene.Scene {
	return sd.s
}

//场景开始
func (sd *guideReplicaSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
}

//刷怪
func (sd *guideReplicaSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {
}

//场景心跳
func (sd *guideReplicaSceneData) OnSceneTick(s scene.Scene) {
	if sd.s.State() == scene.SceneStateFinish {
		now := global.GetGame().GetTimeService().Now()
		elapseFinishTime := now - sd.s.GetFinishTime()
		//完成后 几秒退出
		if elapseFinishTime >= int64(guideReplicaAfterFinishTime) {
			sd.s.Stop(true, false)
			return
		}
	}
}

//生物进入
func (sd *guideReplicaSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {
}
func (sd *guideReplicaSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *guideReplicaSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("guidereplica:引导副本应该是同一个场景"))
	}

	//判断是否剩余最后一只npc
	if sd.s.GetNumOfNPC() == 1 {
		sd.s.Finish(true)
		return
	}
}

func (sd *guideReplicaSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *guideReplicaSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}
func (sd *guideReplicaSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *guideReplicaSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("guidereplica:引导副本应该是同一个场景"))
	}
	startTime := s.GetStartTime()
	pl := p.(player.Player)
	onPushGuideReplicaSceneInfo(pl, startTime, int32(sd.s.MapTemplate().TemplateId()), int32(sd.guidereplicaTemp.GetGuideType()), sd.questId)
}

//玩家重生
func (sd *guideReplicaSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

//玩家死亡
func (sd *guideReplicaSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("guidereplica:引导副本应该是同一个场景"))
	}
}

//玩家退出
func (sd *guideReplicaSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("guidereplica:引导副本应该是同一个场景"))
	}

	//主动退出 结束副本
	if active {
		sd.s.Stop(true, false)
	}
}

//场景完成
func (sd *guideReplicaSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("guidereplica:引导副本应该是同一个场景"))
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
	onGuideReplicaFinish(pl, newItemList, success)
}

//场景退出了
func (sd *guideReplicaSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("guidereplica:引导副本应该是同一个场景"))
	}
}

//场景获取物品
func (sd *guideReplicaSceneData) OnScenePlayerGetItem(s scene.Scene, p scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("guidereplica:引导副本应该是同一个场景"))
	}

	sd.itemList = append(sd.itemList, itemData)
}

//玩家获得经验
func (sd *guideReplicaSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("guidereplica:引导副本应该是同一个场景"))
	}
}

func createGuideReplicaSceneData(pl player.Player, temp *gametemplate.GuideReplicaTemplate, questId int32) scene.SceneDelegate {
	csd := &guideReplicaSceneData{
		ownerId:          pl.GetId(),
		guidereplicaTemp: temp,
		questId:          questId,
	}
	csd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return csd
}

//挑战结束
func onGuideReplicaFinish(p player.Player, itemList []*droptemplate.DropItemData, isSuccess bool) {
	scMsg := pbutil.BuildSCGuideReplicaChallengeResult(isSuccess, itemList)
	p.SendMsg(scMsg)
}

//下发场景信息
func onPushGuideReplicaSceneInfo(p player.Player, startTime int64, mapId int32, guideType int32, questId int32) {
	scMsg := pbutil.BuildSCGuideReplicaSceneInfo(startTime, mapId, guideType, questId)
	p.SendMsg(scMsg)
}
