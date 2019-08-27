package tower_scene

import (
	activityguaji "fgame/fgame/game/activity/guaji/guaji"
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	towerlogic "fgame/fgame/game/tower/logic"
	playertower "fgame/fgame/game/tower/player"
	towertypes "fgame/fgame/game/tower/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeTowerScene, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
}

//打宝塔副本挂机中
func (a *idleAction) GuaJi(p scene.Player) {
	pl, ok := p.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("tower:挂机者不是真实玩家")
		p.BackLastScene()
		return
	}

	activityTemplate := activityguaji.GetGuaJiActivityTemplate(pl)
	if activityTemplate == nil {
		p.BackLastScene()
		return
	}
	//检查打宝时间
	towerManager := pl.GetPlayerDataManager(playertypes.PlayerTowerDataManagerType).(*playertower.PlayerTowerDataManager)
	if !pl.IsOnDabao() {
		remainTime := towerManager.GetRemainTime()
		if remainTime <= 0 {
			log.WithFields(
				log.Fields{
					"playerId": p.GetId(),
				}).Warn("tower:挂机者没有打宝时间")
			p.BackLastScene()
			return
		}
		towerlogic.HandleOperateTower(pl, towertypes.TowerOperationTypeBegin)
	}

	//TODO 往合适的楼层爬 ai

	e := scenelogic.FindHatestEnemy(p)

	if e == nil {
		return
	}

	//追击
	pl.SetAttackTarget(e.BattleObject)
	pl.GuaJiTrace()
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
