package logic

import (
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	majoreventtypes "fgame/fgame/game/major/event/types"
	majortemplate "fgame/fgame/game/major/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

type MajorSceneData struct {
	*scene.SceneDelegateBase
	s scene.Scene
	//玩家id
	ownerId int64
	//配偶id
	spouseId int64
	//双修副本数据
	currentMajorTemplate majortemplate.MajorTemplate
}

func (sd *MajorSceneData) GetScene() scene.Scene {
	return sd.s
}

//场景开始
func (sd *MajorSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
}

//刷怪
func (sd *MajorSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {

}

//场景心跳
func (sd *MajorSceneData) OnSceneTick(s scene.Scene) {

}

//生物进入
func (sd *MajorSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {

}

func (sd *MajorSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *MajorSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("major:双修副本应该是同一个场景"))
	}
	biology := npc.GetBiologyTemplate()
	if sd.currentMajorTemplate.GetBossId() == int32(biology.TemplateId()) {
		sd.s.Finish(true)
	}
}

//怪物死亡
func (sd *MajorSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}

//玩家重生
func (sd *MajorSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

//玩家死亡
func (sd *MajorSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("major:双修副本应该是同一个场景"))
	}
}
func (sd *MajorSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *MajorSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("major:双修副本应该是同一个场景"))
	}
	pl := p.(player.Player)
	if pl == nil {
		return
	}

	onPushSceneInfo(pl, s.GetStartTime(), sd.ownerId, sd.spouseId, int32(sd.currentMajorTemplate.GetMajorType()), int32(sd.currentMajorTemplate.TemplateId()))
	gameevent.Emit(majoreventtypes.EventTypePlayerEnterMajorScene, p, sd.currentMajorTemplate)
}

//玩家退出
func (sd *MajorSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("major:双修副本应该是同一个场景"))
	}
	//主动退出
	if active {
		playerId := p.GetId()
		if playerId == sd.ownerId {
			sd.ownerId = 0
		}
		if playerId == sd.spouseId {
			sd.spouseId = 0
		}

		//双方都主动
		if sd.ownerId == 0 && sd.spouseId == 0 {
			sd.s.Stop(true, false)
		}
	}
}

//场景完成
func (sd *MajorSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("major:双修副本应该是同一个场景"))
	}
	onMajorFinish(sd, success)
}

//场景退出了
func (sd *MajorSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("major:双修副本应该是同一个场景"))
	}
}

//场景获取物品
func (sd *MajorSceneData) OnScenePlayerGetItem(s scene.Scene, pl scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("major:双修副本应该是同一个场景"))
	}
}

func (sd *MajorSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//玩家获得经验
func (sd *MajorSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("major:双修副本应该是同一个场景"))
	}
}

func CreateMajorSceneData(ownerId int64, spouseId int64, currentMajorTemplate majortemplate.MajorTemplate) *MajorSceneData {
	sd := &MajorSceneData{
		ownerId:              ownerId,
		spouseId:             spouseId,
		currentMajorTemplate: currentMajorTemplate,
	}
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return sd
}
