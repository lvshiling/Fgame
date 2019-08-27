package tower

import (
	activityguaji "fgame/fgame/game/activity/guaji/guaji"
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	towerlogic "fgame/fgame/game/tower/logic"
	playertower "fgame/fgame/game/tower/player"
	towertempalte "fgame/fgame/game/tower/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeTower, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
}

const (
	defaultFloor = 1
	exitDeadTime = 1
)

//打宝塔副本挂机中
func (a *idleAction) GuaJi(p scene.Player) {
	pl, ok := p.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("tower:挂机者不是真实玩家")
		p.ExitGuaJi()
		return
	}

	if pl.GetGuaJiDeadTimes() >= exitDeadTime {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
				"deadTime": pl.GetGuaJiDeadTimes(),
			}).Warn("tower:挂机者死亡超过次数")
		p.ExitGuaJi()
		return
	}

	activityTemplate := activityguaji.GetGuaJiActivityTemplate(pl)
	if activityTemplate == nil {
		p.ExitGuaJi()
		return
	}

	//检查打宝时间
	towerManager := pl.GetPlayerDataManager(playertypes.PlayerTowerDataManagerType).(*playertower.PlayerTowerDataManager)

	remainTime := towerManager.GetRemainTime()
	if remainTime <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("tower:挂机者没有打宝时间")
		p.ExitGuaJi()
		return
	}

	towerTemplate := towertempalte.GetTowerTemplateService().GetRecommentTower(pl.GetLevel())
	if towerTemplate == nil {
		p.ExitGuaJi()
		return
	}
	floor := int32(towerTemplate.TemplateId())
	if !towerlogic.CheckIfCanEnterTower(pl, floor) {
		if !towerlogic.CheckIfCanEnterTower(pl, defaultFloor) {
			log.WithFields(
				log.Fields{
					"playerId": p.GetId(),
					"floor":    defaultFloor,
				}).Warn("tower:挂机者不能进入层数")
			p.ExitGuaJi()
			return
		}
		towerlogic.HandleEnterTower(pl, defaultFloor)
		return
	}

	towerlogic.HandleEnterTower(pl, floor)
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
