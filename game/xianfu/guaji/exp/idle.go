package quest

import (
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	xianfulogic "fgame/fgame/game/xianfu/logic"
	xianfutypes "fgame/fgame/game/xianfu/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeXianFuExp, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
}

//仙府银两挂机中
func (a *idleAction) GuaJi(p scene.Player) {
	pl, ok := p.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("xianfu:挂机者不是真实玩家")
		p.ExitGuaJi()
		return
	}

	xianfuType := xianfutypes.XianfuTypeExp
	flag := xianfulogic.CheckIfPlayerCanEnterXianFu(pl, xianfuType)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("xianfu:经验副本进不去")
		p.ExitGuaJi()
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("xianfu:进入经验副本")
	xianfulogic.HandleXianfuChallenge(pl, xianfuType)
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
