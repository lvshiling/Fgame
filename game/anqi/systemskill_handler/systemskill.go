package sysskill_handler

import (
	"fgame/fgame/game/additionsys/additionsys"
	additionsystypes "fgame/fgame/game/additionsys/types"
	playeranqi "fgame/fgame/game/anqi/player"
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
	systemskill.RegisterSystemSkillHandler(sysskilltypes.SystemSkillTypeAnQi, systemskill.SystemSkillHandlerFunc(getAnQiSystemNumber))
	additionsys.RegisterSystemAdvancedHandler(additionsystypes.AdditionSysTypeAnqiJiguan, additionsys.SystemAdvancedHandlerFunc(getAnQiSystemNumber))
	welfare.RegisterSystemAdvancedHandler(welfaretypes.AdvancedTypeAnqi, welfare.SystemAdvancedHandlerFunc(getAnQiSystemNumber))
	systemcompensate.RegisterSystemCompensateHandler(systemcompensatetypes.SystemCompensateTypeAnQi, systemcompensate.SystemCompensateHandlerFunc(getAnQiSystemNumber))

}

//获取暗器系统阶数
func getAnQiSystemNumber(pl player.Player) (number int32) {
	manager := pl.GetPlayerDataManager(types.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	return manager.GetAnqiAdvanced()
}
