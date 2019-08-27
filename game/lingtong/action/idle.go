package action

import (
	coretypes "fgame/fgame/core/types"
	coreutils "fgame/fgame/core/utils"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	"math"
)

func init() {
	scene.RegisterLingTongActionFactory(scene.LingTongStateInit, scene.LingTongActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyLingTongAction
}

const (
	maxDistance       = 2.5
	minDistance       = 2
	maxDistanceSquare = maxDistance * maxDistance
	minDistanceSquare = minDistance * minDistance
)

func (a *idleAction) Action(lingTong scene.LingTong) {
	//跟随
	if !scenelogic.CheckIfLingTongAndPlayerSameScene(lingTong) {
		return
	}
	owner := lingTong.GetOwner()
	if owner.GetAttackTarget() != nil {
		lingTong.Trace()
		return
	}

	//判断是不是超过近身距离
	if coreutils.DistanceSquare(owner.GetPosition(), lingTong.GetPosition()) > maxDistanceSquare {
		lingTong.Trace()
		return
	}
	//TODO 随便走
	return
}

func newIdleAction() scene.LingTongAction {
	a := &idleAction{}
	a.DummyLingTongAction = scene.NewDummyLingTongAction()
	return a
}

func getLingTongPosition(owner scene.Player) (pos coretypes.Position) {
	s := owner.GetScene()
	mapTemplate := s.MapTemplate()
	centerPosition := owner.GetPosition()
	angle := -owner.GetAngle()
	pos = coretypes.Position{
		X: centerPosition.X + math.Cos(angle)*minDistance,
		Y: centerPosition.Y,
		Z: centerPosition.Z + math.Sin(angle)*minDistance,
	}
	if mapTemplate.GetMap().IsMask(pos.X, pos.Z) {
		pos.Y = mapTemplate.GetMap().GetHeight(pos.X, pos.Z)
		return
	}
	return centerPosition
}
