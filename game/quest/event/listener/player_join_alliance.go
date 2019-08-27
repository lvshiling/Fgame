package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家加入仙盟
func playerJoinAlliance(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeJoinAlliance, 0, 1)
	if err != nil {
		return
	}

	//玩家加入仙盟随仙盟日常任务
	dailyTag := questtypes.QuestDailyTagAlliance
	manager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	quest, flag := manager.AcceptDailyQuest(dailyTag)
	if !flag {
		return
	}
	dailyObj := manager.GetDailyObj(dailyTag)
	if dailyObj == nil {
		return
	}
	scQuestDailSeq := pbutil.BuildSCQuestDailySeq(int32(dailyTag), dailyObj.GetSeqId(), dailyObj.GetTimes())
	pl.SendMsg(scQuestDailSeq)
	scQuestUpdate := pbutil.BuildSCQuestUpdate(quest)
	pl.SendMsg(scQuestUpdate)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypePlayerAllianceJoin, event.EventListenerFunc(playerJoinAlliance))
}
