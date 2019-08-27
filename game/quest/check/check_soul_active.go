package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	playersoul "fgame/fgame/game/soul/player"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeSoulActive, quest.CheckHandlerFunc(handleSoulActive))
}

//check 帝魂激活
func handleSoulActive(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理帝魂激活")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for soulTag, _ := range questDemandMap {
		manager := pl.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
		soulAllMap := manager.GetSoulInfoAll()

		for curSoulTag, _ := range soulAllMap {
			if int32(curSoulTag) == soulTag {
				questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
				flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), soulTag, 1)
				if !flag {
					panic("quest:设置 SetQuestData 应该成功")
				}
				break
			}
		}

		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理帝魂激活,完成")
	return nil
}
