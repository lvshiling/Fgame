package check

import (
	playerlingtong "fgame/fgame/game/lingtong/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeLingTongActivateNum, quest.CheckHandlerFunc(handleQuestLingTongActivateNum))
}

//check 灵童激活数量
func handleQuestLingTongActivateNum(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理灵童激活数量")

	lingTongManager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongInfo := lingTongManager.GetLingTong()
	if lingTongInfo == nil || !lingTongInfo.IsActivateSys() {
		return
	}
	lingTongMap := lingTongManager.GetLingTongMap()
	num := int32(len(lingTongMap))
	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), 0, num)
	if !flag {
		panic("quest:设置 SetQuestData 应该成功")
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理灵童激活数量,完成")
	return nil
}
