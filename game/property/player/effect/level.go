package effect

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	gametemplate "fgame/fgame/game/template"
)

func init() {

	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeLevel, LevelPropertyEffect)

}

//等级作用器
func LevelPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	propertyManager := p.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	level := propertyManager.GetLevel()
	//获取等级模板
	levelTemplate, _ := template.GetTemplateService().Get(int(level), (*gametemplate.CharacterLevelTemplate)(nil)).(*gametemplate.CharacterLevelTemplate)

	prop.SetBase(propertytypes.BattlePropertyTypeMoveSpeed, int64(levelTemplate.SpeedMove))
	prop.SetBase(propertytypes.BattlePropertyTypeAttack, int64(levelTemplate.BaseAttack))
	prop.SetBase(propertytypes.BattlePropertyTypeDefend, int64(levelTemplate.BaseDefense))
	prop.SetBase(propertytypes.BattlePropertyTypeCrit, int64(levelTemplate.BaseCritical))
	prop.SetBase(propertytypes.BattlePropertyTypeTough, int64(levelTemplate.BaseTough))
	prop.SetBase(propertytypes.BattlePropertyTypeBlock, int64(levelTemplate.BaseBlock))
	prop.SetBase(propertytypes.BattlePropertyTypeAbnormality, int64(levelTemplate.BaseBreak))
	prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, int64(levelTemplate.BaseMaxhp))
	prop.SetBase(propertytypes.BattlePropertyTypeMaxTP, int64(levelTemplate.BaseMaxtp))

}
