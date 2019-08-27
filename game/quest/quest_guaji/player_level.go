package quest_guaji

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	xianfutypes "fgame/fgame/game/xianfu/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypePlayerLevel, guaji.QuestGuaJiFunc(playerLevel))
}

func playerLevel(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

	questManager := p.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	if !questManager.IfDailyQuestFinish(questtypes.QuestDailyTagPerson) {
		//进入日环任务
		p.EnterGuaJi(scenetypes.GuaJiTypeDailyQuest)
		return true
	}
	log.Info("quest_guaji:日环任务已经做完")
	flag := doXianFu(p, xianfutypes.XianfuTypeExp)
	if flag {
		return true
	}
	log.Info("quest_guaji:经验副本已经做完")
	return false
}
