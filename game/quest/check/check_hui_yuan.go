package check

import (
	playerhuiyuan "fgame/fgame/game/huiyuan/player"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeBuyHuiYuan, quest.CheckHandlerFunc(handleQuestFinishBuyHuiYuan))
}

//check 购买至尊会员
func handleQuestFinishBuyHuiYuan(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理完成购买至尊会员")

	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	huiyuanManager := pl.GetPlayerDataManager(types.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	if huiyuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypePlus) {
		flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), 0, 1)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理完成购买至尊会员,完成")
	return nil
}
