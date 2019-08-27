package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	skilltemplate "fgame/fgame/game/skill/template"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeActiveSkill, quest.CheckHandlerFunc(handleQuestFinishActiveSkill))
}

//check 激活技能
func handleQuestFinishActiveSkill(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理激活技能")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	for demandId, needNum := range questDemandMap {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(demandId)
		if skillTemplate == nil {
			return
		}
		skillTypeId := skillTemplate.TypeId
		needLevel := skillTemplate.Lev

		skillMap := pl.GetAllSkills()
		skillObj, exist := skillMap[skillTypeId]
		if !exist {
			return
		}
		curLevel := skillObj.GetLevel()
		if curLevel < needLevel {
			return
		}

		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), demandId, needNum)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理激活技能,完成")
	return nil
}
