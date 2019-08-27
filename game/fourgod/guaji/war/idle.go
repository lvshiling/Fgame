package war

import (
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/common/common"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeFourGodWar, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
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
	e := scenelogic.FindHatestEnemy(p)

	if e != nil {
		p.SetAttackTarget(e.BattleObject)
		p.GuaJiTrace()
		return
	}

	//查找默认目标
	bo := p.GetDefaultAttackTarget()
	if bo != nil {
		p.SetAttackTarget(bo)
		p.GuaJiTrace()
		return
	}
	//判断是否正在移动
	if p.IsMove() {
		return
	}
	pos := s.MapTemplate().GetMap().RandomPosition()
	flag := p.SetDestPosition(pos)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("fourgod:挂机找不到路")
		return
	}
	return
}

func (a *idleAction) OnExit() {

}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
