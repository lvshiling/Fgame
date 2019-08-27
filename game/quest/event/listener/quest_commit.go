package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
	"fgame/fgame/game/quest/quest"
	questtemplate "fgame/fgame/game/quest/template"
)

//任务提交事件
func questCommit(target event.EventTarget, data event.EventData) (err error) {

	//TODO 添加任务
	questId := data.(int32)
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	questList := questlogic.CheckQuestIfInit(pl, questId)

	if len(questList) != 0 {
		// for _, qu := range questList {
		// 	log.WithFields(
		// 		log.Fields{
		// 			"playerId": pl.GetId(),
		// 			"questId":  qu.QuestId,
		// 			"state":    qu.QuestState,
		// 		}).Info("quest:任务更新")
		// }
		scQuestListUpdate := pbutil.BuildSCQuestListUpdate(questList)
		pl.SendMsg(scQuestListUpdate)
	}
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	//任务嵌套任务
	questType := questTemplate.GetQuestType()
	questSubType, exist := questType.QuestNestedSubType()
	if exist {
		//免费次数用完默认完成
		questlogic.IncreaseQuestData(pl, questSubType, 0, 1)
	}

	//后续处理
	quest.CommitFlowHandle(pl, questTemplate)
	return nil
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestCommit, event.EventListenerFunc(questCommit))
}
