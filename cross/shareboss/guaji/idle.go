package action

import (
	"fgame/fgame/cross/shareboss/shareboss"
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	worldbosstypes "fgame/fgame/game/worldboss/types"

	log "github.com/Sirupsen/logrus"
)

// 世界boss 一直挂机
func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeCrossWorldBoss, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
}

func (a *idleAction) GuaJi(p scene.Player) {
	log.Infof("玩家查找boss")
	//查找可以攻击的bossId
	worldBossList := shareboss.GetShareBossService().GetGuaiJiShareBossList(worldbosstypes.BossTypeShareBoss, p.GetForce())
	lenOfWorldBossList := len(worldBossList)
	if lenOfWorldBossList <= 0 {
		p.BackLastScene()
		return
	}
	var boss scene.NPC
	for i := lenOfWorldBossList - 1; i >= 0; i-- {
		boss = worldBossList[i]
		if boss.IsDead() {
			continue
		}
		break
	}
	//找不到boss
	if boss == nil {
		log.Infof("玩家没有可以打的boss,退出战斗")
		p.BackLastScene()
		return
	}

	log.Infof("玩家追杀boss[%s]", boss.GetBiologyTemplate().Name)
	p.SetAttackTarget(boss)
	p.GuaJiTrace()
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
