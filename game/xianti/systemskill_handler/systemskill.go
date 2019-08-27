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
	playerxianti "fgame/fgame/game/xianti/player"
)

func init() {
	systemskill.RegisterSystemSkillHandler(sysskilltypes.SystemSkillTypeXianTi, systemskill.SystemSkillHandlerFunc(getXianTiSystemNumber))
	additionsys.RegisterSystemAdvancedHandler(additionsystypes.AdditionSysTypeXianTi, additionsys.SystemAdvancedHandlerFunc(getXianTiSystemNumber))
	welfare.RegisterSystemAdvancedHandler(welfaretypes.AdvancedTypeXianTi, welfare.SystemAdvancedHandlerFunc(getXianTiSystemNumber))
	systemcompensate.RegisterSystemCompensateHandler(systemcompensatetypes.SystemCompensateTypeXianTi, systemcompensate.SystemCompensateHandlerFunc(getXianTiSystemNumber))

}

//获取仙体系统阶数
func getXianTiSystemNumber(pl player.Player) (number int32) {
	manager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	return manager.GetXianTiAdvancedId()
}
