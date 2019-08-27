package sysskill_handler

import (
	"fgame/fgame/game/additionsys/additionsys"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/systemcompensate/systemcompensate"
	systemcompensatetypes "fgame/fgame/game/systemcompensate/types"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltypes "fgame/fgame/game/systemskill/types"
	playertianmo "fgame/fgame/game/tianmo/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	systemskill.RegisterSystemSkillHandler(sysskilltypes.SystemSkillTypeTianMo, systemskill.SystemSkillHandlerFunc(getTianMoSystemNumber))
	additionsys.RegisterSystemAdvancedHandler(additionsystypes.AdditionSysTypeTianMoTi, additionsys.SystemAdvancedHandlerFunc(getTianMoSystemNumber))
	welfare.RegisterSystemAdvancedHandler(welfaretypes.AdvancedTypeTianMoTi, welfare.SystemAdvancedHandlerFunc(getTianMoSystemNumber))
	systemcompensate.RegisterSystemCompensateHandler(systemcompensatetypes.SystemCompensateTypeTianMo, systemcompensate.SystemCompensateHandlerFunc(getTianMoSystemNumber))

}

//获取天魔系统阶数
func getTianMoSystemNumber(pl player.Player) (number int32) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerTianMoDataManagerType).(*playertianmo.PlayerTianMoDataManager)
	return manager.GetTianMoAdvanced()
}
