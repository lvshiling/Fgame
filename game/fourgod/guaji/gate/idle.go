package gate

import (
	"fgame/fgame/game/battle/battle"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeFourGodGate, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction

	portalTemplate *gametemplate.PortalTemplate
	lastTime       int64
}

func (a *idleAction) GuaJi(p scene.Player) {
	s := p.GetScene()
	if s == nil {
		return
	}

	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeFourGodGate {
		if p.IsMove() {
			return
		}
		if a.portalTemplate == nil {
			//随机一个点
			allPortal := scenetemplate.GetSceneTemplateService().GetPortalTemplateMapByMapId(s.MapId())
			randIndex := mathutils.RandomRange(0, len(allPortal))
			i := 0
			for _, portalTemplate := range allPortal {
				if i == randIndex {
					a.portalTemplate = portalTemplate
					return
				}
				i++
			}
		} else {
			scenelogic.MoveToPortal(p, a.portalTemplate)
		}
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
