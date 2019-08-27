package system_handler

import (
	playerbodyshield "fgame/fgame/game/bodyshield/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/systemcompensate/systemcompensate"
	systemcompensatetypes "fgame/fgame/game/systemcompensate/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterSystemAdvancedHandler(welfaretypes.AdvancedTypeBodyshield, welfare.SystemAdvancedHandlerFunc(getBodyshieldSystemNumber))
	systemcompensate.RegisterSystemCompensateHandler(systemcompensatetypes.SystemCompensateTypeBodyShield, systemcompensate.SystemCompensateHandlerFunc(getBodyshieldSystemNumber))

	welfare.RegisterSystemAdvancedHandler(welfaretypes.AdvancedTypeShield, welfare.SystemAdvancedHandlerFunc(getShieldSystemNumber))
	systemcompensate.RegisterSystemCompensateHandler(systemcompensatetypes.SystemCompensateTypeShield, systemcompensate.SystemCompensateHandlerFunc(getShieldSystemNumber))

}

//获取护体盾系统阶数
func getBodyshieldSystemNumber(pl player.Player) (number int32) {
	manager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
	return int32(manager.GetBodyShiedInfo().AdvanceId)
}

//获取盾刺系统阶数
func getShieldSystemNumber(pl player.Player) (number int32) {
	manager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
	return int32(manager.GetBodyShiedInfo().ShieldId)
}
