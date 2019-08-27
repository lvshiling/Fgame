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
	quest.RegisterCheck(questtypes.QuestSubTypeSkillXinFa, quest.CheckHandlerFunc(checkActivateSkillXinFaNum))
}

// 处理激活技能心法的数量
func checkActivateSkillXinFaNum(pl player.Player, questTemp *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest: 处理激活技能心法的数量")

	questDemandMap := questTemp.GetQuestDemandMap(pl.GetRole())
	manager := pl.GetPlayerDataManager(types.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	skillMap := manager.GetRoleSkillMap()
	skillXinFaNum := 0
	for _, obj := range skillMap {
		if len(obj.TianFuMap) != 0 {
			skillXinFaNum += len(obj.TianFuMap)
		}
	}
	for demandId, _ := range questDemandMap {
		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		flag := questManager.SetQuestData(int32(questTemp.TemplateId()), demandId, int32(skillXinFaNum))
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest: 处理激活技能心法的数量,完成")

	return
}
