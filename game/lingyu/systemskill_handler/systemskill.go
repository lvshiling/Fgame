package sysskill_handler

import (
	"fgame/fgame/game/additionsys/additionsys"
	additionsystypes "fgame/fgame/game/additionsys/types"
	playerlingyu "fgame/fgame/game/lingyu/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/systemcompensate/systemcompensate"
	systemcompensatetypes "fgame/fgame/game/systemcompensate/types"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltypes "fgame/fgame/game/systemskill/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	systemskill.RegisterSystemSkillHandler(sysskilltypes.SystemSkillTypeLingYu, systemskill.SystemSkillHandlerFunc(getLingYuSystemNumber))
	additionsys.RegisterSystemAdvancedHandler(additionsystypes.AdditionSysTypeLingYu, additionsys.SystemAdvancedHandlerFunc(getLingYuSystemNumber))
	welfare.RegisterSystemAdvancedHandler(welfaretypes.AdvancedTypeLingyu, welfare.SystemAdvancedHandlerFunc(getLingYuSystemNumber))
	systemcompensate.RegisterSystemCompensateHandler(systemcompensatetypes.SystemCompensateTypeLingYu, systemcompensate.SystemCompensateHandlerFunc(getLingYuSystemNumber))

}

//获取领域系统阶数
func getLingYuSystemNumber(pl player.Player) (number int32) {
	manager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	return manager.GetLingyuAdvanced()
}
