package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questypes "fgame/fgame/game/quest/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	//同步任务数据
	p := target.(player.Player)
	manager := p.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)

	//有取玩家等级 放afterLoad会报错
	err = manager.AfterLoadDailyQuest()
	if err != nil {
		return
	}

	for dailyTag := questypes.QuestDailyTagPerson; dailyTag <= questypes.QuestDailyTagAlliance; dailyTag++ {
		manager.AcceptDailyQuest(dailyTag)
	}

	questlogic.CheckInitQuestList(p)
	questlogic.CheckAcceptQuestList(p)

	err = questlogic.CheckCommitQuestList(p)
	questList := make([]*playerquest.PlayerQuestObject, 0, 16)
	activeQuests := manager.GetQuestMap(questypes.QuestStateActive)
	for _, quest := range activeQuests {
		questList = append(questList, quest)
	}
	acceptQuests := manager.GetQuestMap(questypes.QuestStateAccept)
	for _, quest := range acceptQuests {
		questList = append(questList, quest)
	}
	finishQuests := manager.GetQuestMap(questypes.QuestStateFinish)

	for _, quest := range finishQuests {
		questList = append(questList, quest)
	}

	for dailyTag := questypes.QuestDailyTagPerson;  dailyTag <= questypes.QuestDailyTagAlliance; dailyTag++ {
		dailyObj := manager.GetDailyObj(dailyTag)
		scQuestDailSeq := pbutil.BuildSCQuestDailySeq(int32(dailyTag), dailyObj.GetSeqId(), dailyObj.GetTimes())
		p.SendMsg(scQuestDailSeq)
	}

	finish := false
	mainQuestId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMainQuestId)
	mainQuest := manager.GetQuestById(mainQuestId)
	if mainQuest != nil && mainQuest.QuestState == questypes.QuestStateCommit {
		finish = true
	}
	scQuestList := pbutil.BuildSCQuestList(questList, finish)
	p.SendMsg(scQuestList)

	//屠魔次数
	num, buyNum, leftNum := manager.GetTuMoNum()
	scQuestTuMoNumGet := pbutil.BuildSCQuestTuMoNumGet(num, buyNum, leftNum)
	p.SendMsg(scQuestTuMoNumGet)

	//开服目标
	kaiFuMuBiaoMap := manager.GetKaiFuMuBiaoMap()
	scQuestKaiFuMuBiaoGet := pbutil.BuildSCQuestKaiFuMuBiaoGet(kaiFuMuBiaoMap)
	p.SendMsg(scQuestKaiFuMuBiaoGet)

	//奇遇任务
	qiyuMap := manager.GetQiYuMap()
	scMsg := pbutil.BuildSCQuestQiYuNoticeAll(qiyuMap)
	p.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
