package sysskill_handler

import (
	"fgame/fgame/game/additionsys/additionsys"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playershihunfan "fgame/fgame/game/shihunfan/player"
	"fgame/fgame/game/systemcompensate/systemcompensate"
	systemcompensatetypes "fgame/fgame/game/systemcompensate/types"
	"fgame/fgame/game/systemskill/systemskill"
	sysskilltypes "fgame/fgame/game/systemskill/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	systemskill.RegisterSystemSkillHandler(sysskilltypes.SystemSkillTypeShiHunFan, systemskill.SystemSkillHandlerFunc(getShiHunFanSystemNumber))
	additionsys.RegisterSystemAdvancedHandler(additionsystypes.AdditionSysTypeShiHunFan, additionsys.SystemAdvancedHandlerFunc(getShiHunFanSystemNumber))
	welfare.RegisterSystemAdvancedHandler(welfaretypes.AdvancedTypeShiHunFan, welfare.SystemAdvancedHandlerFunc(getShiHunFanSystemNumber))
	systemcompensate.RegisterSystemCompensateHandler(systemcompensatetypes.SystemCompensateTypeShiHunFan, systemcompensate.SystemCompensateHandlerFunc(getShiHunFanSystemNumber))

}

//获取噬魂幡系统阶数
func getShiHunFanSystemNumber(pl player.Player) (number int32) {
	manager := pl.GetPlayerDataManager(types.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	return manager.GetShiHunFanAdvanced()
}
