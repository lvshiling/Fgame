package check

import (
	playeralliance "fgame/fgame/game/alliance/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeJoinAlliance, quest.CheckHandlerFunc(handleQuestJoinAlliance))
}

//check 加入仙盟
func handleQuestJoinAlliance(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理加入仙盟")

	allianceManager := pl.GetPlayerDataManager(types.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	allianceId := allianceManager.GetAllianceId()
	if allianceId == 0 {
		return
	}
	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), 0, 1)
	if !flag {
		panic("quest:设置 SetQuestData 应该成功")
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理加入仙盟,完成")
	return nil
}
