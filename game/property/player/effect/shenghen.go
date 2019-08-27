package effect

import (
	additionsystypes "fgame/fgame/game/additionsys/types"
	// funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	skilltypes "fgame/fgame/game/skill/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeShengHen, shengHenPropertyEffect)
}

//圣痕作用器
func shengHenPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {

	//附加系统属性
	additionSysPropertyEffect(p, additionsystypes.AdditionSysTypeShengHenQingLong, prop)
	additionSysPropertyEffect(p, additionsystypes.AdditionSysTypeShengHenBaiHu, prop)
	additionSysPropertyEffect(p, additionsystypes.AdditionSysTypeShengHenZhuQue, prop)
	additionSysPropertyEffect(p, additionsystypes.AdditionSysTypeShengHenXuanWu, prop)

	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeShengHenQingLong, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeShengHenBaiHu, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeShengHenZhuQue, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeShengHenXuanWu, prop)
}
