package action

import (
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeMoonLove, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

var (
	sleepTime = 3 * common.SECOND
)

type idleAction struct {
	*scene.DummyGuaJiAction
}

func (a *idleAction) GuaJi(p scene.Player) {
	s := p.GetScene()
	if s == nil {
		return
	}

	if p.IsMove() {
		return
	}

	pos := s.MapTemplate().GetMap().RandomPosition()
	flag := p.SetDestPosition(pos)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("moonlove:挂机找不到路")
		return
	}
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
