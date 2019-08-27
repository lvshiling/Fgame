package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	goaltemplate "fgame/fgame/game/welfare/huhu/goal/template"
	goaltypes "fgame/fgame/game/welfare/huhu/goal/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeGoal, welfare.InfoGetHandlerFunc(getGoalInfo))
}

//获取活动目标请求逻辑
func getGoalInfo(pl player.Player, groupId int32) (err error) {
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare: 模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	groupTemp := groupInterface.(*goaltemplate.GroupTemplateGoal)

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := welfareManager.GetOpenActivity(groupId)
	var record []int32
	goalCount := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*goaltypes.GoalInfo)
		goalCount = info.GoalCount
		record = info.GetRewRecord()
	}
	goalId := groupTemp.GetGolaId()

	//初始化运营目标任务
	if welfarelogic.IsOnActivityTime(groupId) {
		questlogic.InitYunYingGoalQuest(pl, groupTemp.GetGolaId())
	}

	scMsg := pbutil.BuildSCOpenActivityGetInfoGoal(groupId, startTime, endTime, record, goalId, goalCount)
	pl.SendMsg(scMsg)
	return
}
