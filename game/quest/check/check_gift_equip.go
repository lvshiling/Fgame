package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	discountzhuanshengtypes "fgame/fgame/game/welfare/discount/zhuansheng/types"
	playerwelfare "fgame/fgame/game/welfare/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeBuyEquipGift, quest.CheckHandlerFunc(handleQuestFinishBuyEquipGift))
}

//check 购买绝版首饰
func handleQuestFinishBuyEquipGift(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理完成购买投资计划")

	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(questtypes.QiYuEquipGiftGroupId)
	if obj == nil {
		return
	}

	info := obj.GetActivityData().(*discountzhuanshengtypes.DiscountZhuanShengInfo)
	if info.IsBuy(questtypes.QiYuEquipGiftIndex) {
		flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), 0, 1)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}

	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理完成购买投资计划,完成")
	return nil
}
