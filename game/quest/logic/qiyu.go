package logic

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
)

func InitQiYuQuest(pl player.Player, fei int32) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)

	// 是否可以激活奇遇任务
	qiyuTempMap := questtemplate.GetQuestTemplateService().GetQiYuTemplateAll()
	for _, temp := range qiyuTempMap {

		qiyuId := int32(temp.Id)
		if manager.IsFinishQiYu(qiyuId) {
			continue
		}

		// 最大满足条件的
		condition := temp.GetMatchCondition(pl.GetLevel(), pl.GetZhuanSheng(), fei)
		if condition == nil {
			continue
		}

		qiyu, flag := manager.AddQiYuQuest(qiyuId, condition.Level, condition.Zhuan, condition.Fei)
		if !flag {
			continue
		}

		// 奇遇子任务
		var questList []*playerquest.PlayerQuestObject
		for questId, _ := range temp.GetQuestMap() {
			qu := manager.GetQuestById(questId)
			if qu != nil {
				continue
			}

			manager.AddQuest(questId)
			questList = append(questList, CheckInitQuestList(pl)...)
		}
		if len(questList) > 0 {
			scMsgUpdate := pbutil.BuildSCQuestListUpdate(questList)
			pl.SendMsg(scMsgUpdate)
		}

		scMsg := pbutil.BuildSCQuestQiYuNotice(qiyu)
		pl.SendMsg(scMsg)

		manager.NoticeQiYu(qiyuId)
	}
}
