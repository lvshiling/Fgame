package realm

import (
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	realmlogic "fgame/fgame/game/realm/logic"
	playerrealm "fgame/fgame/game/realm/player"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeTianJieTa, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

const (
	exitDeadTimes = 1
)

type idleAction struct {
	*scene.DummyGuaJiAction
	currentLevel int32
}

//天劫塔副本挂机中
func (a *idleAction) GuaJi(p scene.Player) {
	pl, ok := p.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("realm:挂机者不是真实玩家")
		p.ExitGuaJi()
		return
	}
	if p.GetGuaJiDeadTimes() >= exitDeadTimes {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
				"deadTime": p.GetGuaJiDeadTimes(),
			}).Info("realm:挂机者死亡超过次数")
		p.ExitGuaJi()
		return
	}

	if !realmlogic.CheckIfCanEnterTianJieTa(pl) {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("realm:挂机不能进入天劫塔")
		p.ExitGuaJi()
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	currentLevel := manager.GetTianJieTaLevel()
	if a.currentLevel == currentLevel {
		//退出场景
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("realm:挂机者挑战失败")
		p.ExitGuaJi()
		return
	}
	a.currentLevel = currentLevel
	realmlogic.HandleTianJieTa(pl)
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	a.currentLevel = -1
	return a
}
