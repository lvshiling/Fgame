package transpotation

import (
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	transportationlogic "fgame/fgame/game/transportation/logic"
	"fgame/fgame/game/transportation/transpotation"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeBiaoChe, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
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
			}).Warn("transportation_guaji:镖车挂机,不是用户")
		p.ExitGuaJi()
		return
	}
	//镖车存在
	biaoCheNpc := transpotation.GetTransportService().GetTransportation(pl.GetId())
	if biaoCheNpc != nil {
		return
	}
	//判断是否功能开启
	typ, flag := transportationlogic.CheckIfCanTransportaion(pl)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("transportation_guaji:镖车挂机,没有镖车可以押")
		p.ExitGuaJi()
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
			"typ":      typ.String(),
		}).Info("transportation_guaji:镖车挂机,押镖")
	transportationlogic.HandlePersonalTransportation(pl, typ)

	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
