package bagua

import (
	bagualogic "fgame/fgame/game/bagua/logic"
	playerbagua "fgame/fgame/game/bagua/player"
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeBaGua, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
	currentLevel int32
}

const (
	exitDeadTime = 1
)

//八卦副本挂机中
func (a *idleAction) GuaJi(p scene.Player) {
	pl, ok := p.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("bagua:挂机者不是真实玩家")
		p.ExitGuaJi()
		return
	}

	if p.GetGuaJiDeadTimes() >= exitDeadTime {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
				"deaTimes": p.GetGuaJiDeadTimes(),
			}).Info("bagua:死亡次数已经超过")
		p.ExitGuaJi()
		return
	}

	if !bagualogic.CheckPlayerIfCanEnterBaGua(pl) {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("bagua:挂机者不能进入八卦")
		p.ExitGuaJi()
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerBaGuaDataManagerType).(*playerbagua.PlayerBaGuaDataManager)
	currentLevel := manager.GetLevel()
	if a.currentLevel == currentLevel {
		//退出场景
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("bagua:挂机者挑战失败")
		p.ExitGuaJi()
		return
	}

	bagualogic.HandleBaGuaToKill(pl)
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	a.currentLevel = -1
	return a
}
