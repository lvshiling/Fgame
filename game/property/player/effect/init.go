package effect

import (
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/player"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeInit, InitPropertyEffect)
}

//初始作用器
func InitPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	hit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeInitalHit)
	prop.SetBase(propertytypes.BattlePropertyTypeHit, int64(hit))
	moveSpeed := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeInitMoveSpeed)
	prop.SetBase(propertytypes.BattlePropertyTypeMoveSpeed, int64(moveSpeed))
}
