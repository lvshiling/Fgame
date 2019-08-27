package battle

import (
	coretypes "fgame/fgame/core/types"
	chuangshiscene "fgame/fgame/cross/chuangshi/scene"
	chuangshitemplate "fgame/fgame/game/chuangshi/template"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterCheckAttackedMoveHandler(scenetypes.SceneTypeChuangShiZhiZhanFuShu, scene.CheckAttackedMoveHandlerFunc(chenZhanCheckAttackedMove))
}

const (
	numOfDoor = 1
)

// pos:被攻击后的位置
func chenZhanCheckAttackedMove(attackObject scene.BattleObject, defenceObject scene.BattleObject, pos coretypes.Position) bool {
	s := attackObject.GetScene()
	if s == nil {
		return false
	}

	sd, ok := s.SceneDelegate().(chuangshiscene.FuShuSceneData)
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
func checkDoorArea(sd chuangshiscene.FuShuSceneData, sourcePos, afterPos coretypes.Position) bool {
	//TODO 优化3个门
	//城门全破了
	currentDoor := sd.GetCurrentDoor()
	if currentDoor >= numOfDoor {
		return true
	}
	initDefenCamp := sd.GetInitDefendCampType()
	sourceArea := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiWarTemp(initDefenCamp).GetArea(sourcePos)
	destArea := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiWarTemp(initDefenCamp).GetArea(afterPos)
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
func checkAttackedProtectNpcArea(sd chuangshiscene.FuShuSceneData, sourcePos, afterPos coretypes.Position) bool {
	if sd.IsProtectBroken() {
		return true
	}

	initDefenCamp := sd.GetInitDefendCampType()
	sourceFlag := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiWarTemp(initDefenCamp).IsOnProtectArea(sourcePos)
	afterFlag := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiWarTemp(initDefenCamp).IsOnProtectArea(afterPos)

	if sourceFlag || afterFlag {
		return false
	}

	return true
}
