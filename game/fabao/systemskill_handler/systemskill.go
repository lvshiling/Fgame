package sysskill_handler

import (
	"fgame/fgame/game/additionsys/additionsys"
	additionsystypes "fgame/fgame/game/additionsys/types"
	playerfabao "fgame/fgame/game/fabao/player"
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
	systemskill.RegisterSystemSkillHandler(sysskilltypes.SystemSkillTypeFaBao, systemskill.SystemSkillHandlerFunc(getFaBaoSystemNumber))
	additionsys.RegisterSystemAdvancedHandler(additionsystypes.AdditionSysTypeFaBao, additionsys.SystemAdvancedHandlerFunc(getFaBaoSystemNumber))
	welfare.RegisterSystemAdvancedHandler(welfaretypes.AdvancedTypeFaBao, welfare.SystemAdvancedHandlerFunc(getFaBaoSystemNumber))
	systemcompensate.RegisterSystemCompensateHandler(systemcompensatetypes.SystemCompensateTypeFaBao, systemcompensate.SystemCompensateHandlerFunc(getFaBaoSystemNumber))

}

//获取法宝系统阶数
func getFaBaoSystemNumber(pl player.Player) (number int32) {
	manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	return manager.GetFaBaoAdvancedId()
}
