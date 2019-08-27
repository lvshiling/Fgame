package action

import (
	activityguaji "fgame/fgame/game/activity/guaji/guaji"
	activitylogic "fgame/fgame/game/activity/logic"
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeActivity, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
}

func (a *idleAction) GuaJi(p scene.Player) {
	pl, ok := p.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("alliance_guaji:活动挂机,不是用户")
		p.ExitGuaJi()
		return
	}
	activityTemplate := activityguaji.GetGuaJiActivityTemplate(pl)
	if activityTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("activity_guaji:活动挂机,找不到活动")
		p.ExitGuaJi()
		return
	}
	activitylogic.HandleActiveAttend(pl, activityTemplate.GetActivityType(), "")
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
