package sysskill_handler

import (
	"fgame/fgame/game/additionsys/additionsys"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playershenfa "fgame/fgame/game/shenfa/player"
	"fgame/fgame/game/systemcompensate/systemcompensate"
	systemcompensatetypes "fgame/fgame/game/systemcompensate/types"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltypes "fgame/fgame/game/systemskill/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	systemskill.RegisterSystemSkillHandler(sysskilltypes.SystemSkillTypeShenFa, systemskill.SystemSkillHandlerFunc(getShenFaSystemNumber))
	additionsys.RegisterSystemAdvancedHandler(additionsystypes.AdditionSysTypeShenFa, additionsys.SystemAdvancedHandlerFunc(getShenFaSystemNumber))
	welfare.RegisterSystemAdvancedHandler(welfaretypes.AdvancedTypeShenfa, welfare.SystemAdvancedHandlerFunc(getShenFaSystemNumber))
	systemcompensate.RegisterSystemCompensateHandler(systemcompensatetypes.SystemCompensateTypeShenFa, systemcompensate.SystemCompensateHandlerFunc(getShenFaSystemNumber))

}

//获取身法系统阶数
func getShenFaSystemNumber(pl player.Player) (number int32) {
	manager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	return manager.GetShenfaAdvanced()
}
