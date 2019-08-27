package check

import (
	playermaterial "fgame/fgame/game/material/player"
	materialtemplate "fgame/fgame/game/material/template"
	materialtypes "fgame/fgame/game/material/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypechallengeSpecialMaterialFuBen, quest.CheckHandlerFunc(handleQuestFinishSpecialMaterial))
}

//check 完成X次指定材料副本
func handleQuestFinishSpecialMaterial(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理完成X次指定材料副本")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for typ, needNum := range questDemandMap {
		materialType := materialtypes.MaterialType(typ)
		if !materialType.Valid() {
			return
		}
		manager := pl.GetPlayerDataManager(types.PlayerMaterialDataManagerType).(*playermaterial.PlayerMaterialDataManager)
		materialObj := manager.GetPlayerMaterialInfo(materialType)
		if materialObj == nil {
			return
		}
		useTimes := materialObj.GetUseTimes()
		materialTemplate := materialtemplate.GetMaterialTemplateService().GetMaterialTemplate(materialType)
		if materialTemplate == nil {
			return
		}
		allTimes := materialTemplate.AllTimes
		freeTimes := materialTemplate.Free
		leftNum := allTimes - useTimes
		leftFreeNum := freeTimes - useTimes

		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		if leftFreeNum < needNum {
			if !questTemplate.IsAutoFinishByUsedFree() {
				return
			}
			flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), typ, needNum)
			if !flag {
				panic("quest:设置 SetQuestData 应该成功")
			}
			return
		} else {
			if !questTemplate.IsAutoFinishByUsedFree() {
				return
			}
			flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), typ, allTimes-leftNum)
			if !flag {
				panic("quest:设置 SetQuestData 应该成功")
			}
		}

		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理完成X次指定材料副本,完成")
	return nil
}
