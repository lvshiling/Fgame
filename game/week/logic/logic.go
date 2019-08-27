package logic

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerweek "fgame/fgame/game/week/player"
	weektypes "fgame/fgame/game/week/types"
)

// 周卡特权
func IsSeniorWeek(pl player.Player) bool {
	weekkManager := pl.GetPlayerDataManager(playertypes.PlayerWeekDataManagerType).(*playerweek.PlayerWeekManager)
	weekInfo := weekkManager.GetWeekInfo(weektypes.WeekTypeSenior)
	if weekInfo == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	return weekInfo.IsWeek(now)
}
