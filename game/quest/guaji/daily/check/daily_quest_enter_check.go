package check

import (
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
)

func init() {
	guaji.RegisterGuaJiEnterCheckHandler(guajitypes.GuaJiTypeDailyQuest, guaji.GuaJiEnterCheckHandlerFunc(dailyQuestEnterCheck))
}

func dailyQuestEnterCheck(pl player.Player) bool {
	m := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	q := m.GetCurrentDailyQuest()
	if q == nil {
		return false
	}
	return true
}
