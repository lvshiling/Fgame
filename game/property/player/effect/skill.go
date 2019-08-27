package effect

import (
	"fgame/fgame/game/player"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	skilltypes "fgame/fgame/game/skill/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeSkill, SkillPropertyEffect)
}

//技能作用器
func SkillPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {

	for tSkill := range skilltypes.GetSkillModulePropertyEffectoryMap() {
		skillByModulePropertyEffect(p, tSkill, prop)
	}

}
