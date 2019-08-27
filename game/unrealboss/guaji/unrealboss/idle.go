package outlangboss

import (
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/unrealboss/unrealboss"

	log "github.com/Sirupsen/logrus"
)

// 世界boss 一直挂机
func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeUnrealBoss, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
}

const (
	exitDeadTime = 1
)

func (a *idleAction) GuaJi(p scene.Player) {
	pl, ok := p.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("unrealboss:挂机者不是玩家")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("unrealboss:挂机查找外域boss")

	//玩家死亡次数超过
	if p.GetGuaJiDeadTimes() >= exitDeadTime {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("unrealboss:挂机超过死亡次数")
		p.ExitGuaJi()
		return
	}

	//判断疲劳值

	//查找可以攻击的bossId
	unrealBossList := unrealboss.GetUnrealBossService().GetGuaiJiUnrealBossList(pl.GetForce())
	lenOfOutlandBossList := len(unrealBossList)
	if lenOfOutlandBossList <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("unrealboss:玩家找不到boss")
		p.ExitGuaJi()
		return
	}
	var boss scene.NPC
	for i := lenOfOutlandBossList - 1; i >= 0; i-- {
		tempBoss := unrealBossList[i]
		if tempBoss.IsDead() {
			continue
		}
		//不够疲劳值
		if !pl.IsEnoughPilao(tempBoss.GetBiologyTemplate().NeedPilao) {
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
			}).Info("unrealboss:玩家找不到boss")
		p.ExitGuaJi()
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
			"boss":     boss.GetBiologyTemplate().Name,
		}).Info("unrealboss:玩家追杀boss")
	p.SetAttackTarget(boss)
	p.GuaJiTrace()
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
