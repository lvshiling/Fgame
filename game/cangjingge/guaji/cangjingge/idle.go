package cangjingge

import (
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/cangjingge/cangjingge"
	playerhuiyuan "fgame/fgame/game/huiyuan/player"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

// boss 一直挂机
func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeCangJingGe, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
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
			}).Info("cangjingge:挂机者不是玩家")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("cangjingge:挂机查找boss")

	//玩家死亡次数超过
	if p.GetGuaJiDeadTimes() >= exitDeadTime {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("cangjingge:挂机超过死亡次数")
		p.ExitGuaJi()
		return
	}

	huiYuanManager := pl.GetPlayerDataManager(types.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	isHuiYuan := huiYuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypePlus)
	tempHuiYuan := huiYuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypeInterim)
	if !(isHuiYuan || tempHuiYuan) {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("cangjingge:不是至尊会员")
		p.ExitGuaJi()
		return
	}

	//查找可以攻击的bossId
	bossList := cangjingge.GetCangJingGeService().GetGuaiJiCangJingGeBossList(pl.GetForce())
	lenOfBossList := len(bossList)
	if lenOfBossList <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("cangjingge:玩家找不到boss")
		p.ExitGuaJi()
		return
	}
	var boss scene.NPC
	for i := lenOfBossList - 1; i >= 0; i-- {
		tempBoss := bossList[i]
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
			}).Info("cangjingge:玩家找不到boss")
		p.ExitGuaJi()
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
			"boss":     boss.GetBiologyTemplate().Name,
		}).Info("cangjingge:玩家追杀boss")
	p.SetAttackTarget(boss)
	p.GuaJiTrace()
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
