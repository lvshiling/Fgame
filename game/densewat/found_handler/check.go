package found_handler

import (
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/found/found"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	"fgame/fgame/pkg/timeutils"
)

func init() {
	found.RegistFoundCheckHandler(foundtypes.FoundResourceTypeDenseWat, found.FoundCheckHandlerFunc(checkFoundBack))
}

func checkFoundBack(pl player.Player) bool {
	now := global.GetGame().GetTimeService().Now()
	preDay, _ := timeutils.PreDayOfTime(now)
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	timeTemp := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeDenseWat).GetOnDateTimeTemplate(preDay, openTime, mergeTime)
	if timeTemp == nil {
		return false
	}
	return true
}
