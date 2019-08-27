package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	goaltemplate "fgame/fgame/game/welfare/huhu/goal/template"
	goaltypes "fgame/fgame/game/welfare/huhu/goal/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

//运营活动目标任务
func questCommit(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	questId, ok := data.(int32)
	if !ok {
		return
	}

	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}

	if questTemplate.GetQuestType() != questtypes.QuestTypeYunYingGoal {
		return
	}

	typ := welfaretypes.OpenActivityTypeHuHu
	subType := welfaretypes.OpenActivitySpecialSubTypeGoal
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp := groupInterface.(*goaltemplate.GroupTemplateGoal)
		if !groupTemp.IsGoalQuest(questId) {
			continue
		}

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*goaltypes.GoalInfo)
		info.ReachGoal()
		record := info.GetRewRecord()
		goalCount := info.GoalCount
		welfareManager.UpdateObj(obj)

		goalId := groupTemp.GetGolaId()
		startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
		scMsg := pbutil.BuildSCOpenActivityGetInfoGoal(groupId, startTime, endTime, record, goalId, goalCount)
		pl.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestCommit, event.EventListenerFunc(questCommit))
}
