package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopeneventtypes "fgame/fgame/game/funcopen/event/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家功能开启变化
func playerFuncOpenChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	funcOpenType := data.(funcopentypes.FuncOpenType)

	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	//活跃度任务
	manager.CheckLivenessQuest()
	questList := questlogic.CheckInitQuestList(pl)
	//日环任务
	quest := manager.CheckDailyQuest(funcOpenType)
	if quest != nil {
		questList = append(questList, quest)
	}
	if len(questList) == 0 {
		return
	}

	if funcOpenType == funcopentypes.FuncOpenTypeDailyQuest {
		dailyTag := questtypes.QuestDailyTagPerson
		dailyObj := manager.GetDailyObj(dailyTag)
		scQuestDailSeq := pbutil.BuildSCQuestDailySeq(int32(dailyTag), dailyObj.GetSeqId(), dailyObj.GetTimes())
		pl.SendMsg(scQuestDailSeq)
	} else if funcOpenType == funcopentypes.FuncOpenTypeAllianceDaily {

		dailyTag := questtypes.QuestDailyTagAlliance
		dailyObj := manager.GetDailyObj(dailyTag)
		scQuestDailSeq := pbutil.BuildSCQuestDailySeq(int32(dailyTag), dailyObj.GetSeqId(), dailyObj.GetTimes())
		pl.SendMsg(scQuestDailSeq)
	}

	scQuestListUpdate := pbutil.BuildSCQuestListUpdate(questList)
	pl.SendMsg(scQuestListUpdate)
	return
}

func init() {
	gameevent.AddEventListener(funcopeneventtypes.EventTypeFuncOpen, event.EventListenerFunc(playerFuncOpenChanged))
}
