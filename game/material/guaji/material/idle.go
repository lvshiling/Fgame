package quest

import (
	"fgame/fgame/game/battle/battle"
	materiallogic "fgame/fgame/game/material/logic"
	materialtypes "fgame/fgame/game/material/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeMaterial, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
}

//材料副本挂机中
func (a *idleAction) GuaJi(p scene.Player) {
	pl, ok := p.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("material:挂机者不是真实玩家")
		p.ExitGuaJi()
		return
	}
	for t, _ := range materialtypes.GetMaterialMap() {
		if materiallogic.CheckIfCanEnterMaterial(pl, t) {
			log.WithFields(
				log.Fields{
					"playerId": p.GetId(),
					"typ":      t.String(),
				}).Info("material:挂机者进入材料副本")
			materiallogic.HandlePlayerMaterialChallenge(pl, t)
			return
		}
	}
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Warn("material:挂机者没有可以进入材料副本")
	p.ExitGuaJi()
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
