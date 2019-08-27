package worldboss

import (
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/worldboss/worldboss"

	log "github.com/Sirupsen/logrus"
)

// 世界boss 一直挂机
func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeWorldBoss, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
}

const (
	exitDeadTime = 1
)

func (a *idleAction) GuaJi(p scene.Player) {
	// log.WithFields(
	// 	log.Fields{
	// 		"playerId": p.GetId(),
	// 	}).Info("玩家查找boss")
	//玩家死亡次数超过
	if p.GetGuaJiDeadTimes() >= exitDeadTime {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("worldboss:挂机超过死亡次数")
		p.ExitGuaJi()
		return
	}
	//查找可以攻击的bossId
	worldBossList := worldboss.GetWorldBossService().GetGuaiJiWorldBossList(p.GetForce())
	lenOfWorldBossList := len(worldBossList)
	if lenOfWorldBossList <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("worldboss:玩家找不到boss")
		p.ExitGuaJi()
		return
	}
	var boss scene.NPC
	for i := lenOfWorldBossList - 1; i >= 0; i-- {
		tempBoss := worldBossList[i]
		if tempBoss.IsDead() {
			continue
		}
		boss = tempBoss
		break
	}
	//找不到boss
	if boss == nil {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("worldboss:玩家找不到boss")
		p.ExitGuaJi()
		return
	}
	// log.WithFields(
	// 	log.Fields{
	// 		"playerId": p.GetId(),
	// 		"boss":     boss.GetBiologyTemplate().Name,
	// 	}).Info("worldboss:玩家追杀boss")
	p.SetAttackTarget(boss)
	p.GuaJiTrace()
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
