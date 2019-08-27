package sysskill_handler

import (
	"fgame/fgame/game/additionsys/additionsys"
	additionsystypes "fgame/fgame/game/additionsys/types"
	playermount "fgame/fgame/game/mount/player"
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
	systemskill.RegisterSystemSkillHandler(sysskilltypes.SystemSkillTypeMount, systemskill.SystemSkillHandlerFunc(getMountSystemNumber))
	additionsys.RegisterSystemAdvancedHandler(additionsystypes.AdditionSysTypeMountEquip, additionsys.SystemAdvancedHandlerFunc(getMountSystemNumber))
	welfare.RegisterSystemAdvancedHandler(welfaretypes.AdvancedTypeMount, welfare.SystemAdvancedHandlerFunc(getMountSystemNumber))
	systemcompensate.RegisterSystemCompensateHandler(systemcompensatetypes.SystemCompensateTypeMount, systemcompensate.SystemCompensateHandlerFunc(getMountSystemNumber))
}

//获取坐骑系统阶数
func getMountSystemNumber(pl player.Player) (number int32) {
	manager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	return manager.GetMountAdvancedId()
}
