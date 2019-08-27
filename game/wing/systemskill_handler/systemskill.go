package sysskill_handler

import (
	"fgame/fgame/game/additionsys/additionsys"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/systemcompensate/systemcompensate"
	systemcompensatetypes "fgame/fgame/game/systemcompensate/types"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltypes "fgame/fgame/game/systemskill/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
	playerwing "fgame/fgame/game/wing/player"
)

func init() {
	systemskill.RegisterSystemSkillHandler(sysskilltypes.SystemSkillTypeWing, systemskill.SystemSkillHandlerFunc(getWingSystemNumber))
	additionsys.RegisterSystemAdvancedHandler(additionsystypes.AdditionSysTypeWingStone, additionsys.SystemAdvancedHandlerFunc(getWingSystemNumber))
	welfare.RegisterSystemAdvancedHandler(welfaretypes.AdvancedTypeWing, welfare.SystemAdvancedHandlerFunc(getWingSystemNumber))
	systemcompensate.RegisterSystemCompensateHandler(systemcompensatetypes.SystemCompensateTypeWing, systemcompensate.SystemCompensateHandlerFunc(getWingSystemNumber))

	welfare.RegisterSystemAdvancedHandler(welfaretypes.AdvancedTypeFeather, welfare.SystemAdvancedHandlerFunc(getFeatherSystemNumber))
	systemcompensate.RegisterSystemCompensateHandler(systemcompensatetypes.SystemCompensateTypeFeather, systemcompensate.SystemCompensateHandlerFunc(getFeatherSystemNumber))

}

//获取战翼系统阶数
func getWingSystemNumber(pl player.Player) (number int32) {
	manager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	return manager.GetWingAdvancedId()
}

//获取仙羽系统阶数
func getFeatherSystemNumber(pl player.Player) (number int32) {
	manager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	return manager.GetWingInfo().FeatherId
}
