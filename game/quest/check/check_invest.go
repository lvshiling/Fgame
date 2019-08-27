package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	investnewleveltypes "fgame/fgame/game/welfare/invest/new_level/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeBuyInvest, quest.CheckHandlerFunc(handleQuestFinishBuyInvest))
}

//check 购买投资计划
func handleQuestFinishBuyInvest(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理完成购买投资计划")

	typ := welfaretypes.OpenActivityTypeInvest
	subType := welfaretypes.OpenActivityInvestSubTypeNewLevel

	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		obj := welfareManager.GetOpenActivity(timeTemp.Group)
		if obj == nil {
			continue
		}

		info := obj.GetActivityData().(*investnewleveltypes.InvestNewLevelInfo)
		if info.IsBuy() {
			flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), 0, 1)
			if !flag {
				panic("quest:设置 SetQuestData 应该成功")
			}

			break
		}
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理完成购买投资计划,完成")
	return nil
}
