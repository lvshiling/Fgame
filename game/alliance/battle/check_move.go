package battle

import (
	coretypes "fgame/fgame/core/types"
	alliancelogic "fgame/fgame/game/alliance/logic"
	alliancescene "fgame/fgame/game/alliance/scene"
	alliancetemplate "fgame/fgame/game/alliance/template"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterCheckMoveHandler(scenetypes.SceneTypeChengZhan, scene.CheckMoveHandlerFunc(chenZhanCheckMove))
}

// pos：移动的位置
func chenZhanCheckMove(p scene.Player, pos coretypes.Position) (flag bool, fixPos coretypes.Position) {
	// 城门限制区域检查
	flag, fixPos = alliancelogic.ChenZhanCheckMove(p, pos)
	if !flag {
		return
	}

	// 保护罩区域检查
	flag, fixPos = checkProtectNpcArea(p, pos)
	if !flag {
		return
	}

	return
}

func checkProtectNpcArea(p scene.Player, pos coretypes.Position) (flag bool, fixPos coretypes.Position) {
	s := p.GetScene()
	if s == nil {
		flag = true
		return
	}
	sd, ok := s.SceneDelegate().(alliancescene.AllianceSceneData)
	if !ok {
		flag = true
		return
	}

	if sd.IsProtectBroken() {
		flag = true
		return
	}

	warTemp := alliancetemplate.GetAllianceTemplateService().GetWarTemplate()
	sourceFlag := warTemp.IsOnProtectArea(pos)
	if sourceFlag {
		flag = false
		fixPos = warTemp.GetProtectFixPos()
		return
	}

	flag = true
	return
}
