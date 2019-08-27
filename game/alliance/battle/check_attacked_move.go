package battle

import (
	coretypes "fgame/fgame/core/types"
	alliancescene "fgame/fgame/game/alliance/scene"
	alliancetemplate "fgame/fgame/game/alliance/template"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterCheckAttackedMoveHandler(scenetypes.SceneTypeChengZhan, scene.CheckAttackedMoveHandlerFunc(chenZhanCheckAttackedMove))
}

const (
	numOfDoor = 3
)

// pos:被攻击后的位置
func chenZhanCheckAttackedMove(attackObject scene.BattleObject, defenceObject scene.BattleObject, pos coretypes.Position) bool {
	s := attackObject.GetScene()
	if s == nil {
		return false
	}

	sd, ok := s.SceneDelegate().(alliancescene.AllianceSceneData)
	if !ok {
		return false
	}

	sourcePos := defenceObject.GetPosition()
	// 城门限制区域检查
	if !checkDoorArea(sd, sourcePos, pos) {
		return false
	}

	// 保护罩区域检查
	if !checkAttackedProtectNpcArea(sd, sourcePos, pos) {
		return false
	}

	return true
}

// 要在区域内
func checkDoorArea(sd alliancescene.AllianceSceneData, sourcePos, afterPos coretypes.Position) bool {
	//TODO 优化3个门
	//城门全破了
	currentDoor := sd.GetCurrentDoor()
	if currentDoor >= numOfDoor {
		return true
	}
	sourceArea := alliancetemplate.GetAllianceTemplateService().GetWarTemplate().GetArea(sourcePos)
	destArea := alliancetemplate.GetAllianceTemplateService().GetWarTemplate().GetArea(afterPos)
	//不在限制区域
	if sourceArea < 0 || destArea < 0 {
		return false
	}
	//同一个区域
	if sourceArea == destArea {
		return true
	}

	if sourceArea <= currentDoor && destArea <= currentDoor {
		return true
	}

	return false
}

// 不能在区域内
func checkAttackedProtectNpcArea(sd alliancescene.AllianceSceneData, sourcePos, afterPos coretypes.Position) bool {
	if sd.IsProtectBroken() {
		return true
	}

	sourceFlag := alliancetemplate.GetAllianceTemplateService().GetWarTemplate().IsOnProtectArea(sourcePos)
	afterFlag := alliancetemplate.GetAllianceTemplateService().GetWarTemplate().IsOnProtectArea(afterPos)

	if sourceFlag || afterFlag {
		return false
	}

	return true
}
