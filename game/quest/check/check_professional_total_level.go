package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	playerskill "fgame/fgame/game/skill/player"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeProfessionalSkillTotalLevel, quest.CheckHandlerFunc(handleProfessionalTotalLevel))
}

//check 职业技能总等级
func handleProfessionalTotalLevel(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理职业技能总等级")

	manager := pl.GetPlayerDataManager(types.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	totalLevel := manager.GetTotaolLevel()
	if totalLevel == 0 {
		return
	}
	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), 0, totalLevel)
	if !flag {
		panic("quest:设置 SetQuestData 应该成功")
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理职业技能总等级,完成")
	return nil
}
