package guaji

import (
	playeractivity "fgame/fgame/game/activity/player"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

type ActivityCheckGuaJi interface {
	CheckActivity(p player.Player, activityTemplate *gametemplate.ActivityTemplate) bool
}

type ActivityCheckGuaJiFunc func(p player.Player, activityTemplate *gametemplate.ActivityTemplate) bool

func (f ActivityCheckGuaJiFunc) CheckActivity(p player.Player, activityTemplate *gametemplate.ActivityTemplate) bool {
	return f(p, activityTemplate)
}

var (
	activityCheckGuaJiMap = map[activitytypes.ActivityType]ActivityCheckGuaJi{}
)

func getActivityCheckGuaJi(subType activitytypes.ActivityType) ActivityCheckGuaJi {
	guaJi, ok := activityCheckGuaJiMap[subType]
	if !ok {
		return nil
	}
	return guaJi
}

func RegisterActivityCheckGuaJi(subType activitytypes.ActivityType, guaJi ActivityCheckGuaJi) {
	_, ok := activityCheckGuaJiMap[subType]
	if ok {
		panic(fmt.Errorf("重复注册%s活动检测挂机", subType.String()))
	}
	activityCheckGuaJiMap[subType] = guaJi
}

//获取挂机活动模板
func GetGuaJiActivityTemplate(pl player.Player) *gametemplate.ActivityTemplate {
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	for _, activityTemplate := range activitytemplate.GetActivityTemplateService().GetActiveAll() {
		now := global.GetGame().GetTimeService().Now()
		activityTimeTemplate, err := activityTemplate.GetActivityTimeTemplate(now, openTime, mergeTime)
		if err != nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"err":      err,
				}).Error("activity_guaji:活动挂机,错误")
			continue
		}
		if activityTimeTemplate == nil {
			continue
		}
		activeType := activityTemplate.GetActivityType()
		activityCheckGuaJi := getActivityCheckGuaJi(activeType)
		if activityCheckGuaJi == nil {
			continue
		}

		//TODO 检查其它条件
		activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activeType)
		activityManager := pl.GetPlayerDataManager(playertypes.PlayerActivityDataManagerType).(*playeractivity.PlayerActivityDataManager)
		activityManager.RefreshData()

		//每天限制参与次数
		if flag := activityManager.IsHaveTimes(activeType); !flag {
			continue
		}

		//是否满足开启条件
		if !pl.IsFuncOpen(activityTemplate.GetFuncOpenType()) {
			continue
		}

		//判断活动是否在活动时间内
		flag, err := activityTemplate.IsAtActivityTime(now, openTime, mergeTime)
		if err != nil {
			continue
		}

		if !flag {
			continue
		}

		//是否满足等级条件
		if !activityTemplate.IsReacheLevel(pl.GetLevel()) {
			continue
		}

		flag = activityCheckGuaJi.CheckActivity(pl, activityTemplate)
		if !flag {
			continue
		}

		return activityTemplate
	}
	return nil
}
