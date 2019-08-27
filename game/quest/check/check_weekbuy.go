package check

import (
	global "fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	playerweek "fgame/fgame/game/week/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeWeekBuy, quest.CheckHandlerFunc(handleQuestFinishWeekBuy))
}

//check 购买周卡
func handleQuestFinishWeekBuy(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理完成购买周卡")

	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	weekManager := pl.GetPlayerDataManager(types.PlayerWeekDataManagerType).(*playerweek.PlayerWeekManager)
	weekInfoMap := weekManager.GetWeekInfoMap()
	now := global.GetGame().GetTimeService().Now()
	for _, info := range weekInfoMap {
		if info.IsWeek(now) {
			flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), 0, 1)
			if !flag {
				panic("quest:设置 SetQuestData 应该成功")
			}
		}
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理完成购买周卡,完成")
	return nil
}
