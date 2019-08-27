package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	playerxianfu "fgame/fgame/game/xianfu/player"
	xianfutemplate "fgame/fgame/game/xianfu/template"
	xianfutypes "fgame/fgame/game/xianfu/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeSpecialXianFu, quest.CheckHandlerFunc(handleQuestFinishSpecialXianFu))
}

//check 通关指定秘境仙府X次
func handleQuestFinishSpecialXianFu(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理通关指定秘境仙府X次")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for typ, needNum := range questDemandMap {
		xianfuType := xianfutypes.XianfuType(typ)
		if !xianfuType.Valid() {
			return
		}
		leftNum := int32(0)
		freeTimes := int32(0)
		switch xianfuType {
		case xianfutypes.XianfuTypeSilver,
			xianfutypes.XianfuTypeExp:
			manager := pl.GetPlayerDataManager(types.PlayerXianfuDtatManagerType).(*playerxianfu.PlayerXinafuDataManager)
			xianFuObj := manager.GetPlayerXianfuInfo(xianfuType)
			if xianFuObj == nil {
				return
			}
			xianFuId := xianFuObj.GetXianfuId()
			useTimes := xianFuObj.GetUseTimes()
			freeTimes = xianfutemplate.GetXianfuTemplateService().GetFreePlayTimes(xianfuType, xianFuId)
			leftNum = freeTimes - useTimes
			break
		case xianfutypes.XianfuTypeItem:
			return
		}

		if leftNum < needNum {
			if !questTemplate.IsAutoFinishByUsedFree() {
				return
			}
			questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
			flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), typ, needNum)
			if !flag {
				panic("quest:设置 SetQuestData 应该成功")
			}
			return
		}

		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), typ, freeTimes-leftNum)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理通关指定秘境仙府X次,完成")
	return nil
}
