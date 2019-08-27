package effect

import (
	"fgame/fgame/game/alliance/alliance"
	playeralliance "fgame/fgame/game/alliance/player"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeHuFu, HuFuPropertyEffect)
}

//虎符
func HuFuPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {

	allianceManager := p.GetPlayerDataManager(types.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	allianceId := allianceManager.GetAllianceId()
	if allianceId == 0 {
		return
	}
	mem := alliance.GetAllianceService().GetAllianceMember(p.GetId())
	if mem == nil {
		return
	}
	if p.GetScene() == nil {
		return
	}
	mapType := p.GetScene().MapTemplate().GetMapType()
	//判断是否在城战
	if mapType != scenetypes.SceneTypeChengZhan && mapType != scenetypes.SceneTypeHuangGong {
		return
	}
	huFu := mem.GetAlliance().GetAllianceObject().GetHuFu()
	percentPerHuFu := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAlliancePropertyPerHuFu))

	huFuPercent := huFu * percentPerHuFu
	prop.SetGlobalPercent(propertytypes.BattlePropertyTypeMaxHP, huFuPercent)
	prop.SetGlobalPercent(propertytypes.BattlePropertyTypeAttack, huFuPercent)
	prop.SetGlobalPercent(propertytypes.BattlePropertyTypeDefend, huFuPercent)

}
