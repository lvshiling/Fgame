package logic

import (
	baguaeventtypes "fgame/fgame/game/bagua/event/types"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

type BaGuaSceneData struct {
	*scene.SceneDelegateBase
	s scene.Scene
	//玩家id
	ownerId int64
	//配偶id
	spouseId int64
	//是否结束
	isFinish bool
	//天截塔数据
	curTemplate *gametemplate.BaGuaMiJingTemplate
	//是否是下一关
	isNextLevel bool
}

func (sd *BaGuaSceneData) GetScene() scene.Scene {
	return sd.s
}

//场景开始
func (sd *BaGuaSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
	sd.isNextLevel = false
}

//刷怪
func (sd *BaGuaSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {

}

//场景心跳
func (sd *BaGuaSceneData) OnSceneTick(s scene.Scene) {

}

//生物进入
func (sd *BaGuaSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {

}

func (sd *BaGuaSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}
func (sd *BaGuaSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *BaGuaSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("bagua:八卦秘境应该是同一个场景"))
	}
	biology := npc.GetBiologyTemplate()
	if sd.curTemplate.BossId == int32(biology.TemplateId()) {
		sd.s.Finish(true)
	}
}

//怪物死亡
func (sd *BaGuaSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}

//玩家重生
func (sd *BaGuaSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

//玩家死亡
func (sd *BaGuaSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("bagua:八卦秘境应该是同一个场景"))
	}
	sd.s.Finish(false)
}

func (sd *BaGuaSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *BaGuaSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("bagua:八卦秘境应该是同一个场景"))
	}
	pl := p.(player.Player)
	if pl == nil {
		return
	}
	onPushSceneInfo(pl, s.GetStartTime(), sd.ownerId, sd.spouseId, sd.curTemplate.Level)
}

//玩家退出
func (sd *BaGuaSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("bagua:八卦秘境应该是同一个场景"))
	}

	//配偶退出(含掉线和主动退出)
	if p.GetId() == sd.spouseId && !sd.isFinish {
		//主动退出
		if active {
			sd.spouseId = 0
		}
		gameevent.Emit(baguaeventtypes.EventTypeBaGuaPairSpouseExit, p, nil)
		return
	}

	if p.GetId() == sd.ownerId {
		if !active { //闯关者掉线
			//夫妻助战 默认失败
			if sd.spouseId != 0 {
				gameevent.Emit(baguaeventtypes.EventTypeBaGuaPairInviteOffonline, sd.spouseId, p.GetName())
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
func (sd *BaGuaSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("bagua:八卦秘境应该是同一个场景"))
	}

	sp := s.GetPlayer(sd.ownerId)
	if sp == nil {
		return
	}
	p, ok := sp.(player.Player)
	if !ok {
		return
	}
	sd.isFinish = true

	err := onBaGuaFinish(p, sd.spouseId, sd.curTemplate, success)
	if err != nil {
		p.Close(err)
	}
}

//场景退出了
func (sd *BaGuaSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("bagua:八卦秘境应该是同一个场景"))
	}
}

//场景获取物品
func (sd *BaGuaSceneData) OnScenePlayerGetItem(s scene.Scene, pl scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("bagua:八卦秘境应该是同一个场景"))
	}

}

//玩家获得经验
func (sd *BaGuaSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("bagua:八卦秘境应该是同一个场景"))
	}
}

func (sd *BaGuaSceneData) OnSetNextLevel() {
	sd.isNextLevel = true
}

func (sd *BaGuaSceneData) GetOwerId() int64 {
	return sd.ownerId
}

func CreateBaGuaSceneData(ownerId int64, spouseId int64, currentTemplate *gametemplate.BaGuaMiJingTemplate) *BaGuaSceneData {
	sd := &BaGuaSceneData{
		ownerId:     ownerId,
		spouseId:    spouseId,
		curTemplate: currentTemplate,
		isFinish:    false,
	}
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return sd
}
