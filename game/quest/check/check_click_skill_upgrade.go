package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	playerskill "fgame/fgame/game/skill/player"
	gametemplate "fgame/fgame/game/template"

	clicktypes "fgame/fgame/game/click/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeClickSkillUpgradeButton, quest.CheckHandlerFunc(handleSkillUpgradeButton))
}

//check 处理技能升级点击事件
func handleSkillUpgradeButton(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理技能升级点击事件")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for demandId, num := range questDemandMap {
		clickType := clicktypes.ClickSubTypeSkill(demandId)
		if clickType != clicktypes.ClickSubTypeSkillUpgrade {
			return
		}
		manager := pl.GetPlayerDataManager(types.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
		isAllFull := manager.IfAllRoleFull()
		if !isAllFull {
			return
		}
		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), demandId, num)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理技能升级点击事件,完成")
	return nil
}
