package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	advancedblessfeedbacktypes "fgame/fgame/game/welfare/advanced/bless_feedback/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeAdvanced, welfaretypes.OpenActivityAdvancedSubTypeBlessFeedback, reddot.HandlerFunc(handleRedDotAdvancedBless))
}

//升阶祝福大放送红点(废弃)
func handleRedDotAdvancedBless(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*advancedblessfeedbacktypes.BlessAdvancedInfo)
	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	maxRewNum := int32(0)
	for _, temp := range tempList {
		//对应的升阶日
		curDay := temp.Value1
		if curDay != info.BlessDay {
			continue
		}

		if temp.Value2 > info.AdvancedNum {
			continue
		}
		maxRewNum += 1
	}

	if len(info.RewRecord) >= int(maxRewNum) {
		return
	}

	isNotice = true
	return
}
