package outlangboss

import (
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/outlandboss/outlandboss"
	playeroutlandboss "fgame/fgame/game/outlandboss/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

// 世界boss 一直挂机
func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeOutlandBoss, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
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
			}).Info("outlandboss:挂机者不是玩家")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("outlandboss:挂机查找外域boss")

	//玩家死亡次数超过
	if p.GetGuaJiDeadTimes() >= exitDeadTime {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("outlandboss:挂机超过死亡次数")
		p.ExitGuaJi()
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerOutlandBossDataManagerType).(*playeroutlandboss.PlayerOutlandBossDataManager)
	manager.RefreshZhuoQi()
	if pl.IsZhuoQiLimit() {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("outlandboss:挂机超过浊气值")
		p.ExitGuaJi()
		return
	}

	//查找可以攻击的bossId
	outlandBossList := outlandboss.GetOutlandBossService().GetGuaiJiOutlandBossList(pl.GetForce())
	lenOfOutlandBossList := len(outlandBossList)
	if lenOfOutlandBossList <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("outlandboss:玩家找不到boss")
		p.ExitGuaJi()
		return
	}
	var boss scene.NPC
	for i := lenOfOutlandBossList - 1; i >= 0; i-- {
		tempBoss := outlandBossList[i]
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
			}).Info("outlandboss:玩家找不到boss")
		p.ExitGuaJi()
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
			"boss":     boss.GetBiologyTemplate().Name,
		}).Info("outlandboss:玩家追杀boss")
	p.SetAttackTarget(boss)
	p.GuaJiTrace()
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
