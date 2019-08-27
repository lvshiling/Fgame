package logic

import (
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	realmeventtypes "fgame/fgame/game/realm/event/types"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

type TianJieTaSceneData struct {
	*scene.SceneDelegateBase
	s scene.Scene
	//玩家id
	ownerId int64
	//配偶id
	spouseId int64
	//是否结束
	isFinish bool
	//天截塔数据
	currentTianJieTaTemplate *gametemplate.TianJieTaTemplate
	//是否是下一关
	isNextLevel bool
}

func (sd *TianJieTaSceneData) GetScene() scene.Scene {
	return sd.s
}

//场景开始
func (sd *TianJieTaSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
	sd.isNextLevel = false
}

//刷怪
func (sd *TianJieTaSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {

}

//场景心跳
func (sd *TianJieTaSceneData) OnSceneTick(s scene.Scene) {

}

//生物进入
func (sd *TianJieTaSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {

}

func (sd *TianJieTaSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}
func (sd *TianJieTaSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *TianJieTaSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("realm:天截塔应该是同一个场景"))
	}
	biology := npc.GetBiologyTemplate()
	if sd.currentTianJieTaTemplate.BossId == int32(biology.TemplateId()) {
		sd.s.Finish(true)
	}
}

//怪物死亡
func (sd *TianJieTaSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}

//玩家重生
func (sd *TianJieTaSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

//玩家死亡
func (sd *TianJieTaSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("realm:天截塔应该是同一个场景"))
	}
	sd.s.Finish(false)
}

func (sd *TianJieTaSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *TianJieTaSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("realm:天截塔应该是同一个场景"))
	}
	pl := p.(player.Player)
	if pl == nil {
		return
	}
	onPushSceneInfo(pl, s.GetStartTime(), sd.ownerId, sd.spouseId, sd.currentTianJieTaTemplate.Level)
}

//玩家退出
func (sd *TianJieTaSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("realm:天截塔应该是同一个场景"))
	}

	//配偶退出(含掉线和主动退出)
	if p.GetId() == sd.spouseId && !sd.isFinish {
		//主动退出
		if active {
			sd.spouseId = 0
		}
		gameevent.Emit(realmeventtypes.EventTypeRealmPairSpouseExit, p, nil)
		return
	}

	if p.GetId() == sd.ownerId {
		if !active { //闯关者掉线
			//夫妻助战 默认失败
			if sd.spouseId != 0 {
				gameevent.Emit(realmeventtypes.EventTypeRealmPairInviteOffonline, sd.spouseId, p.GetName())
				sd.spouseId = 0
				sd.s.Stop(true, false)
			}
		} else { //主动退出 结束副本
			//打完
			if sd.isFinish && sd.spouseId != 0 {
				//点下一关
				if sd.isNextLevel {
					return
				}
				//点离开副本
				sd.s.Stop(true, false)
			} else {
				sd.s.Stop(true, false)
			}
			sd.spouseId = 0
		}
	}
}

//场景完成
func (sd *TianJieTaSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("realm:天截塔应该是同一个场景"))
	}
	//TODO: ylz：通过场景获取
	// p := player.GetOnlinePlayerManager().GetPlayerById(sd.ownerId)
	// if p == nil {
	// 	return
	// }
	sp := s.GetPlayer(sd.ownerId)
	if sp == nil {
		return
	}
	p, ok := sp.(player.Player)
	if !ok {
		return
	}
	sd.isFinish = true

	err := onTianJieTaFinish(p, sd.spouseId, sd.currentTianJieTaTemplate, success)
	if err != nil {
		p.Close(err)
	}
}

//场景退出了
func (sd *TianJieTaSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("realm:天截塔应该是同一个场景"))
	}
}

//场景获取物品
func (sd *TianJieTaSceneData) OnScenePlayerGetItem(s scene.Scene, pl scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("realm:天截塔应该是同一个场景"))
	}

}

//玩家获得经验
func (sd *TianJieTaSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("realm:天截塔应该是同一个场景"))
	}
}

func (sd *TianJieTaSceneData) OnSetNextLevel() {
	sd.isNextLevel = true
}

func (sd *TianJieTaSceneData) GetOwerId() int64 {
	return sd.ownerId
}

func CreateTienJieTaSceneData(ownerId int64, spouseId int64, currentTianJieTaTemplate *gametemplate.TianJieTaTemplate) *TianJieTaSceneData {
	sd := &TianJieTaSceneData{
		ownerId:                  ownerId,
		spouseId:                 spouseId,
		currentTianJieTaTemplate: currentTianJieTaTemplate,
		isFinish:                 false,
	}
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return sd
}
